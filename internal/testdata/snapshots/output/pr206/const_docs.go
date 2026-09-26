  package pr206
//        ^^^^^ definition 0.1.test `sg/pr206`/
//              kind Package
//              display_name pr206
//              signature_documentation
//              > package pr206
  
  // Block doc for the const group.
  const (
// ⌄ enclosing_range_start 0.1.test `sg/pr206`/BlockConst1.
   // BlockConst1 is a multi-line doc.
   // It spans two lines.
   BlockConst1 = 1
// ^^^^^^^^^^^ definition 0.1.test `sg/pr206`/BlockConst1.
//             kind Constant
//             display_name BlockConst1
//             signature_documentation
//             > const BlockConst1 untyped int = 1
//             documentation
//             > BlockConst1 is a multi-line doc.
//             > It spans two lines.
//             documentation
//             > Block doc for the const group.
//               ⌃ enclosing_range_end 0.1.test `sg/pr206`/BlockConst1.
  
// ⌄ enclosing_range_start 0.1.test `sg/pr206`/BlockConstNoDoc.
   BlockConstNoDoc = 2
// ^^^^^^^^^^^^^^^ definition 0.1.test `sg/pr206`/BlockConstNoDoc.
//                 kind Constant
//                 display_name BlockConstNoDoc
//                 signature_documentation
//                 > const BlockConstNoDoc untyped int = 2
//                 documentation
//                 > Block doc for the const group.
//                   ⌃ enclosing_range_end 0.1.test `sg/pr206`/BlockConstNoDoc.
  
// ⌄ enclosing_range_start 0.1.test `sg/pr206`/BlockConstTrailing.
   BlockConstTrailing = 3 // trailing comment on const
// ^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/pr206`/BlockConstTrailing.
//                    kind Constant
//                    display_name BlockConstTrailing
//                    signature_documentation
//                    > const BlockConstTrailing untyped int = 3
//                    documentation
//                    > trailing comment on const
//                    documentation
//                    > Block doc for the const group.
//                      ⌃ enclosing_range_end 0.1.test `sg/pr206`/BlockConstTrailing.
  )
  
//⌄ enclosing_range_start 0.1.test `sg/pr206`/OrphanConst.
  const (
   // OrphanConst lives in a block with no block-level doc.
   OrphanConst = 99
// ^^^^^^^^^^^ definition 0.1.test `sg/pr206`/OrphanConst.
//             kind Constant
//             display_name OrphanConst
//             signature_documentation
//             > const OrphanConst untyped int = 99
//             documentation
//             > OrphanConst lives in a block with no block-level doc.
  )
//⌃ enclosing_range_end 0.1.test `sg/pr206`/OrphanConst.
  
