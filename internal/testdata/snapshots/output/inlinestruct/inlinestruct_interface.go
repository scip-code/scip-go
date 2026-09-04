  package inlinestruct
//        ^^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/
  
  import "context"
//        ^^^^^^^ reference github.com/golang/go/src go1.22 context/
  
//⌄ enclosing_range_start 0.1.test `sg/inlinestruct`/Target().
  func Target() interface {
//     ^^^^^^ definition 0.1.test `sg/inlinestruct`/Target().
//            kind Function
//            display_name Target
//            signature_documentation
//            > func Target() interface{AbbreviatedOID(context.Context) (string, error); Commit(context.Context) (string, error); OID(context.Context) (int, error); Type(context.Context) (int, error)}
//            documentation
//            > ```go
//            > func Target() interface{AbbreviatedOID(context.Context) (string, error); Commit(context.Context) (string, error); OID(context.Context) (int, error); Type(context.Context) (int, error)}
//            > ```
   OID(context.Context) (int, error)
// ^^^ definition 0.1.test `sg/inlinestruct`/func:Target:OID().
//     kind MethodSpecification
//     display_name OID
//     signature_documentation
//     > func (interface).OID(context.Context) (int, error)
//     documentation
//     > ```go
//     > func (interface).OID(context.Context) (int, error)
//     > ```
//     ^^^^^^^ reference github.com/golang/go/src go1.22 context/
//             ^^^^^^^ reference github.com/golang/go/src go1.22 context/Context#
   AbbreviatedOID(context.Context) (string, error)
// ^^^^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/func:Target:AbbreviatedOID().
//                kind MethodSpecification
//                display_name AbbreviatedOID
//                signature_documentation
//                > func (interface).AbbreviatedOID(context.Context) (string, error)
//                documentation
//                > ```go
//                > func (interface).AbbreviatedOID(context.Context) (string, error)
//                > ```
//                ^^^^^^^ reference github.com/golang/go/src go1.22 context/
//                        ^^^^^^^ reference github.com/golang/go/src go1.22 context/Context#
   Commit(context.Context) (string, error)
// ^^^^^^ definition 0.1.test `sg/inlinestruct`/func:Target:Commit().
//        kind MethodSpecification
//        display_name Commit
//        signature_documentation
//        > func (interface).Commit(context.Context) (string, error)
//        documentation
//        > ```go
//        > func (interface).Commit(context.Context) (string, error)
//        > ```
//        ^^^^^^^ reference github.com/golang/go/src go1.22 context/
//                ^^^^^^^ reference github.com/golang/go/src go1.22 context/Context#
   Type(context.Context) (int, error)
// ^^^^ definition 0.1.test `sg/inlinestruct`/func:Target:Type().
//      kind MethodSpecification
//      display_name Type
//      signature_documentation
//      > func (interface).Type(context.Context) (int, error)
//      documentation
//      > ```go
//      > func (interface).Type(context.Context) (int, error)
//      > ```
//      ^^^^^^^ reference github.com/golang/go/src go1.22 context/
//              ^^^^^^^ reference github.com/golang/go/src go1.22 context/Context#
  } {
   panic("not implemented")
  }
//⌃ enclosing_range_end 0.1.test `sg/inlinestruct`/Target().
  
//⌄ enclosing_range_start 0.1.test `sg/inlinestruct`/something().
  func something() {
//     ^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/something().
//               kind Function
//               display_name something
//               signature_documentation
//               > func something()
//               documentation
//               > ```go
//               > func something()
//               > ```
   x := Target()
// ^ definition local 0
//   kind Variable
//   display_name x
//   signature_documentation
//   > var x interface{AbbreviatedOID(Context) (string, error); Commit(Context) (string, error); OID(Context) (int, error); Type(Context) (int, error)}
//      ^^^^^^ reference 0.1.test `sg/inlinestruct`/Target().
   x.OID(context.Background())
// ^ reference local 0
//   ^^^ reference 0.1.test `sg/inlinestruct`/func:Target:OID().
//       ^^^^^^^ reference github.com/golang/go/src go1.22 context/
//               ^^^^^^^^^^ reference github.com/golang/go/src go1.22 context/Background().
  }
//⌃ enclosing_range_end 0.1.test `sg/inlinestruct`/something().
  
