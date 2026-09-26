  package pr206
//        ^^^^^ definition 0.1.test `sg/pr206`/
  
//⌄ enclosing_range_start 0.1.test `sg/pr206`/Doer#
  type Doer interface {
//     ^^^^ definition 0.1.test `sg/pr206`/Doer#
//          kind Interface
//          display_name Doer
//          signature_documentation
//          > type Doer interface{ Do() error }
// ⌄ enclosing_range_start 0.1.test `sg/pr206`/Doer#Do.
   // Do performs the action and returns an error if it fails.
   Do() error
// ^^ definition 0.1.test `sg/pr206`/Doer#Do.
//    kind MethodSpecification
//    display_name Do
//    signature_documentation
//    > func (Doer).Do() error
//    documentation
//    > Do performs the action and returns an error if it fails.
//          ⌃ enclosing_range_end 0.1.test `sg/pr206`/Doer#Do.
  }
//⌃ enclosing_range_end 0.1.test `sg/pr206`/Doer#
  
