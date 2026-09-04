  package testdata
//        ^^^^^^^^ definition 0.1.test `sg/testdata`/
//                 kind Package
//                 display_name testdata
//                 signature_documentation
//                 > package testdata
//                 documentation
//                 > ```go
//                 > package testdata
//                 > ```
  
  import "fmt"
//        ^^^ reference github.com/golang/go/src go1.22 fmt/
  
  type TypeContainingAnonymousStructs struct {
//     ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#
//                                    kind Struct
//                                    display_name TypeContainingAnonymousStructs
//                                    signature_documentation
//                                    > type TypeContainingAnonymousStructs struct {
//                                    >     a struct {
//                                    >         x int
//                                    >         y string
//                                    >     }
//                                    >     b struct {
//                                    >         x int
//                                    >         y string
//                                    >     }
//                                    >     c struct {
//                                    >         X int
//                                    >         Y string
//                                    >     }
//                                    > }
//                                    documentation
//                                    > ```go
//                                    > type TypeContainingAnonymousStructs struct {
//                                    >     a struct {
//                                    >         x int
//                                    >         y string
//                                    >     }
//                                    >     b struct {
//                                    >         x int
//                                    >         y string
//                                    >     }
//                                    >     c struct {
//                                    >         X int
//                                    >         Y string
//                                    >     }
//                                    > }
//                                    > ```
   a, b struct {
// ^ definition 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#a.
//   kind Field
//   display_name a
//   signature_documentation
//   > struct field a struct{x int; y string}
//   documentation
//   > ```go
//   > struct field a struct{x int; y string}
//   > ```
//    ^ definition 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#b.
//      kind Field
//      display_name b
//      signature_documentation
//      > struct field b struct{x int; y string}
//      documentation
//      > ```go
//      > struct field b struct{x int; y string}
//      > ```
    x int
//  ^ definition 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#$anon_c0a8952b3a214f68#x.
//    kind Field
//    display_name x
//    signature_documentation
//    > struct field x int
//    documentation
//    > ```go
//    > struct field x int
//    > ```
    y string
//  ^ definition 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#$anon_c0a8952b3a214f68#y.
//    kind Field
//    display_name y
//    signature_documentation
//    > struct field y string
//    documentation
//    > ```go
//    > struct field y string
//    > ```
   }
   c struct {
// ^ definition 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#c.
//   kind Field
//   display_name c
//   signature_documentation
//   > struct field c struct{X int; Y string}
//   documentation
//   > ```go
//   > struct field c struct{X int; Y string}
//   > ```
    X int
//  ^ definition 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#$anon_2f238678626c0da1#X.
//    kind Field
//    display_name X
//    signature_documentation
//    > struct field X int
//    documentation
//    > ```go
//    > struct field X int
//    > ```
    Y string
//  ^ definition 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#$anon_2f238678626c0da1#Y.
//    kind Field
//    display_name Y
//    signature_documentation
//    > struct field Y string
//    documentation
//    > ```go
//    > struct field Y string
//    > ```
   }
  }
  
//⌄ enclosing_range_start 0.1.test `sg/testdata`/funcContainingAnonymousStructs().
  func funcContainingAnonymousStructs() {
//     ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/funcContainingAnonymousStructs().
//                                    kind Function
//                                    display_name funcContainingAnonymousStructs
//                                    signature_documentation
//                                    > func funcContainingAnonymousStructs()
//                                    documentation
//                                    > ```go
//                                    > func funcContainingAnonymousStructs()
//                                    > ```
   d := struct {
// ^ definition local 0
//   kind Variable
//   display_name d
//   signature_documentation
//   > var d struct{x int; y string}
    x int
//  ^ definition local 1
//    kind Field
//    display_name x
//    signature_documentation
//    > field x int
    y string
//  ^ definition local 2
//    kind Field
//    display_name y
//    signature_documentation
//    > field y string
   }{
    x: 1,
//  ^ reference local 1
    y: "one",
//  ^ reference local 2
   }
  
   var e struct {
//     ^ definition local 3
//       kind Variable
//       display_name e
//       signature_documentation
//       > var e struct{x int; y string}
    x int
//  ^ definition local 4
//    kind Field
//    display_name x
//    signature_documentation
//    > field x int
    y string
//  ^ definition local 5
//    kind Field
//    display_name y
//    signature_documentation
//    > field y string
   }
  
   e.x = 2
// ^ reference local 3
//   ^ reference local 4
   e.y = "two"
// ^ reference local 3
//   ^ reference local 5
  
   var f TypeContainingAnonymousStructs
//     ^ definition local 6
//       kind Variable
//       display_name f
//       signature_documentation
//       > var f TypeContainingAnonymousStructs
//       ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#
   f.a.x = 3
// ^ reference local 6
//   ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#a.
//     ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#$anon_c0a8952b3a214f68#x.
   f.a.y = "three"
// ^ reference local 6
//   ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#a.
//     ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#$anon_c0a8952b3a214f68#y.
   f.b.x = 4
// ^ reference local 6
//   ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#b.
//     ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#$anon_c0a8952b3a214f68#x.
   f.b.y = "four"
// ^ reference local 6
//   ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#b.
//     ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#$anon_c0a8952b3a214f68#y.
   f.c.X = 5
// ^ reference local 6
//   ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#c.
//     ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#$anon_2f238678626c0da1#X.
   f.c.Y = "five"
// ^ reference local 6
//   ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#c.
//     ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#$anon_2f238678626c0da1#Y.
  
   fmt.Printf("> %s, %s\n", d.x, d.y)
// ^^^ reference github.com/golang/go/src go1.22 fmt/
//     ^^^^^^ reference github.com/golang/go/src go1.22 fmt/Printf().
//                          ^ reference local 0
//                            ^ reference local 1
//                               ^ reference local 0
//                                 ^ reference local 2
   fmt.Printf("> %s, %s\n", e.x, e.y)
// ^^^ reference github.com/golang/go/src go1.22 fmt/
//     ^^^^^^ reference github.com/golang/go/src go1.22 fmt/Printf().
//                          ^ reference local 3
//                            ^ reference local 4
//                               ^ reference local 3
//                                 ^ reference local 5
  
   fmt.Printf("> %s, %s\n", f.a.x, f.a.y)
// ^^^ reference github.com/golang/go/src go1.22 fmt/
//     ^^^^^^ reference github.com/golang/go/src go1.22 fmt/Printf().
//                          ^ reference local 6
//                            ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#a.
//                              ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#$anon_c0a8952b3a214f68#x.
//                                 ^ reference local 6
//                                   ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#a.
//                                     ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#$anon_c0a8952b3a214f68#y.
   fmt.Printf("> %s, %s\n", f.b.x, f.b.y)
// ^^^ reference github.com/golang/go/src go1.22 fmt/
//     ^^^^^^ reference github.com/golang/go/src go1.22 fmt/Printf().
//                          ^ reference local 6
//                            ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#b.
//                              ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#$anon_c0a8952b3a214f68#x.
//                                 ^ reference local 6
//                                   ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#b.
//                                     ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#$anon_c0a8952b3a214f68#y.
   fmt.Printf("> %s, %s\n", f.c.X, f.c.Y)
// ^^^ reference github.com/golang/go/src go1.22 fmt/
//     ^^^^^^ reference github.com/golang/go/src go1.22 fmt/Printf().
//                          ^ reference local 6
//                            ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#c.
//                              ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#$anon_2f238678626c0da1#X.
//                                 ^ reference local 6
//                                   ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#c.
//                                     ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#$anon_2f238678626c0da1#Y.
  }
//⌃ enclosing_range_end 0.1.test `sg/testdata`/funcContainingAnonymousStructs().
  
