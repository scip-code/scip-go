  package testdata
//        ^^^^^^^^ definition 0.1.test `sg/testdata`/
  
  import (
   . "fmt"
//    ^^^ reference github.com/golang/go/src go1.22 fmt/
   h "net/http"
// ^ reference github.com/golang/go/src go1.22 `net/http`/
//    ^^^^^^^^ reference github.com/golang/go/src go1.22 `net/http`/
  )
  
//⌄ enclosing_range_start 0.1.test `sg/testdata`/Example().
  func Example() {
//     ^^^^^^^ definition 0.1.test `sg/testdata`/Example().
//             kind Function
//             display_name Example
//             signature_documentation
//             > func Example()
//             documentation
//             > ```go
//             > func Example()
//             > ```
   Println(h.CanonicalHeaderKey("accept-encoding"))
// ^^^^^^^ reference github.com/golang/go/src go1.22 fmt/Println().
//         ^ reference github.com/golang/go/src go1.22 `net/http`/
//           ^^^^^^^^^^^^^^^^^^ reference github.com/golang/go/src go1.22 `net/http`/CanonicalHeaderKey().
  }
//⌃ enclosing_range_end 0.1.test `sg/testdata`/Example().
  
