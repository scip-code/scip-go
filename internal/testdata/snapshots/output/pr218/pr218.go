  package pr218
//        ^^^^^ definition 0.1.test `sg/pr218`/
//              kind Package
//              display_name pr218
//              signature_documentation
//              > package pr218
//              documentation
//              > ```go
//              > package pr218
//              > ```
  
  // Deprecated: Use NewGreeting instead.
  const OldGreeting = "hello"
//      ^^^^^^^^^^^ definition 0.1.test `sg/pr218`/OldGreeting.
//                  kind Constant
//                  display_name OldGreeting
//                  signature_documentation
//                  > const OldGreeting untyped string = "hello"
//                  documentation
//                  > ```go
//                  > const OldGreeting untyped string = "hello"
//                  > ```
//                  documentation
//                  > Deprecated: Use NewGreeting instead.
//                  diagnostic Warning:
//                  > Deprecated
  
  const NewGreeting = "hi"
//      ^^^^^^^^^^^ definition 0.1.test `sg/pr218`/NewGreeting.
//                  kind Constant
//                  display_name NewGreeting
//                  signature_documentation
//                  > const NewGreeting untyped string = "hi"
//                  documentation
//                  > ```go
//                  > const NewGreeting untyped string = "hi"
//                  > ```
  
  // Deprecated: Use Add instead.
//⌄ enclosing_range_start 0.1.test `sg/pr218`/OldAdd().
  func OldAdd(a, b int) int {
//     ^^^^^^ definition 0.1.test `sg/pr218`/OldAdd().
//            kind Function
//            display_name OldAdd
//            signature_documentation
//            > func OldAdd(a int, b int) int
//            documentation
//            > ```go
//            > func OldAdd(a int, b int) int
//            > ```
//            documentation
//            > Deprecated: Use Add instead.
//            diagnostic Warning:
//            > Deprecated
//            ^ definition local 0
//              kind Variable
//              display_name a
//              signature_documentation
//              > var a int
//               ^ definition local 1
//                 kind Variable
//                 display_name b
//                 signature_documentation
//                 > var b int
   return a + b
//        ^ reference local 0
//            ^ reference local 1
  }
//⌃ enclosing_range_end 0.1.test `sg/pr218`/OldAdd().
  
//⌄ enclosing_range_start 0.1.test `sg/pr218`/Add().
  func Add(a, b int) int {
//     ^^^ definition 0.1.test `sg/pr218`/Add().
//         kind Function
//         display_name Add
//         signature_documentation
//         > func Add(a int, b int) int
//         documentation
//         > ```go
//         > func Add(a int, b int) int
//         > ```
//         ^ definition local 2
//           kind Variable
//           display_name a
//           signature_documentation
//           > var a int
//            ^ definition local 3
//              kind Variable
//              display_name b
//              signature_documentation
//              > var b int
   return a + b
//        ^ reference local 2
//            ^ reference local 3
  }
//⌃ enclosing_range_end 0.1.test `sg/pr218`/Add().
  
  // Deprecated: Use Server instead.
  type OldServer struct {
//     ^^^^^^^^^ definition 0.1.test `sg/pr218`/OldServer#
//               kind Struct
//               display_name OldServer
//               signature_documentation
//               > type OldServer struct {
//               >     Host string
//               >     Addr string
//               > }
//               documentation
//               > ```go
//               > type OldServer struct {
//               >     Host string
//               >     Addr string
//               > }
//               > ```
//               documentation
//               > Deprecated: Use Server instead.
//               diagnostic Warning:
//               > Deprecated
   // Deprecated: Use Addr instead.
   Host string
// ^^^^ definition 0.1.test `sg/pr218`/OldServer#Host.
//      kind Field
//      display_name Host
//      signature_documentation
//      > struct field Host string
//      documentation
//      > ```go
//      > struct field Host string
//      > ```
//      documentation
//      > Deprecated: Use Addr instead.
//      diagnostic Warning:
//      > Deprecated
   Addr string
// ^^^^ definition 0.1.test `sg/pr218`/OldServer#Addr.
//      kind Field
//      display_name Addr
//      signature_documentation
//      > struct field Addr string
//      documentation
//      > ```go
//      > struct field Addr string
//      > ```
  }
  
  type Server struct {
//     ^^^^^^ definition 0.1.test `sg/pr218`/Server#
//            kind Struct
//            display_name Server
//            signature_documentation
//            > type Server struct{ Addr string }
//            documentation
//            > ```go
//            > type Server struct{ Addr string }
//            > ```
   Addr string
// ^^^^ definition 0.1.test `sg/pr218`/Server#Addr.
//      kind Field
//      display_name Addr
//      signature_documentation
//      > struct field Addr string
//      documentation
//      > ```go
//      > struct field Addr string
//      > ```
  }
  
//⌄ enclosing_range_start 0.1.test `sg/pr218`/UseDeprecated().
  func UseDeprecated() {
//     ^^^^^^^^^^^^^ definition 0.1.test `sg/pr218`/UseDeprecated().
//                   kind Function
//                   display_name UseDeprecated
//                   signature_documentation
//                   > func UseDeprecated()
//                   documentation
//                   > ```go
//                   > func UseDeprecated()
//                   > ```
   _ = OldGreeting
//     ^^^^^^^^^^^ reference 0.1.test `sg/pr218`/OldGreeting.
//                 diagnostic Warning:
//                 > Deprecated
   _ = NewGreeting
//     ^^^^^^^^^^^ reference 0.1.test `sg/pr218`/NewGreeting.
  
   _ = OldAdd(1, 2)
//     ^^^^^^ reference 0.1.test `sg/pr218`/OldAdd().
//            diagnostic Warning:
//            > Deprecated
   _ = Add(1, 2)
//     ^^^ reference 0.1.test `sg/pr218`/Add().
  
   _ = OldServer{}
//     ^^^^^^^^^ reference 0.1.test `sg/pr218`/OldServer#
//               diagnostic Warning:
//               > Deprecated
   _ = Server{}
//     ^^^^^^ reference 0.1.test `sg/pr218`/Server#
  }
//⌃ enclosing_range_end 0.1.test `sg/pr218`/UseDeprecated().
  
