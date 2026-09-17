  package testdata
//        ^^^^^^^^ definition 0.1.test `sg/testdata`/
  
  import (
   "context"
//  ^^^^^^^ reference github.com/golang/go/src go1.22 context/
   "sync"
//  ^^^^ reference github.com/golang/go/src go1.22 sync/
  )
  
  // ParallelizableFunc is a function that can be called concurrently with other instances
  // of this function type.
  type ParallelizableFunc func(ctx context.Context) error
//     ^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/ParallelizableFunc#
//                        kind Type
//                        display_name ParallelizableFunc
//                        signature_documentation
//                        > type ParallelizableFunc func(ctx context.Context) error
//                        documentation
//                        > ```go
//                        > type ParallelizableFunc func(ctx context.Context) error
//                        > ```
//                        documentation
//                        > ParallelizableFunc is a function that can be called concurrently with other instances
//                        > of this function type.
//                             ^^^ definition local 0
//                                 kind Variable
//                                 display_name ctx
//                                 signature_documentation
//                                 > var ctx Context
//                                 ^^^^^^^ reference github.com/golang/go/src go1.22 context/
//                                         ^^^^^^^ reference github.com/golang/go/src go1.22 context/Context#
  
  // Parallel invokes each of the given parallelizable functions in their own goroutines and
  // returns the first error to occur. This method will block until all goroutines have returned.
//⌄ enclosing_range_start 0.1.test `sg/testdata`/Parallel().
  func Parallel(ctx context.Context, fns ...ParallelizableFunc) error {
//     ^^^^^^^^ definition 0.1.test `sg/testdata`/Parallel().
//              kind Function
//              display_name Parallel
//              signature_documentation
//              > func Parallel(ctx context.Context, fns ...ParallelizableFunc) error
//              documentation
//              > ```go
//              > func Parallel(ctx context.Context, fns ...ParallelizableFunc) error
//              > ```
//              documentation
//              > Parallel invokes each of the given parallelizable functions in their own goroutines and
//              > returns the first error to occur. This method will block until all goroutines have returned.
//              ^^^ definition local 1
//                  kind Variable
//                  display_name ctx
//                  signature_documentation
//                  > var ctx Context
//                  ^^^^^^^ reference github.com/golang/go/src go1.22 context/
//                          ^^^^^^^ reference github.com/golang/go/src go1.22 context/Context#
//                                   ^^^ definition local 2
//                                       kind Variable
//                                       display_name fns
//                                       signature_documentation
//                                       > var fns []ParallelizableFunc
//                                          ^^^^^^^^^^^^^^^^^^ reference 0.1.test `sg/testdata`/ParallelizableFunc#
   var wg sync.WaitGroup
//     ^^ definition local 3
//        kind Variable
//        display_name wg
//        signature_documentation
//        > var wg WaitGroup
//        ^^^^ reference github.com/golang/go/src go1.22 sync/
//             ^^^^^^^^^ reference github.com/golang/go/src go1.22 sync/WaitGroup#
   errs := make(chan error, len(fns))
// ^^^^ definition local 4
//      kind Variable
//      display_name errs
//      signature_documentation
//      > var errs chan error
//                              ^^^ reference local 2
  
   for _, fn := range fns {
//        ^^ definition local 5
//           kind Variable
//           display_name fn
//           signature_documentation
//           > var fn ParallelizableFunc
//                    ^^^ reference local 2
    wg.Add(1)
//  ^^ reference local 3
//     ^^^ reference github.com/golang/go/src go1.22 sync/WaitGroup#Add().
  
    go func(fn ParallelizableFunc) {
//          ^^ definition local 6
//             kind Variable
//             display_name fn
//             signature_documentation
//             > var fn ParallelizableFunc
//             ^^^^^^^^^^^^^^^^^^ reference 0.1.test `sg/testdata`/ParallelizableFunc#
     errs <- fn(ctx)
//   ^^^^ reference local 4
//           ^^ reference local 6
//              ^^^ reference local 1
     wg.Done()
//   ^^ reference local 3
//      ^^^^ reference github.com/golang/go/src go1.22 sync/WaitGroup#Done().
    }(fn)
//    ^^ reference local 5
   }
  
   wg.Wait()
// ^^ reference local 3
//    ^^^^ reference github.com/golang/go/src go1.22 sync/WaitGroup#Wait().
  
   for err := range errs {
//     ^^^ definition local 7
//         kind Variable
//         display_name err
//         signature_documentation
//         > var err error
//                  ^^^^ reference local 4
    if err != nil {
//     ^^^ reference local 7
     return err
//          ^^^ reference local 7
    }
   }
  
   return nil
  }
//⌃ enclosing_range_end 0.1.test `sg/testdata`/Parallel().
  
