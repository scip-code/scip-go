  //line grammar.y:100
  package line_directives
//        ^^^^^^^^^^^^^^^ definition 0.1.test `sg/line_directives`/
//                        kind Package
//                        display_name line_directives
//                        signature_documentation
//                        > package line_directives
//                        documentation
//                        > 
  
  import "fmt"
//        ^^^ reference github.com/golang/go/src go1.22 fmt/
  
//⌄ enclosing_range_start 0.1.test `sg/line_directives`/Generated().
  func Generated(value string) {
//     ^^^^^^^^^ definition 0.1.test `sg/line_directives`/Generated().
//               kind Function
//               display_name Generated
//               signature_documentation
//               > func Generated(value string)
//               ^^^^^ definition local 0
//                     kind Variable
//                     display_name value
//                     signature_documentation
//                     > var value string
   fmt.Println(value)
// ^^^ reference github.com/golang/go/src go1.22 fmt/
//     ^^^^^^^ reference github.com/golang/go/src go1.22 fmt/Println().
//             ^^^^^ reference local 0
  }
//⌃ enclosing_range_end 0.1.test `sg/line_directives`/Generated().
  
