  package consumer
//        ^^^^^^^^ definition 0.1.test `sg/pr258`/
//                 kind Package
//                 display_name consumer
//                 signature_documentation
//                 > package consumer
  
  import "github.com/example/deplib"
//        ^^^^^^^^^^^^^^^^^^^^^^^^^ reference github.com/example/deplib 0.1.test `github.com/example/deplib`/
  
  var Sentinel deplib.CustomErr
//    ^^^^^^^^ definition 0.1.test `sg/pr258`/Sentinel.
//             kind Variable
//             display_name Sentinel
//             signature_documentation
//             > var Sentinel deplib.CustomErr
//             ^^^^^^ reference github.com/example/deplib 0.1.test `github.com/example/deplib`/
//                    ^^^^^^^^^ reference github.com/example/deplib 0.1.test `github.com/example/deplib`/CustomErr#
  
//⌄ enclosing_range_start 0.1.test `sg/pr258`/New().
  func New() deplib.CustomErr { return nil }
//     ^^^ definition 0.1.test `sg/pr258`/New().
//         kind Function
//         display_name New
//         signature_documentation
//         > func New() deplib.CustomErr
//           ^^^^^^ reference github.com/example/deplib 0.1.test `github.com/example/deplib`/
//                  ^^^^^^^^^ reference github.com/example/deplib 0.1.test `github.com/example/deplib`/CustomErr#
//                                         ⌃ enclosing_range_end 0.1.test `sg/pr258`/New().
  
  type Wrapper struct {
//     ^^^^^^^ definition 0.1.test `sg/pr258`/Wrapper#
//             kind Struct
//             display_name Wrapper
//             signature_documentation
//             > type Wrapper struct{ Err deplib.CustomErr }
   Err deplib.CustomErr
// ^^^ definition 0.1.test `sg/pr258`/Wrapper#Err.
//     kind Field
//     display_name Err
//     signature_documentation
//     > struct field Err deplib.CustomErr
//     ^^^^^^ reference github.com/example/deplib 0.1.test `github.com/example/deplib`/
//            ^^^^^^^^^ reference github.com/example/deplib 0.1.test `github.com/example/deplib`/CustomErr#
  }
  
