  package pr206
//        ^^^^^ definition 0.1.test `sg/pr206`/
  
  type Base struct{}
//     ^^^^ definition 0.1.test `sg/pr206`/Base#
//          kind Struct
//          display_name Base
//          signature_documentation
//          > type Base struct{}
//          documentation
//          > ```go
//          > type Base struct{}
//          > ```
  
  type Container struct {
//     ^^^^^^^^^ definition 0.1.test `sg/pr206`/Container#
//               kind Struct
//               display_name Container
//               signature_documentation
//               > type Container struct{ Base }
//               documentation
//               > ```go
//               > type Container struct{ Base }
//               > ```
   // Base is embedded to inherit shared fields.
   Base
// ^^^^ definition 0.1.test `sg/pr206`/Container#Base.
//      kind Field
//      display_name Base
//      signature_documentation
//      > struct field Base Base
//      documentation
//      > ```go
//      > struct field Base Base
//      > ```
//      documentation
//      > Base is embedded to inherit shared fields.
// ^^^^ reference 0.1.test `sg/pr206`/Base#
  }
  
