  package pr202
//        ^^^^^ definition 0.1.test `sg/pr202`/
//              kind Package
//              display_name pr202
//              signature_documentation
//              > package pr202
//              documentation
//              > ```go
//              > package pr202
//              > ```
  
  // Identical anonymous struct types should share symbols for nested fields.
  
  type IdenticalAnonFields struct {
//     ^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/pr202`/IdenticalAnonFields#
//                         kind Struct
//                         display_name IdenticalAnonFields
//                         signature_documentation
//                         > type IdenticalAnonFields struct {
//                         >     x struct{ t int }
//                         >     z struct{ t int }
//                         > }
//                         documentation
//                         > ```go
//                         > type IdenticalAnonFields struct {
//                         >     x struct{ t int }
//                         >     z struct{ t int }
//                         > }
//                         > ```
   x struct{ t int }
// ^ definition 0.1.test `sg/pr202`/IdenticalAnonFields#x.
//   kind Field
//   display_name x
//   signature_documentation
//   > struct field x struct{t int}
//   documentation
//   > ```go
//   > struct field x struct{t int}
//   > ```
//           ^ definition 0.1.test `sg/pr202`/IdenticalAnonFields#$anon_44af0565eb406c15#t.
//             kind Field
//             display_name t
//             signature_documentation
//             > struct field t int
//             documentation
//             > ```go
//             > struct field t int
//             > ```
   z struct{ t int }
// ^ definition 0.1.test `sg/pr202`/IdenticalAnonFields#z.
//   kind Field
//   display_name z
//   signature_documentation
//   > struct field z struct{t int}
//   documentation
//   > ```go
//   > struct field z struct{t int}
//   > ```
//           ^ definition 0.1.test `sg/pr202`/IdenticalAnonFields#$anon_44af0565eb406c15#t.
//             kind Field
//             display_name t
//             signature_documentation
//             > struct field t int
//             documentation
//             > ```go
//             > struct field t int
//             > ```
  }
  
//⌄ enclosing_range_start 0.1.test `sg/pr202`/useIdenticalAnonFields().
  func useIdenticalAnonFields() {
//     ^^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/pr202`/useIdenticalAnonFields().
//                            kind Function
//                            display_name useIdenticalAnonFields
//                            signature_documentation
//                            > func useIdenticalAnonFields()
//                            documentation
//                            > ```go
//                            > func useIdenticalAnonFields()
//                            > ```
   var y IdenticalAnonFields
//     ^ definition local 0
//       kind Variable
//       display_name y
//       signature_documentation
//       > var y IdenticalAnonFields
//       ^^^^^^^^^^^^^^^^^^^ reference 0.1.test `sg/pr202`/IdenticalAnonFields#
   y.x = y.z
// ^ reference local 0
//   ^ reference 0.1.test `sg/pr202`/IdenticalAnonFields#x.
//       ^ reference local 0
//         ^ reference 0.1.test `sg/pr202`/IdenticalAnonFields#z.
   y.x.t = 1
// ^ reference local 0
//   ^ reference 0.1.test `sg/pr202`/IdenticalAnonFields#x.
//     ^ reference 0.1.test `sg/pr202`/IdenticalAnonFields#$anon_44af0565eb406c15#t.
   y.z.t = 2
// ^ reference local 0
//   ^ reference 0.1.test `sg/pr202`/IdenticalAnonFields#z.
//     ^ reference 0.1.test `sg/pr202`/IdenticalAnonFields#$anon_44af0565eb406c15#t.
  }
//⌃ enclosing_range_end 0.1.test `sg/pr202`/useIdenticalAnonFields().
  
  // Different field order means different type — symbols must NOT unify.
  type FieldOrderMatters struct {
//     ^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/pr202`/FieldOrderMatters#
//                       kind Struct
//                       display_name FieldOrderMatters
//                       signature_documentation
//                       > type FieldOrderMatters struct {
//                       >     a struct {
//                       >         x int
//                       >         y string
//                       >     }
//                       >     b struct {
//                       >         y string
//                       >         x int
//                       >     }
//                       > }
//                       documentation
//                       > ```go
//                       > type FieldOrderMatters struct {
//                       >     a struct {
//                       >         x int
//                       >         y string
//                       >     }
//                       >     b struct {
//                       >         y string
//                       >         x int
//                       >     }
//                       > }
//                       > ```
//                       documentation
//                       > Different field order means different type — symbols must NOT unify.
   a struct {
// ^ definition 0.1.test `sg/pr202`/FieldOrderMatters#a.
//   kind Field
//   display_name a
//   signature_documentation
//   > struct field a struct{x int; y string}
//   documentation
//   > ```go
//   > struct field a struct{x int; y string}
//   > ```
    x int
//  ^ definition 0.1.test `sg/pr202`/FieldOrderMatters#$anon_c0a8952b3a214f68#x.
//    kind Field
//    display_name x
//    signature_documentation
//    > struct field x int
//    documentation
//    > ```go
//    > struct field x int
//    > ```
    y string
//  ^ definition 0.1.test `sg/pr202`/FieldOrderMatters#$anon_c0a8952b3a214f68#y.
//    kind Field
//    display_name y
//    signature_documentation
//    > struct field y string
//    documentation
//    > ```go
//    > struct field y string
//    > ```
   }
   b struct {
// ^ definition 0.1.test `sg/pr202`/FieldOrderMatters#b.
//   kind Field
//   display_name b
//   signature_documentation
//   > struct field b struct{y string; x int}
//   documentation
//   > ```go
//   > struct field b struct{y string; x int}
//   > ```
    y string
//  ^ definition 0.1.test `sg/pr202`/FieldOrderMatters#$anon_b8d88f3211c0d7a4#y.
//    kind Field
//    display_name y
//    signature_documentation
//    > struct field y string
//    documentation
//    > ```go
//    > struct field y string
//    > ```
    x int
//  ^ definition 0.1.test `sg/pr202`/FieldOrderMatters#$anon_b8d88f3211c0d7a4#x.
//    kind Field
//    display_name x
//    signature_documentation
//    > struct field x int
//    documentation
//    > ```go
//    > struct field x int
//    > ```
   }
  }
  
  // Different struct tags — symbols must NOT unify.
  type DifferentTags struct {
//     ^^^^^^^^^^^^^ definition 0.1.test `sg/pr202`/DifferentTags#
//                   kind Struct
//                   display_name DifferentTags
//                   signature_documentation
//                   > type DifferentTags struct {
//                   >     a struct {
//                   >         Name string "json:\"name\""
//                   >     }
//                   >     b struct {
//                   >         Name string "json:\"full_name\""
//                   >     }
//                   > }
//                   documentation
//                   > ```go
//                   > type DifferentTags struct {
//                   >     a struct {
//                   >         Name string "json:\"name\""
//                   >     }
//                   >     b struct {
//                   >         Name string "json:\"full_name\""
//                   >     }
//                   > }
//                   > ```
//                   documentation
//                   > Different struct tags — symbols must NOT unify.
   a struct {
// ^ definition 0.1.test `sg/pr202`/DifferentTags#a.
//   kind Field
//   display_name a
//   signature_documentation
//   > struct field a struct{Name string `json:"name"`}
//   documentation
//   > ```go
//   > struct field a struct{Name string `json:"name"`}
//   > ```
    Name string `json:"name"`
//  ^^^^ definition 0.1.test `sg/pr202`/DifferentTags#$anon_ed545d904f2246eb#Name.
//       kind Field
//       display_name Name
//       signature_documentation
//       > struct field Name string
//       documentation
//       > ```go
//       > struct field Name string
//       > ```
   }
   b struct {
// ^ definition 0.1.test `sg/pr202`/DifferentTags#b.
//   kind Field
//   display_name b
//   signature_documentation
//   > struct field b struct{Name string `json:"full_name"`}
//   documentation
//   > ```go
//   > struct field b struct{Name string `json:"full_name"`}
//   > ```
    Name string `json:"full_name"`
//  ^^^^ definition 0.1.test `sg/pr202`/DifferentTags#$anon_29f1ad2683b11ed0#Name.
//       kind Field
//       display_name Name
//       signature_documentation
//       > struct field Name string
//       documentation
//       > ```go
//       > struct field Name string
//       > ```
   }
  }
  
