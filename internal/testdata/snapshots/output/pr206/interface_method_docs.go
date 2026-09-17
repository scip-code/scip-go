  package pr206
//        ^^^^^ definition 0.1.test `sg/pr206`/
  
  type Doer interface {
//     ^^^^ definition 0.1.test `sg/pr206`/Doer#
//          kind Interface
//          display_name Doer
//          signature_documentation
//          > type Doer interface{ Do() error }
//          documentation
//          > ```go
//          > type Doer interface{ Do() error }
//          > ```
   // Do performs the action and returns an error if it fails.
   Do() error
// ^^ definition 0.1.test `sg/pr206`/Doer#Do.
//    kind MethodSpecification
//    display_name Do
//    signature_documentation
//    > func (Doer).Do() error
//    documentation
//    > ```go
//    > func (Doer).Do() error
//    > ```
//    documentation
//    > Do performs the action and returns an error if it fails.
  }
  
