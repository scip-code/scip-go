  package testdata
//        ^^^^^^^^ definition 0.1.test `sg/testdata`/
  
  import "net/http"
//        ^^^^^^^^ reference github.com/golang/go/src go1.22 `net/http`/
  
  type implementsWriter struct{}
//     ^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/implementsWriter#
//                      kind Struct
//                      display_name implementsWriter
//                      signature_documentation
//                      > type implementsWriter struct{}
//                      relationship github.com/golang/go/src go1.22 `net/http`/ResponseWriter# implementation
//                      relationship github.com/golang/go/src go1.22 io/Writer# implementation
  
//⌄ enclosing_range_start 0.1.test `sg/testdata`/implementsWriter#Header().
  func (implementsWriter) Header() http.Header        { panic("Just for how") }
//      ^^^^^^^^^^^^^^^^ reference 0.1.test `sg/testdata`/implementsWriter#
//                        ^^^^^^ definition 0.1.test `sg/testdata`/implementsWriter#Header().
//                               kind Method
//                               display_name Header
//                               signature_documentation
//                               > func (implementsWriter).Header() http.Header
//                               relationship github.com/golang/go/src go1.22 `net/http`/ResponseWriter#Header(). implementation
//                                 ^^^^ reference github.com/golang/go/src go1.22 `net/http`/
//                                      ^^^^^^ reference github.com/golang/go/src go1.22 `net/http`/Header#
//                                                                            ⌃ enclosing_range_end 0.1.test `sg/testdata`/implementsWriter#Header().
//⌄ enclosing_range_start 0.1.test `sg/testdata`/implementsWriter#Write().
  func (implementsWriter) Write([]byte) (int, error)  { panic("Just for show") }
//      ^^^^^^^^^^^^^^^^ reference 0.1.test `sg/testdata`/implementsWriter#
//                        ^^^^^ definition 0.1.test `sg/testdata`/implementsWriter#Write().
//                              kind Method
//                              display_name Write
//                              signature_documentation
//                              > func (implementsWriter).Write([]byte) (int, error)
//                              relationship github.com/golang/go/src go1.22 `net/http`/ResponseWriter#Write(). implementation
//                              relationship github.com/golang/go/src go1.22 io/Writer#Write(). implementation
//                                                                             ⌃ enclosing_range_end 0.1.test `sg/testdata`/implementsWriter#Write().
//⌄ enclosing_range_start 0.1.test `sg/testdata`/implementsWriter#WriteHeader().
  func (implementsWriter) WriteHeader(statusCode int) {}
//      ^^^^^^^^^^^^^^^^ reference 0.1.test `sg/testdata`/implementsWriter#
//                        ^^^^^^^^^^^ definition 0.1.test `sg/testdata`/implementsWriter#WriteHeader().
//                                    kind Method
//                                    display_name WriteHeader
//                                    signature_documentation
//                                    > func (implementsWriter).WriteHeader(statusCode int)
//                                    relationship github.com/golang/go/src go1.22 `net/http`/ResponseWriter#WriteHeader(). implementation
//                                    ^^^^^^^^^^ definition local 0
//                                               kind Variable
//                                               display_name statusCode
//                                               signature_documentation
//                                               > var statusCode int
//                                                     ⌃ enclosing_range_end 0.1.test `sg/testdata`/implementsWriter#WriteHeader().
  
//⌄ enclosing_range_start 0.1.test `sg/testdata`/ShowsInSignature().
  func ShowsInSignature(respWriter http.ResponseWriter) {
//     ^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/ShowsInSignature().
//                      kind Function
//                      display_name ShowsInSignature
//                      signature_documentation
//                      > func ShowsInSignature(respWriter http.ResponseWriter)
//                      ^^^^^^^^^^ definition local 1
//                                 kind Variable
//                                 display_name respWriter
//                                 signature_documentation
//                                 > var respWriter ResponseWriter
//                                 ^^^^ reference github.com/golang/go/src go1.22 `net/http`/
//                                      ^^^^^^^^^^^^^^ reference github.com/golang/go/src go1.22 `net/http`/ResponseWriter#
   respWriter.WriteHeader(1)
// ^^^^^^^^^^ reference local 1
//            ^^^^^^^^^^^ reference github.com/golang/go/src go1.22 `net/http`/ResponseWriter#WriteHeader().
  }
//⌃ enclosing_range_end 0.1.test `sg/testdata`/ShowsInSignature().
  
