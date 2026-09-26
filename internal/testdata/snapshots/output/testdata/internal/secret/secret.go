  package secret
//        ^^^^^^ definition 0.1.test `sg/testdata/internal/secret`/
  
//⌄ enclosing_range_start 0.1.test `sg/testdata/internal/secret`/SecretScore.
  // SecretScore is like score but _secret_.
  const SecretScore = uint64(43)
//      ^^^^^^^^^^^ definition 0.1.test `sg/testdata/internal/secret`/SecretScore.
//                  kind Constant
//                  display_name SecretScore
//                  signature_documentation
//                  > const SecretScore uint64 = 43
//                  documentation
//                  > SecretScore is like score but _secret_.
//                             ⌃ enclosing_range_end 0.1.test `sg/testdata/internal/secret`/SecretScore.
  
//⌄ enclosing_range_start 0.1.test `sg/testdata/internal/secret`/Burger#
  // Original doc
  type Burger struct {
//     ^^^^^^ definition 0.1.test `sg/testdata/internal/secret`/Burger#
//            kind Struct
//            display_name Burger
//            signature_documentation
//            > type Burger struct{ Field int }
//            documentation
//            > Original doc
// ⌄ enclosing_range_start 0.1.test `sg/testdata/internal/secret`/Burger#Field.
   Field int
// ^^^^^ definition 0.1.test `sg/testdata/internal/secret`/Burger#Field.
//       kind Field
//       display_name Field
//       signature_documentation
//       > struct field Field int
//         ⌃ enclosing_range_end 0.1.test `sg/testdata/internal/secret`/Burger#Field.
  }
//⌃ enclosing_range_end 0.1.test `sg/testdata/internal/secret`/Burger#
  
