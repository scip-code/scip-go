  package testspecial_test
//        ^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testspecial_test`/
//                         kind Package
//                         display_name testspecial_test
//                         signature_documentation
//                         > package testspecial_test
//                         documentation
//                         > ```go
//                         > package testspecial_test
//                         > ```
  
  import (
   "testing"
//  ^^^^^^^ reference github.com/golang/go/src go1.22 testing/
  
   "sg/testspecial"
//  ^^^^^^^^^^^^^^ reference 0.1.test `sg/testspecial`/
  )
  
//⌄ enclosing_range_start 0.1.test `sg/testspecial_test`/TestFoo_Blackbox().
  func TestFoo_Blackbox(*testing.T) { testspecial.Foo() }
//     ^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testspecial_test`/TestFoo_Blackbox().
//                      kind Function
//                      display_name TestFoo_Blackbox
//                      signature_documentation
//                      > func TestFoo_Blackbox(*testing.T)
//                      documentation
//                      > ```go
//                      > func TestFoo_Blackbox(*testing.T)
//                      > ```
//                       ^^^^^^^ reference github.com/golang/go/src go1.22 testing/
//                               ^ reference github.com/golang/go/src go1.22 testing/T#
//                                    ^^^^^^^^^^^ reference 0.1.test `sg/testspecial`/
//                                                ^^^ reference 0.1.test `sg/testspecial`/Foo().
//                                                      ⌃ enclosing_range_end 0.1.test `sg/testspecial_test`/TestFoo_Blackbox().
  
