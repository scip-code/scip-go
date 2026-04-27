  package embedded
//        ^^^^^^^^ definition 0.1.test `sg/embedded`/
  
  import "net/http"
//        ^^^^^^^^ reference github.com/golang/go/src go1.22 `net/http`/
  
  type NestedHandler struct {
//     ^^^^^^^^^^^^^ definition 0.1.test `sg/embedded`/NestedHandler#
//                   kind Struct
//                   display_name NestedHandler
//                   signature_documentation
//                   > type NestedHandler struct {
//                   >     http.Handler
//                   >     Other int
//                   > }
//                   relationship github.com/golang/go/src go1.22 `net/http`/Handler# implementation
   http.Handler
// ^^^^ reference github.com/golang/go/src go1.22 `net/http`/
//      ^^^^^^^ definition 0.1.test `sg/embedded`/NestedHandler#Handler.
//              kind Field
//              display_name Handler
//              signature_documentation
//              > struct field Handler http.Handler
//      ^^^^^^^ reference github.com/golang/go/src go1.22 `net/http`/Handler#
  
   // Wow, a great thing for integers
   Other int
// ^^^^^ definition 0.1.test `sg/embedded`/NestedHandler#Other.
//       kind Field
//       display_name Other
//       signature_documentation
//       > struct field Other int
//       documentation
//       > Wow, a great thing for integers
  }
  
//⌄ enclosing_range_start 0.1.test `sg/embedded`/NestedExample().
  func NestedExample(n NestedHandler) {
//     ^^^^^^^^^^^^^ definition 0.1.test `sg/embedded`/NestedExample().
//                   kind Function
//                   display_name NestedExample
//                   signature_documentation
//                   > func NestedExample(n NestedHandler)
//                   ^ definition local 0
//                     kind Variable
//                     display_name n
//                     signature_documentation
//                     > var n NestedHandler
//                     ^^^^^^^^^^^^^ reference 0.1.test `sg/embedded`/NestedHandler#
   _ = n.Handler.ServeHTTP
//     ^ reference local 0
//       ^^^^^^^ reference 0.1.test `sg/embedded`/NestedHandler#Handler.
//               ^^^^^^^^^ reference github.com/golang/go/src go1.22 `net/http`/Handler#ServeHTTP().
   _ = n.ServeHTTP
//     ^ reference local 0
//       ^^^^^^^^^ reference github.com/golang/go/src go1.22 `net/http`/Handler#ServeHTTP().
   _ = n.Other
//     ^ reference local 0
//       ^^^^^ reference 0.1.test `sg/embedded`/NestedHandler#Other.
  }
//⌃ enclosing_range_end 0.1.test `sg/embedded`/NestedExample().
  
