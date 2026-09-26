  package pr216
//        ^^^^^ definition 0.1.test `sg/pr216`/
//              kind Package
//              display_name pr216
//              signature_documentation
//              > package pr216
  
  import "fmt"
//        ^^^ reference 0.1.test fmt/
  
//⌄ enclosing_range_start 0.1.test `sg/pr216`/Greeting.
  const Greeting = "hello"
//      ^^^^^^^^ definition 0.1.test `sg/pr216`/Greeting.
//               kind Constant
//               display_name Greeting
//               signature_documentation
//               > const Greeting untyped string = "hello"
//                       ⌃ enclosing_range_end 0.1.test `sg/pr216`/Greeting.
  
//⌄ enclosing_range_start 0.1.test `sg/pr216`/UseFmt().
  func UseFmt() {
//     ^^^^^^ definition 0.1.test `sg/pr216`/UseFmt().
//            kind Function
//            display_name UseFmt
//            signature_documentation
//            > func UseFmt()
   fmt.Println(Greeting)
// ^^^ reference 0.1.test fmt/
//     ^^^^^^^ reference 0.1.test fmt/Println().
//             ^^^^^^^^ reference 0.1.test `sg/pr216`/Greeting.
  }
//⌃ enclosing_range_end 0.1.test `sg/pr216`/UseFmt().
  
