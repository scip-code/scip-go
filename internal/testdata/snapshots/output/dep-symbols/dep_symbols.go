  package depsymbols
//        ^^^^^^^^^^ definition 0.1.test `sg/dep-symbols`/
//                   kind Package
//                   display_name depsymbols
//                   signature_documentation
//                   > package depsymbols
  
  import "github.com/example/deplib"
//        ^^^^^^^^^^^^^^^^^^^^^^^^^ reference github.com/example/deplib 0.1.test `github.com/example/deplib`/
  
//⌄ enclosing_range_start 0.1.test `sg/dep-symbols`/UseGenericField().
  func UseGenericField() int {
//     ^^^^^^^^^^^^^^^ definition 0.1.test `sg/dep-symbols`/UseGenericField().
//                     kind Function
//                     display_name UseGenericField
//                     signature_documentation
//                     > func UseGenericField() int
   b := deplib.Box[int]{Value: 42}
// ^ definition local 0
//   kind Variable
//   display_name b
//   signature_documentation
//   > var b Box[int]
//      ^^^^^^ reference github.com/example/deplib 0.1.test `github.com/example/deplib`/
//             ^^^ reference github.com/example/deplib 0.1.test `github.com/example/deplib`/Box#
//                      ^^^^^ reference github.com/example/deplib 0.1.test `github.com/example/deplib`/Box#Value.
   return b.Value
//        ^ reference local 0
//          ^^^^^ reference github.com/example/deplib 0.1.test `github.com/example/deplib`/Box#Value.
  }
//⌃ enclosing_range_end 0.1.test `sg/dep-symbols`/UseGenericField().
  
//⌄ enclosing_range_start 0.1.test `sg/dep-symbols`/UseGenericMethod().
  func UseGenericMethod() string {
//     ^^^^^^^^^^^^^^^^ definition 0.1.test `sg/dep-symbols`/UseGenericMethod().
//                      kind Function
//                      display_name UseGenericMethod
//                      signature_documentation
//                      > func UseGenericMethod() string
   b := deplib.Box[string]{Value: "hello"}
// ^ definition local 1
//   kind Variable
//   display_name b
//   signature_documentation
//   > var b Box[string]
//      ^^^^^^ reference github.com/example/deplib 0.1.test `github.com/example/deplib`/
//             ^^^ reference github.com/example/deplib 0.1.test `github.com/example/deplib`/Box#
//                         ^^^^^ reference github.com/example/deplib 0.1.test `github.com/example/deplib`/Box#Value.
   return b.Get()
//        ^ reference local 1
//          ^^^ reference github.com/example/deplib 0.1.test `github.com/example/deplib`/Box#Get().
  }
//⌃ enclosing_range_end 0.1.test `sg/dep-symbols`/UseGenericMethod().
  
//⌄ enclosing_range_start 0.1.test `sg/dep-symbols`/UseNonGenericField().
  func UseNonGenericField() string {
//     ^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/dep-symbols`/UseNonGenericField().
//                        kind Function
//                        display_name UseNonGenericField
//                        signature_documentation
//                        > func UseNonGenericField() string
   c := deplib.Config{Name: "test", Verbose: true}
// ^ definition local 2
//   kind Variable
//   display_name c
//   signature_documentation
//   > var c Config
//      ^^^^^^ reference github.com/example/deplib 0.1.test `github.com/example/deplib`/
//             ^^^^^^ reference github.com/example/deplib 0.1.test `github.com/example/deplib`/Config#
//                    ^^^^ reference github.com/example/deplib 0.1.test `github.com/example/deplib`/Config#Name.
//                                  ^^^^^^^ reference github.com/example/deplib 0.1.test `github.com/example/deplib`/Config#Verbose.
   return c.Name
//        ^ reference local 2
//          ^^^^ reference github.com/example/deplib 0.1.test `github.com/example/deplib`/Config#Name.
  }
//⌃ enclosing_range_end 0.1.test `sg/dep-symbols`/UseNonGenericField().
  
//⌄ enclosing_range_start 0.1.test `sg/dep-symbols`/UseConst().
  func UseConst() string {
//     ^^^^^^^^ definition 0.1.test `sg/dep-symbols`/UseConst().
//              kind Function
//              display_name UseConst
//              signature_documentation
//              > func UseConst() string
   return deplib.DefaultName
//        ^^^^^^ reference github.com/example/deplib 0.1.test `github.com/example/deplib`/
//               ^^^^^^^^^^^ reference github.com/example/deplib 0.1.test `github.com/example/deplib`/DefaultName.
  }
//⌃ enclosing_range_end 0.1.test `sg/dep-symbols`/UseConst().
  
//⌄ enclosing_range_start 0.1.test `sg/dep-symbols`/UseVar().
  func UseVar() int {
//     ^^^^^^ definition 0.1.test `sg/dep-symbols`/UseVar().
//            kind Function
//            display_name UseVar
//            signature_documentation
//            > func UseVar() int
   return deplib.GlobalCounter
//        ^^^^^^ reference github.com/example/deplib 0.1.test `github.com/example/deplib`/
//               ^^^^^^^^^^^^^ reference github.com/example/deplib 0.1.test `github.com/example/deplib`/GlobalCounter.
  }
//⌃ enclosing_range_end 0.1.test `sg/dep-symbols`/UseVar().
  
  type LocalType struct{}
//     ^^^^^^^^^ definition 0.1.test `sg/dep-symbols`/LocalType#
//               kind Struct
//               display_name LocalType
//               signature_documentation
//               > type LocalType struct{}
//               relationship github.com/example/deplib 0.1.test `github.com/example/deplib`/Stringer# implementation
  
//⌄ enclosing_range_start 0.1.test `sg/dep-symbols`/LocalType#String().
  func (l LocalType) String() string { return "local" }
//      ^ definition local 3
//        kind Variable
//        display_name l
//        signature_documentation
//        > var l LocalType
//        ^^^^^^^^^ reference 0.1.test `sg/dep-symbols`/LocalType#
//                   ^^^^^^ definition 0.1.test `sg/dep-symbols`/LocalType#String().
//                          kind Method
//                          display_name String
//                          signature_documentation
//                          > func (LocalType).String() string
//                          relationship github.com/example/deplib 0.1.test `github.com/example/deplib`/Stringer#String(). implementation
//                                                    ⌃ enclosing_range_end 0.1.test `sg/dep-symbols`/LocalType#String().
  
  type EmbeddedStringer struct {
//     ^^^^^^^^^^^^^^^^ definition 0.1.test `sg/dep-symbols`/EmbeddedStringer#
//                      kind Struct
//                      display_name EmbeddedStringer
//                      signature_documentation
//                      > type EmbeddedStringer struct{ LocalType }
//                      relationship github.com/example/deplib 0.1.test `github.com/example/deplib`/Stringer# implementation
   LocalType
// ^^^^^^^^^ definition 0.1.test `sg/dep-symbols`/EmbeddedStringer#LocalType.
//           kind Field
//           display_name LocalType
//           signature_documentation
//           > struct field LocalType LocalType
// ^^^^^^^^^ reference 0.1.test `sg/dep-symbols`/LocalType#
  }
  
  type LocalInterface interface {
//     ^^^^^^^^^^^^^^ definition 0.1.test `sg/dep-symbols`/LocalInterface#
//                    kind Interface
//                    display_name LocalInterface
//                    signature_documentation
//                    > type LocalInterface interface{ Get() int }
   Get() int
// ^^^ definition 0.1.test `sg/dep-symbols`/LocalInterface#Get.
//     kind MethodSpecification
//     display_name Get
//     signature_documentation
//     > func (LocalInterface).Get() int
  }
  
//⌄ enclosing_range_start 0.1.test `sg/dep-symbols`/UseDepWriter().
  func UseDepWriter(w deplib.Writer) {
//     ^^^^^^^^^^^^ definition 0.1.test `sg/dep-symbols`/UseDepWriter().
//                  kind Function
//                  display_name UseDepWriter
//                  signature_documentation
//                  > func UseDepWriter(w deplib.Writer)
//                  ^ definition local 4
//                    kind Variable
//                    display_name w
//                    signature_documentation
//                    > var w Writer
//                    ^^^^^^ reference github.com/example/deplib 0.1.test `github.com/example/deplib`/
//                           ^^^^^^ reference github.com/example/deplib 0.1.test `github.com/example/deplib`/Writer#
   w.Write(nil)
// ^ reference local 4
//   ^^^^^ reference github.com/example/deplib 0.1.test `github.com/example/deplib`/Writer#Write().
  }
//⌃ enclosing_range_end 0.1.test `sg/dep-symbols`/UseDepWriter().
  
