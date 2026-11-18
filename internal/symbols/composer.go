package symbols

import (
	"go/types"
	"sync"

	"github.com/scip-code/scip/bindings/go/scip"
	"golang.org/x/tools/go/packages"
)

type Composer struct {
	defaultModulePath    string
	defaultModuleVersion string

	mu           sync.Mutex
	compositions map[types.Object]string
}

func NewComposer(defaultModulePath, defaultModuleVersion string) *Composer {
	return &Composer{
		defaultModulePath:    defaultModulePath,
		defaultModuleVersion: defaultModuleVersion,

		compositions: make(map[types.Object]string),
	}
}

func (c *Composer) Compose(pkg *packages.Package, obj types.Object) string {
	if obj == nil || obj.Pkg() == nil {
		return ""
	}

	switch o := obj.(type) {
	case *types.Var:
		obj = o.Origin()
	case *types.Func:
		obj = o.Origin()
	}

	c.mu.Lock()
	defer c.mu.Unlock()

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
		Scheme:      "scip-go",
		Package:     c.pack(pkg),
		Descriptors: descriptors,
	})
	return c.compositions[obj]
}

func (c *Composer) pack(pkg *packages.Package) *scip.Package {
	if pkg == nil || pkg.Module == nil {
		return &scip.Package{
			Manager: "gomod",
			Name:    c.defaultModulePath,
			Version: c.defaultModuleVersion,
		}
	}

	return &scip.Package{
		Manager: "gomod",
		Name:    pkg.Module.Path,
		Version: pkg.Module.Version,
	}
}

func (c *Composer) describePkgName(pkgName *types.PkgName) []*scip.Descriptor {
	return []*scip.Descriptor{
		{
			Name:   pkgName.Imported().Path(),
			Suffix: scip.Descriptor_Namespace,
		},
	}
}

func (c *Composer) describeTypeName(typeName *types.TypeName) []*scip.Descriptor {
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

func (c *Composer) describeConst(constant *types.Const) []*scip.Descriptor {
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

func (c *Composer) describeFunc(fn *types.Func) []*scip.Descriptor {
	sig, ok := fn.Type().(*types.Signature)
	if !ok {
		return nil
	}

	if recv := sig.Recv(); recv != nil {
		recvTypeName := c.nameType(recv.Type())
		if recvTypeName == "" {
			return nil
		}
		suffix := scip.Descriptor_Method
		if _, ok := recv.Type().Underlying().(*types.Interface); ok {
			suffix = scip.Descriptor_Term
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
				Suffix: suffix,
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

func (c *Composer) nameType(t types.Type) string {
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

func (c *Composer) describeVar(variable *types.Var) []*scip.Descriptor {
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

func (c *Composer) locateOwner(field *types.Var) *types.TypeName {
	scope := field.Pkg().Scope()
	for _, name := range scope.Names() {
		typeName, ok := scope.Lookup(name).(*types.TypeName)
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

		for f := range structType.Fields() {
			if f == field {
				return typeName
			}
		}
	}

	return nil
}
