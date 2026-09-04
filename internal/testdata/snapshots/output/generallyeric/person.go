  package generallyeric
//        ^^^^^^^^^^^^^ definition 0.1.test `sg/generallyeric`/
  
  import "fmt"
//        ^^^ reference github.com/golang/go/src go1.22 fmt/
  
  type Person interface {
//     ^^^^^^ definition 0.1.test `sg/generallyeric`/Person#
//            kind Interface
//            display_name Person
//            signature_documentation
//            > type Person interface{ Work() }
//            documentation
//            > ```go
//            > type Person interface{ Work() }
//            > ```
   Work()
// ^^^^ definition 0.1.test `sg/generallyeric`/Person#Work.
//      kind MethodSpecification
//      display_name Work
//      signature_documentation
//      > func (Person).Work()
//      documentation
//      > ```go
//      > func (Person).Work()
//      > ```
  }
  
  type worker string
//     ^^^^^^ definition 0.1.test `sg/generallyeric`/worker#
//            kind Type
//            display_name worker
//            signature_documentation
//            > type worker string
//            documentation
//            > ```go
//            > type worker string
//            > ```
//            relationship 0.1.test `sg/generallyeric`/Person# implementation
  
//⌄ enclosing_range_start 0.1.test `sg/generallyeric`/worker#Work().
  func (w worker) Work() {
//      ^ definition local 0
//        kind Variable
//        display_name w
//        signature_documentation
//        > var w worker
//        ^^^^^^ reference 0.1.test `sg/generallyeric`/worker#
//                ^^^^ definition 0.1.test `sg/generallyeric`/worker#Work().
//                     kind Method
//                     display_name Work
//                     signature_documentation
//                     > func (worker).Work()
//                     documentation
//                     > ```go
//                     > func (worker).Work()
//                     > ```
//                     relationship 0.1.test `sg/generallyeric`/Person#Work. implementation
   fmt.Printf("%s is working\n", w)
// ^^^ reference github.com/golang/go/src go1.22 fmt/
//     ^^^^^^ reference github.com/golang/go/src go1.22 fmt/Printf().
//                               ^ reference local 0
  }
//⌃ enclosing_range_end 0.1.test `sg/generallyeric`/worker#Work().
  
//⌄ enclosing_range_start 0.1.test `sg/generallyeric`/DoWork().
  func DoWork[T Person](things []T) {
//     ^^^^^^ definition 0.1.test `sg/generallyeric`/DoWork().
//            kind Function
//            display_name DoWork
//            signature_documentation
//            > func DoWork[T Person](things []T)
//            documentation
//            > ```go
//            > func DoWork[T Person](things []T)
//            > ```
//            ^ definition local 1
//              kind Interface
//              display_name T
//              signature_documentation
//              > type parameter T Person
//              ^^^^^^ reference 0.1.test `sg/generallyeric`/Person#
//                      ^^^^^^ definition local 2
//                             kind Variable
//                             display_name things
//                             signature_documentation
//                             > var things []T
//                               ^ reference local 1
   for _, v := range things {
//        ^ definition local 3
//          kind Variable
//          display_name v
//          signature_documentation
//          > var v T
//                   ^^^^^^ reference local 2
    v.Work()
//  ^ reference local 3
//    ^^^^ reference 0.1.test `sg/generallyeric`/Person#Work.
   }
  }
//⌃ enclosing_range_end 0.1.test `sg/generallyeric`/DoWork().
  
//⌄ enclosing_range_start 0.1.test `sg/generallyeric`/main().
  func main() {
//     ^^^^ definition 0.1.test `sg/generallyeric`/main().
//          kind Function
//          display_name main
//          signature_documentation
//          > func main()
//          documentation
//          > ```go
//          > func main()
//          > ```
   var a, b, c worker
//     ^ definition local 4
//       kind Variable
//       display_name a
//       signature_documentation
//       > var a worker
//        ^ definition local 5
//          kind Variable
//          display_name b
//          signature_documentation
//          > var b worker
//           ^ definition local 6
//             kind Variable
//             display_name c
//             signature_documentation
//             > var c worker
//             ^^^^^^ reference 0.1.test `sg/generallyeric`/worker#
   a = "A"
// ^ reference local 4
   b = "B"
// ^ reference local 5
   c = "C"
// ^ reference local 6
   DoWork([]worker{a, b, c})
// ^^^^^^ reference 0.1.test `sg/generallyeric`/DoWork().
//          ^^^^^^ reference 0.1.test `sg/generallyeric`/worker#
//                 ^ reference local 4
//                    ^ reference local 5
//                       ^ reference local 6
  }
//⌃ enclosing_range_end 0.1.test `sg/generallyeric`/main().
  
