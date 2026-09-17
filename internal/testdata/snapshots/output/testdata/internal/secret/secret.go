  package secret
//        ^^^^^^ definition 0.1.test `sg/testdata/internal/secret`/
  
  // SecretScore is like score but _secret_.
  const SecretScore = uint64(43)
//      ^^^^^^^^^^^ definition 0.1.test `sg/testdata/internal/secret`/SecretScore.
//                  kind Constant
//                  display_name SecretScore
//                  signature_documentation
//                  > const SecretScore uint64 = 43
//                  documentation
//                  > ```go
//                  > const SecretScore uint64 = 43
//                  > ```
//                  documentation
//                  > SecretScore is like score but _secret_.
  
  // Original doc
  type Burger struct {
//     ^^^^^^ definition 0.1.test `sg/testdata/internal/secret`/Burger#
//            kind Struct
//            display_name Burger
//            signature_documentation
//            > type Burger struct{ Field int }
//            documentation
//            > ```go
//            > type Burger struct{ Field int }
//            > ```
//            documentation
//            > Original doc
   Field int
// ^^^^^ definition 0.1.test `sg/testdata/internal/secret`/Burger#Field.
//       kind Field
//       display_name Field
//       signature_documentation
//       > struct field Field int
//       documentation
//       > ```go
//       > struct field Field int
//       > ```
  }
  
