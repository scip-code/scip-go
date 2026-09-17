  package initial
//        ^^^^^^^ definition 0.1.test `sg/initial`/
//                kind Package
//                display_name initial
//                signature_documentation
//                > package initial
//                documentation
//                > ```go
//                > package initial
//                > ```
//                documentation
//                > This is a module for testing purposes.
//                > This should now be the place that has a definition
  
//⌄ enclosing_range_start 0.1.test `sg/initial`/UsesBuiltin().
  func UsesBuiltin() int {
//     ^^^^^^^^^^^ definition 0.1.test `sg/initial`/UsesBuiltin().
//                 kind Function
//                 display_name UsesBuiltin
//                 signature_documentation
//                 > func UsesBuiltin() int
//                 documentation
//                 > ```go
//                 > func UsesBuiltin() int
//                 > ```
   var x int = 5
//     ^ definition local 0
//       kind Variable
//       display_name x
//       signature_documentation
//       > var x int
   return x
//        ^ reference local 0
  }
//⌃ enclosing_range_end 0.1.test `sg/initial`/UsesBuiltin().
  
