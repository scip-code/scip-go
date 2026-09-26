  package inlinestruct
//        ^^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/
  
//⌄ enclosing_range_start 0.1.test `sg/inlinestruct`/Params#
  type Params struct{}
//     ^^^^^^ definition 0.1.test `sg/inlinestruct`/Params#
//            kind Struct
//            display_name Params
//            signature_documentation
//            > type Params struct{}
//                   ⌃ enclosing_range_end 0.1.test `sg/inlinestruct`/Params#
//⌄ enclosing_range_start 0.1.test `sg/inlinestruct`/HighlightedCode#
  type HighlightedCode struct{}
//     ^^^^^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/HighlightedCode#
//                     kind Struct
//                     display_name HighlightedCode
//                     signature_documentation
//                     > type HighlightedCode struct{}
//                            ⌃ enclosing_range_end 0.1.test `sg/inlinestruct`/HighlightedCode#
  
//⌄ enclosing_range_start 0.1.test `sg/inlinestruct`/Mocks.
//⌄ enclosing_range_start 0.1.test `sg/inlinestruct`/emptyMocks.
  var Mocks, emptyMocks struct {
//    ^^^^^ definition 0.1.test `sg/inlinestruct`/Mocks.
//          kind Variable
//          display_name Mocks
//          signature_documentation
//          > var Mocks struct{Code func(p Params) (response *HighlightedCode, aborted bool, err error)}
//           ^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/emptyMocks.
//                      kind Variable
//                      display_name emptyMocks
//                      signature_documentation
//                      > var emptyMocks struct{Code func(p Params) (response *HighlightedCode, aborted bool, err error)}
// ⌄ enclosing_range_start 0.1.test `sg/inlinestruct`/inline-6-5:Code.
   Code func(p Params) (response *HighlightedCode, aborted bool, err error)
// ^^^^ definition 0.1.test `sg/inlinestruct`/inline-6-5:Code.
//      kind Field
//      display_name Code
//      signature_documentation
//      > struct field Code func(p Params) (response *HighlightedCode, aborted bool, err error)
//           ^ definition local 0
//             kind Variable
//             display_name p
//             signature_documentation
//             > var p Params
//             ^^^^^^ reference 0.1.test `sg/inlinestruct`/Params#
//                      ^^^^^^^^ definition local 1
//                               kind Variable
//                               display_name response
//                               signature_documentation
//                               > var response *HighlightedCode
//                                ^^^^^^^^^^^^^^^ reference 0.1.test `sg/inlinestruct`/HighlightedCode#
//                                                 ^^^^^^^ definition local 2
//                                                         kind Variable
//                                                         display_name aborted
//                                                         signature_documentation
//                                                         > var aborted bool
//                                                               ^^^ definition local 3
//                                                                   kind Variable
//                                                                   display_name err
//                                                                   signature_documentation
//                                                                   > var err error
//                                                                        ⌃ enclosing_range_end 0.1.test `sg/inlinestruct`/inline-6-5:Code.
  }
//⌃ enclosing_range_end 0.1.test `sg/inlinestruct`/Mocks.
//⌃ enclosing_range_end 0.1.test `sg/inlinestruct`/emptyMocks.
  
//⌄ enclosing_range_start 0.1.test `sg/inlinestruct`/MocksSingle.
  var MocksSingle struct {
//    ^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/MocksSingle.
//                kind Variable
//                display_name MocksSingle
//                signature_documentation
//                > var MocksSingle struct{Code func(p Params) (response *HighlightedCode, aborted bool, err error)}
// ⌄ enclosing_range_start 0.1.test `sg/inlinestruct`/MocksSingle:Code.
   Code func(p Params) (response *HighlightedCode, aborted bool, err error)
// ^^^^ definition 0.1.test `sg/inlinestruct`/MocksSingle:Code.
//      kind Field
//      display_name Code
//      signature_documentation
//      > struct field Code func(p Params) (response *HighlightedCode, aborted bool, err error)
//           ^ definition local 4
//             kind Variable
//             display_name p
//             signature_documentation
//             > var p Params
//             ^^^^^^ reference 0.1.test `sg/inlinestruct`/Params#
//                      ^^^^^^^^ definition local 5
//                               kind Variable
//                               display_name response
//                               signature_documentation
//                               > var response *HighlightedCode
//                                ^^^^^^^^^^^^^^^ reference 0.1.test `sg/inlinestruct`/HighlightedCode#
//                                                 ^^^^^^^ definition local 6
//                                                         kind Variable
//                                                         display_name aborted
//                                                         signature_documentation
//                                                         > var aborted bool
//                                                               ^^^ definition local 7
//                                                                   kind Variable
//                                                                   display_name err
//                                                                   signature_documentation
//                                                                   > var err error
//                                                                        ⌃ enclosing_range_end 0.1.test `sg/inlinestruct`/MocksSingle:Code.
  }
//⌃ enclosing_range_end 0.1.test `sg/inlinestruct`/MocksSingle.
  
  var (
// ⌄ enclosing_range_start 0.1.test `sg/inlinestruct`/okReply.
   okReply   interface{} = "OK"
// ^^^^^^^ definition 0.1.test `sg/inlinestruct`/okReply.
//         kind Variable
//         display_name okReply
//         signature_documentation
//         > var okReply interface{}
//                            ⌃ enclosing_range_end 0.1.test `sg/inlinestruct`/okReply.
// ⌄ enclosing_range_start 0.1.test `sg/inlinestruct`/pongReply.
   pongReply interface{} = "PONG"
// ^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/pongReply.
//           kind Variable
//           display_name pongReply
//           signature_documentation
//           > var pongReply interface{}
//                              ⌃ enclosing_range_end 0.1.test `sg/inlinestruct`/pongReply.
  )
  
