  package pr256_test
//        ^^^^^^^^^^ definition 0.1.test `sg/pr256_test`/
//                   kind Package
//                   display_name pr256_test
//                   signature_documentation
//                   > package pr256_test
//                   documentation
//                   > ```go
//                   > package pr256_test
//                   > ```
  
  import "testing"
//        ^^^^^^^ reference github.com/golang/go/src go1.22 testing/
  
//⌄ enclosing_range_start 0.1.test `sg/pr256_test`/TestComment().
  func TestComment(t *testing.T) {}
//     ^^^^^^^^^^^ definition 0.1.test `sg/pr256_test`/TestComment().
//                 kind Function
//                 display_name TestComment
//                 signature_documentation
//                 > func TestComment(t *testing.T)
//                 documentation
//                 > ```go
//                 > func TestComment(t *testing.T)
//                 > ```
//                 ^ definition local 0
//                   kind Variable
//                   display_name t
//                   signature_documentation
//                   > var t *T
//                    ^^^^^^^ reference github.com/golang/go/src go1.22 testing/
//                            ^ reference github.com/golang/go/src go1.22 testing/T#
//                                ⌃ enclosing_range_end 0.1.test `sg/pr256_test`/TestComment().
  
