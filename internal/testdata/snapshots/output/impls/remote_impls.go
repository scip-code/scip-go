  package impls
//        ^^^^^ definition 0.1.test `sg/impls`/
  
  import "net/http"
//        ^^^^^^^^ reference github.com/golang/go/src go1.22 `net/http`/
  
//⌄ enclosing_range_start 0.1.test `sg/impls`/Something().
  func Something(r http.ResponseWriter) {}
//     ^^^^^^^^^ definition 0.1.test `sg/impls`/Something().
//               kind Function
//               display_name Something
//               signature_documentation
//               > func Something(r http.ResponseWriter)
//               ^ definition local 0
//                 kind Variable
//                 display_name r
//                 signature_documentation
//                 > var r ResponseWriter
//                 ^^^^ reference github.com/golang/go/src go1.22 `net/http`/
//                      ^^^^^^^^^^^^^^ reference github.com/golang/go/src go1.22 `net/http`/ResponseWriter#
//                                       ⌃ enclosing_range_end 0.1.test `sg/impls`/Something().
  
  type MyWriter struct{}
//     ^^^^^^^^ definition 0.1.test `sg/impls`/MyWriter#
//              kind Struct
//              display_name MyWriter
//              signature_documentation
//              > type MyWriter struct{}
//              relationship github.com/golang/go/src go1.22 `net/http`/ResponseWriter# implementation
//              relationship github.com/golang/go/src go1.22 io/Writer# implementation
  
//⌄ enclosing_range_start 0.1.test `sg/impls`/MyWriter#Header().
  func (w MyWriter) Header() http.Header        { panic("") }
//      ^ definition local 1
//        kind Variable
//        display_name w
//        signature_documentation
//        > var w MyWriter
//        ^^^^^^^^ reference 0.1.test `sg/impls`/MyWriter#
//                  ^^^^^^ definition 0.1.test `sg/impls`/MyWriter#Header().
//                         kind Method
//                         display_name Header
//                         signature_documentation
//                         > func (MyWriter).Header() http.Header
//                         relationship github.com/golang/go/src go1.22 `net/http`/ResponseWriter#Header(). implementation
//                           ^^^^ reference github.com/golang/go/src go1.22 `net/http`/
//                                ^^^^^^ reference github.com/golang/go/src go1.22 `net/http`/Header#
//                                                          ⌃ enclosing_range_end 0.1.test `sg/impls`/MyWriter#Header().
//⌄ enclosing_range_start 0.1.test `sg/impls`/MyWriter#Write().
  func (w MyWriter) Write([]byte) (int, error)  { panic("") }
//      ^ definition local 2
//        kind Variable
//        display_name w
//        signature_documentation
//        > var w MyWriter
//        ^^^^^^^^ reference 0.1.test `sg/impls`/MyWriter#
//                  ^^^^^ definition 0.1.test `sg/impls`/MyWriter#Write().
//                        kind Method
//                        display_name Write
//                        signature_documentation
//                        > func (MyWriter).Write([]byte) (int, error)
//                        relationship github.com/golang/go/src go1.22 `net/http`/ResponseWriter#Write(). implementation
//                        relationship github.com/golang/go/src go1.22 io/Writer#Write(). implementation
//                                                          ⌃ enclosing_range_end 0.1.test `sg/impls`/MyWriter#Write().
//⌄ enclosing_range_start 0.1.test `sg/impls`/MyWriter#WriteHeader().
  func (w MyWriter) WriteHeader(statusCode int) { panic("") }
//      ^ definition local 3
//        kind Variable
//        display_name w
//        signature_documentation
//        > var w MyWriter
//        ^^^^^^^^ reference 0.1.test `sg/impls`/MyWriter#
//                  ^^^^^^^^^^^ definition 0.1.test `sg/impls`/MyWriter#WriteHeader().
//                              kind Method
//                              display_name WriteHeader
//                              signature_documentation
//                              > func (MyWriter).WriteHeader(statusCode int)
//                              relationship github.com/golang/go/src go1.22 `net/http`/ResponseWriter#WriteHeader(). implementation
//                              ^^^^^^^^^^ definition local 4
//                                         kind Variable
//                                         display_name statusCode
//                                         signature_documentation
//                                         > var statusCode int
//                                                          ⌃ enclosing_range_end 0.1.test `sg/impls`/MyWriter#WriteHeader().
  
//⌄ enclosing_range_start 0.1.test `sg/impls`/Another().
  func Another() {
//     ^^^^^^^ definition 0.1.test `sg/impls`/Another().
//             kind Function
//             display_name Another
//             signature_documentation
//             > func Another()
   Something(MyWriter{})
// ^^^^^^^^^ reference 0.1.test `sg/impls`/Something().
//           ^^^^^^^^ reference 0.1.test `sg/impls`/MyWriter#
  }
//⌃ enclosing_range_end 0.1.test `sg/impls`/Another().
  
