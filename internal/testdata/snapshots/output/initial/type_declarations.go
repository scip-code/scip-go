  package initial
//        ^^^^^^^ definition 0.1.test `sg/initial`/
  
  type LiteralType int
//     ^^^^^^^^^^^ definition 0.1.test `sg/initial`/LiteralType#
//                 kind Type
//                 display_name LiteralType
//                 signature_documentation
//                 > type LiteralType int
//                 documentation
//                 > ```go
//                 > type LiteralType int
//                 > ```
  
  type FuncType func(LiteralType, int) bool
//     ^^^^^^^^ definition 0.1.test `sg/initial`/FuncType#
//              kind Type
//              display_name FuncType
//              signature_documentation
//              > type FuncType func(LiteralType, int) bool
//              documentation
//              > ```go
//              > type FuncType func(LiteralType, int) bool
//              > ```
//                   ^^^^^^^^^^^ reference 0.1.test `sg/initial`/LiteralType#
  
  type IfaceType interface {
//     ^^^^^^^^^ definition 0.1.test `sg/initial`/IfaceType#
//               kind Interface
//               display_name IfaceType
//               signature_documentation
//               > type IfaceType interface{ Method() LiteralType }
//               documentation
//               > ```go
//               > type IfaceType interface{ Method() LiteralType }
//               > ```
   Method() LiteralType
// ^^^^^^ definition 0.1.test `sg/initial`/IfaceType#Method.
//        kind MethodSpecification
//        display_name Method
//        signature_documentation
//        > func (IfaceType).Method() LiteralType
//        documentation
//        > ```go
//        > func (IfaceType).Method() LiteralType
//        > ```
//          ^^^^^^^^^^^ reference 0.1.test `sg/initial`/LiteralType#
  }
  
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
//                documentation
//                > ```go
//                > type StructType struct {
//                >     m    IfaceType
//                >     f    LiteralType
//                >     anon struct{ sub int }
//                >     i    interface{ AnonMethod() bool }
//                > }
//                > ```
   m IfaceType
// ^ definition 0.1.test `sg/initial`/StructType#m.
//   kind Field
//   display_name m
//   signature_documentation
//   > struct field m IfaceType
//   documentation
//   > ```go
//   > struct field m IfaceType
//   > ```
//   ^^^^^^^^^ reference 0.1.test `sg/initial`/IfaceType#
   f LiteralType
// ^ definition 0.1.test `sg/initial`/StructType#f.
//   kind Field
//   display_name f
//   signature_documentation
//   > struct field f LiteralType
//   documentation
//   > ```go
//   > struct field f LiteralType
//   > ```
//   ^^^^^^^^^^^ reference 0.1.test `sg/initial`/LiteralType#
  
   // anonymous struct
   anon struct {
// ^^^^ definition 0.1.test `sg/initial`/StructType#anon.
//      kind Field
//      display_name anon
//      signature_documentation
//      > struct field anon struct{sub int}
//      documentation
//      > ```go
//      > struct field anon struct{sub int}
//      > ```
//      documentation
//      > anonymous struct
    sub int
//  ^^^ definition 0.1.test `sg/initial`/StructType#$anon_0ba9ace1dcfd6761#sub.
//      kind Field
//      display_name sub
//      signature_documentation
//      > struct field sub int
//      documentation
//      > ```go
//      > struct field sub int
//      > ```
   }
  
   // interface within struct
   i interface {
// ^ definition 0.1.test `sg/initial`/StructType#i.
//   kind Field
//   display_name i
//   signature_documentation
//   > struct field i interface{AnonMethod() bool}
//   documentation
//   > ```go
//   > struct field i interface{AnonMethod() bool}
//   > ```
//   documentation
//   > interface within struct
    AnonMethod() bool
//  ^^^^^^^^^^ definition 0.1.test `sg/initial`/StructType#$anon_97e7de633e3ef8e8#AnonMethod.
//             kind MethodSpecification
//             display_name AnonMethod
//             signature_documentation
//             > func (interface).AnonMethod() bool
//             documentation
//             > ```go
//             > func (interface).AnonMethod() bool
//             > ```
   }
  }
  
  type DeclaredBefore struct{ DeclaredAfter }
//     ^^^^^^^^^^^^^^ definition 0.1.test `sg/initial`/DeclaredBefore#
//                    kind Struct
//                    display_name DeclaredBefore
//                    signature_documentation
//                    > type DeclaredBefore struct{ DeclaredAfter }
//                    documentation
//                    > ```go
//                    > type DeclaredBefore struct{ DeclaredAfter }
//                    > ```
//                            ^^^^^^^^^^^^^ definition 0.1.test `sg/initial`/DeclaredBefore#DeclaredAfter.
//                                          kind Field
//                                          display_name DeclaredAfter
//                                          signature_documentation
//                                          > struct field DeclaredAfter DeclaredAfter
//                                          documentation
//                                          > ```go
//                                          > struct field DeclaredAfter DeclaredAfter
//                                          > ```
//                            ^^^^^^^^^^^^^ reference 0.1.test `sg/initial`/DeclaredAfter#
  type DeclaredAfter struct{}
//     ^^^^^^^^^^^^^ definition 0.1.test `sg/initial`/DeclaredAfter#
//                   kind Struct
//                   display_name DeclaredAfter
//                   signature_documentation
//                   > type DeclaredAfter struct{}
//                   documentation
//                   > ```go
//                   > type DeclaredAfter struct{}
//                   > ```
  
