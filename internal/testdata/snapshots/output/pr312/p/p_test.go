  package p
//        ^ definition 0.1.test `sg/pr312/p`/
  
  import "testing"
//        ^^^^^^^ reference github.com/golang/go/src go1.22 testing/
  
//⌄ enclosing_range_start 0.1.test `sg/pr312/p`/TestOpts().
  func TestOpts(t *testing.T) { _ = Opts{} }
//     ^^^^^^^^ definition 0.1.test `sg/pr312/p`/TestOpts().
//              kind Function
//              display_name TestOpts
//              signature_documentation
//              > func TestOpts(t *testing.T)
//              ^ definition local 0
//                kind Variable
//                display_name t
//                signature_documentation
//                > var t *T
//                 ^^^^^^^ reference github.com/golang/go/src go1.22 testing/
//                         ^ reference github.com/golang/go/src go1.22 testing/T#
//                                  ^^^^ reference 0.1.test `sg/pr312/p`/Opts#
//                                         ⌃ enclosing_range_end 0.1.test `sg/pr312/p`/TestOpts().
  
