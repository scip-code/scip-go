  package initial
//        ^^^^^^^ definition 0.1.test `sg/initial`/
  
//⌄ enclosing_range_start 0.1.test `sg/initial`/LiteralType#
  type LiteralType int
//     ^^^^^^^^^^^ definition 0.1.test `sg/initial`/LiteralType#
//                 kind Type
//                 display_name LiteralType
//                 signature_documentation
//                 > type LiteralType int
//                   ⌃ enclosing_range_end 0.1.test `sg/initial`/LiteralType#
  
//⌄ enclosing_range_start 0.1.test `sg/initial`/FuncType#
  type FuncType func(LiteralType, int) bool
//     ^^^^^^^^ definition 0.1.test `sg/initial`/FuncType#
//              kind Type
//              display_name FuncType
//              signature_documentation
//              > type FuncType func(LiteralType, int) bool
//                   ^^^^^^^^^^^ reference 0.1.test `sg/initial`/LiteralType#
//                                        ⌃ enclosing_range_end 0.1.test `sg/initial`/FuncType#
  
//⌄ enclosing_range_start 0.1.test `sg/initial`/IfaceType#
  type IfaceType interface {
//     ^^^^^^^^^ definition 0.1.test `sg/initial`/IfaceType#
//               kind Interface
//               display_name IfaceType
//               signature_documentation
//               > type IfaceType interface{ Method() LiteralType }
// ⌄ enclosing_range_start 0.1.test `sg/initial`/IfaceType#Method.
   Method() LiteralType
// ^^^^^^ definition 0.1.test `sg/initial`/IfaceType#Method.
//        kind MethodSpecification
//        display_name Method
//        signature_documentation
//        > func (IfaceType).Method() LiteralType
//          ^^^^^^^^^^^ reference 0.1.test `sg/initial`/LiteralType#
//                    ⌃ enclosing_range_end 0.1.test `sg/initial`/IfaceType#Method.
  }
//⌃ enclosing_range_end 0.1.test `sg/initial`/IfaceType#
  
//⌄ enclosing_range_start 0.1.test `sg/initial`/StructType#
  type StructType struct {
//     ^^^^^^^^^^ definition 0.1.test `sg/initial`/StructType#
//                kind Struct
//                display_name StructType
//                signature_documentation
//                > type StructType struct {
//                >     m    IfaceType
//                >     f    LiteralType
//                >     anon struct{ sub int }
//                >     i    interface{ AnonMethod() bool }
//                > }
// ⌄ enclosing_range_start 0.1.test `sg/initial`/StructType#m.
   m IfaceType
// ^ definition 0.1.test `sg/initial`/StructType#m.
//   kind Field
//   display_name m
//   signature_documentation
//   > struct field m IfaceType
//   ^^^^^^^^^ reference 0.1.test `sg/initial`/IfaceType#
//           ⌃ enclosing_range_end 0.1.test `sg/initial`/StructType#m.
// ⌄ enclosing_range_start 0.1.test `sg/initial`/StructType#f.
   f LiteralType
// ^ definition 0.1.test `sg/initial`/StructType#f.
//   kind Field
//   display_name f
//   signature_documentation
//   > struct field f LiteralType
//   ^^^^^^^^^^^ reference 0.1.test `sg/initial`/LiteralType#
//             ⌃ enclosing_range_end 0.1.test `sg/initial`/StructType#f.
  
// ⌄ enclosing_range_start 0.1.test `sg/initial`/StructType#anon.
   // anonymous struct
   anon struct {
// ^^^^ definition 0.1.test `sg/initial`/StructType#anon.
//      kind Field
//      display_name anon
//      signature_documentation
//      > struct field anon struct{sub int}
//      documentation
//      > anonymous struct
//  ⌄ enclosing_range_start 0.1.test `sg/initial`/StructType#$anon_0ba9ace1dcfd6761#sub.
    sub int
//  ^^^ definition 0.1.test `sg/initial`/StructType#$anon_0ba9ace1dcfd6761#sub.
//      kind Field
//      display_name sub
//      signature_documentation
//      > struct field sub int
//        ⌃ enclosing_range_end 0.1.test `sg/initial`/StructType#$anon_0ba9ace1dcfd6761#sub.
   }
// ⌃ enclosing_range_end 0.1.test `sg/initial`/StructType#anon.
  
// ⌄ enclosing_range_start 0.1.test `sg/initial`/StructType#i.
   // interface within struct
   i interface {
// ^ definition 0.1.test `sg/initial`/StructType#i.
//   kind Field
//   display_name i
//   signature_documentation
//   > struct field i interface{AnonMethod() bool}
//   documentation
//   > interface within struct
//  ⌄ enclosing_range_start 0.1.test `sg/initial`/StructType#$anon_97e7de633e3ef8e8#AnonMethod.
    AnonMethod() bool
//  ^^^^^^^^^^ definition 0.1.test `sg/initial`/StructType#$anon_97e7de633e3ef8e8#AnonMethod.
//             kind MethodSpecification
//             display_name AnonMethod
//             signature_documentation
//             > func (interface).AnonMethod() bool
//                  ⌃ enclosing_range_end 0.1.test `sg/initial`/StructType#$anon_97e7de633e3ef8e8#AnonMethod.
   }
// ⌃ enclosing_range_end 0.1.test `sg/initial`/StructType#i.
  }
//⌃ enclosing_range_end 0.1.test `sg/initial`/StructType#
  
//⌄ enclosing_range_start 0.1.test `sg/initial`/DeclaredBefore#
//                            ⌄ enclosing_range_start 0.1.test `sg/initial`/DeclaredBefore#DeclaredAfter.
  type DeclaredBefore struct{ DeclaredAfter }
//     ^^^^^^^^^^^^^^ definition 0.1.test `sg/initial`/DeclaredBefore#
//                    kind Struct
//                    display_name DeclaredBefore
//                    signature_documentation
//                    > type DeclaredBefore struct{ DeclaredAfter }
//                            ^^^^^^^^^^^^^ definition 0.1.test `sg/initial`/DeclaredBefore#DeclaredAfter.
//                                          kind Field
//                                          display_name DeclaredAfter
//                                          signature_documentation
//                                          > struct field DeclaredAfter DeclaredAfter
//                            ^^^^^^^^^^^^^ reference 0.1.test `sg/initial`/DeclaredAfter#
//                                        ⌃ enclosing_range_end 0.1.test `sg/initial`/DeclaredBefore#DeclaredAfter.
//                                          ⌃ enclosing_range_end 0.1.test `sg/initial`/DeclaredBefore#
//⌄ enclosing_range_start 0.1.test `sg/initial`/DeclaredAfter#
  type DeclaredAfter struct{}
//     ^^^^^^^^^^^^^ definition 0.1.test `sg/initial`/DeclaredAfter#
//                   kind Struct
//                   display_name DeclaredAfter
//                   signature_documentation
//                   > type DeclaredAfter struct{}
//                          ⌃ enclosing_range_end 0.1.test `sg/initial`/DeclaredAfter#
  
