  package generallyeric
//        ^^^^^^^^^^^^^ definition 0.1.test `sg/generallyeric`/
  
  type Number interface {
//     ^^^^^^ definition 0.1.test `sg/generallyeric`/Number#
//            kind Interface
//            display_name Number
//            signature_documentation
//            > type Number interface {
//            >     ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
//            > }
//            documentation
//            > ```go
//            > type Number interface {
//            >     ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
//            > }
//            > ```
   ~int | ~int8 | ~int16 | ~int32 | ~int64 |
    ~float32 | ~float64
  }
  
//⌄ enclosing_range_start 0.1.test `sg/generallyeric`/Double().
  func Double[T Number](value T) T {
//     ^^^^^^ definition 0.1.test `sg/generallyeric`/Double().
//            kind Function
//            display_name Double
//            signature_documentation
//            > func Double[T Number](value T) T
//            documentation
//            > ```go
//            > func Double[T Number](value T) T
//            > ```
//            ^ definition local 0
//              kind Interface
//              display_name T
//              signature_documentation
//              > type parameter T Number
//              ^^^^^^ reference 0.1.test `sg/generallyeric`/Number#
//                      ^^^^^ definition local 1
//                            kind Variable
//                            display_name value
//                            signature_documentation
//                            > var value T
//                            ^ reference local 0
//                               ^ reference local 0
   return value * 2
//        ^^^^^ reference local 1
  }
//⌃ enclosing_range_end 0.1.test `sg/generallyeric`/Double().
  
  type Box[T any] struct {
//     ^^^ definition 0.1.test `sg/generallyeric`/Box#
//         kind Struct
//         display_name Box
//         signature_documentation
//         > type Box struct{ Something T }
//         documentation
//         > ```go
//         > type Box struct{ Something T }
//         > ```
//         ^ definition local 2
//           kind Interface
//           display_name T
//           signature_documentation
//           > type parameter T any
   Something T
// ^^^^^^^^^ definition 0.1.test `sg/generallyeric`/Box#Something.
//           kind Field
//           display_name Something
//           signature_documentation
//           > struct field Something T
//           documentation
//           > ```go
//           > struct field Something T
//           > ```
//           ^ reference local 2
  }
  
  type handler[T any] struct {
//     ^^^^^^^ definition 0.1.test `sg/generallyeric`/handler#
//             kind Struct
//             display_name handler
//             signature_documentation
//             > type handler struct {
//             >     Box[T]
//             >     Another string
//             > }
//             documentation
//             > ```go
//             > type handler struct {
//             >     Box[T]
//             >     Another string
//             > }
//             > ```
//             ^ definition local 3
//               kind Interface
//               display_name T
//               signature_documentation
//               > type parameter T any
   Box[T]
// ^^^ definition 0.1.test `sg/generallyeric`/handler#Box.
//     kind Field
//     display_name Box
//     signature_documentation
//     > struct field Box Box[T]
//     documentation
//     > ```go
//     > struct field Box Box[T]
//     > ```
// ^^^ reference 0.1.test `sg/generallyeric`/Box#
//     ^ reference local 3
   Another string
// ^^^^^^^ definition 0.1.test `sg/generallyeric`/handler#Another.
//         kind Field
//         display_name Another
//         signature_documentation
//         > struct field Another string
//         documentation
//         > ```go
//         > struct field Another string
//         > ```
  }
  
