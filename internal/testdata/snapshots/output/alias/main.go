  package main
//        ^^^^ definition github.com/scip-code/scip-go 0.1.test `github.com/scip-code/scip-go/internal/testdata/snapshots/input/alias`/
//             kind Package
//             display_name main
//             signature_documentation
//             > package main
//             documentation
//             > ```go
//             > package main
//             > ```
  
  // Check that we don't panic
  // Copied from https://github.com/golang/go/issues/68877#issuecomment-2290000187
  type (
   T struct{}
// ^ definition github.com/scip-code/scip-go 0.1.test `github.com/scip-code/scip-go/internal/testdata/snapshots/input/alias`/T#
//   kind Struct
//   display_name T
//   signature_documentation
//   > type T struct{}
//   documentation
//   > ```go
//   > type T struct{}
//   > ```
//   documentation
//   > Check that we don't panic
//   > Copied from https://github.com/golang/go/issues/68877#issuecomment-2290000187
   U = T
// ^ definition github.com/scip-code/scip-go 0.1.test `github.com/scip-code/scip-go/internal/testdata/snapshots/input/alias`/U#
//   kind TypeAlias
//   display_name U
//   signature_documentation
//   > type U = T
//   documentation
//   > ```go
//   > type U = T
//   > ```
//   documentation
//   > Check that we don't panic
//   > Copied from https://github.com/golang/go/issues/68877#issuecomment-2290000187
//     ^ reference github.com/scip-code/scip-go 0.1.test `github.com/scip-code/scip-go/internal/testdata/snapshots/input/alias`/T#
   V = U
// ^ definition github.com/scip-code/scip-go 0.1.test `github.com/scip-code/scip-go/internal/testdata/snapshots/input/alias`/V#
//   kind TypeAlias
//   display_name V
//   signature_documentation
//   > type V = U
//   documentation
//   > ```go
//   > type V = U
//   > ```
//   documentation
//   > Check that we don't panic
//   > Copied from https://github.com/golang/go/issues/68877#issuecomment-2290000187
//     ^ reference github.com/scip-code/scip-go 0.1.test `github.com/scip-code/scip-go/internal/testdata/snapshots/input/alias`/U#
   S U
// ^ definition github.com/scip-code/scip-go 0.1.test `github.com/scip-code/scip-go/internal/testdata/snapshots/input/alias`/S#
//   kind Struct
//   display_name S
//   signature_documentation
//   > type S struct{}
//   documentation
//   > ```go
//   > type S struct{}
//   > ```
//   documentation
//   > Check that we don't panic
//   > Copied from https://github.com/golang/go/issues/68877#issuecomment-2290000187
//   ^ reference github.com/scip-code/scip-go 0.1.test `github.com/scip-code/scip-go/internal/testdata/snapshots/input/alias`/U#
   Z int32
// ^ definition github.com/scip-code/scip-go 0.1.test `github.com/scip-code/scip-go/internal/testdata/snapshots/input/alias`/Z#
//   kind Type
//   display_name Z
//   signature_documentation
//   > type Z int32
//   documentation
//   > ```go
//   > type Z int32
//   > ```
//   documentation
//   > Check that we don't panic
//   > Copied from https://github.com/golang/go/issues/68877#issuecomment-2290000187
  )
  
//⌄ enclosing_range_start github.com/scip-code/scip-go 0.1.test `github.com/scip-code/scip-go/internal/testdata/snapshots/input/alias`/f().
  func f(u U) {}
//     ^ definition github.com/scip-code/scip-go 0.1.test `github.com/scip-code/scip-go/internal/testdata/snapshots/input/alias`/f().
//       kind Function
//       display_name f
//       signature_documentation
//       > func f(u U)
//       documentation
//       > ```go
//       > func f(u U)
//       > ```
//       ^ definition local 0
//         kind Variable
//         display_name u
//         signature_documentation
//         > var u U
//         ^ reference github.com/scip-code/scip-go 0.1.test `github.com/scip-code/scip-go/internal/testdata/snapshots/input/alias`/U#
//             ⌃ enclosing_range_end github.com/scip-code/scip-go 0.1.test `github.com/scip-code/scip-go/internal/testdata/snapshots/input/alias`/f().
  
