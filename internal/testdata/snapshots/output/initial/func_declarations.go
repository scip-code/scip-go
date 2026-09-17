  package initial
//        ^^^^^^^ definition 0.1.test `sg/initial`/
  
//⌄ enclosing_range_start 0.1.test `sg/initial`/UsesLater().
  func UsesLater() {
//     ^^^^^^^^^ definition 0.1.test `sg/initial`/UsesLater().
//               kind Function
//               display_name UsesLater
//               signature_documentation
//               > func UsesLater()
//               documentation
//               > ```go
//               > func UsesLater()
//               > ```
   DefinedLater()
// ^^^^^^^^^^^^ reference 0.1.test `sg/initial`/DefinedLater().
  }
//⌃ enclosing_range_end 0.1.test `sg/initial`/UsesLater().
  
//⌄ enclosing_range_start 0.1.test `sg/initial`/DefinedLater().
  func DefinedLater() {}
//     ^^^^^^^^^^^^ definition 0.1.test `sg/initial`/DefinedLater().
//                  kind Function
//                  display_name DefinedLater
//                  signature_documentation
//                  > func DefinedLater()
//                  documentation
//                  > ```go
//                  > func DefinedLater()
//                  > ```
//                     ⌃ enclosing_range_end 0.1.test `sg/initial`/DefinedLater().
  
