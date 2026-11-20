  package deplib
//        ^^^^^^ definition github.com/example/deplib 0.1.test `github.com/example/deplib`/
//               kind Package
//               display_name deplib
//               signature_documentation
//               > package deplib
  
//⌄ enclosing_range_start github.com/example/deplib 0.1.test `github.com/example/deplib`/Box#
  type Box[T any] struct {
//     ^^^ definition github.com/example/deplib 0.1.test `github.com/example/deplib`/Box#
//         kind Struct
//         display_name Box
//         signature_documentation
//         > type Box struct{ Value T }
//         ^ definition local 0
//           kind Interface
//           display_name T
//           signature_documentation
//           > type parameter T any
// ⌄ enclosing_range_start github.com/example/deplib 0.1.test `github.com/example/deplib`/Box#Value.
   Value T
// ^^^^^ definition github.com/example/deplib 0.1.test `github.com/example/deplib`/Box#Value.
//       kind Field
//       display_name Value
//       signature_documentation
//       > struct field Value T
//       ^ reference local 0
//       ⌃ enclosing_range_end github.com/example/deplib 0.1.test `github.com/example/deplib`/Box#Value.
  }
//⌃ enclosing_range_end github.com/example/deplib 0.1.test `github.com/example/deplib`/Box#
  
//⌄ enclosing_range_start github.com/example/deplib 0.1.test `github.com/example/deplib`/Box#Get().
  func (b Box[T]) Get() T {
//      ^ definition local 1
//        kind Variable
//        display_name b
//        signature_documentation
//        > var b Box[T]
//        ^^^ reference github.com/example/deplib 0.1.test `github.com/example/deplib`/Box#
//            ^ definition local 2
//              kind Interface
//              display_name T
//              signature_documentation
//              > type parameter T any
//            ^ reference local 2
//                ^^^ definition github.com/example/deplib 0.1.test `github.com/example/deplib`/Box#Get().
//                    kind Method
//                    display_name Get
//                    signature_documentation
//                    > func (Box[T]).Get() T
//                      ^ reference local 2
   return b.Value
//        ^ reference local 1
//          ^^^^^ reference github.com/example/deplib 0.1.test `github.com/example/deplib`/Box#Value.
  }
//⌃ enclosing_range_end github.com/example/deplib 0.1.test `github.com/example/deplib`/Box#Get().
  
//⌄ enclosing_range_start github.com/example/deplib 0.1.test `github.com/example/deplib`/Config#
  type Config struct {
//     ^^^^^^ definition github.com/example/deplib 0.1.test `github.com/example/deplib`/Config#
//            kind Struct
//            display_name Config
//            signature_documentation
//            > type Config struct {
//            >     Name    string
//            >     Verbose bool
//            > }
// ⌄ enclosing_range_start github.com/example/deplib 0.1.test `github.com/example/deplib`/Config#Name.
   Name    string
// ^^^^ definition github.com/example/deplib 0.1.test `github.com/example/deplib`/Config#Name.
//      kind Field
//      display_name Name
//      signature_documentation
//      > struct field Name string
//              ⌃ enclosing_range_end github.com/example/deplib 0.1.test `github.com/example/deplib`/Config#Name.
// ⌄ enclosing_range_start github.com/example/deplib 0.1.test `github.com/example/deplib`/Config#Verbose.
   Verbose bool
// ^^^^^^^ definition github.com/example/deplib 0.1.test `github.com/example/deplib`/Config#Verbose.
//         kind Field
//         display_name Verbose
//         signature_documentation
//         > struct field Verbose bool
//            ⌃ enclosing_range_end github.com/example/deplib 0.1.test `github.com/example/deplib`/Config#Verbose.
  }
//⌃ enclosing_range_end github.com/example/deplib 0.1.test `github.com/example/deplib`/Config#
  
//⌄ enclosing_range_start github.com/example/deplib 0.1.test `github.com/example/deplib`/DefaultName.
  const DefaultName = "default"
//      ^^^^^^^^^^^ definition github.com/example/deplib 0.1.test `github.com/example/deplib`/DefaultName.
//                  kind Constant
//                  display_name DefaultName
//                  signature_documentation
//                  > const DefaultName untyped string = "default"
//                            ⌃ enclosing_range_end github.com/example/deplib 0.1.test `github.com/example/deplib`/DefaultName.
  
//⌄ enclosing_range_start github.com/example/deplib 0.1.test `github.com/example/deplib`/GlobalCounter.
  var GlobalCounter int
//    ^^^^^^^^^^^^^ definition github.com/example/deplib 0.1.test `github.com/example/deplib`/GlobalCounter.
//                  kind Variable
//                  display_name GlobalCounter
//                  signature_documentation
//                  > var GlobalCounter int
//                    ⌃ enclosing_range_end github.com/example/deplib 0.1.test `github.com/example/deplib`/GlobalCounter.
  
//⌄ enclosing_range_start github.com/example/deplib 0.1.test `github.com/example/deplib`/Stringer#
  type Stringer interface {
//     ^^^^^^^^ definition github.com/example/deplib 0.1.test `github.com/example/deplib`/Stringer#
//              kind Interface
//              display_name Stringer
//              signature_documentation
//              > type Stringer interface{ String() string }
// ⌄ enclosing_range_start github.com/example/deplib 0.1.test `github.com/example/deplib`/Stringer#String.
   String() string
// ^^^^^^ definition github.com/example/deplib 0.1.test `github.com/example/deplib`/Stringer#String.
//        kind MethodSpecification
//        display_name String
//        signature_documentation
//        > func (Stringer).String() string
//               ⌃ enclosing_range_end github.com/example/deplib 0.1.test `github.com/example/deplib`/Stringer#String.
  }
//⌃ enclosing_range_end github.com/example/deplib 0.1.test `github.com/example/deplib`/Stringer#
  
//⌄ enclosing_range_start github.com/example/deplib 0.1.test `github.com/example/deplib`/Writer#
  type Writer interface {
//     ^^^^^^ definition github.com/example/deplib 0.1.test `github.com/example/deplib`/Writer#
//            kind Interface
//            display_name Writer
//            signature_documentation
//            > type Writer interface {
//            >     Write(p []byte) (n int, err error)
//            > }
// ⌄ enclosing_range_start github.com/example/deplib 0.1.test `github.com/example/deplib`/Writer#Write.
   Write(p []byte) (n int, err error)
// ^^^^^ definition github.com/example/deplib 0.1.test `github.com/example/deplib`/Writer#Write.
//       kind MethodSpecification
//       display_name Write
//       signature_documentation
//       > func (Writer).Write(p []byte) (n int, err error)
//       ^ definition local 3
//         kind Variable
//         display_name p
//         signature_documentation
//         > var p []byte
//                  ^ definition local 4
//                    kind Variable
//                    display_name n
//                    signature_documentation
//                    > var n int
//                         ^^^ definition local 5
//                             kind Variable
//                             display_name err
//                             signature_documentation
//                             > var err error
//                                  ⌃ enclosing_range_end github.com/example/deplib 0.1.test `github.com/example/deplib`/Writer#Write.
  }
//⌃ enclosing_range_end github.com/example/deplib 0.1.test `github.com/example/deplib`/Writer#
  
