  package pr201
//        ^^^^^ definition 0.1.test `sg/pr201`/
//              kind Package
//              display_name pr201
//              signature_documentation
//              > package pr201
//              documentation
//              > ```go
//              > package pr201
//              > ```
  
  import (
   "context"
//  ^^^^^^^ reference github.com/golang/go/src go1.22 context/
   "net/http"
//  ^^^^^^^^ reference github.com/golang/go/src go1.22 `net/http`/
  
   _ "embed"
//    ^^^^^ reference github.com/golang/go/src go1.22 embed/
  
   . "strings"
//    ^^^^^^^ reference github.com/golang/go/src go1.22 strings/
  
   fmt "fmt"
// ^^^ reference github.com/golang/go/src go1.22 fmt/
//      ^^^ reference github.com/golang/go/src go1.22 fmt/
   h "net/http"
// ^ reference github.com/golang/go/src go1.22 `net/http`/
//    ^^^^^^^^ reference github.com/golang/go/src go1.22 `net/http`/
  )
  
//⌄ enclosing_range_start 0.1.test `sg/pr201`/UseImports().
  func UseImports() {
//     ^^^^^^^^^^ definition 0.1.test `sg/pr201`/UseImports().
//                kind Function
//                display_name UseImports
//                signature_documentation
//                > func UseImports()
//                documentation
//                > ```go
//                > func UseImports()
//                > ```
   fmt.Println(context.Background())
// ^^^ reference github.com/golang/go/src go1.22 fmt/
//     ^^^^^^^ reference github.com/golang/go/src go1.22 fmt/Println().
//             ^^^^^^^ reference github.com/golang/go/src go1.22 context/
//                     ^^^^^^^^^^ reference github.com/golang/go/src go1.22 context/Background().
   _ = http.StatusOK
//     ^^^^ reference github.com/golang/go/src go1.22 `net/http`/
//          ^^^^^^^^ reference github.com/golang/go/src go1.22 `net/http`/StatusOK.
   _ = h.DefaultClient
//     ^ reference github.com/golang/go/src go1.22 `net/http`/
//       ^^^^^^^^^^^^^ reference github.com/golang/go/src go1.22 `net/http`/DefaultClient.
   _ = Contains("hello", "ell")
//     ^^^^^^^^ reference github.com/golang/go/src go1.22 strings/Contains().
  }
//⌃ enclosing_range_end 0.1.test `sg/pr201`/UseImports().
  
