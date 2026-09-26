  package testdata
//        ^^^^^^^^ definition 0.1.test `sg/testdata`/
  
//⌄ enclosing_range_start 0.1.test `sg/testdata`/I0#
  type I0 interface{}
//     ^^ definition 0.1.test `sg/testdata`/I0#
//        kind Interface
//        display_name I0
//        signature_documentation
//        > type I0 interface{}
//                  ⌃ enclosing_range_end 0.1.test `sg/testdata`/I0#
  
//⌄ enclosing_range_start 0.1.test `sg/testdata`/I1#
  type I1 interface {
//     ^^ definition 0.1.test `sg/testdata`/I1#
//        kind Interface
//        display_name I1
//        signature_documentation
//        > type I1 interface{ F1() }
// ⌄ enclosing_range_start 0.1.test `sg/testdata`/I1#F1.
   F1()
// ^^ definition 0.1.test `sg/testdata`/I1#F1.
//    kind MethodSpecification
//    display_name F1
//    signature_documentation
//    > func (I1).F1()
//    ⌃ enclosing_range_end 0.1.test `sg/testdata`/I1#F1.
  }
//⌃ enclosing_range_end 0.1.test `sg/testdata`/I1#
  
//⌄ enclosing_range_start 0.1.test `sg/testdata`/I2#
  type I2 interface {
//     ^^ definition 0.1.test `sg/testdata`/I2#
//        kind Interface
//        display_name I2
//        signature_documentation
//        > type I2 interface{ F2() }
// ⌄ enclosing_range_start 0.1.test `sg/testdata`/I2#F2.
   F2()
// ^^ definition 0.1.test `sg/testdata`/I2#F2.
//    kind MethodSpecification
//    display_name F2
//    signature_documentation
//    > func (I2).F2()
//    ⌃ enclosing_range_end 0.1.test `sg/testdata`/I2#F2.
  }
//⌃ enclosing_range_end 0.1.test `sg/testdata`/I2#
  
//⌄ enclosing_range_start 0.1.test `sg/testdata`/T1#
  type T1 int
//     ^^ definition 0.1.test `sg/testdata`/T1#
//        kind Type
//        display_name T1
//        signature_documentation
//        > type T1 int
//        relationship 0.1.test `sg/testdata`/I1# implementation
//          ⌃ enclosing_range_end 0.1.test `sg/testdata`/T1#
  
//⌄ enclosing_range_start 0.1.test `sg/testdata`/T1#F1().
  func (r T1) F1() {}
//      ^ definition local 0
//        kind Variable
//        display_name r
//        signature_documentation
//        > var r T1
//        ^^ reference 0.1.test `sg/testdata`/T1#
//            ^^ definition 0.1.test `sg/testdata`/T1#F1().
//               kind Method
//               display_name F1
//               signature_documentation
//               > func (T1).F1()
//               relationship 0.1.test `sg/testdata`/I1#F1. implementation
//                  ⌃ enclosing_range_end 0.1.test `sg/testdata`/T1#F1().
  
//⌄ enclosing_range_start 0.1.test `sg/testdata`/T2#
  type T2 int
//     ^^ definition 0.1.test `sg/testdata`/T2#
//        kind Type
//        display_name T2
//        signature_documentation
//        > type T2 int
//        relationship 0.1.test `sg/testdata`/I1# implementation
//        relationship 0.1.test `sg/testdata`/I2# implementation
//          ⌃ enclosing_range_end 0.1.test `sg/testdata`/T2#
  
//⌄ enclosing_range_start 0.1.test `sg/testdata`/T2#F1().
  func (r T2) F1() {}
//      ^ definition local 1
//        kind Variable
//        display_name r
//        signature_documentation
//        > var r T2
//        ^^ reference 0.1.test `sg/testdata`/T2#
//            ^^ definition 0.1.test `sg/testdata`/T2#F1().
//               kind Method
//               display_name F1
//               signature_documentation
//               > func (T2).F1()
//               relationship 0.1.test `sg/testdata`/I1#F1. implementation
//                  ⌃ enclosing_range_end 0.1.test `sg/testdata`/T2#F1().
//⌄ enclosing_range_start 0.1.test `sg/testdata`/T2#F2().
  func (r T2) F2() {}
//      ^ definition local 2
//        kind Variable
//        display_name r
//        signature_documentation
//        > var r T2
//        ^^ reference 0.1.test `sg/testdata`/T2#
//            ^^ definition 0.1.test `sg/testdata`/T2#F2().
//               kind Method
//               display_name F2
//               signature_documentation
//               > func (T2).F2()
//               relationship 0.1.test `sg/testdata`/I2#F2. implementation
//                  ⌃ enclosing_range_end 0.1.test `sg/testdata`/T2#F2().
  
//⌄ enclosing_range_start 0.1.test `sg/testdata`/A1#
  type A1 = T1
//     ^^ definition 0.1.test `sg/testdata`/A1#
//        kind TypeAlias
//        display_name A1
//        signature_documentation
//        > type A1 = T1
//          ^^ reference 0.1.test `sg/testdata`/T1#
//           ⌃ enclosing_range_end 0.1.test `sg/testdata`/A1#
//⌄ enclosing_range_start 0.1.test `sg/testdata`/A12#
  type A12 = A1
//     ^^^ definition 0.1.test `sg/testdata`/A12#
//         kind TypeAlias
//         display_name A12
//         signature_documentation
//         > type A12 = A1
//           ^^ reference 0.1.test `sg/testdata`/A1#
//            ⌃ enclosing_range_end 0.1.test `sg/testdata`/A12#
  
//⌄ enclosing_range_start 0.1.test `sg/testdata`/InterfaceWithNonExportedMethod#
  type InterfaceWithNonExportedMethod interface {
//     ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/InterfaceWithNonExportedMethod#
//                                    kind Interface
//                                    display_name InterfaceWithNonExportedMethod
//                                    signature_documentation
//                                    > type InterfaceWithNonExportedMethod interface{ nonExportedMethod() }
// ⌄ enclosing_range_start 0.1.test `sg/testdata`/InterfaceWithNonExportedMethod#nonExportedMethod.
   nonExportedMethod()
// ^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/InterfaceWithNonExportedMethod#nonExportedMethod.
//                   kind MethodSpecification
//                   display_name nonExportedMethod
//                   signature_documentation
//                   > func (InterfaceWithNonExportedMethod).nonExportedMethod()
//                   ⌃ enclosing_range_end 0.1.test `sg/testdata`/InterfaceWithNonExportedMethod#nonExportedMethod.
  }
//⌃ enclosing_range_end 0.1.test `sg/testdata`/InterfaceWithNonExportedMethod#
  
//⌄ enclosing_range_start 0.1.test `sg/testdata`/InterfaceWithExportedMethod#
  type InterfaceWithExportedMethod interface {
//     ^^^^^^^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/InterfaceWithExportedMethod#
//                                 kind Interface
//                                 display_name InterfaceWithExportedMethod
//                                 signature_documentation
//                                 > type InterfaceWithExportedMethod interface{ ExportedMethod() }
// ⌄ enclosing_range_start 0.1.test `sg/testdata`/InterfaceWithExportedMethod#ExportedMethod.
   ExportedMethod()
// ^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/InterfaceWithExportedMethod#ExportedMethod.
//                kind MethodSpecification
//                display_name ExportedMethod
//                signature_documentation
//                > func (InterfaceWithExportedMethod).ExportedMethod()
//                ⌃ enclosing_range_end 0.1.test `sg/testdata`/InterfaceWithExportedMethod#ExportedMethod.
  }
//⌃ enclosing_range_end 0.1.test `sg/testdata`/InterfaceWithExportedMethod#
  
//⌄ enclosing_range_start 0.1.test `sg/testdata`/Foo#
  type Foo int
//     ^^^ definition 0.1.test `sg/testdata`/Foo#
//         kind Type
//         display_name Foo
//         signature_documentation
//         > type Foo int
//         relationship 0.1.test `sg/testdata`/EmbeddedI3# implementation
//         relationship 0.1.test `sg/testdata`/I3# implementation
//         relationship 0.1.test `sg/testdata`/InterfaceWithExportedMethod# implementation
//         relationship 0.1.test `sg/testdata`/InterfaceWithNonExportedMethod# implementation
//           ⌃ enclosing_range_end 0.1.test `sg/testdata`/Foo#
  
//⌄ enclosing_range_start 0.1.test `sg/testdata`/Foo#nonExportedMethod().
  func (r Foo) nonExportedMethod() {}
//      ^ definition local 3
//        kind Variable
//        display_name r
//        signature_documentation
//        > var r Foo
//        ^^^ reference 0.1.test `sg/testdata`/Foo#
//             ^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/Foo#nonExportedMethod().
//                               kind Method
//                               display_name nonExportedMethod
//                               signature_documentation
//                               > func (Foo).nonExportedMethod()
//                               relationship 0.1.test `sg/testdata`/InterfaceWithNonExportedMethod#nonExportedMethod. implementation
//                                  ⌃ enclosing_range_end 0.1.test `sg/testdata`/Foo#nonExportedMethod().
//⌄ enclosing_range_start 0.1.test `sg/testdata`/Foo#ExportedMethod().
  func (r Foo) ExportedMethod()    {}
//      ^ definition local 4
//        kind Variable
//        display_name r
//        signature_documentation
//        > var r Foo
//        ^^^ reference 0.1.test `sg/testdata`/Foo#
//             ^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/Foo#ExportedMethod().
//                            kind Method
//                            display_name ExportedMethod
//                            signature_documentation
//                            > func (Foo).ExportedMethod()
//                            relationship 0.1.test `sg/testdata`/InterfaceWithExportedMethod#ExportedMethod. implementation
//                                  ⌃ enclosing_range_end 0.1.test `sg/testdata`/Foo#ExportedMethod().
//⌄ enclosing_range_start 0.1.test `sg/testdata`/Foo#ScipTestMethod().
  func (r Foo) ScipTestMethod()    {}
//      ^ definition local 5
//        kind Variable
//        display_name r
//        signature_documentation
//        > var r Foo
//        ^^^ reference 0.1.test `sg/testdata`/Foo#
//             ^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/Foo#ScipTestMethod().
//                            kind Method
//                            display_name ScipTestMethod
//                            signature_documentation
//                            > func (Foo).ScipTestMethod()
//                            relationship 0.1.test `sg/testdata`/EmbeddedI3#ScipTestMethod. implementation
//                            relationship 0.1.test `sg/testdata`/I3#ScipTestMethod. implementation
//                                  ⌃ enclosing_range_end 0.1.test `sg/testdata`/Foo#ScipTestMethod().
  
//⌄ enclosing_range_start 0.1.test `sg/testdata`/SharedOne#
  type SharedOne interface {
//     ^^^^^^^^^ definition 0.1.test `sg/testdata`/SharedOne#
//               kind Interface
//               display_name SharedOne
//               signature_documentation
//               > type SharedOne interface {
//               >     Distinct()
//               >     Shared()
//               > }
// ⌄ enclosing_range_start 0.1.test `sg/testdata`/SharedOne#Shared.
   Shared()
// ^^^^^^ definition 0.1.test `sg/testdata`/SharedOne#Shared.
//        kind MethodSpecification
//        display_name Shared
//        signature_documentation
//        > func (SharedOne).Shared()
//        ⌃ enclosing_range_end 0.1.test `sg/testdata`/SharedOne#Shared.
// ⌄ enclosing_range_start 0.1.test `sg/testdata`/SharedOne#Distinct.
   Distinct()
// ^^^^^^^^ definition 0.1.test `sg/testdata`/SharedOne#Distinct.
//          kind MethodSpecification
//          display_name Distinct
//          signature_documentation
//          > func (SharedOne).Distinct()
//          ⌃ enclosing_range_end 0.1.test `sg/testdata`/SharedOne#Distinct.
  }
//⌃ enclosing_range_end 0.1.test `sg/testdata`/SharedOne#
  
//⌄ enclosing_range_start 0.1.test `sg/testdata`/SharedTwo#
  type SharedTwo interface {
//     ^^^^^^^^^ definition 0.1.test `sg/testdata`/SharedTwo#
//               kind Interface
//               display_name SharedTwo
//               signature_documentation
//               > type SharedTwo interface {
//               >     Shared()
//               >     Unique()
//               > }
// ⌄ enclosing_range_start 0.1.test `sg/testdata`/SharedTwo#Shared.
   Shared()
// ^^^^^^ definition 0.1.test `sg/testdata`/SharedTwo#Shared.
//        kind MethodSpecification
//        display_name Shared
//        signature_documentation
//        > func (SharedTwo).Shared()
//        ⌃ enclosing_range_end 0.1.test `sg/testdata`/SharedTwo#Shared.
// ⌄ enclosing_range_start 0.1.test `sg/testdata`/SharedTwo#Unique.
   Unique()
// ^^^^^^ definition 0.1.test `sg/testdata`/SharedTwo#Unique.
//        kind MethodSpecification
//        display_name Unique
//        signature_documentation
//        > func (SharedTwo).Unique()
//        ⌃ enclosing_range_end 0.1.test `sg/testdata`/SharedTwo#Unique.
  }
//⌃ enclosing_range_end 0.1.test `sg/testdata`/SharedTwo#
  
//⌄ enclosing_range_start 0.1.test `sg/testdata`/Between#
  type Between struct{}
//     ^^^^^^^ definition 0.1.test `sg/testdata`/Between#
//             kind Struct
//             display_name Between
//             signature_documentation
//             > type Between struct{}
//             relationship 0.1.test `sg/testdata`/SharedOne# implementation
//             relationship 0.1.test `sg/testdata`/SharedTwo# implementation
//                    ⌃ enclosing_range_end 0.1.test `sg/testdata`/Between#
  
//⌄ enclosing_range_start 0.1.test `sg/testdata`/Between#Shared().
  func (Between) Shared()   {}
//      ^^^^^^^ reference 0.1.test `sg/testdata`/Between#
//               ^^^^^^ definition 0.1.test `sg/testdata`/Between#Shared().
//                      kind Method
//                      display_name Shared
//                      signature_documentation
//                      > func (Between).Shared()
//                      relationship 0.1.test `sg/testdata`/SharedOne#Shared. implementation
//                      relationship 0.1.test `sg/testdata`/SharedTwo#Shared. implementation
//                           ⌃ enclosing_range_end 0.1.test `sg/testdata`/Between#Shared().
//⌄ enclosing_range_start 0.1.test `sg/testdata`/Between#Distinct().
  func (Between) Distinct() {}
//      ^^^^^^^ reference 0.1.test `sg/testdata`/Between#
//               ^^^^^^^^ definition 0.1.test `sg/testdata`/Between#Distinct().
//                        kind Method
//                        display_name Distinct
//                        signature_documentation
//                        > func (Between).Distinct()
//                        relationship 0.1.test `sg/testdata`/SharedOne#Distinct. implementation
//                           ⌃ enclosing_range_end 0.1.test `sg/testdata`/Between#Distinct().
//⌄ enclosing_range_start 0.1.test `sg/testdata`/Between#Unique().
  func (Between) Unique()   {}
//      ^^^^^^^ reference 0.1.test `sg/testdata`/Between#
//               ^^^^^^ definition 0.1.test `sg/testdata`/Between#Unique().
//                      kind Method
//                      display_name Unique
//                      signature_documentation
//                      > func (Between).Unique()
//                      relationship 0.1.test `sg/testdata`/SharedTwo#Unique. implementation
//                           ⌃ enclosing_range_end 0.1.test `sg/testdata`/Between#Unique().
  
//⌄ enclosing_range_start 0.1.test `sg/testdata`/shouldShow().
  func shouldShow(shared SharedOne) {
//     ^^^^^^^^^^ definition 0.1.test `sg/testdata`/shouldShow().
//                kind Function
//                display_name shouldShow
//                signature_documentation
//                > func shouldShow(shared SharedOne)
//                ^^^^^^ definition local 6
//                       kind Variable
//                       display_name shared
//                       signature_documentation
//                       > var shared SharedOne
//                       ^^^^^^^^^ reference 0.1.test `sg/testdata`/SharedOne#
   shared.Shared()
// ^^^^^^ reference local 6
//        ^^^^^^ reference 0.1.test `sg/testdata`/SharedOne#Shared.
  }
//⌃ enclosing_range_end 0.1.test `sg/testdata`/shouldShow().
  
