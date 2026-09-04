  package impls
//        ^^^^^ definition 0.1.test `sg/impls`/
//              kind Package
//              display_name impls
//              signature_documentation
//              > package impls
//              documentation
//              > ```go
//              > package impls
//              > ```
  
  type I1 interface {
//     ^^ definition 0.1.test `sg/impls`/I1#
//        kind Interface
//        display_name I1
//        signature_documentation
//        > type I1 interface{ F1() }
//        documentation
//        > ```go
//        > type I1 interface{ F1() }
//        > ```
   F1()
// ^^ definition 0.1.test `sg/impls`/I1#F1.
//    kind MethodSpecification
//    display_name F1
//    signature_documentation
//    > func (I1).F1()
//    documentation
//    > ```go
//    > func (I1).F1()
//    > ```
  }
  
  type I1Clone interface {
//     ^^^^^^^ definition 0.1.test `sg/impls`/I1Clone#
//             kind Interface
//             display_name I1Clone
//             signature_documentation
//             > type I1Clone interface{ F1() }
//             documentation
//             > ```go
//             > type I1Clone interface{ F1() }
//             > ```
   F1()
// ^^ definition 0.1.test `sg/impls`/I1Clone#F1.
//    kind MethodSpecification
//    display_name F1
//    signature_documentation
//    > func (I1Clone).F1()
//    documentation
//    > ```go
//    > func (I1Clone).F1()
//    > ```
  }
  
  type IfaceOther interface {
//     ^^^^^^^^^^ definition 0.1.test `sg/impls`/IfaceOther#
//                kind Interface
//                display_name IfaceOther
//                signature_documentation
//                > type IfaceOther interface {
//                >     Another()
//                >     Something()
//                > }
//                documentation
//                > ```go
//                > type IfaceOther interface {
//                >     Another()
//                >     Something()
//                > }
//                > ```
   Something()
// ^^^^^^^^^ definition 0.1.test `sg/impls`/IfaceOther#Something.
//           kind MethodSpecification
//           display_name Something
//           signature_documentation
//           > func (IfaceOther).Something()
//           documentation
//           > ```go
//           > func (IfaceOther).Something()
//           > ```
   Another()
// ^^^^^^^ definition 0.1.test `sg/impls`/IfaceOther#Another.
//         kind MethodSpecification
//         display_name Another
//         signature_documentation
//         > func (IfaceOther).Another()
//         documentation
//         > ```go
//         > func (IfaceOther).Another()
//         > ```
  }
  
  type T1 int
//     ^^ definition 0.1.test `sg/impls`/T1#
//        kind Type
//        display_name T1
//        signature_documentation
//        > type T1 int
//        documentation
//        > ```go
//        > type T1 int
//        > ```
//        relationship 0.1.test `sg/impls`/I1# implementation
//        relationship 0.1.test `sg/impls`/I1Clone# implementation
  
//⌄ enclosing_range_start 0.1.test `sg/impls`/T1#F1().
  func (r T1) F1() {}
//      ^ definition local 0
//        kind Variable
//        display_name r
//        signature_documentation
//        > var r T1
//        ^^ reference 0.1.test `sg/impls`/T1#
//            ^^ definition 0.1.test `sg/impls`/T1#F1().
//               kind Method
//               display_name F1
//               signature_documentation
//               > func (T1).F1()
//               documentation
//               > ```go
//               > func (T1).F1()
//               > ```
//               relationship 0.1.test `sg/impls`/I1#F1. implementation
//               relationship 0.1.test `sg/impls`/I1Clone#F1. implementation
//                  ⌃ enclosing_range_end 0.1.test `sg/impls`/T1#F1().
  
  type T2 int
//     ^^ definition 0.1.test `sg/impls`/T2#
//        kind Type
//        display_name T2
//        signature_documentation
//        > type T2 int
//        documentation
//        > ```go
//        > type T2 int
//        > ```
//        relationship 0.1.test `sg/impls`/I1# implementation
//        relationship 0.1.test `sg/impls`/I1Clone# implementation
  
//⌄ enclosing_range_start 0.1.test `sg/impls`/T2#F1().
  func (r T2) F1() {}
//      ^ definition local 1
//        kind Variable
//        display_name r
//        signature_documentation
//        > var r T2
//        ^^ reference 0.1.test `sg/impls`/T2#
//            ^^ definition 0.1.test `sg/impls`/T2#F1().
//               kind Method
//               display_name F1
//               signature_documentation
//               > func (T2).F1()
//               documentation
//               > ```go
//               > func (T2).F1()
//               > ```
//               relationship 0.1.test `sg/impls`/I1#F1. implementation
//               relationship 0.1.test `sg/impls`/I1Clone#F1. implementation
//                  ⌃ enclosing_range_end 0.1.test `sg/impls`/T2#F1().
//⌄ enclosing_range_start 0.1.test `sg/impls`/T2#F2().
  func (r T2) F2() {}
//      ^ definition local 2
//        kind Variable
//        display_name r
//        signature_documentation
//        > var r T2
//        ^^ reference 0.1.test `sg/impls`/T2#
//            ^^ definition 0.1.test `sg/impls`/T2#F2().
//               kind Method
//               display_name F2
//               signature_documentation
//               > func (T2).F2()
//               documentation
//               > ```go
//               > func (T2).F2()
//               > ```
//                  ⌃ enclosing_range_end 0.1.test `sg/impls`/T2#F2().
  
