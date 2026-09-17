package document

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/doc"
	"go/format"
	"go/token"
	"go/types"
	"log/slog"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/scip-code/scip-go/internal/lookup"
	"github.com/scip-code/scip-go/internal/symbols"
	"github.com/scip-code/scip/bindings/go/scip"
	"golang.org/x/tools/go/packages"
)

// indent is used to format struct fields.
const indent = "    "

// IsDocDeprecated reports whether the documentation strings contain a paragraph
// starting with "Deprecated:" per Go convention.
func IsDocDeprecated(docs []string) bool {
	for _, d := range docs {
		for line := range strings.SplitSeq(d, "\n") {
			if strings.HasPrefix(line, "Deprecated:") {
				return true
			}
		}
	}
	return false
}

func NewDocument(
	relative string,
	originAbs string,
	pkg *packages.Package,
	pkgSymbols *lookup.Package,
) *Document {
	return &Document{
		RelativePath: relative,
		originAbs:    originAbs,
		lineLen:      loadLineLengths(originAbs),
		pkg:          pkg,
		pkgSymbols:   pkgSymbols,

		docPkg: &doc.Package{},
	}
}

type Document struct {
	// Document relative path. To be used for scip.Document
	RelativePath string

	// The occurrence for `package foo` at the top of a Go file.
	//   It could be a definition or a reference, depending on the package structure.
	//   It doesn't get traversed in the same way as other parts of the tree,
	//   so we special case it here. It must get added to the occurences when
	//   creating a visitors.fileVisitor
	PackageOccurrence *scip.Occurrence

	// The package this document is contained in
	pkg *packages.Package

	// Hold information for docstrings and pretty linking
	docPkg *doc.Package

	// pkgSymbols maps positions to symbol names within
	// this document.
	pkgSymbols *lookup.Package

	// originAbs is the cleaned, symlink-resolved absolute path of the source file
	// this document represents; lineLen is the byte length of each of its lines
	// (0-indexed), or nil if it couldn't be read. Together they let AppendOccurrence
	// callers reject occurrences whose //line-adjusted range escapes the real file.
	originAbs string
	lineLen   []int

	// lines is the document's source split by line, loaded lazily by lineText:
	// only documents that need a range repair pay to keep their source resident.
	// linesLoaded distinguishes "not tried yet" from "tried and unreadable".
	lines       []string
	linesLoaded bool

	// occurrences accumulates every occurrence routed to this document. Usually
	// that is only its own file's, but a generated file may attribute occurrences
	// here via //line directives (e.g. cgo's rewritten source). extraSymbols
	// accumulates SymbolInformation from the file(s) that map here.
	occurrences  []*scip.Occurrence
	extraSymbols []*scip.SymbolInformation
}

// InBounds reports whether r is a well-formed range that fits within this
// document's source file. It rejects:
//   - negative line/column (a `//line file:N` directive with no column collapses
//     positions to column 0 -> scip -1, which is malformed);
//   - lines past EOF or columns past the line (cgo's `defer C.f(x)` rewrites and
//     inserted `_cgoCheckPointer` thunks land here);
//   - reversed ranges (end before start).
//
// Such occurrences cannot be faithfully represented and are dropped rather than
// emitted with a bogus location (which downstream SCIP consumers reject). A
// document whose source could not be read admits everything (no over-dropping).
func (d *Document) InBounds(r scip.Range) bool {
	sl, sc, el, ec := int(r.Start.Line), int(r.Start.Character), int(r.End.Line), int(r.End.Character)
	if sl < 0 || sc < 0 || el < 0 || ec < 0 {
		return false
	}
	if el < sl || (el == sl && ec < sc) {
		return false
	}
	if d.lineLen == nil {
		return true
	}
	if sl >= len(d.lineLen) || el >= len(d.lineLen) {
		return false
	}
	return sc <= d.lineLen[sl] && ec <= d.lineLen[el]
}

// cgoSourceName matches the text cgo rewrote into a mangled identifier: a
// `C.foo` selector, or a bare identifier for manglings that drop the prefix.
// Anchored, so it only matches at the column the range starts on.
var cgoSourceName = regexp.MustCompile(`^(?:C\.)?[\p{L}_][\p{L}\p{Nd}_]*`)

// RepairRange rebuilds an out-of-bounds range from the source text it actually
// covers, returning the replacement and true when one is available.
//
// A range's width comes from the identifier the type checker sees, and cgo's
// identifiers are mangled: `C.puts` is rewritten to `_Cfunc_puts`, so a range
// anchored at the right column is emitted five characters too wide and can spill
// past the end of the real line. The position is fine; only the width is wrong.
// Measuring the identifier that is genuinely at that column recovers the
// occurrence instead of discarding it.
//
// Only a single-line range starting inside real source can be repaired. A
// synthesized position -- cgo's `defer C.f(x)` wrapper, which lands at
// end-of-line where there is no identifier to measure -- matches nothing and is
// still dropped, so this never invents a location.
func (d *Document) RepairRange(r scip.Range) (scip.Range, bool) {
	if r.Start.Line != r.End.Line || r.Start.Line < 0 || r.Start.Character < 0 {
		return r, false
	}
	line, ok := d.lineText(int(r.Start.Line))
	if !ok || int(r.Start.Character) >= len(line) {
		return r, false
	}
	name := cgoSourceName.FindString(line[r.Start.Character:])
	if name == "" {
		return r, false
	}
	repaired := scip.Range{
		Start: r.Start,
		End: scip.Position{
			Line:      r.Start.Line,
			Character: r.Start.Character + int32(len(name)),
		},
	}
	if !d.InBounds(repaired) {
		return r, false
	}
	return repaired, true
}

// lineText returns 0-indexed line l of this document's source. The source is
// read on first use and cached, so documents that never need a repair keep only
// their line lengths.
func (d *Document) lineText(l int) (string, bool) {
	if !d.linesLoaded {
		d.linesLoaded = true
		if b, err := os.ReadFile(d.originAbs); err == nil {
			d.lines = strings.Split(string(b), "\n")
		}
	}
	if l < 0 || l >= len(d.lines) {
		return "", false
	}
	return strings.TrimSuffix(d.lines[l], "\r"), true
}

// AppendOccurrence records occ against this document. Called from the single
// file-walking goroutine, so no synchronization is required.
func (d *Document) AppendOccurrence(occ *scip.Occurrence) {
	d.occurrences = append(d.occurrences, occ)
}

// AddSymbols records SymbolInformation contributed by a file that maps here.
func (d *Document) AddSymbols(syms []*scip.SymbolInformation) {
	d.extraSymbols = append(d.extraSymbols, syms...)
}

// ToScip renders the accumulated occurrences and symbols as a scip.Document.
func (d *Document) ToScip() *scip.Document {
	occurrences := d.occurrences
	if d.PackageOccurrence != nil {
		occurrences = append([]*scip.Occurrence{d.PackageOccurrence}, occurrences...)
	}
	return &scip.Document{
		Language:     "go",
		RelativePath: d.RelativePath,
		Occurrences:  occurrences,
		Symbols:      d.extraSymbols,
	}
}

// loadLineLengths returns the byte length of each line of path (0-indexed), or
// nil if it can't be read. Used to bounds-check occurrence ranges against the
// real source (cgo rewrites some constructs to positions past the original
// line/EOF; those can't be faithfully represented and are dropped).
func loadLineLengths(path string) []int {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	lines := strings.Split(string(b), "\n")
	lengths := make([]int, len(lines))
	for i, line := range lines {
		lengths[i] = len(strings.TrimSuffix(line, "\r"))
	}
	return lengths
}

func (d *Document) GetSymbol(pos token.Pos) (string, bool) {
	return d.pkgSymbols.GetSymbol(pos)
}

// SetSymbolInformation registers a pre-built SymbolInformation at the given position.
func (d *Document) SetSymbolInformation(
	pos token.Pos, info *scip.SymbolInformation,
) {
	d.pkgSymbols.Set(pos, info)
}

// SetNewSymbol declares a new symbol and tracks it within a Document.
//
// NOTE: Does NOT emit a new occurrence
func (d *Document) SetNewSymbol(
	symbol string,
	parent ast.Node,
	ident *ast.Ident,
) {
	d.SetNewSymbolForPos(symbol, parent, ident, ident.Pos())
}

// SetNewSymbolForPos declares a new symbol and tracks it within a Document
// but allows for an override of the position. Generally speaking, you should use
// DeclareNewSymbol instead (since it will calculate the pos for most cases)
//
// NOTE: Does NOT emit a new occurrence
func (d *Document) SetNewSymbolForPos(
	symbol string,
	parent ast.Node,
	ident *ast.Ident,
	pos token.Pos,
) {
	var displayName string
	var documentation []string
	var sigDoc *scip.Signature
	var def types.Object

	if ident != nil {
		displayName = ident.Name

		def = d.pkg.TypesInfo.Defs[ident]
		if def != nil {
			if signature := typeStringForObject(def); signature != "" {
				sigDoc = &scip.Signature{
					Language: "go",
					Text:     signature,
				}
				// Also render the signature into `documentation`: consumers that predate
				// `signature_documentation` read only that field, and would otherwise lose
				// the signature entirely.
				documentation = append(documentation, symbols.FormatCode(signature))
			}
		}
		var hoverText string
		if hover := d.extractHoverText(parent, ident); hover != "" {
			hoverText = hover
			documentation = append(documentation, hover)
		}
		if genDecl, ok := parent.(*ast.GenDecl); ok && genDecl.Doc != nil {
			blockDoc := strings.TrimSpace(genDecl.Doc.Text())
			if blockDoc != "" && blockDoc != hoverText {
				documentation = append(documentation, blockDoc)
			}
		}
	}

	d.pkgSymbols.Set(pos, &scip.SymbolInformation{
		Symbol:                 symbol,
		Kind:                   symbols.KindForObject(def),
		DisplayName:            displayName,
		Documentation:          documentation,
		SignatureDocumentation: sigDoc,
		Relationships:          []*scip.Relationship{},
	})
}

func (d *Document) extractHoverText(parent ast.Node, node ast.Node) string {
	switch v := node.(type) {
	case *ast.File:
		if v.Doc != nil {
			return strings.TrimSpace(v.Doc.Text())
		} else {
			return fmt.Sprintf("package %s", v.Name.Name)
		}
	case *ast.FuncDecl:
		return strings.TrimSpace(v.Doc.Text())
	case *ast.GenDecl:
		return strings.TrimSpace(v.Doc.Text())
	case *ast.TypeSpec:
		// Typespecs do not have the doc associated with them much
		// of the time. They are often associated with the `type`
		// token itself.
		//
		// This is why we have to pass the declaration node
		doc := strings.TrimSpace(v.Doc.Text())
		if doc == "" && parent != nil {
			doc = d.extractHoverText(nil, parent)
		}

		return doc
	case *ast.ValueSpec:
		doc := strings.TrimSpace(v.Doc.Text() + "\n" + v.Comment.Text())
		if doc == "" && parent != nil {
			doc = d.extractHoverText(nil, parent)
		}

		return doc
	case *ast.Field:
		return strings.TrimSpace(v.Doc.Text() + "\n" + v.Comment.Text())
	case *ast.Ident:
		if genDecl, ok := parent.(*ast.GenDecl); ok {
			for _, spec := range genDecl.Specs {
				switch s := spec.(type) {
				case *ast.TypeSpec:
					if s.Name == v {
						return d.extractHoverText(parent, s)
					}
				case *ast.ValueSpec:
					for _, name := range s.Names {
						if name == v {
							return d.extractHoverText(parent, s)
						}
					}
				}
			}
		}
		if parent != nil {
			return d.extractHoverText(nil, parent)
		}
	}

	return ""
}

// relativeQualifier returns a types.Qualifier that omits the package name for
// types in pkg (the "current" package) but includes it for all other packages.
// This ensures that e.g. embedded fields like io.Reader retain their package prefix.
func relativeQualifier(pkg *types.Package) types.Qualifier {
	return func(other *types.Package) string {
		if pkg == other {
			return ""
		}
		return other.Name()
	}
}

func typeStringForObject(obj types.Object) string {
	qual := relativeQualifier(obj.Pkg())

	switch v := obj.(type) {
	case *types.PkgName:
		return fmt.Sprintf("package %s", v.Name())

	case *types.TypeName:
		return formatTypeDeclaration(v)

	case *types.Var:
		if v.IsField() {
			// TODO(tjdevries) - make this be "(T).F" instead of "struct field F string"
			return fmt.Sprintf("struct %s", quotedTagsToBacktick(types.ObjectString(obj, qual)))
		}

	case *types.Const:
		return fmt.Sprintf("%s = %s", types.ObjectString(v, qual), v.Val())
	}

	return types.ObjectString(obj, qual)
}

var loggedGODEBUGWarning sync.Once

// formatTypeDeclaration returns the full type declaration string for a type,
// e.g. "type T struct{}", "type I interface { ... }", "type U = T", "type Z int32".
func formatTypeDeclaration(obj *types.TypeName) string {
	if obj.IsAlias() {
		return formatAliasDeclaration(obj)
	}

	return fmt.Sprintf("type %s %s", obj.Name(), expandTypeExpr(obj.Pkg(), obj.Type().Underlying()))
}

// formatAliasDeclaration returns the type declaration for alias types.
func formatAliasDeclaration(obj *types.TypeName) string {
	switch ty := obj.Type().(type) {
	case *types.Alias:
		qual := relativeQualifier(obj.Pkg())
		lhs := obj.Name() + formatTypeParamList(ty.TypeParams())
		rhs := ty.Rhs()
		switch rhs.(type) {
		case *types.Alias, *types.Named:
			return fmt.Sprintf("type %s = %s", lhs, types.TypeString(rhs, qual))
		default:
			return fmt.Sprintf("type %s = %s", lhs, expandTypeExpr(obj.Pkg(), rhs))
		}
	default:
		if val := os.Getenv("GODEBUG"); strings.Contains(val, "gotypealias=0") {
			loggedGODEBUGWarning.Do(func() {
				slog.Warn(
					"Running with GODEBUG=gotypealias=0, this may cause incorrect hover docs")
			})
		} else {
			slog.Warn(
				"IsAlias() is true but Type() is not Alias; please report this as a bug",
				"obj", obj.String(), "obj.Type()", ty.String())
		}
	}

	// Fallback for when GODEBUG=gotypealias=0 or unexpected types.
	return fmt.Sprintf("type %s %s", obj.Name(), expandTypeExpr(obj.Pkg(), obj.Type().Underlying()))
}

// formatTypeParamList renders a TypeParamList as "[K C1, V C2]".
// Returns "" when tparams is nil or empty.
func formatTypeParamList(tparams *types.TypeParamList) string {
	if tparams == nil || tparams.Len() == 0 {
		return ""
	}
	var buf strings.Builder
	buf.WriteByte('[')
	for i := range tparams.Len() {
		if i > 0 {
			buf.WriteString(", ")
		}
		tp := tparams.At(i)
		buf.WriteString(tp.Obj().Name())
		buf.WriteByte(' ')
		buf.WriteString(tp.Constraint().String())
	}
	buf.WriteByte(']')
	return buf.String()
}

// expandTypeExpr renders a type expression, formatting struct and interface
// types with aligned fields using go/format.
func expandTypeExpr(pkg *types.Package, t types.Type) string {
	raw := types.TypeString(t, relativeQualifier(pkg))

	switch t.(type) {
	case *types.Struct, *types.Interface:
		if formatted, err := formatGoDecl("type _ " + raw); err == nil {
			return strings.TrimPrefix(formatted, "type _ ")
		}
	}

	return raw
}

// formatGoDecl formats a Go type declaration using go/format, returning
// the formatted result with tabs replaced by spaces.
func formatGoDecl(decl string) (string, error) {
	src := "package p\n\n" + decl + "\n"
	formatted, err := format.Source([]byte(src))
	if err != nil {
		return "", err
	}
	result := strings.TrimPrefix(string(formatted), "package p\n\n")
	result = strings.TrimSpace(result)
	result = strings.ReplaceAll(result, "\t", indent)
	return result, nil
}

// quotedTagsToBacktick replaces double-quoted struct tag strings with
// backtick-quoted equivalents in the output of types.ObjectString.
func quotedTagsToBacktick(s string) string {
	buf := bytes.NewBuffer(make([]byte, 0, len(s)))
	for i := 0; i < len(s); i++ {
		if s[i] != '"' {
			buf.WriteByte(s[i])
			continue
		}
		for j := i + 1; j < len(s); j++ {
			if s[j] == '\\' {
				j++
				continue
			}
			if s[j] == '"' {
				quoted := s[i : j+1]
				if unquoted, err := strconv.Unquote(quoted); err == nil {
					buf.WriteByte('`')
					buf.WriteString(unquoted)
					buf.WriteByte('`')
				} else {
					buf.WriteString(quoted)
				}
				i = j
				break
			}
		}
	}
	return buf.String()
}
