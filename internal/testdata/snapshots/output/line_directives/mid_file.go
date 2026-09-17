  package line_directives
//        ^^^^^^^^^^^^^^^ definition 0.1.test `sg/line_directives`/
  
  import "strings"
//        ^^^^^^^ reference github.com/golang/go/src go1.22 strings/
  
//⌄ enclosing_range_start 0.1.test `sg/line_directives`/MidFile().
  func MidFile(value string) string {
//     ^^^^^^^ definition 0.1.test `sg/line_directives`/MidFile().
//             kind Function
//             display_name MidFile
//             signature_documentation
//             > func MidFile(value string) string
//             ^^^^^ definition local 0
//                   kind Variable
//                   display_name value
//                   signature_documentation
//                   > var value string
   before := strings.TrimSpace(value)
// ^^^^^^ definition local 1
//        kind Variable
//        display_name before
//        signature_documentation
//        > var before string
//           ^^^^^^^ reference github.com/golang/go/src go1.22 strings/
//                   ^^^^^^^^^ reference github.com/golang/go/src go1.22 strings/TrimSpace().
//                             ^^^^^ reference local 0
  //line target.go:3:1
   after := strings.TrimSpace(value)
// ^^^^^ definition local 2
//       kind Variable
//       display_name after
//       signature_documentation
//       > var after string
//          ^^^^^^^ reference github.com/golang/go/src go1.22 strings/
//                  ^^^^^^^^^ reference github.com/golang/go/src go1.22 strings/TrimSpace().
//                            ^^^^^ reference local 0
   return before + after
//        ^^^^^^ reference local 1
//                 ^^^^^ reference local 2
  }
//⌃ enclosing_range_end 0.1.test `sg/line_directives`/MidFile().
  
