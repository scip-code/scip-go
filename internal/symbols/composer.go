package symbols

import (
	"go/types"

	"github.com/scip-code/scip/bindings/go/scip"
	"golang.org/x/tools/go/packages"
)

const (
	_scheme  = "scip-go"
	_manager = "gomod"
)

type Composer interface {
	Compose(pkg *packages.Package, obj types.Object) string
}

type ComposerConfig struct {
	DefaultModulePath    string
	DefaultModuleVersion string
}

type composer struct {
	defaultModulePath    string
	defaultModuleVersion string

	compositions map[types.Object]string
	fieldOwners  map[*types.Var]*types.TypeName
}

func NewComposer(cfg ComposerConfig) Composer {
	return &composer{
		defaultModulePath:    cfg.DefaultModulePath,
		defaultModuleVersion: cfg.DefaultModuleVersion,

		compositions: make(map[types.Object]string),
		fieldOwners:  make(map[*types.Var]*types.TypeName),
	}
}

func (c *composer) Compose(pkg *packages.Package, obj types.Object) string {
	if obj == nil || obj.Pkg() == nil {
		return ""
	}

	if composition, ok := c.compositions[obj]; ok {
		return composition
	}

	var descriptors []*scip.Descriptor
	switch obj := obj.(type) {
	case *types.PkgName:
		descriptors = c.describePkgName(obj)
	case *types.TypeName:
		descriptors = c.describeTypeName(obj)
	case *types.Const:
		descriptors = c.describeConst(obj)
	case *types.Func:
		descriptors = c.describeFunc(obj)
	case *types.Var:
		descriptors = c.describeVar(obj)
	}
	if len(descriptors) == 0 {
		return ""
	}

	c.compositions[obj] = scip.VerboseSymbolFormatter.FormatSymbol(&scip.Symbol{
		Scheme:      _scheme,
		Package:     c.pack(pkg),
		Descriptors: descriptors,
	})
	return c.compositions[obj]
}

func (c *composer) pack(pkg *packages.Package) *scip.Package {
	if pkg == nil || pkg.Module == nil {
		return &scip.Package{
			Manager: _manager,
			Name:    c.defaultModulePath,
			Version: c.defaultModuleVersion,
		}
	}

	return &scip.Package{
		Manager: _manager,
		Name:    pkg.Module.Path,
		Version: pkg.Module.Version,
	}
}

func (c *composer) describePkgName(pkgName *types.PkgName) []*scip.Descriptor {
	return []*scip.Descriptor{
		{
			Name:   pkgName.Imported().Path(),
			Suffix: scip.Descriptor_Namespace,
		},
	}
}

func (c *composer) describeTypeName(typeName *types.TypeName) []*scip.Descriptor {
	return []*scip.Descriptor{
		{
			Name:   typeName.Pkg().Path(),
			Suffix: scip.Descriptor_Namespace,
		},
		{
			Name:   typeName.Name(),
			Suffix: scip.Descriptor_Type,
		},
	}
}

func (c *composer) describeConst(constant *types.Const) []*scip.Descriptor {
	return []*scip.Descriptor{
		{
			Name:   constant.Pkg().Path(),
			Suffix: scip.Descriptor_Namespace,
		},
		{
			Name:   constant.Name(),
			Suffix: scip.Descriptor_Term,
		},
	}
}

func (c *composer) describeFunc(fn *types.Func) []*scip.Descriptor {
	sig, ok := fn.Type().(*types.Signature)
	if !ok {
		return nil
	}

	if recv := sig.Recv(); recv != nil {
		recvTypeName := c.nameType(recv.Type())
		if recvTypeName == "" {
			return nil
		}

		return []*scip.Descriptor{
			{
				Name:   fn.Pkg().Path(),
				Suffix: scip.Descriptor_Namespace,
			},
			{
				Name:   recvTypeName,
				Suffix: scip.Descriptor_Type,
			},
			{
				Name:   fn.Name(),
				Suffix: scip.Descriptor_Method,
			},
		}
	}

	return []*scip.Descriptor{
		{
			Name:   fn.Pkg().Path(),
			Suffix: scip.Descriptor_Namespace,
		},
		{
			Name:   fn.Name(),
			Suffix: scip.Descriptor_Method,
		},
	}
}

func (c *composer) nameType(t types.Type) string {
	for {
		switch u := types.Unalias(t).(type) {
		case *types.Pointer:
			t = u.Elem()
		case *types.Named:
			return u.Obj().Name()
		default:
			return ""
		}
	}
}

func (c *composer) describeVar(variable *types.Var) []*scip.Descriptor {
	if !variable.IsField() {
		return []*scip.Descriptor{
			{
				Name:   variable.Pkg().Path(),
				Suffix: scip.Descriptor_Namespace,
			},
			{
				Name:   variable.Name(),
				Suffix: scip.Descriptor_Term,
			},
		}
	}

	owner := c.locateOwner(variable)
	if owner == nil {
		return nil
	}

	return []*scip.Descriptor{
		{
			Name:   variable.Pkg().Path(),
			Suffix: scip.Descriptor_Namespace,
		},
		{
			Name:   owner.Name(),
			Suffix: scip.Descriptor_Type,
		},
		{
			Name:   variable.Name(),
			Suffix: scip.Descriptor_Term,
		},
	}
}

func (c *composer) locateOwner(field *types.Var) *types.TypeName {
	origin := field.Origin()

	if owner, ok := c.fieldOwners[origin]; ok {
		return owner
	}

	scope := field.Pkg().Scope()
	for _, name := range scope.Names() {
		obj := scope.Lookup(name)

		typeName, ok := obj.(*types.TypeName)
		if !ok {
			continue
		}

		named, ok := typeName.Type().(*types.Named)
		if !ok {
			continue
		}

		structType, ok := named.Underlying().(*types.Struct)
		if !ok {
			continue
		}

		for field := range structType.Fields() {
			c.fieldOwners[field.Origin()] = typeName
		}
	}

	return c.fieldOwners[origin]
}
