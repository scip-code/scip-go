  package embedded
//        ^^^^^^^^ definition 0.1.test `sg/embedded`/
//                 kind Package
//                 display_name embedded
//                 signature_documentation
//                 > package embedded
  
  import (
   "fmt"
//  ^^^ reference github.com/golang/go/src go1.22 fmt/
   "os/exec"
//  ^^^^^^^ reference github.com/golang/go/src go1.22 `os/exec`/
  )
  
  type osExecCommand struct {
//     ^^^^^^^^^^^^^ definition 0.1.test `sg/embedded`/osExecCommand#
//                   kind Struct
//                   display_name osExecCommand
//                   signature_documentation
//                   > type osExecCommand struct{ *exec.Cmd }
//                   relationship github.com/golang/go/src go1.22 fmt/Stringer# implementation
   *exec.Cmd
//  ^^^^ reference github.com/golang/go/src go1.22 `os/exec`/
//       ^^^ definition 0.1.test `sg/embedded`/osExecCommand#Cmd.
//           kind Field
//           display_name Cmd
//           signature_documentation
//           > struct field Cmd *exec.Cmd
//       ^^^ reference github.com/golang/go/src go1.22 `os/exec`/Cmd#
  }
  
//⌄ enclosing_range_start 0.1.test `sg/embedded`/wrapExecCommand().
  func wrapExecCommand(c *exec.Cmd) {
//     ^^^^^^^^^^^^^^^ definition 0.1.test `sg/embedded`/wrapExecCommand().
//                     kind Function
//                     display_name wrapExecCommand
//                     signature_documentation
//                     > func wrapExecCommand(c *exec.Cmd)
//                     ^ definition local 0
//                       kind Variable
//                       display_name c
//                       signature_documentation
//                       > var c *Cmd
//                        ^^^^ reference github.com/golang/go/src go1.22 `os/exec`/
//                             ^^^ reference github.com/golang/go/src go1.22 `os/exec`/Cmd#
   _ = &osExecCommand{Cmd: c}
//      ^^^^^^^^^^^^^ reference 0.1.test `sg/embedded`/osExecCommand#
//                    ^^^ reference 0.1.test `sg/embedded`/osExecCommand#Cmd.
//                         ^ reference local 0
  }
//⌃ enclosing_range_end 0.1.test `sg/embedded`/wrapExecCommand().
  
  type Inner struct {
//     ^^^^^ definition 0.1.test `sg/embedded`/Inner#
//           kind Struct
//           display_name Inner
//           signature_documentation
//           > type Inner struct {
//           >     X int
//           >     Y int
//           >     Z int
//           > }
   X int
// ^ definition 0.1.test `sg/embedded`/Inner#X.
//   kind Field
//   display_name X
//   signature_documentation
//   > struct field X int
   Y int
// ^ definition 0.1.test `sg/embedded`/Inner#Y.
//   kind Field
//   display_name Y
//   signature_documentation
//   > struct field Y int
   Z int
// ^ definition 0.1.test `sg/embedded`/Inner#Z.
//   kind Field
//   display_name Z
//   signature_documentation
//   > struct field Z int
  }
  
  type Outer struct {
//     ^^^^^ definition 0.1.test `sg/embedded`/Outer#
//           kind Struct
//           display_name Outer
//           signature_documentation
//           > type Outer struct {
//           >     Inner
//           >     W int
//           > }
   Inner
// ^^^^^ definition 0.1.test `sg/embedded`/Outer#Inner.
//       kind Field
//       display_name Inner
//       signature_documentation
//       > struct field Inner Inner
// ^^^^^ reference 0.1.test `sg/embedded`/Inner#
   W int
// ^ definition 0.1.test `sg/embedded`/Outer#W.
//   kind Field
//   display_name W
//   signature_documentation
//   > struct field W int
  }
  
//⌄ enclosing_range_start 0.1.test `sg/embedded`/useOfCompositeStructs().
  func useOfCompositeStructs() {
//     ^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/embedded`/useOfCompositeStructs().
//                           kind Function
//                           display_name useOfCompositeStructs
//                           signature_documentation
//                           > func useOfCompositeStructs()
   o := Outer{
// ^ definition local 1
//   kind Variable
//   display_name o
//   signature_documentation
//   > var o Outer
//      ^^^^^ reference 0.1.test `sg/embedded`/Outer#
    Inner: Inner{
//  ^^^^^ reference 0.1.test `sg/embedded`/Outer#Inner.
//         ^^^^^ reference 0.1.test `sg/embedded`/Inner#
     X: 1,
//   ^ reference 0.1.test `sg/embedded`/Inner#X.
     Y: 2,
//   ^ reference 0.1.test `sg/embedded`/Inner#Y.
     Z: 3,
//   ^ reference 0.1.test `sg/embedded`/Inner#Z.
    },
    W: 4,
//  ^ reference 0.1.test `sg/embedded`/Outer#W.
   }
  
   fmt.Printf("> %d\n", o.X)
// ^^^ reference github.com/golang/go/src go1.22 fmt/
//     ^^^^^^ reference github.com/golang/go/src go1.22 fmt/Printf().
//                      ^ reference local 1
//                        ^ reference 0.1.test `sg/embedded`/Inner#X.
   fmt.Println(o.Inner.Y)
// ^^^ reference github.com/golang/go/src go1.22 fmt/
//     ^^^^^^^ reference github.com/golang/go/src go1.22 fmt/Println().
//             ^ reference local 1
//               ^^^^^ reference 0.1.test `sg/embedded`/Outer#Inner.
//                     ^ reference 0.1.test `sg/embedded`/Inner#Y.
  }
//⌃ enclosing_range_end 0.1.test `sg/embedded`/useOfCompositeStructs().
  
