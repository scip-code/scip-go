  package dep
//        ^^^ definition github.com/example/dep 0.1.test `github.com/example/dep`/
//            kind Package
//            display_name dep
//            signature_documentation
//            > package dep
  
  import "fmt"
//        ^^^ reference github.com/golang/go/src go1.22 fmt/
  
//⌄ enclosing_range_start github.com/example/dep 0.1.test `github.com/example/dep`/T#
  type T struct{}
//     ^ definition github.com/example/dep 0.1.test `github.com/example/dep`/T#
//       kind Struct
//       display_name T
//       signature_documentation
//       > type T struct{}
//              ⌃ enclosing_range_end github.com/example/dep 0.1.test `github.com/example/dep`/T#
  
//⌄ enclosing_range_start github.com/example/dep 0.1.test `github.com/example/dep`/T#Bar().
  func (t *T) Bar() {
//      ^ definition local 0
//        kind Variable
//        display_name t
//        signature_documentation
//        > var t *T
//         ^ reference github.com/example/dep 0.1.test `github.com/example/dep`/T#
//            ^^^ definition github.com/example/dep 0.1.test `github.com/example/dep`/T#Bar().
//                kind Method
//                display_name Bar
//                signature_documentation
//                > func (*T).Bar()
   fmt.Println("Bar")
// ^^^ reference github.com/golang/go/src go1.22 fmt/
//     ^^^^^^^ reference github.com/golang/go/src go1.22 fmt/Println().
  }
//⌃ enclosing_range_end github.com/example/dep 0.1.test `github.com/example/dep`/T#Bar().
  
