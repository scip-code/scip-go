  package pr202
//        ^^^^^ definition 0.1.test `sg/pr202`/
  
  // Anonymous structs inside container types and nested structs.
  
  type ContainerAnon struct {
//     ^^^^^^^^^^^^^ definition 0.1.test `sg/pr202`/ContainerAnon#
//                   kind Struct
//                   display_name ContainerAnon
//                   signature_documentation
//                   > type ContainerAnon struct {
//                   >     items   []struct{ id int }
//                   >     entries map[string]struct{ count int }
//                   >     ptr     *struct{ data int }
//                   > }
//                   documentation
//                   > ```go
//                   > type ContainerAnon struct {
//                   >     items   []struct{ id int }
//                   >     entries map[string]struct{ count int }
//                   >     ptr     *struct{ data int }
//                   > }
//                   > ```
   items   []struct{ id int }
// ^^^^^ definition 0.1.test `sg/pr202`/ContainerAnon#items.
//       kind Field
//       display_name items
//       signature_documentation
//       > struct field items []struct{id int}
//       documentation
//       > ```go
//       > struct field items []struct{id int}
//       > ```
//                   ^^ definition 0.1.test `sg/pr202`/ContainerAnon#$anon_71c5ea8d9342795c#id.
//                      kind Field
//                      display_name id
//                      signature_documentation
//                      > struct field id int
//                      documentation
//                      > ```go
//                      > struct field id int
//                      > ```
   entries map[string]struct{ count int }
// ^^^^^^^ definition 0.1.test `sg/pr202`/ContainerAnon#entries.
//         kind Field
//         display_name entries
//         signature_documentation
//         > struct field entries map[string]struct{count int}
//         documentation
//         > ```go
//         > struct field entries map[string]struct{count int}
//         > ```
//                            ^^^^^ definition 0.1.test `sg/pr202`/ContainerAnon#$anon_721f9800014370ac#count.
//                                  kind Field
//                                  display_name count
//                                  signature_documentation
//                                  > struct field count int
//                                  documentation
//                                  > ```go
//                                  > struct field count int
//                                  > ```
   ptr     *struct{ data int }
// ^^^ definition 0.1.test `sg/pr202`/ContainerAnon#ptr.
//     kind Field
//     display_name ptr
//     signature_documentation
//     > struct field ptr *struct{data int}
//     documentation
//     > ```go
//     > struct field ptr *struct{data int}
//     > ```
//                  ^^^^ definition 0.1.test `sg/pr202`/ContainerAnon#$anon_944f727740dfb75d#data.
//                       kind Field
//                       display_name data
//                       signature_documentation
//                       > struct field data int
//                       documentation
//                       > ```go
//                       > struct field data int
//                       > ```
  }
  
  type DeepNested struct {
//     ^^^^^^^^^^ definition 0.1.test `sg/pr202`/DeepNested#
//                kind Struct
//                display_name DeepNested
//                signature_documentation
//                > type DeepNested struct {
//                >     outer struct{ inner struct{ value int } }
//                > }
//                documentation
//                > ```go
//                > type DeepNested struct {
//                >     outer struct{ inner struct{ value int } }
//                > }
//                > ```
   outer struct {
// ^^^^^ definition 0.1.test `sg/pr202`/DeepNested#outer.
//       kind Field
//       display_name outer
//       signature_documentation
//       > struct field outer struct{inner struct{value int}}
//       documentation
//       > ```go
//       > struct field outer struct{inner struct{value int}}
//       > ```
    inner struct {
//  ^^^^^ definition 0.1.test `sg/pr202`/DeepNested#$anon_5ee0364e53e1abd6#inner.
//        kind Field
//        display_name inner
//        signature_documentation
//        > struct field inner struct{value int}
//        documentation
//        > ```go
//        > struct field inner struct{value int}
//        > ```
     value int
//   ^^^^^ definition 0.1.test `sg/pr202`/DeepNested#$anon_5ee0364e53e1abd6#$anon_77e42bf2e5c84d1a#value.
//         kind Field
//         display_name value
//         signature_documentation
//         > struct field value int
//         documentation
//         > ```go
//         > struct field value int
//         > ```
    }
   }
  }
  
  // Two fields with identical slice-of-anonymous-struct type.
  type SliceAnonShared struct {
//     ^^^^^^^^^^^^^^^ definition 0.1.test `sg/pr202`/SliceAnonShared#
//                     kind Struct
//                     display_name SliceAnonShared
//                     signature_documentation
//                     > type SliceAnonShared struct {
//                     >     a []struct{ v int }
//                     >     b []struct{ v int }
//                     > }
//                     documentation
//                     > ```go
//                     > type SliceAnonShared struct {
//                     >     a []struct{ v int }
//                     >     b []struct{ v int }
//                     > }
//                     > ```
//                     documentation
//                     > Two fields with identical slice-of-anonymous-struct type.
   a []struct{ v int }
// ^ definition 0.1.test `sg/pr202`/SliceAnonShared#a.
//   kind Field
//   display_name a
//   signature_documentation
//   > struct field a []struct{v int}
//   documentation
//   > ```go
//   > struct field a []struct{v int}
//   > ```
//             ^ definition 0.1.test `sg/pr202`/SliceAnonShared#$anon_358bfde4cba1ecae#v.
//               kind Field
//               display_name v
//               signature_documentation
//               > struct field v int
//               documentation
//               > ```go
//               > struct field v int
//               > ```
   b []struct{ v int }
// ^ definition 0.1.test `sg/pr202`/SliceAnonShared#b.
//   kind Field
//   display_name b
//   signature_documentation
//   > struct field b []struct{v int}
//   documentation
//   > ```go
//   > struct field b []struct{v int}
//   > ```
//             ^ definition 0.1.test `sg/pr202`/SliceAnonShared#$anon_358bfde4cba1ecae#v.
//               kind Field
//               display_name v
//               signature_documentation
//               > struct field v int
//               documentation
//               > ```go
//               > struct field v int
//               > ```
  }
  
//⌄ enclosing_range_start 0.1.test `sg/pr202`/useContainerAnon().
  func useContainerAnon() {
//     ^^^^^^^^^^^^^^^^ definition 0.1.test `sg/pr202`/useContainerAnon().
//                      kind Function
//                      display_name useContainerAnon
//                      signature_documentation
//                      > func useContainerAnon()
//                      documentation
//                      > ```go
//                      > func useContainerAnon()
//                      > ```
   var c ContainerAnon
//     ^ definition local 0
//       kind Variable
//       display_name c
//       signature_documentation
//       > var c ContainerAnon
//       ^^^^^^^^^^^^^ reference 0.1.test `sg/pr202`/ContainerAnon#
   if len(c.items) > 0 {
//        ^ reference local 0
//          ^^^^^ reference 0.1.test `sg/pr202`/ContainerAnon#items.
    _ = c.items[0].id
//      ^ reference local 0
//        ^^^^^ reference 0.1.test `sg/pr202`/ContainerAnon#items.
//                 ^^ reference 0.1.test `sg/pr202`/ContainerAnon#$anon_71c5ea8d9342795c#id.
   }
   entry := c.entries["key"]
// ^^^^^ definition local 1
//       kind Variable
//       display_name entry
//       signature_documentation
//       > var entry struct{count int}
//          ^ reference local 0
//            ^^^^^^^ reference 0.1.test `sg/pr202`/ContainerAnon#entries.
   _ = entry.count
//     ^^^^^ reference local 1
//           ^^^^^ reference 0.1.test `sg/pr202`/ContainerAnon#$anon_721f9800014370ac#count.
   if c.ptr != nil {
//    ^ reference local 0
//      ^^^ reference 0.1.test `sg/pr202`/ContainerAnon#ptr.
    _ = c.ptr.data
//      ^ reference local 0
//        ^^^ reference 0.1.test `sg/pr202`/ContainerAnon#ptr.
//            ^^^^ reference 0.1.test `sg/pr202`/ContainerAnon#$anon_944f727740dfb75d#data.
   }
  
   var d DeepNested
//     ^ definition local 2
//       kind Variable
//       display_name d
//       signature_documentation
//       > var d DeepNested
//       ^^^^^^^^^^ reference 0.1.test `sg/pr202`/DeepNested#
   d.outer.inner.value = 42
// ^ reference local 2
//   ^^^^^ reference 0.1.test `sg/pr202`/DeepNested#outer.
//         ^^^^^ reference 0.1.test `sg/pr202`/DeepNested#$anon_5ee0364e53e1abd6#inner.
//               ^^^^^ reference 0.1.test `sg/pr202`/DeepNested#$anon_5ee0364e53e1abd6#$anon_77e42bf2e5c84d1a#value.
  }
//⌃ enclosing_range_end 0.1.test `sg/pr202`/useContainerAnon().
  
