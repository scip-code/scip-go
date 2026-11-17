  package q
//        ^ definition 0.1.test `sg/pr312/q`/
//          kind Package
//          display_name q
//          signature_documentation
//          > package q
  
  import "sg/pr312/p"
//        ^^^^^^^^^^ reference 0.1.test `sg/pr312/p`/
  
//⌄ enclosing_range_start 0.1.test `sg/pr312/q`/Impl#
  type Impl struct{}
//     ^^^^ definition 0.1.test `sg/pr312/q`/Impl#
//          kind Struct
//          display_name Impl
//          signature_documentation
//          > type Impl struct{}
//          relationship 0.1.test `sg/pr312/p`/Named# implementation
//                 ⌃ enclosing_range_end 0.1.test `sg/pr312/q`/Impl#
  
//⌄ enclosing_range_start 0.1.test `sg/pr312/q`/Impl#Run().
  func (Impl) Run(o *p.Opts) error { return nil }
//      ^^^^ reference 0.1.test `sg/pr312/q`/Impl#
//            ^^^ definition 0.1.test `sg/pr312/q`/Impl#Run().
//                kind Method
//                display_name Run
//                signature_documentation
//                > func (Impl).Run(o *p.Opts) error
//                ^ definition local 0
//                  kind Variable
//                  display_name o
//                  signature_documentation
//                  > var o *Opts
//                   ^ reference 0.1.test `sg/pr312/p`/
//                     ^^^^ reference 0.1.test `sg/pr312/p`/Opts#
//                                              ⌃ enclosing_range_end 0.1.test `sg/pr312/q`/Impl#Run().
  
//⌄ enclosing_range_start 0.1.test `sg/pr312/q`/Impl#Nickname().
  func (Impl) Nickname() string { return "" }
//      ^^^^ reference 0.1.test `sg/pr312/q`/Impl#
//            ^^^^^^^^ definition 0.1.test `sg/pr312/q`/Impl#Nickname().
//                     kind Method
//                     display_name Nickname
//                     signature_documentation
//                     > func (Impl).Nickname() string
//                     relationship 0.1.test `sg/pr312/p`/Named#Nickname. implementation
//                                          ⌃ enclosing_range_end 0.1.test `sg/pr312/q`/Impl#Nickname().
  
