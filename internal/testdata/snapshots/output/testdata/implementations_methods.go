  package testdata
//        ^^^^^^^^ definition 0.1.test `sg/testdata`/
  
  type InterfaceWithSingleMethod interface {
//     ^^^^^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/InterfaceWithSingleMethod#
//                               kind Interface
//                               display_name InterfaceWithSingleMethod
//                               signature_documentation
//                               > type InterfaceWithSingleMethod interface{ SingleMethod() float64 }
//                               documentation
//                               > ```go
//                               > type InterfaceWithSingleMethod interface{ SingleMethod() float64 }
//                               > ```
   SingleMethod() float64
// ^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/InterfaceWithSingleMethod#SingleMethod.
//              kind MethodSpecification
//              display_name SingleMethod
//              signature_documentation
//              > func (InterfaceWithSingleMethod).SingleMethod() float64
//              documentation
//              > ```go
//              > func (InterfaceWithSingleMethod).SingleMethod() float64
//              > ```
  }
  
  type StructWithMethods struct{}
//     ^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/StructWithMethods#
//                       kind Struct
//                       display_name StructWithMethods
//                       signature_documentation
//                       > type StructWithMethods struct{}
//                       documentation
//                       > ```go
//                       > type StructWithMethods struct{}
//                       > ```
//                       relationship 0.1.test `sg/testdata`/InterfaceWithSingleMethod# implementation
  
//⌄ enclosing_range_start 0.1.test `sg/testdata`/StructWithMethods#SingleMethod().
  func (StructWithMethods) SingleMethod() float64 { return 5.0 }
//      ^^^^^^^^^^^^^^^^^ reference 0.1.test `sg/testdata`/StructWithMethods#
//                         ^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/StructWithMethods#SingleMethod().
//                                      kind Method
//                                      display_name SingleMethod
//                                      signature_documentation
//                                      > func (StructWithMethods).SingleMethod() float64
//                                      documentation
//                                      > ```go
//                                      > func (StructWithMethods).SingleMethod() float64
//                                      > ```
//                                      relationship 0.1.test `sg/testdata`/InterfaceWithSingleMethod#SingleMethod. implementation
//                                                             ⌃ enclosing_range_end 0.1.test `sg/testdata`/StructWithMethods#SingleMethod().
  
  type InterfaceWithSingleMethodTwoImplementers interface {
//     ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/InterfaceWithSingleMethodTwoImplementers#
//                                              kind Interface
//                                              display_name InterfaceWithSingleMethodTwoImplementers
//                                              signature_documentation
//                                              > type InterfaceWithSingleMethodTwoImplementers interface{ SingleMethodTwoImpl() float64 }
//                                              documentation
//                                              > ```go
//                                              > type InterfaceWithSingleMethodTwoImplementers interface{ SingleMethodTwoImpl() float64 }
//                                              > ```
   SingleMethodTwoImpl() float64
// ^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/InterfaceWithSingleMethodTwoImplementers#SingleMethodTwoImpl.
//                     kind MethodSpecification
//                     display_name SingleMethodTwoImpl
//                     signature_documentation
//                     > func (InterfaceWithSingleMethodTwoImplementers).SingleMethodTwoImpl() float64
//                     documentation
//                     > ```go
//                     > func (InterfaceWithSingleMethodTwoImplementers).SingleMethodTwoImpl() float64
//                     > ```
  }
  
  type TwoImplOne struct{}
//     ^^^^^^^^^^ definition 0.1.test `sg/testdata`/TwoImplOne#
//                kind Struct
//                display_name TwoImplOne
//                signature_documentation
//                > type TwoImplOne struct{}
//                documentation
//                > ```go
//                > type TwoImplOne struct{}
//                > ```
//                relationship 0.1.test `sg/testdata`/InterfaceWithSingleMethodTwoImplementers# implementation
  
//⌄ enclosing_range_start 0.1.test `sg/testdata`/TwoImplOne#SingleMethodTwoImpl().
  func (TwoImplOne) SingleMethodTwoImpl() float64 { return 5.0 }
//      ^^^^^^^^^^ reference 0.1.test `sg/testdata`/TwoImplOne#
//                  ^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/TwoImplOne#SingleMethodTwoImpl().
//                                      kind Method
//                                      display_name SingleMethodTwoImpl
//                                      signature_documentation
//                                      > func (TwoImplOne).SingleMethodTwoImpl() float64
//                                      documentation
//                                      > ```go
//                                      > func (TwoImplOne).SingleMethodTwoImpl() float64
//                                      > ```
//                                      relationship 0.1.test `sg/testdata`/InterfaceWithSingleMethodTwoImplementers#SingleMethodTwoImpl. implementation
//                                                             ⌃ enclosing_range_end 0.1.test `sg/testdata`/TwoImplOne#SingleMethodTwoImpl().
  
  type TwoImplTwo struct{}
//     ^^^^^^^^^^ definition 0.1.test `sg/testdata`/TwoImplTwo#
//                kind Struct
//                display_name TwoImplTwo
//                signature_documentation
//                > type TwoImplTwo struct{}
//                documentation
//                > ```go
//                > type TwoImplTwo struct{}
//                > ```
//                relationship 0.1.test `sg/testdata`/InterfaceWithSingleMethodTwoImplementers# implementation
  
//⌄ enclosing_range_start 0.1.test `sg/testdata`/TwoImplTwo#SingleMethodTwoImpl().
  func (TwoImplTwo) SingleMethodTwoImpl() float64         { return 5.0 }
//      ^^^^^^^^^^ reference 0.1.test `sg/testdata`/TwoImplTwo#
//                  ^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/TwoImplTwo#SingleMethodTwoImpl().
//                                      kind Method
//                                      display_name SingleMethodTwoImpl
//                                      signature_documentation
//                                      > func (TwoImplTwo).SingleMethodTwoImpl() float64
//                                      documentation
//                                      > ```go
//                                      > func (TwoImplTwo).SingleMethodTwoImpl() float64
//                                      > ```
//                                      relationship 0.1.test `sg/testdata`/InterfaceWithSingleMethodTwoImplementers#SingleMethodTwoImpl. implementation
//                                                                     ⌃ enclosing_range_end 0.1.test `sg/testdata`/TwoImplTwo#SingleMethodTwoImpl().
//⌄ enclosing_range_start 0.1.test `sg/testdata`/TwoImplTwo#RandomThingThatDoesntMatter().
  func (TwoImplTwo) RandomThingThatDoesntMatter() float64 { return 5.0 }
//      ^^^^^^^^^^ reference 0.1.test `sg/testdata`/TwoImplTwo#
//                  ^^^^^^^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/TwoImplTwo#RandomThingThatDoesntMatter().
//                                              kind Method
//                                              display_name RandomThingThatDoesntMatter
//                                              signature_documentation
//                                              > func (TwoImplTwo).RandomThingThatDoesntMatter() float64
//                                              documentation
//                                              > ```go
//                                              > func (TwoImplTwo).RandomThingThatDoesntMatter() float64
//                                              > ```
//                                                                     ⌃ enclosing_range_end 0.1.test `sg/testdata`/TwoImplTwo#RandomThingThatDoesntMatter().
  
