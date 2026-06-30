package symbols

import (
	"fmt"
	"go/token"

	"github.com/scip-code/scip/bindings/go/scip"
	"golang.org/x/tools/go/packages"
)

func FromDescriptors(pkg *packages.Package, descriptors ...*scip.Descriptor) string {
	if pkg.Module == nil {
		panic(fmt.Sprintf("Failed to find package for: %+v\n", pkg.PkgPath))
	}

	return scip.VerboseSymbolFormatter.FormatSymbol(&scip.Symbol{
		Scheme: "scip-go",
		Package: &scip.Package{
			Manager: "gomod",
			Name:    pkg.Module.Path,
			Version: pkg.Module.Version,
		},
		Descriptors: descriptors,
	})
}

func RangeFromName(position token.Position, name string, adjust bool) scip.Range {
	var adjustment int32 = 0
	if adjust {
		adjustment = 1
	}

	line := int32(position.Line - 1)
	column := int32(position.Column - 1)
	n := int32(len(name))

	return scip.Range{
		Start: scip.Position{Line: line, Character: column + adjustment},
		End:   scip.Position{Line: line, Character: column + n + adjustment},
	}
}

func FormatCode(v string) string {
	if v == "" {
		return ""
	}

	return fmt.Sprintf("```go\n%s\n```", v)
}
