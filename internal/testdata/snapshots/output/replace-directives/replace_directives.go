  package replacers
//        ^^^^^^^^^ definition 0.1.test `sg/replace-directives`/
//                  kind Package
//                  display_name replacers
//                  signature_documentation
//                  > package replacers
//                  documentation
//                  > ```go
//                  > package replacers
//                  > ```
  
  import (
   "fmt"
//  ^^^ reference github.com/golang/go/src go1.22 fmt/
  
   replaced "github.com/example/original"
// ^^^^^^^^ reference github.com/example/replaced 0.1.test `github.com/example/original`/
//           ^^^^^^^^^^^^^^^^^^^^^^^^^^^ reference github.com/example/replaced 0.1.test `github.com/example/original`/
  )
  
//⌄ enclosing_range_start 0.1.test `sg/replace-directives`/Something().
  func Something() {
//     ^^^^^^^^^ definition 0.1.test `sg/replace-directives`/Something().
//               kind Function
//               display_name Something
//               signature_documentation
//               > func Something()
//               documentation
//               > ```go
//               > func Something()
//               > ```
   fmt.Println(replaced.DefaultConfig)
// ^^^ reference github.com/golang/go/src go1.22 fmt/
//     ^^^^^^^ reference github.com/golang/go/src go1.22 fmt/Println().
//             ^^^^^^^^ reference github.com/example/replaced 0.1.test `github.com/example/original`/
//                      ^^^^^^^^^^^^^ reference github.com/example/replaced 0.1.test `github.com/example/original`/DefaultConfig.
  }
//⌃ enclosing_range_end 0.1.test `sg/replace-directives`/Something().
  
