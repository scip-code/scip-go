  package pr206
//        ^^^^^ definition 0.1.test `sg/pr206`/
  
//⌄ enclosing_range_start 0.1.test `sg/pr206`/Base#
  type Base struct{}
//     ^^^^ definition 0.1.test `sg/pr206`/Base#
//          kind Struct
//          display_name Base
//          signature_documentation
//          > type Base struct{}
//                 ⌃ enclosing_range_end 0.1.test `sg/pr206`/Base#
  
//⌄ enclosing_range_start 0.1.test `sg/pr206`/Container#
  type Container struct {
//     ^^^^^^^^^ definition 0.1.test `sg/pr206`/Container#
//               kind Struct
//               display_name Container
//               signature_documentation
//               > type Container struct{ Base }
// ⌄ enclosing_range_start 0.1.test `sg/pr206`/Container#Base.
   // Base is embedded to inherit shared fields.
   Base
// ^^^^ definition 0.1.test `sg/pr206`/Container#Base.
//      kind Field
//      display_name Base
//      signature_documentation
//      > struct field Base Base
//      documentation
//      > Base is embedded to inherit shared fields.
// ^^^^ reference 0.1.test `sg/pr206`/Base#
//    ⌃ enclosing_range_end 0.1.test `sg/pr206`/Container#Base.
  }
//⌃ enclosing_range_end 0.1.test `sg/pr206`/Container#
  
