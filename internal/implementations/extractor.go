package impls

import (
	"fmt"
	"go/types"
	"log/slog"

	"github.com/scip-code/scip-go/internal/loader"
	"github.com/scip-code/scip-go/internal/lookup"
	"github.com/scip-code/scip/bindings/go/scip"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/types/typeutil"
)

type Extractor interface {
	Extract(pkgLookup loader.PackageLookup) (interfaces, concretes map[string]ImplDef)
}

type extractor struct {
	global         *lookup.Global
	methodSetCache typeutil.MethodSetCache
}

func NewExtractor(global *lookup.Global) Extractor {
	return &extractor{
		global: global,
	}
}

func (e *extractor) Extract(pkgLookup loader.PackageLookup) (map[string]ImplDef, map[string]ImplDef) {
	interfaces := map[string]ImplDef{}
	concretes := map[string]ImplDef{}

	for _, pkg := range pkgLookup {
		if pkg.Name == "builtin" {
			continue
		}

		if pkg.TypesInfo != nil {
			e.extractLocal(pkg, interfaces, concretes)
		} else if pkg.Types != nil {
			e.extractRemote(pkg, interfaces, concretes)
		} else {
			slog.Warn("No types for package", "path", pkg.PkgPath)
		}
	}

	return interfaces, concretes
}

func (e *extractor) extractLocal(pkg *packages.Package, interfaces, concretes map[string]ImplDef) {
	pkgSymbols := e.global.GetPackage(pkg)
	if pkgSymbols == nil {
		slog.Warn("No symbols for package", "path", pkg.PkgPath)
		return
	}

	for ident, obj := range pkg.TypesInfo.Defs {
		if obj == nil {
			continue
		}

		typeName, ok := obj.(*types.TypeName)
		if !ok {
			continue
		}

		if pkg.Types != nil && typeName.Parent() != pkg.Types.Scope() {
			continue
		}

		named, ok := obj.Type().(*types.Named)
		if !ok {
			continue
		}

		sym, ok := pkgSymbols.Get(typeName.Pos())
		if !ok {
			slog.Debug(
				"No symbol for package-level named type",
				"identifier", ident.Name,
				"package", pkg.PkgPath,
				"id", obj.Id(),
			)
			continue
		}

		e.classify(named, sym, pkg.PkgPath, interfaces, concretes)
	}
}

func (e *extractor) extractRemote(pkg *packages.Package, interfaces, concretes map[string]ImplDef) {
	scope := pkg.Types.Scope()

	for _, name := range scope.Names() {
		typeName, ok := scope.Lookup(name).(*types.TypeName)
		if !ok || !typeName.Exported() {
			continue
		}

		named, ok := typeName.Type().(*types.Named)
		if !ok {
			continue
		}

		sym := e.global.Composer().Compose(pkg, typeName)
		if sym == "" {
			continue
		}

		e.classify(named, &scip.SymbolInformation{Symbol: sym}, pkg.PkgPath, interfaces, concretes)
	}
}

func (e *extractor) classify(
	named *types.Named,
	sym *scip.SymbolInformation,
	pkgPath string,
	interfaces, concretes map[string]ImplDef,
) {
	methods := typeutil.IntuitiveMethodSet(named, &e.methodSetCache)
	if len(methods) == 0 {
		return
	}

	methodSymbols := map[methodID]*scip.SymbolInformation{}
	for _, method := range methods {
		sym, ok, err := e.global.GetSymbolOfObject(method.Obj())
		if err != nil {
			slog.Debug(fmt.Sprintf("Error while looking for symbol %s | %s", err, method.Obj()))
			continue
		}
		if !ok {
			continue
		}

		methodSymbols[methodID(method.Obj().Id())] = sym
	}

	impl := ImplDef{
		Symbol:        sym,
		Named:         named,
		Methods:       methodSymbols,
		Mask:          methodMask(methods),
		MethodCount:   len(methods),
		HasUnexported: hasUnexportedMethods(methods),
		PkgPath:       pkgPath,
	}
	if types.IsInterface(named) {
		interfaces[impl.Symbol.Symbol] = impl
	} else {
		concretes[impl.Symbol.Symbol] = impl
	}
}
