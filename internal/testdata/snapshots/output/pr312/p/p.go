  package p
//        ^ definition 0.1.test `sg/pr312/p`/
  
//⌄ enclosing_range_start 0.1.test `sg/pr312/p`/Opts#
  type Opts struct{}
//     ^^^^ definition 0.1.test `sg/pr312/p`/Opts#
//          kind Struct
//          display_name Opts
//          signature_documentation
//          > type Opts struct{}
//                 ⌃ enclosing_range_end 0.1.test `sg/pr312/p`/Opts#
  
//⌄ enclosing_range_start 0.1.test `sg/pr312/p`/Svc#
  type Svc interface {
//     ^^^ definition 0.1.test `sg/pr312/p`/Svc#
//         kind Interface
//         display_name Svc
//         signature_documentation
//         > type Svc interface{ Run(o *Opts) error }
// ⌄ enclosing_range_start 0.1.test `sg/pr312/p`/Svc#Run.
   Run(o *Opts) error
// ^^^ definition 0.1.test `sg/pr312/p`/Svc#Run.
//     kind MethodSpecification
//     display_name Run
//     signature_documentation
//     > func (Svc).Run(o *Opts) error
//     ^ definition local 0
//       kind Variable
//       display_name o
//       signature_documentation
//       > var o *Opts
//        ^^^^ reference 0.1.test `sg/pr312/p`/Opts#
//                  ⌃ enclosing_range_end 0.1.test `sg/pr312/p`/Svc#Run.
  }
//⌃ enclosing_range_end 0.1.test `sg/pr312/p`/Svc#
  
