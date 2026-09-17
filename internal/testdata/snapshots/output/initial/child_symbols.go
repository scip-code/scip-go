  package initial
//        ^^^^^^^ definition 0.1.test `sg/initial`/
  
  // Const is a constant equal to 5. It's the best constant I've ever written. 😹
  const Const = 5
//      ^^^^^ definition 0.1.test `sg/initial`/Const.
//            kind Constant
//            display_name Const
//            signature_documentation
//            > const Const untyped int = 5
//            documentation
//            > ```go
//            > const Const untyped int = 5
//            > ```
//            documentation
//            > Const is a constant equal to 5. It's the best constant I've ever written. 😹
  
  // Docs for the const block itself.
  const (
   // ConstBlock1 is a constant in a block.
   ConstBlock1 = 1
// ^^^^^^^^^^^ definition 0.1.test `sg/initial`/ConstBlock1.
//             kind Constant
//             display_name ConstBlock1
//             signature_documentation
//             > const ConstBlock1 untyped int = 1
//             documentation
//             > ```go
//             > const ConstBlock1 untyped int = 1
//             > ```
//             documentation
//             > ConstBlock1 is a constant in a block.
//             documentation
//             > Docs for the const block itself.
  
   // ConstBlock2 is a constant in a block.
   ConstBlock2 = 2
// ^^^^^^^^^^^ definition 0.1.test `sg/initial`/ConstBlock2.
//             kind Constant
//             display_name ConstBlock2
//             signature_documentation
//             > const ConstBlock2 untyped int = 2
//             documentation
//             > ```go
//             > const ConstBlock2 untyped int = 2
//             > ```
//             documentation
//             > ConstBlock2 is a constant in a block.
//             documentation
//             > Docs for the const block itself.
  )
  
  // Var is a variable interface.
  var Var Interface = &Struct{Field: "bar!"}
//    ^^^ definition 0.1.test `sg/initial`/Var.
//        kind Variable
//        display_name Var
//        signature_documentation
//        > var Var Interface
//        documentation
//        > ```go
//        > var Var Interface
//        > ```
//        documentation
//        > Var is a variable interface.
//        ^^^^^^^^^ reference 0.1.test `sg/initial`/Interface#
//                     ^^^^^^ reference 0.1.test `sg/initial`/Struct#
//                            ^^^^^ reference 0.1.test `sg/initial`/Struct#Field.
  
  // unexportedVar is an unexported variable interface.
  var unexportedVar Interface = &Struct{Field: "bar!"}
//    ^^^^^^^^^^^^^ definition 0.1.test `sg/initial`/unexportedVar.
//                  kind Variable
//                  display_name unexportedVar
//                  signature_documentation
//                  > var unexportedVar Interface
//                  documentation
//                  > ```go
//                  > var unexportedVar Interface
//                  > ```
//                  documentation
//                  > unexportedVar is an unexported variable interface.
//                  ^^^^^^^^^ reference 0.1.test `sg/initial`/Interface#
//                               ^^^^^^ reference 0.1.test `sg/initial`/Struct#
//                                      ^^^^^ reference 0.1.test `sg/initial`/Struct#Field.
  
  // x has a builtin error type
  var x error
//    ^ definition 0.1.test `sg/initial`/x.
//      kind Variable
//      display_name x
//      signature_documentation
//      > var x error
//      documentation
//      > ```go
//      > var x error
//      > ```
//      documentation
//      > x has a builtin error type
  
  var BigVar Interface = &Struct{
//    ^^^^^^ definition 0.1.test `sg/initial`/BigVar.
//           kind Variable
//           display_name BigVar
//           signature_documentation
//           > var BigVar Interface
//           documentation
//           > ```go
//           > var BigVar Interface
//           > ```
//           ^^^^^^^^^ reference 0.1.test `sg/initial`/Interface#
//                        ^^^^^^ reference 0.1.test `sg/initial`/Struct#
   Field: "bar!",
// ^^^^^ reference 0.1.test `sg/initial`/Struct#Field.
   Anonymous: struct {
// ^^^^^^^^^ reference 0.1.test `sg/initial`/Struct#Anonymous.
    FieldA int
//  ^^^^^^ definition 0.1.test `sg/initial`/BigVar:FieldA.
//         kind Field
//         display_name FieldA
//         signature_documentation
//         > struct field FieldA int
//         documentation
//         > ```go
//         > struct field FieldA int
//         > ```
    FieldB int
//  ^^^^^^ definition 0.1.test `sg/initial`/BigVar:FieldB.
//         kind Field
//         display_name FieldB
//         signature_documentation
//         > struct field FieldB int
//         documentation
//         > ```go
//         > struct field FieldB int
//         > ```
    FieldC int
//  ^^^^^^ definition 0.1.test `sg/initial`/BigVar:FieldC.
//         kind Field
//         display_name FieldC
//         signature_documentation
//         > struct field FieldC int
//         documentation
//         > ```go
//         > struct field FieldC int
//         > ```
   }{FieldA: 1337},
//   ^^^^^^ reference 0.1.test `sg/initial`/BigVar:FieldA.
  }
  
  // What are docs, really?
  // I can't say for sure, I don't write any.
  // But look, a CAT!
  //
  //       |\      _,,,---,,_
  // ZZZzz /,`.-'`'    -.  ;-;;,_
  //      |,4-  ) )-,_. ,\ (  `'-'
  //     '---''(_/--'  `-'\_)
  //
  // It's sleeping! Some people write that as `sleeping` but Markdown
  // isn't allowed in Go docstrings, right? right?!
  var (
   // This has some docs
   VarBlock1 = "if you're reading this"
// ^^^^^^^^^ definition 0.1.test `sg/initial`/VarBlock1.
//           kind Variable
//           display_name VarBlock1
//           signature_documentation
//           > var VarBlock1 string
//           documentation
//           > ```go
//           > var VarBlock1 string
//           > ```
//           documentation
//           > This has some docs
//           documentation
//           > What are docs, really?
//           > I can't say for sure, I don't write any.
//           > But look, a CAT!
//           > 
//           > 	      |\      _,,,---,,_
//           > 	ZZZzz /,`.-'`'    -.  ;-;;,_
//           > 	     |,4-  ) )-,_. ,\ (  `'-'
//           > 	    '---''(_/--'  `-'\_)
//           > 
//           > It's sleeping! Some people write that as `sleeping` but Markdown
//           > isn't allowed in Go docstrings, right? right?!
  
   VarBlock2 = "hi"
// ^^^^^^^^^ definition 0.1.test `sg/initial`/VarBlock2.
//           kind Variable
//           display_name VarBlock2
//           signature_documentation
//           > var VarBlock2 string
//           documentation
//           > ```go
//           > var VarBlock2 string
//           > ```
//           documentation
//           > What are docs, really?
//           > I can't say for sure, I don't write any.
//           > But look, a CAT!
//           > 
//           > 	      |\      _,,,---,,_
//           > 	ZZZzz /,`.-'`'    -.  ;-;;,_
//           > 	     |,4-  ) )-,_. ,\ (  `'-'
//           > 	    '---''(_/--'  `-'\_)
//           > 
//           > It's sleeping! Some people write that as `sleeping` but Markdown
//           > isn't allowed in Go docstrings, right? right?!
  )
  
  // Embedded is a struct, to be embedded in another struct.
  type Embedded struct {
//     ^^^^^^^^ definition 0.1.test `sg/initial`/Embedded#
//              kind Struct
//              display_name Embedded
//              signature_documentation
//              > type Embedded struct {
//              >     EmbeddedField string
//              >     Field         string
//              > }
//              documentation
//              > ```go
//              > type Embedded struct {
//              >     EmbeddedField string
//              >     Field         string
//              > }
//              > ```
//              documentation
//              > Embedded is a struct, to be embedded in another struct.
   // EmbeddedField has some docs!
   EmbeddedField string
// ^^^^^^^^^^^^^ definition 0.1.test `sg/initial`/Embedded#EmbeddedField.
//               kind Field
//               display_name EmbeddedField
//               signature_documentation
//               > struct field EmbeddedField string
//               documentation
//               > ```go
//               > struct field EmbeddedField string
//               > ```
//               documentation
//               > EmbeddedField has some docs!
   Field         string // conflicts with parent "Field"
// ^^^^^ definition 0.1.test `sg/initial`/Embedded#Field.
//       kind Field
//       display_name Field
//       signature_documentation
//       > struct field Field string
//       documentation
//       > ```go
//       > struct field Field string
//       > ```
//       documentation
//       > conflicts with parent "Field"
  }
  
  type Struct struct {
//     ^^^^^^ definition 0.1.test `sg/initial`/Struct#
//            kind Struct
//            display_name Struct
//            signature_documentation
//            > type Struct struct {
//            >     *Embedded
//            >     Field     string
//            >     Anonymous struct {
//            >         FieldA int
//            >         FieldB int
//            >         FieldC int
//            >     }
//            > }
//            documentation
//            > ```go
//            > type Struct struct {
//            >     *Embedded
//            >     Field     string
//            >     Anonymous struct {
//            >         FieldA int
//            >         FieldB int
//            >         FieldC int
//            >     }
//            > }
//            > ```
//            relationship 0.1.test `sg/initial`/Interface# implementation
   *Embedded
//  ^^^^^^^^ definition 0.1.test `sg/initial`/Struct#Embedded.
//           kind Field
//           display_name Embedded
//           signature_documentation
//           > struct field Embedded *Embedded
//           documentation
//           > ```go
//           > struct field Embedded *Embedded
//           > ```
//  ^^^^^^^^ reference 0.1.test `sg/initial`/Embedded#
   Field     string
// ^^^^^ definition 0.1.test `sg/initial`/Struct#Field.
//       kind Field
//       display_name Field
//       signature_documentation
//       > struct field Field string
//       documentation
//       > ```go
//       > struct field Field string
//       > ```
   Anonymous struct {
// ^^^^^^^^^ definition 0.1.test `sg/initial`/Struct#Anonymous.
//           kind Field
//           display_name Anonymous
//           signature_documentation
//           > struct field Anonymous struct{FieldA int; FieldB int; FieldC int}
//           documentation
//           > ```go
//           > struct field Anonymous struct{FieldA int; FieldB int; FieldC int}
//           > ```
    FieldA int
//  ^^^^^^ definition 0.1.test `sg/initial`/Struct#$anon_81475a76ba757de7#FieldA.
//         kind Field
//         display_name FieldA
//         signature_documentation
//         > struct field FieldA int
//         documentation
//         > ```go
//         > struct field FieldA int
//         > ```
    FieldB int
//  ^^^^^^ definition 0.1.test `sg/initial`/Struct#$anon_81475a76ba757de7#FieldB.
//         kind Field
//         display_name FieldB
//         signature_documentation
//         > struct field FieldB int
//         documentation
//         > ```go
//         > struct field FieldB int
//         > ```
    FieldC int
//  ^^^^^^ definition 0.1.test `sg/initial`/Struct#$anon_81475a76ba757de7#FieldC.
//         kind Field
//         display_name FieldC
//         signature_documentation
//         > struct field FieldC int
//         documentation
//         > ```go
//         > struct field FieldC int
//         > ```
   }
  }
  
  // StructMethod has some docs!
//⌄ enclosing_range_start 0.1.test `sg/initial`/Struct#StructMethod().
  func (s *Struct) StructMethod() {}
//      ^ definition local 0
//        kind Variable
//        display_name s
//        signature_documentation
//        > var s *Struct
//         ^^^^^^ reference 0.1.test `sg/initial`/Struct#
//                 ^^^^^^^^^^^^ definition 0.1.test `sg/initial`/Struct#StructMethod().
//                              kind Method
//                              display_name StructMethod
//                              signature_documentation
//                              > func (*Struct).StructMethod()
//                              documentation
//                              > ```go
//                              > func (*Struct).StructMethod()
//                              > ```
//                              documentation
//                              > StructMethod has some docs!
//                                 ⌃ enclosing_range_end 0.1.test `sg/initial`/Struct#StructMethod().
  
//⌄ enclosing_range_start 0.1.test `sg/initial`/Struct#ImplementsInterface().
  func (s *Struct) ImplementsInterface() string { return "hi!" }
//      ^ definition local 1
//        kind Variable
//        display_name s
//        signature_documentation
//        > var s *Struct
//         ^^^^^^ reference 0.1.test `sg/initial`/Struct#
//                 ^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/initial`/Struct#ImplementsInterface().
//                                     kind Method
//                                     display_name ImplementsInterface
//                                     signature_documentation
//                                     > func (*Struct).ImplementsInterface() string
//                                     documentation
//                                     > ```go
//                                     > func (*Struct).ImplementsInterface() string
//                                     > ```
//                                     relationship 0.1.test `sg/initial`/Interface#ImplementsInterface. implementation
//                                                             ⌃ enclosing_range_end 0.1.test `sg/initial`/Struct#ImplementsInterface().
  
//⌄ enclosing_range_start 0.1.test `sg/initial`/Struct#MachineLearning().
  func (s *Struct) MachineLearning(
//      ^ definition local 2
//        kind Variable
//        display_name s
//        signature_documentation
//        > var s *Struct
//         ^^^^^^ reference 0.1.test `sg/initial`/Struct#
//                 ^^^^^^^^^^^^^^^ definition 0.1.test `sg/initial`/Struct#MachineLearning().
//                                 kind Method
//                                 display_name MachineLearning
//                                 signature_documentation
//                                 > func (*Struct).MachineLearning(param1 float32, hyperparam2 float32, hyperparam3 float32) float32
//                                 documentation
//                                 > ```go
//                                 > func (*Struct).MachineLearning(param1 float32, hyperparam2 float32, hyperparam3 float32) float32
//                                 > ```
   param1 float32, // It's ML, I can't describe what this param is.
// ^^^^^^ definition local 3
//        kind Variable
//        display_name param1
//        signature_documentation
//        > var param1 float32
  
   // We call the below hyperparameters because, uhh, well:
   //
   //    ,-.       _,---._ __  / \
   //   /  )    .-'       `./ /   \
   //   (  (   ,'            `/    /|
   //    \  `-"             \'\   / |
   //     `.              ,  \ \ /  |
   //   /`.          ,'-`----Y   |
   //     (            ;        |   '
   //     |  ,-.    ,-'         |  /
   //     |  | (   |        hjw | /
   //     )  |  \  `.___________|/
   //     `--'   `--'
   //
   hyperparam2 float32,
// ^^^^^^^^^^^ definition local 4
//             kind Variable
//             display_name hyperparam2
//             signature_documentation
//             > var hyperparam2 float32
   hyperparam3 float32,
// ^^^^^^^^^^^ definition local 5
//             kind Variable
//             display_name hyperparam3
//             signature_documentation
//             > var hyperparam3 float32
  ) float32 {
   // varShouldNotHaveDocs is in a function, should not have docs emitted.
   var varShouldNotHaveDocs int32
//     ^^^^^^^^^^^^^^^^^^^^ definition local 6
//                          kind Variable
//                          display_name varShouldNotHaveDocs
//                          signature_documentation
//                          > var varShouldNotHaveDocs int32
  
   // constShouldNotHaveDocs is in a function, should not have docs emitted.
   const constShouldNotHaveDocs = 5
//       ^^^^^^^^^^^^^^^^^^^^^^ definition local 7
//                              kind Constant
//                              display_name constShouldNotHaveDocs
//                              signature_documentation
//                              > const constShouldNotHaveDocs untyped int
  
   // typeShouldNotHaveDocs is in a function, should not have docs emitted.
   type typeShouldNotHaveDocs struct{ a string }
//      ^^^^^^^^^^^^^^^^^^^^^ definition local 8
//                            kind Struct
//                            display_name typeShouldNotHaveDocs
//                            signature_documentation
//                            > type typeShouldNotHaveDocs struct{a string}
//                                    ^ definition local 9
//                                      kind Field
//                                      display_name a
//                                      signature_documentation
//                                      > field a string
  
   // funcShouldNotHaveDocs is in a function, should not have docs emitted.
   funcShouldNotHaveDocs := func(a string) string { return "hello" }
// ^^^^^^^^^^^^^^^^^^^^^ definition local 10
//                       kind Variable
//                       display_name funcShouldNotHaveDocs
//                       signature_documentation
//                       > var funcShouldNotHaveDocs func(a string) string
//                               ^ definition local 11
//                                 kind Variable
//                                 display_name a
//                                 signature_documentation
//                                 > var a string
  
   return param1 + (hyperparam2 * *hyperparam3) // lol is this all ML is? I'm gonna be rich
//        ^^^^^^ reference local 3
//                  ^^^^^^^^^^^ reference local 4
//                                 ^^^^^^^^^^^ reference local 5
  }
//⌃ enclosing_range_end 0.1.test `sg/initial`/Struct#MachineLearning().
  
  // Interface has docs too
  type Interface interface {
//     ^^^^^^^^^ definition 0.1.test `sg/initial`/Interface#
//               kind Interface
//               display_name Interface
//               signature_documentation
//               > type Interface interface{ ImplementsInterface() string }
//               documentation
//               > ```go
//               > type Interface interface{ ImplementsInterface() string }
//               > ```
//               documentation
//               > Interface has docs too
   ImplementsInterface() string
// ^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/initial`/Interface#ImplementsInterface.
//                     kind MethodSpecification
//                     display_name ImplementsInterface
//                     signature_documentation
//                     > func (Interface).ImplementsInterface() string
//                     documentation
//                     > ```go
//                     > func (Interface).ImplementsInterface() string
//                     > ```
  }
  
//⌄ enclosing_range_start 0.1.test `sg/initial`/NewInterface().
  func NewInterface() Interface { return nil }
//     ^^^^^^^^^^^^ definition 0.1.test `sg/initial`/NewInterface().
//                  kind Function
//                  display_name NewInterface
//                  signature_documentation
//                  > func NewInterface() Interface
//                  documentation
//                  > ```go
//                  > func NewInterface() Interface
//                  > ```
//                    ^^^^^^^^^ reference 0.1.test `sg/initial`/Interface#
//                                           ⌃ enclosing_range_end 0.1.test `sg/initial`/NewInterface().
  
  var SortExportedFirst = 1
//    ^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/initial`/SortExportedFirst.
//                      kind Variable
//                      display_name SortExportedFirst
//                      signature_documentation
//                      > var SortExportedFirst int
//                      documentation
//                      > ```go
//                      > var SortExportedFirst int
//                      > ```
  
  var sortUnexportedSecond = 2
//    ^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/initial`/sortUnexportedSecond.
//                         kind Variable
//                         display_name sortUnexportedSecond
//                         signature_documentation
//                         > var sortUnexportedSecond int
//                         documentation
//                         > ```go
//                         > var sortUnexportedSecond int
//                         > ```
  
  var _sortUnderscoreLast = 3
//    ^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/initial`/_sortUnderscoreLast.
//                        kind Variable
//                        display_name _sortUnderscoreLast
//                        signature_documentation
//                        > var _sortUnderscoreLast int
//                        documentation
//                        > ```go
//                        > var _sortUnderscoreLast int
//                        > ```
  
  // Yeah this is some Go magic incantation which is common.
  //
  //  ,_     _
  //  |\\_,-~/
  //  / _  _ |    ,--.
  // (  @  @ )   / ,-'
  //  \  _T_/-._( (
  // /         `. \
  // |         _  \ |
  // \ \ ,  /      |
  //  || |-_\__   /
  // ((_/`(____,-'
  var _ = Interface(&Struct{})
//        ^^^^^^^^^ reference 0.1.test `sg/initial`/Interface#
//                   ^^^^^^ reference 0.1.test `sg/initial`/Struct#
  
  type _ = struct{}
  
  // crypto/tls/common_string.go uses this pattern..
  func _() {
  }
  
  // Go can be fun
  type (
   // And confusing
   X struct {
// ^ definition 0.1.test `sg/initial`/X#
//   kind Struct
//   display_name X
//   signature_documentation
//   > type X struct{ bar string }
//   documentation
//   > ```go
//   > type X struct{ bar string }
//   > ```
//   documentation
//   > And confusing
//   documentation
//   > Go can be fun
    bar string
//  ^^^ definition 0.1.test `sg/initial`/X#bar.
//      kind Field
//      display_name bar
//      signature_documentation
//      > struct field bar string
//      documentation
//      > ```go
//      > struct field bar string
//      > ```
   }
  
   Y struct {
// ^ definition 0.1.test `sg/initial`/Y#
//   kind Struct
//   display_name Y
//   signature_documentation
//   > type Y struct{ baz float64 }
//   documentation
//   > ```go
//   > type Y struct{ baz float64 }
//   > ```
//   documentation
//   > Go can be fun
    baz float64
//  ^^^ definition 0.1.test `sg/initial`/Y#baz.
//      kind Field
//      display_name baz
//      signature_documentation
//      > struct field baz float64
//      documentation
//      > ```go
//      > struct field baz float64
//      > ```
   }
  )
  
