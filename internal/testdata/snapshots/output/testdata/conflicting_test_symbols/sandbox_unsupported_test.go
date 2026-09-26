  //go:build !linux && !windows && !freebsd
  // +build !linux,!windows,!freebsd
  
  package osl
//        ^^^ definition 0.1.test `sg/testdata/conflicting_test_symbols`/
//            kind Package
//            display_name osl
//            signature_documentation
//            > package osl
  
  import (
   "errors"
//  ^^^^^^ reference github.com/golang/go/src go1.22 errors/
   "testing"
//  ^^^^^^^ reference github.com/golang/go/src go1.22 testing/
  )
  
//⌄ enclosing_range_start 0.1.test `sg/testdata/conflicting_test_symbols`/ErrNotImplemented.
  var ErrNotImplemented = errors.New("not implemented")
//    ^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata/conflicting_test_symbols`/ErrNotImplemented.
//                      kind Variable
//                      display_name ErrNotImplemented
//                      signature_documentation
//                      > var ErrNotImplemented error
//                        ^^^^^^ reference github.com/golang/go/src go1.22 errors/
//                               ^^^ reference github.com/golang/go/src go1.22 errors/New().
//                                                    ⌃ enclosing_range_end 0.1.test `sg/testdata/conflicting_test_symbols`/ErrNotImplemented.
  
//⌄ enclosing_range_start 0.1.test `sg/testdata/conflicting_test_symbols`/newKey().
  func newKey(t *testing.T) (string, error) {
//     ^^^^^^ definition 0.1.test `sg/testdata/conflicting_test_symbols`/newKey().
//            kind Function
//            display_name newKey
//            signature_documentation
//            > func newKey(t *testing.T) (string, error)
//            ^ definition local 0
//              kind Variable
//              display_name t
//              signature_documentation
//              > var t *T
//               ^^^^^^^ reference github.com/golang/go/src go1.22 testing/
//                       ^ reference github.com/golang/go/src go1.22 testing/T#
   return "", ErrNotImplemented
//            ^^^^^^^^^^^^^^^^^ reference 0.1.test `sg/testdata/conflicting_test_symbols`/ErrNotImplemented.
  }
//⌃ enclosing_range_end 0.1.test `sg/testdata/conflicting_test_symbols`/newKey().
  
//⌄ enclosing_range_start 0.1.test `sg/testdata/conflicting_test_symbols`/verifySandbox().
  func verifySandbox(t *testing.T, s string) {
//     ^^^^^^^^^^^^^ definition 0.1.test `sg/testdata/conflicting_test_symbols`/verifySandbox().
//                   kind Function
//                   display_name verifySandbox
//                   signature_documentation
//                   > func verifySandbox(t *testing.T, s string)
//                   ^ definition local 1
//                     kind Variable
//                     display_name t
//                     signature_documentation
//                     > var t *T
//                      ^^^^^^^ reference github.com/golang/go/src go1.22 testing/
//                              ^ reference github.com/golang/go/src go1.22 testing/T#
//                                 ^ definition local 2
//                                   kind Variable
//                                   display_name s
//                                   signature_documentation
//                                   > var s string
   return
  }
//⌃ enclosing_range_end 0.1.test `sg/testdata/conflicting_test_symbols`/verifySandbox().
  
