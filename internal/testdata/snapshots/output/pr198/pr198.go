  package pr198
//        ^^^^^ definition 0.1.test `sg/pr198`/
//              kind Package
//              display_name pr198
//              signature_documentation
//              > package pr198
//              documentation
//              > ```go
//              > package pr198
//              > ```
  
  import "github.com/example/dep"
//        ^^^^^^^^^^^^^^^^^^^^^^ reference github.com/example/dep 0.1.test `github.com/example/dep`/
  
  // Foo is an interface defined downstream of the type that implements it.
  // The dep.T type (from a dependency) implements Foo, and scip-go should
  // emit an external symbol for dep.T with an IsImplementation relationship
  // pointing to Foo.
  type Foo interface {
//     ^^^ definition 0.1.test `sg/pr198`/Foo#
//         kind Interface
//         display_name Foo
//         signature_documentation
//         > type Foo interface{ Bar() }
//         documentation
//         > ```go
//         > type Foo interface{ Bar() }
//         > ```
//         documentation
//         > Foo is an interface defined downstream of the type that implements it.
//         > The dep.T type (from a dependency) implements Foo, and scip-go should
//         > emit an external symbol for dep.T with an IsImplementation relationship
//         > pointing to Foo.
   Bar()
// ^^^ definition 0.1.test `sg/pr198`/Foo#Bar.
//     kind MethodSpecification
//     display_name Bar
//     signature_documentation
//     > func (Foo).Bar()
//     documentation
//     > ```go
//     > func (Foo).Bar()
//     > ```
  }
  
//⌄ enclosing_range_start 0.1.test `sg/pr198`/UseFoo().
  func UseFoo(f Foo) {}
//     ^^^^^^ definition 0.1.test `sg/pr198`/UseFoo().
//            kind Function
//            display_name UseFoo
//            signature_documentation
//            > func UseFoo(f Foo)
//            documentation
//            > ```go
//            > func UseFoo(f Foo)
//            > ```
//            ^ definition local 0
//              kind Variable
//              display_name f
//              signature_documentation
//              > var f Foo
//              ^^^ reference 0.1.test `sg/pr198`/Foo#
//                    ⌃ enclosing_range_end 0.1.test `sg/pr198`/UseFoo().
  
//⌄ enclosing_range_start 0.1.test `sg/pr198`/Example().
  func Example() {
//     ^^^^^^^ definition 0.1.test `sg/pr198`/Example().
//             kind Function
//             display_name Example
//             signature_documentation
//             > func Example()
//             documentation
//             > ```go
//             > func Example()
//             > ```
   UseFoo(&dep.T{})
// ^^^^^^ reference 0.1.test `sg/pr198`/UseFoo().
//         ^^^ reference github.com/example/dep 0.1.test `github.com/example/dep`/
//             ^ reference github.com/example/dep 0.1.test `github.com/example/dep`/T#
  }
//⌃ enclosing_range_end 0.1.test `sg/pr198`/Example().
  
