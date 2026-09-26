  package pr202
//        ^^^^^ definition 0.1.test `sg/pr202`/
  
  // Multi-name field declarations: a, b share a type and must be siblings.
  
//⌄ enclosing_range_start 0.1.test `sg/pr202`/MultiNameStruct#
  type MultiNameStruct struct {
//     ^^^^^^^^^^^^^^^ definition 0.1.test `sg/pr202`/MultiNameStruct#
//                     kind Struct
//                     display_name MultiNameStruct
//                     signature_documentation
//                     > type MultiNameStruct struct {
//                     >     a struct {
//                     >         x int
//                     >         y string
//                     >     }
//                     >     b struct {
//                     >         x int
//                     >         y string
//                     >     }
//                     > }
// ⌄ enclosing_range_start 0.1.test `sg/pr202`/MultiNameStruct#a.
// ⌄ enclosing_range_start 0.1.test `sg/pr202`/MultiNameStruct#b.
   a, b struct {
// ^ definition 0.1.test `sg/pr202`/MultiNameStruct#a.
//   kind Field
//   display_name a
//   signature_documentation
//   > struct field a struct{x int; y string}
//    ^ definition 0.1.test `sg/pr202`/MultiNameStruct#b.
//      kind Field
//      display_name b
//      signature_documentation
//      > struct field b struct{x int; y string}
//  ⌄ enclosing_range_start 0.1.test `sg/pr202`/MultiNameStruct#$anon_c0a8952b3a214f68#x.
    x int
//  ^ definition 0.1.test `sg/pr202`/MultiNameStruct#$anon_c0a8952b3a214f68#x.
//    kind Field
//    display_name x
//    signature_documentation
//    > struct field x int
//      ⌃ enclosing_range_end 0.1.test `sg/pr202`/MultiNameStruct#$anon_c0a8952b3a214f68#x.
//  ⌄ enclosing_range_start 0.1.test `sg/pr202`/MultiNameStruct#$anon_c0a8952b3a214f68#y.
    y string
//  ^ definition 0.1.test `sg/pr202`/MultiNameStruct#$anon_c0a8952b3a214f68#y.
//    kind Field
//    display_name y
//    signature_documentation
//    > struct field y string
//         ⌃ enclosing_range_end 0.1.test `sg/pr202`/MultiNameStruct#$anon_c0a8952b3a214f68#y.
   }
// ⌃ enclosing_range_end 0.1.test `sg/pr202`/MultiNameStruct#a.
// ⌃ enclosing_range_end 0.1.test `sg/pr202`/MultiNameStruct#b.
  }
//⌃ enclosing_range_end 0.1.test `sg/pr202`/MultiNameStruct#
  
//⌄ enclosing_range_start 0.1.test `sg/pr202`/useMultiNameFields().
  func useMultiNameFields() {
//     ^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/pr202`/useMultiNameFields().
//                        kind Function
//                        display_name useMultiNameFields
//                        signature_documentation
//                        > func useMultiNameFields()
   var m MultiNameStruct
//     ^ definition local 0
//       kind Variable
//       display_name m
//       signature_documentation
//       > var m MultiNameStruct
//       ^^^^^^^^^^^^^^^ reference 0.1.test `sg/pr202`/MultiNameStruct#
   m.a.x = 1
// ^ reference local 0
//   ^ reference 0.1.test `sg/pr202`/MultiNameStruct#a.
//     ^ reference 0.1.test `sg/pr202`/MultiNameStruct#$anon_c0a8952b3a214f68#x.
   m.b.x = 2
// ^ reference local 0
//   ^ reference 0.1.test `sg/pr202`/MultiNameStruct#b.
//     ^ reference 0.1.test `sg/pr202`/MultiNameStruct#$anon_c0a8952b3a214f68#x.
   m.a = m.b
// ^ reference local 0
//   ^ reference 0.1.test `sg/pr202`/MultiNameStruct#a.
//       ^ reference local 0
//         ^ reference 0.1.test `sg/pr202`/MultiNameStruct#b.
  }
//⌃ enclosing_range_end 0.1.test `sg/pr202`/useMultiNameFields().
  
