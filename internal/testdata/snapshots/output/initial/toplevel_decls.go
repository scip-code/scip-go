  package initial
//        ^^^^^^^ definition 0.1.test `sg/initial`/
  
  const MY_THING = 10
//      ^^^^^^^^ definition 0.1.test `sg/initial`/MY_THING.
//               kind Constant
//               display_name MY_THING
//               signature_documentation
//               > const MY_THING untyped int = 10
//               documentation
//               > ```go
//               > const MY_THING untyped int = 10
//               > ```
  const OTHER_THING = MY_THING
//      ^^^^^^^^^^^ definition 0.1.test `sg/initial`/OTHER_THING.
//                  kind Constant
//                  display_name OTHER_THING
//                  signature_documentation
//                  > const OTHER_THING untyped int = 10
//                  documentation
//                  > ```go
//                  > const OTHER_THING untyped int = 10
//                  > ```
//                    ^^^^^^^^ reference 0.1.test `sg/initial`/MY_THING.
  
//⌄ enclosing_range_start 0.1.test `sg/initial`/usesMyThing().
  func usesMyThing() {
//     ^^^^^^^^^^^ definition 0.1.test `sg/initial`/usesMyThing().
//                 kind Function
//                 display_name usesMyThing
//                 signature_documentation
//                 > func usesMyThing()
//                 documentation
//                 > ```go
//                 > func usesMyThing()
//                 > ```
   _ = MY_THING
//     ^^^^^^^^ reference 0.1.test `sg/initial`/MY_THING.
  }
//⌃ enclosing_range_end 0.1.test `sg/initial`/usesMyThing().
  
  var initFunctions = map[string]int{}
//    ^^^^^^^^^^^^^ definition 0.1.test `sg/initial`/initFunctions.
//                  kind Variable
//                  display_name initFunctions
//                  signature_documentation
//                  > var initFunctions map[string]int
//                  documentation
//                  > ```go
//                  > var initFunctions map[string]int
//                  > ```
  
