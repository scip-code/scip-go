package index

import (
	_ "embed"
	"go/ast"
	"log/slog"
	"maps"
	"net/url"
	"slices"
	"sort"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/scip-code/scip-go/internal/config"
	"github.com/scip-code/scip-go/internal/document"
	impls "github.com/scip-code/scip-go/internal/implementations"
	"github.com/scip-code/scip-go/internal/loader"
	"github.com/scip-code/scip-go/internal/lookup"
	"github.com/scip-code/scip-go/internal/newtypes"
	"github.com/scip-code/scip-go/internal/output"
	"github.com/scip-code/scip-go/internal/symbols"
	"github.com/scip-code/scip-go/internal/visitors"
	"github.com/scip-code/scip/bindings/go/scip"
	"google.golang.org/protobuf/proto"
)

//go:embed version.txt
var versionFile string
var ScipGoVersion = strings.TrimSpace(versionFile)

func GetPackages(opts config.IndexOpts) (current []newtypes.PackageID, deps []newtypes.PackageID, err error) {
	pkgs, pkgLookup, err := loader.LoadPackages(opts, opts.ModuleRoot)
	if err != nil {
		return nil, nil, err
	}

	for name := range pkgs {
		current = append(current, name)
	}

	sort.Slice(current, func(i, j int) bool {
		return current[i] < current[j]
	})

	for name := range pkgLookup {
		deps = append(deps, name)
	}

	sort.Slice(deps, func(i, j int) bool {
		return deps[i] < deps[j]
	})

	return
}

func ListMissing(opts config.IndexOpts) (missing []string, err error) {
	projectPackages, _, err := loader.LoadPackages(opts, opts.ModuleRoot)
	if err != nil {
		return nil, err
	}

	composer := symbols.NewComposer(opts.ModuleRoot, opts.ModuleVersion)
	globalSymbols := lookup.NewGlobalSymbols(composer)

	pathToDocuments := map[string]*document.Document{}
	for _, pkg := range projectPackages {
		visitors.VisitPackageSyntax(
			opts.ModuleRoot, pkg, pathToDocuments, globalSymbols)
	}

	for _, pkg := range projectPackages {
		goFiles := visitors.RealGoFiles(pkg)
		for _, f := range pkg.Syntax {
			origin := visitors.OriginFile(pkg, f.Package)
			if _, isReal := goFiles[origin]; !isReal {
				// Generated file with no real-source origin (e.g. cgo glue).
				continue
			}
			if _, ok := pathToDocuments[origin]; !ok {
				missing = append(missing, origin)
			}
		}
	}

	return missing, nil
}

func Index(writer func(proto.Message) error, opts config.IndexOpts) error {
	// Emit Metadata.
	//   NOTE: Must be the first field emitted
	projectRoot := url.URL{Scheme: "file", Path: opts.ModuleRoot}
	if err := writer(&scip.Metadata{
		Version: scip.ProtocolVersion_UnspecifiedProtocolVersion,
		ToolInfo: &scip.ToolInfo{
			Name:      "scip-go",
			Version:   ScipGoVersion,
			Arguments: opts.Arguments,
		},
		ProjectRoot:          projectRoot.String(),
		TextDocumentEncoding: scip.TextEncoding_UTF8,
	}); err != nil {
		return err
	}

	projectPackages, allPackages, err := loader.LoadPackages(opts, opts.ModuleRoot)
	if err != nil {
		return err
	}

	var externalSymbols []*scip.SymbolInformation
	pathToDocument, globalSymbols := indexVisitPackages(opts, projectPackages, allPackages)
	if !opts.SkipImplementations {
		implSymbols, err := impls.AddImplementationRelationships(
			projectPackages, allPackages, impls.NewExtractor(globalSymbols),
		)
		if err != nil {
			return err
		}
		externalSymbols = implSymbols
	}

	pkgIDs := slices.Sorted(maps.Keys(projectPackages))
	pkgLen := len(pkgIDs)

	var count uint64
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		for _, ID := range pkgIDs {
			pkg := projectPackages[ID]
			pkgSymbols := globalSymbols.GetPackage(pkg)

			for _, file := range pkg.Syntax {
				origin := visitors.OriginFile(pkg, file.Package)
				doc := pathToDocument[origin]
				if doc == nil {
					// No document: a generated file (e.g. cgo's
					// _cgo_gotypes.go) whose occurrences are compiler glue.
					continue
				}

				// The visitor routes each occurrence to the document of its
				// //line-adjusted origin (usually this file, but a generated file
				// can attribute some occurrences to another real source file), so
				// it needs the whole document map, not just this file's document.
				visitor := visitors.NewFileVisitor(
					doc,
					pkg,
					file,
					pkgSymbols,
					globalSymbols,
					pathToDocument,
				)

				// Traverse the file (routing occurrences), then attach this
				// file's symbols to its document.
				ast.Walk(visitor, file)
				visitor.Finish()
			}

			atomic.AddUint64(&count, 1)
		}
	}()

	output.WithProgressParallel(&wg, "Visiting Project Files", &count, uint64(pkgLen))

	// Emit one document per source file -- occurrences routed from every file
	// that maps here are now accumulated -- in a stable, path-sorted order.
	for _, origin := range slices.Sorted(maps.Keys(pathToDocument)) {
		if err := writer(pathToDocument[origin].ToScip()); err != nil {
			return err
		}
	}

	// Emit external symbols for remote types that implement local interfaces
	for _, sym := range externalSymbols {
		if err := writer(sym); err != nil {
			return err
		}
	}

	return nil
}

func indexVisitPackages(
	opts config.IndexOpts,
	projectPackages loader.PackageLookup,
	allPackages loader.PackageLookup,
) (map[string]*document.Document, *lookup.Global) {
	pathToDocuments := map[string]*document.Document{}

	composer := symbols.NewComposer(opts.ModuleRoot, opts.ModuleVersion)
	globalSymbols := lookup.NewGlobalSymbols(composer)
	for _, pkg := range allPackages {
		globalSymbols.SetPkgSymbol(pkg)
	}

	var count uint64
	var wg sync.WaitGroup
	wg.Add(1)

	lookupIDs := slices.Sorted(maps.Keys(projectPackages))

	// Visit project packages to collect definition sites. Dependency packages
	// skip syntax visiting; their symbols are composed on demand.
	go func() {
		defer wg.Done()

		for _, pkgID := range lookupIDs {
			pkg := projectPackages[pkgID]
			slog.Debug("Visiting package", "path", pkg.PkgPath)
			visitors.VisitPackageSyntax(opts.ModuleRoot, pkg, pathToDocuments, globalSymbols)

			// A package may have no parsed source files (e.g. a directory
			// that only contains *_test.go files belonging to an external
			// "*_test" package). There is nothing to attach package symbol
			// information or occurrences to, so skip it.
			if len(pkg.Syntax) == 0 {
				atomic.AddUint64(&count, 1)
				continue
			}

			pkgSymbol, _ := globalSymbols.GetPkgSymbol(pkg)

			pkgSignature := "package " + pkg.Name
			symInfo := &scip.SymbolInformation{
				Symbol:      pkgSymbol,
				Kind:        scip.SymbolInformation_Package,
				DisplayName: pkg.Name,
				Documentation: append(
					[]string{symbols.FormatCode(pkgSignature)},
					findPackageDocs(pkg)...,
				),
				SignatureDocumentation: &scip.Signature{
					Language: "go",
					Text:     pkgSignature,
				},
			}
			// Attach the package symbol to the first real document and a
			// package occurrence to each. Generated files (e.g. cgo's
			// _cgo_gotypes.go) have no document, so skip them.
			pkgDeclared := false
			for _, f := range pkg.Syntax {
				doc := pathToDocuments[visitors.OriginFile(pkg, f.Package)]
				if doc == nil {
					continue
				}
				if !pkgDeclared {
					doc.SetSymbolInformation(f.Name.NamePos, symInfo)
					pkgDeclared = true
				}
				position := pkg.Fset.Position(f.Name.NamePos)
				doc.PackageOccurrence = &scip.Occurrence{
					TypedRange:  symbols.RangeFromName(position, f.Name.Name, false).AsTypedRange(),
					Symbol:      pkgSymbol,
					SymbolRoles: int32(scip.SymbolRole_Definition),
				}
			}

			atomic.AddUint64(&count, 1)
		}
	}()

	output.WithProgressParallel(&wg, "Visiting Packages", &count, uint64(len(lookupIDs)))

	return pathToDocuments, globalSymbols
}
