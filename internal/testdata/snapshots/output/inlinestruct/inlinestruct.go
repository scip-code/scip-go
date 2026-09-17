  package inlinestruct
//        ^^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/
//                     kind Package
//                     display_name inlinestruct
//                     signature_documentation
//                     > package inlinestruct
//                     documentation
//                     > ```go
//                     > package inlinestruct
//                     > ```
  
  type FieldInterface interface {
//     ^^^^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/FieldInterface#
//                    kind Interface
//                    display_name FieldInterface
//                    signature_documentation
//                    > type FieldInterface interface{ SomeMethod() string }
//                    documentation
//                    > ```go
//                    > type FieldInterface interface{ SomeMethod() string }
//                    > ```
   SomeMethod() string
// ^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/FieldInterface#SomeMethod.
//            kind MethodSpecification
//            display_name SomeMethod
//            signature_documentation
//            > func (FieldInterface).SomeMethod() string
//            documentation
//            > ```go
//            > func (FieldInterface).SomeMethod() string
//            > ```
  }
  
  var MyInline = struct {
//    ^^^^^^^^ definition 0.1.test `sg/inlinestruct`/MyInline.
//             kind Variable
//             display_name MyInline
//             signature_documentation
//             > var MyInline struct{privateField FieldInterface; PublicField FieldInterface}
//             documentation
//             > ```go
//             > var MyInline struct{privateField FieldInterface; PublicField FieldInterface}
//             > ```
   privateField FieldInterface
// ^^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/MyInline:privateField.
//              kind Field
//              display_name privateField
//              signature_documentation
//              > struct field privateField FieldInterface
//              documentation
//              > ```go
//              > struct field privateField FieldInterface
//              > ```
//              ^^^^^^^^^^^^^^ reference 0.1.test `sg/inlinestruct`/FieldInterface#
   PublicField  FieldInterface
// ^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/MyInline:PublicField.
//             kind Field
//             display_name PublicField
//             signature_documentation
//             > struct field PublicField FieldInterface
//             documentation
//             > ```go
//             > struct field PublicField FieldInterface
//             > ```
//              ^^^^^^^^^^^^^^ reference 0.1.test `sg/inlinestruct`/FieldInterface#
  }{}
  
//⌄ enclosing_range_start 0.1.test `sg/inlinestruct`/MyFunc().
  func MyFunc() {
//     ^^^^^^ definition 0.1.test `sg/inlinestruct`/MyFunc().
//            kind Function
//            display_name MyFunc
//            signature_documentation
//            > func MyFunc()
//            documentation
//            > ```go
//            > func MyFunc()
//            > ```
   _ = MyInline.privateField
//     ^^^^^^^^ reference 0.1.test `sg/inlinestruct`/MyInline.
//              ^^^^^^^^^^^^ reference 0.1.test `sg/inlinestruct`/MyInline:privateField.
   _ = MyInline.PublicField
//     ^^^^^^^^ reference 0.1.test `sg/inlinestruct`/MyInline.
//              ^^^^^^^^^^^ reference 0.1.test `sg/inlinestruct`/MyInline:PublicField.
  }
//⌃ enclosing_range_end 0.1.test `sg/inlinestruct`/MyFunc().
  
