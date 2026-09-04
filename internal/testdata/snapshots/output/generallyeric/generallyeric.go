  // generallyeric -> generic for short
  package generallyeric
//        ^^^^^^^^^^^^^ definition 0.1.test `sg/generallyeric`/
//                      kind Package
//                      display_name generallyeric
//                      signature_documentation
//                      > package generallyeric
//                      documentation
//                      > ```go
//                      > package generallyeric
//                      > ```
//                      documentation
//                      > generallyeric -> generic for short
  
  import "fmt"
//        ^^^ reference github.com/golang/go/src go1.22 fmt/
  
//⌄ enclosing_range_start 0.1.test `sg/generallyeric`/Print().
  func Print[T any](s []T) {
//     ^^^^^ definition 0.1.test `sg/generallyeric`/Print().
//           kind Function
//           display_name Print
//           signature_documentation
//           > func Print[T any](s []T)
//           documentation
//           > ```go
//           > func Print[T any](s []T)
//           > ```
//           ^ definition local 0
//             kind Interface
//             display_name T
//             signature_documentation
//             > type parameter T any
//                  ^ definition local 1
//                    kind Variable
//                    display_name s
//                    signature_documentation
//                    > var s []T
//                      ^ reference local 0
   for _, v := range s {
//        ^ definition local 2
//          kind Variable
//          display_name v
//          signature_documentation
//          > var v T
//                   ^ reference local 1
    fmt.Print(v)
//  ^^^ reference github.com/golang/go/src go1.22 fmt/
//      ^^^^^ reference github.com/golang/go/src go1.22 fmt/Print().
//            ^ reference local 2
   }
  }
//⌃ enclosing_range_end 0.1.test `sg/generallyeric`/Print().
  
