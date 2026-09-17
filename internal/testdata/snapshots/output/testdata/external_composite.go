  package testdata
//        ^^^^^^^^ definition 0.1.test `sg/testdata`/
  
  import "net/http"
//        ^^^^^^^^ reference github.com/golang/go/src go1.22 `net/http`/
  
  type NestedHandler struct {
//     ^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/NestedHandler#
//                   kind Struct
//                   display_name NestedHandler
//                   signature_documentation
//                   > type NestedHandler struct {
//                   >     http.Handler
//                   >     Other int
//                   > }
//                   documentation
//                   > ```go
//                   > type NestedHandler struct {
//                   >     http.Handler
//                   >     Other int
//                   > }
//                   > ```
//                   relationship github.com/golang/go/src go1.22 `net/http`/Handler# implementation
   http.Handler
// ^^^^ reference github.com/golang/go/src go1.22 `net/http`/
//      ^^^^^^^ definition 0.1.test `sg/testdata`/NestedHandler#Handler.
//              kind Field
//              display_name Handler
//              signature_documentation
//              > struct field Handler http.Handler
//              documentation
//              > ```go
//              > struct field Handler http.Handler
//              > ```
//      ^^^^^^^ reference github.com/golang/go/src go1.22 `net/http`/Handler#
   Other int
// ^^^^^ definition 0.1.test `sg/testdata`/NestedHandler#Other.
//       kind Field
//       display_name Other
//       signature_documentation
//       > struct field Other int
//       documentation
//       > ```go
//       > struct field Other int
//       > ```
  }
  
