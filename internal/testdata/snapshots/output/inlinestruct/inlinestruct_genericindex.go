  package inlinestruct
//        ^^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/
  
//⌄ enclosing_range_start 0.1.test `sg/inlinestruct`/Processor#
  type Processor[T any] interface {
//     ^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/Processor#
//               kind Interface
//               display_name Processor
//               signature_documentation
//               > type Processor interface {
//               >     Process(payload T)
//               >     ProcessorType() string
//               > }
//               ^ definition local 0
//                 kind Interface
//                 display_name T
//                 signature_documentation
//                 > type parameter T any
// ⌄ enclosing_range_start 0.1.test `sg/inlinestruct`/Processor#Process.
   Process(payload T)
// ^^^^^^^ definition 0.1.test `sg/inlinestruct`/Processor#Process.
//         kind MethodSpecification
//         display_name Process
//         signature_documentation
//         > func (Processor[T any]).Process(payload T)
//         ^^^^^^^ definition local 1
//                 kind Variable
//                 display_name payload
//                 signature_documentation
//                 > var payload T
//                 ^ reference local 0
//                  ⌃ enclosing_range_end 0.1.test `sg/inlinestruct`/Processor#Process.
// ⌄ enclosing_range_start 0.1.test `sg/inlinestruct`/Processor#ProcessorType.
   ProcessorType() string
// ^^^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/Processor#ProcessorType.
//               kind MethodSpecification
//               display_name ProcessorType
//               signature_documentation
//               > func (Processor[T any]).ProcessorType() string
//                      ⌃ enclosing_range_end 0.1.test `sg/inlinestruct`/Processor#ProcessorType.
  }
//⌃ enclosing_range_end 0.1.test `sg/inlinestruct`/Processor#
  
//⌄ enclosing_range_start 0.1.test `sg/inlinestruct`/Limit#
  type Limit int
//     ^^^^^ definition 0.1.test `sg/inlinestruct`/Limit#
//           kind Type
//           display_name Limit
//           signature_documentation
//           > type Limit int
//             ⌃ enclosing_range_end 0.1.test `sg/inlinestruct`/Limit#
  
//⌄ enclosing_range_start 0.1.test `sg/inlinestruct`/ProcessImpl#
  type ProcessImpl struct{}
//     ^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/ProcessImpl#
//                 kind Struct
//                 display_name ProcessImpl
//                 signature_documentation
//                 > type ProcessImpl struct{}
//                        ⌃ enclosing_range_end 0.1.test `sg/inlinestruct`/ProcessImpl#
  
//⌄ enclosing_range_start 0.1.test `sg/inlinestruct`/ProcessImpl#Process().
  func (p *ProcessImpl) Process(payload Limit) { panic("not implemented") }
//      ^ definition local 2
//        kind Variable
//        display_name p
//        signature_documentation
//        > var p *ProcessImpl
//         ^^^^^^^^^^^ reference 0.1.test `sg/inlinestruct`/ProcessImpl#
//                      ^^^^^^^ definition 0.1.test `sg/inlinestruct`/ProcessImpl#Process().
//                              kind Method
//                              display_name Process
//                              signature_documentation
//                              > func (*ProcessImpl).Process(payload Limit)
//                              ^^^^^^^ definition local 3
//                                      kind Variable
//                                      display_name payload
//                                      signature_documentation
//                                      > var payload Limit
//                                      ^^^^^ reference 0.1.test `sg/inlinestruct`/Limit#
//                                                                        ⌃ enclosing_range_end 0.1.test `sg/inlinestruct`/ProcessImpl#Process().
//⌄ enclosing_range_start 0.1.test `sg/inlinestruct`/ProcessImpl#ProcessorType().
  func (p *ProcessImpl) ProcessorType() string { panic("not implemented") }
//      ^ definition local 4
//        kind Variable
//        display_name p
//        signature_documentation
//        > var p *ProcessImpl
//         ^^^^^^^^^^^ reference 0.1.test `sg/inlinestruct`/ProcessImpl#
//                      ^^^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/ProcessImpl#ProcessorType().
//                                    kind Method
//                                    display_name ProcessorType
//                                    signature_documentation
//                                    > func (*ProcessImpl).ProcessorType() string
//                                                                        ⌃ enclosing_range_end 0.1.test `sg/inlinestruct`/ProcessImpl#ProcessorType().
  
  var _ Processor[Limit] = &ProcessImpl{}
//      ^^^^^^^^^ reference 0.1.test `sg/inlinestruct`/Processor#
//                ^^^^^ reference 0.1.test `sg/inlinestruct`/Limit#
//                          ^^^^^^^^^^^ reference 0.1.test `sg/inlinestruct`/ProcessImpl#
  
