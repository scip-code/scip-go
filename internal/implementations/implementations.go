package impls

import (
	"go/types"
	"hash/crc32"
	"sync"
	"sync/atomic"

	"github.com/scip-code/scip-go/internal/implementations/fingerprint"
	"github.com/scip-code/scip-go/internal/loader"
	"github.com/scip-code/scip-go/internal/output"
	"github.com/scip-code/scip/bindings/go/scip"
)

// methodID is a unique identifier for a method, using types.Id semantics
// (package-path-qualified for unexported methods, just the name for exported).
type methodID string

type ImplDef struct {
	Symbol  *scip.SymbolInformation
	Named   *types.Named
	Methods map[methodID]*scip.SymbolInformation

	// Precomputed fields for fast filtering before calling types.Implements.
	Mask          uint64 // bitmask of hashed method signatures
	MethodCount   int
	HasUnexported bool   // true if any method is unexported
	PkgPath       string // package path, used to skip cross-package checks for unexported methods
}

// methodMask computes a 64-bit bitmask from a set of methods, using
// the same algorithm as gopls's methodsets package. Each method's
// identity (Id + fingerprint) is hashed and mapped to a single bit.
// The resulting mask enables O(1) rejection of non-matching
// type/interface pairs: if ty.Mask & iface.Mask != iface.Mask,
// the type is missing at least one method of the interface.
func methodMask(methods []*types.Selection) uint64 {
	var mask uint64
	var buf []byte
	for _, m := range methods {
		fn := m.Obj().(*types.Func)
		id := fn.Id()
		fp, _ := fingerprint.Encode(fn.Signature())
		buf = append(buf[:0], id...)
		buf = append(buf, fp...)
		sum := crc32.ChecksumIEEE(buf)
		mask |= 1 << uint64(((sum>>24)^(sum>>16)^(sum>>8)^sum)&0x3f)
	}
	return mask
}

// hasUnexportedMethods returns true if any of the given methods is unexported.
func hasUnexportedMethods(methods []*types.Selection) bool {
	for _, m := range methods {
		if !m.Obj().Exported() {
			return true
		}
	}
	return false
}

func findImplementations(concreteTypes map[string]ImplDef, interfaces map[string]ImplDef, count *uint64) {
	for _, ty := range concreteTypes {
		for _, iface := range interfaces {
			ifaceType, ok := iface.Named.Underlying().(*types.Interface)
			if !ok {
				continue
			}

			// Fast filter 1: if the interface has unexported methods,
			// only types in the same package can implement it.
			if iface.HasUnexported && ty.PkgPath != iface.PkgPath {
				continue
			}

			// Fast filter 2: the type must have at least as many methods
			// as the interface requires.
			if ty.MethodCount < iface.MethodCount {
				continue
			}

			// Fast filter 3: bitmask subset check. If the type's mask
			// doesn't cover all bits of the interface's mask, the type
			// is definitely missing at least one required method.
			if ty.Mask&iface.Mask != iface.Mask {
				continue
			}

			// Use types.Implements to correctly check interface satisfaction,
			// handling all edge cases (embedded types, generics, unexported methods).
			if !types.Implements(ty.Named, ifaceType) &&
				!types.Implements(types.NewPointer(ty.Named), ifaceType) {
				continue
			}

			// Add implementation details for the struct & interface relationship
			ty.Symbol.Relationships = append(ty.Symbol.Relationships, &scip.Relationship{
				Symbol:           iface.Symbol.Symbol,
				IsImplementation: true,
			})

			// For all methods, add implementation details as well
			for name, implMethod := range iface.Methods {
				tyMethodInfo, ok := ty.Methods[name]
				if !ok {
					continue
				}

				tyMethodInfo.Relationships = append(tyMethodInfo.Relationships, &scip.Relationship{
					Symbol:           implMethod.Symbol,
					IsImplementation: true,
				})
			}
		}

		ty.Symbol.Relationships = scip.CanonicalizeRelationships(ty.Symbol.Relationships)
		for _, method := range ty.Methods {
			method.Relationships = scip.CanonicalizeRelationships(method.Relationships)
		}

		atomic.AddUint64(count, 1)
	}
}

func AddImplementationRelationships(
	pkgs loader.PackageLookup,
	allPackages loader.PackageLookup,
	extractor Extractor,
) ([]*scip.SymbolInformation, error) {
	var externalSymbols []*scip.SymbolInformation

	localInterfaces, localTypes := extractor.Extract(pkgs)

	remotePackages := make(loader.PackageLookup)
	for pkgID, pkg := range allPackages {
		if _, ok := pkgs[pkgID]; ok {
			continue
		}

		remotePackages[pkgID] = pkg
	}
	remoteInterfaces, remoteTypes := extractor.Extract(remotePackages)

	// Total concrete types to check across the three passes.
	total := uint64(len(localTypes)*2 + len(remoteTypes))
	var count uint64
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		// local type -> local interface
		findImplementations(localTypes, localInterfaces, &count)

		// local type -> remote interface
		findImplementations(localTypes, remoteInterfaces, &count)

		// remote type -> local interface
		// We emit these as external symbols so index consumer can merge them.
		findImplementations(remoteTypes, localInterfaces, &count)
	}()

	output.WithProgressParallel(&wg, "Indexing Implementations", &count, total)

	// Collect remote type symbols that gained relationships
	for _, typ := range remoteTypes {
		if len(typ.Symbol.Relationships) > 0 {
			externalSymbols = append(externalSymbols, typ.Symbol)
		}
	}

	return externalSymbols, nil
}
