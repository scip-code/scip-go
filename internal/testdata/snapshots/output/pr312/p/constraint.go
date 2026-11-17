  package p
//        ^ definition 0.1.test `sg/pr312/p`/
//          kind Package
//          display_name p
//          signature_documentation
//          > package p
  
//⌄ enclosing_range_start 0.1.test `sg/pr312/p`/Named#
  type Named interface {
//     ^^^^^ definition 0.1.test `sg/pr312/p`/Named#
//           kind Interface
//           display_name Named
//           signature_documentation
//           > type Named interface{ Nickname() string }
// ⌄ enclosing_range_start 0.1.test `sg/pr312/p`/Named#Nickname.
   Nickname() string
// ^^^^^^^^ definition 0.1.test `sg/pr312/p`/Named#Nickname.
//          kind MethodSpecification
//          display_name Nickname
//          signature_documentation
//          > func (Named).Nickname() string
//                 ⌃ enclosing_range_end 0.1.test `sg/pr312/p`/Named#Nickname.
  }
//⌃ enclosing_range_end 0.1.test `sg/pr312/p`/Named#
  
//⌄ enclosing_range_start 0.1.test `sg/pr312/p`/NamedPtr#
  // NamedPtr is a constraint: only *T for some T with a Nickname method satisfies it.
  // No named type implements it, so it must not get implementation edges.
  type NamedPtr[T any] interface {
//     ^^^^^^^^ definition 0.1.test `sg/pr312/p`/NamedPtr#
//              kind Interface
//              display_name NamedPtr
//              signature_documentation
//              > type NamedPtr interface {
//              >     Named
//              >     *T
//              > }
//              documentation
//              > NamedPtr is a constraint: only *T for some T with a Nickname method satisfies it.
//              > No named type implements it, so it must not get implementation edges.
//              ^ definition local 0
//                kind Interface
//                display_name T
//                signature_documentation
//                > type parameter T any
   Named
// ^^^^^ reference 0.1.test `sg/pr312/p`/Named#
   *T
//  ^ reference local 0
  }
//⌃ enclosing_range_end 0.1.test `sg/pr312/p`/NamedPtr#
  
