  package testdata
//        ^^^^^^^^ definition 0.1.test `sg/testdata`/
  
  import (
   "sg/testdata/internal/secret"
//  ^^^^^^^^^^^^^^^^^^^^^^^^^^^ reference 0.1.test `sg/testdata/internal/secret`/
  )
  
  // Type aliased doc
  type SecretBurger = secret.Burger
//     ^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/SecretBurger#
//                  kind TypeAlias
//                  display_name SecretBurger
//                  signature_documentation
//                  > type SecretBurger = secret.Burger
//                  documentation
//                  > ```go
//                  > type SecretBurger = secret.Burger
//                  > ```
//                  documentation
//                  > Type aliased doc
//                    ^^^^^^ reference 0.1.test `sg/testdata/internal/secret`/
//                           ^^^^^^ reference 0.1.test `sg/testdata/internal/secret`/Burger#
  
  type BadBurger = struct {
//     ^^^^^^^^^ definition 0.1.test `sg/testdata`/BadBurger#
//               kind TypeAlias
//               display_name BadBurger
//               signature_documentation
//               > type BadBurger = struct{ Field string }
//               documentation
//               > ```go
//               > type BadBurger = struct{ Field string }
//               > ```
   Field string
// ^^^^^ definition 0.1.test `sg/testdata`/BadBurger#Field.
//       kind Field
//       display_name Field
//       signature_documentation
//       > struct field Field string
//       documentation
//       > ```go
//       > struct field Field string
//       > ```
  }
  
