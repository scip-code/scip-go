  package testdata
//        ^^^^^^^^ definition 0.1.test `sg/testdata`/
  
  import (
   "context"
//  ^^^^^^^ reference github.com/golang/go/src go1.22 context/
  
   "sg/testdata/internal/secret"
//  ^^^^^^^^^^^^^^^^^^^^^^^^^^^ reference 0.1.test `sg/testdata/internal/secret`/
  )
  
  // TestInterface is an interface used for testing.
  type TestInterface interface {
//     ^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/TestInterface#
//                   kind Interface
//                   display_name TestInterface
//                   signature_documentation
//                   > type TestInterface interface {
//                   >     Do(ctx context.Context, data string) (score int, _ error)
//                   > }
//                   documentation
//                   > ```go
//                   > type TestInterface interface {
//                   >     Do(ctx context.Context, data string) (score int, _ error)
//                   > }
//                   > ```
//                   documentation
//                   > TestInterface is an interface used for testing.
   // Do does a test thing.
   Do(ctx context.Context, data string) (score int, _ error)
// ^^ definition 0.1.test `sg/testdata`/TestInterface#Do.
//    kind MethodSpecification
//    display_name Do
//    signature_documentation
//    > func (TestInterface).Do(ctx context.Context, data string) (score int, _ error)
//    documentation
//    > ```go
//    > func (TestInterface).Do(ctx context.Context, data string) (score int, _ error)
//    > ```
//    documentation
//    > Do does a test thing.
//    ^^^ definition local 0
//        kind Variable
//        display_name ctx
//        signature_documentation
//        > var ctx Context
//        ^^^^^^^ reference github.com/golang/go/src go1.22 context/
//                ^^^^^^^ reference github.com/golang/go/src go1.22 context/Context#
//                         ^^^^ definition local 1
//                              kind Variable
//                              display_name data
//                              signature_documentation
//                              > var data string
//                                       ^^^^^ definition local 2
//                                             kind Variable
//                                             display_name score
//                                             signature_documentation
//                                             > var score int
  }
  
  type (
   // TestStruct is a struct used for testing.
   TestStruct struct {
// ^^^^^^^^^^ definition 0.1.test `sg/testdata`/TestStruct#
//            kind Struct
//            display_name TestStruct
//            signature_documentation
//            > type TestStruct struct {
//            >     SimpleA                int
//            >     SimpleB                int
//            >     SimpleC                int
//            >     FieldWithTag           string "json:\"tag\""
//            >     FieldWithAnonymousType struct {
//            >         NestedA string
//            >         NestedB string
//            >         NestedC string
//            >     }
//            >     EmptyStructField struct{}
//            > }
//            documentation
//            > ```go
//            > type TestStruct struct {
//            >     SimpleA                int
//            >     SimpleB                int
//            >     SimpleC                int
//            >     FieldWithTag           string "json:\"tag\""
//            >     FieldWithAnonymousType struct {
//            >         NestedA string
//            >         NestedB string
//            >         NestedC string
//            >     }
//            >     EmptyStructField struct{}
//            > }
//            > ```
//            documentation
//            > TestStruct is a struct used for testing.
    // SimpleA docs
    SimpleA int
//  ^^^^^^^ definition 0.1.test `sg/testdata`/TestStruct#SimpleA.
//          kind Field
//          display_name SimpleA
//          signature_documentation
//          > struct field SimpleA int
//          documentation
//          > ```go
//          > struct field SimpleA int
//          > ```
//          documentation
//          > SimpleA docs
    // SimpleB docs
    SimpleB int
//  ^^^^^^^ definition 0.1.test `sg/testdata`/TestStruct#SimpleB.
//          kind Field
//          display_name SimpleB
//          signature_documentation
//          > struct field SimpleB int
//          documentation
//          > ```go
//          > struct field SimpleB int
//          > ```
//          documentation
//          > SimpleB docs
    // SimpleC docs
    SimpleC int
//  ^^^^^^^ definition 0.1.test `sg/testdata`/TestStruct#SimpleC.
//          kind Field
//          display_name SimpleC
//          signature_documentation
//          > struct field SimpleC int
//          documentation
//          > ```go
//          > struct field SimpleC int
//          > ```
//          documentation
//          > SimpleC docs
  
    FieldWithTag           string `json:"tag"`
//  ^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/TestStruct#FieldWithTag.
//               kind Field
//               display_name FieldWithTag
//               signature_documentation
//               > struct field FieldWithTag string
//               documentation
//               > ```go
//               > struct field FieldWithTag string
//               > ```
    FieldWithAnonymousType struct {
//  ^^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/TestStruct#FieldWithAnonymousType.
//                         kind Field
//                         display_name FieldWithAnonymousType
//                         signature_documentation
//                         > struct field FieldWithAnonymousType struct{NestedA string; NestedB string; NestedC string}
//                         documentation
//                         > ```go
//                         > struct field FieldWithAnonymousType struct{NestedA string; NestedB string; NestedC string}
//                         > ```
     NestedA string
//   ^^^^^^^ definition 0.1.test `sg/testdata`/TestStruct#$anon_2bed88e490dc48af#NestedA.
//           kind Field
//           display_name NestedA
//           signature_documentation
//           > struct field NestedA string
//           documentation
//           > ```go
//           > struct field NestedA string
//           > ```
     NestedB string
//   ^^^^^^^ definition 0.1.test `sg/testdata`/TestStruct#$anon_2bed88e490dc48af#NestedB.
//           kind Field
//           display_name NestedB
//           signature_documentation
//           > struct field NestedB string
//           documentation
//           > ```go
//           > struct field NestedB string
//           > ```
     // NestedC docs
     NestedC string
//   ^^^^^^^ definition 0.1.test `sg/testdata`/TestStruct#$anon_2bed88e490dc48af#NestedC.
//           kind Field
//           display_name NestedC
//           signature_documentation
//           > struct field NestedC string
//           documentation
//           > ```go
//           > struct field NestedC string
//           > ```
//           documentation
//           > NestedC docs
    }
  
    EmptyStructField struct{}
//  ^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/TestStruct#EmptyStructField.
//                   kind Field
//                   display_name EmptyStructField
//                   signature_documentation
//                   > struct field EmptyStructField struct{}
//                   documentation
//                   > ```go
//                   > struct field EmptyStructField struct{}
//                   > ```
   }
  
   TestEmptyStruct struct{}
// ^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/TestEmptyStruct#
//                 kind Struct
//                 display_name TestEmptyStruct
//                 signature_documentation
//                 > type TestEmptyStruct struct{}
//                 documentation
//                 > ```go
//                 > type TestEmptyStruct struct{}
//                 > ```
  )
  
  // Score is just a hardcoded number.
  const Score = uint64(42)
//      ^^^^^ definition 0.1.test `sg/testdata`/Score.
//            kind Constant
//            display_name Score
//            signature_documentation
//            > const Score uint64 = 42
//            documentation
//            > ```go
//            > const Score uint64 = 42
//            > ```
//            documentation
//            > Score is just a hardcoded number.
  const secretScore = secret.SecretScore
//      ^^^^^^^^^^^ definition 0.1.test `sg/testdata`/secretScore.
//                  kind Constant
//                  display_name secretScore
//                  signature_documentation
//                  > const secretScore uint64 = 43
//                  documentation
//                  > ```go
//                  > const secretScore uint64 = 43
//                  > ```
//                    ^^^^^^ reference 0.1.test `sg/testdata/internal/secret`/
//                           ^^^^^^^^^^^ reference 0.1.test `sg/testdata/internal/secret`/SecretScore.
  
  const SomeString = "foobar"
//      ^^^^^^^^^^ definition 0.1.test `sg/testdata`/SomeString.
//                 kind Constant
//                 display_name SomeString
//                 signature_documentation
//                 > const SomeString untyped string = "foobar"
//                 documentation
//                 > ```go
//                 > const SomeString untyped string = "foobar"
//                 > ```
  const LongString = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed tincidunt viverra aliquam. Phasellus finibus, arcu eu commodo porta, dui quam dictum ante, nec porta enim leo quis felis. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Curabitur luctus orci tortor, non condimentum arcu bibendum ut. Proin sit amet vulputate lorem, ut egestas arcu. Curabitur quis sagittis mi. Aenean elit sem, imperdiet ut risus eget, varius varius erat.\nNullam lobortis tortor sed sodales consectetur. Aenean condimentum vehicula elit, eget interdum ante finibus nec. Mauris mollis, nulla eu vehicula rhoncus, eros lectus viverra tellus, ac hendrerit quam massa et felis. Nunc vestibulum diam a facilisis sollicitudin. Aenean nec varius metus. Sed nec diam nibh. Ut erat erat, suscipit et ante eget, tincidunt condimentum orci. Aenean nec facilisis augue, ac sodales ex. Nulla dictum hendrerit tempus. Aliquam fringilla tortor in massa molestie, quis bibendum nulla ullamcorper. Suspendisse congue laoreet elit, vitae consectetur orci facilisis non. Aliquam tempus ultricies sapien, rhoncus tincidunt nisl tincidunt eget. Aliquam nisi ante, rutrum eget viverra imperdiet, congue ut nunc. Donec mollis sed tellus vel placerat. Sed mi ex, fringilla a fermentum a, tincidunt eget lectus.\nPellentesque lacus nibh, accumsan eget feugiat nec, gravida eget urna. Donec quam velit, imperdiet in consequat eget, ultricies eget nunc. Curabitur interdum vel sem et euismod. Donec sed vulputate odio, sit amet bibendum tellus. Integer pellentesque nunc eu turpis cursus, vestibulum sodales ipsum posuere. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Ut at vestibulum sapien. In hac habitasse platea dictumst. Nullam sed lobortis urna, non bibendum ipsum. Sed in sapien quis purus semper fringilla. Integer ut egestas nulla, eu ornare lectus. Maecenas quis sapien condimentum, dignissim urna quis, hendrerit neque. Donec cursus sit amet metus eu mollis.\nSed scelerisque vitae odio non egestas. Cras hendrerit tortor mauris. Aenean quis imperdiet nulla, a viverra purus. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Praesent finibus faucibus orci, sed ultrices justo iaculis ut. Ut libero massa, condimentum at elit non, fringilla iaculis quam. Sed sit amet ipsum placerat, tincidunt sem in, efficitur lacus. Curabitur ligula orci, tempus ut magna eget, sodales tristique odio.\nPellentesque in libero ac risus pretium ultrices. In hac habitasse platea dictumst. Curabitur a quam sed orci tempus luctus. Integer commodo nec odio quis consequat. Aenean vitae dapibus augue, nec dictum lectus. Etiam sit amet leo diam. Duis eu ligula venenatis, fermentum lacus vel, interdum odio. Vivamus sit amet libero vitae elit interdum cursus et eu erat. Cras interdum augue sit amet ex aliquet tempor. Praesent dolor nisl, convallis bibendum mauris a, euismod commodo ante. Phasellus non ipsum condimentum, molestie dolor quis, pretium nisi. Mauris augue urna, fermentum ut lacinia a, efficitur vitae odio. Praesent finibus nisl et dolor luctus faucibus. Donec eget lectus sed mi porttitor placerat ac eu odio."
//      ^^^^^^^^^^ definition 0.1.test `sg/testdata`/LongString.
//                 kind Constant
//                 display_name LongString
//                 signature_documentation
//                 > const LongString untyped string = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed tincidu...
//                 documentation
//                 > ```go
//                 > const LongString untyped string = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed tincidu...
//                 > ```
  const ConstMath = 1 + (2+3)*5
//      ^^^^^^^^^ definition 0.1.test `sg/testdata`/ConstMath.
//                kind Constant
//                display_name ConstMath
//                signature_documentation
//                > const ConstMath untyped int = 26
//                documentation
//                > ```go
//                > const ConstMath untyped int = 26
//                > ```
  
  type StringAlias string
//     ^^^^^^^^^^^ definition 0.1.test `sg/testdata`/StringAlias#
//                 kind Type
//                 display_name StringAlias
//                 signature_documentation
//                 > type StringAlias string
//                 documentation
//                 > ```go
//                 > type StringAlias string
//                 > ```
  
  const AliasedString StringAlias = "foobar"
//      ^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/AliasedString.
//                    kind Constant
//                    display_name AliasedString
//                    signature_documentation
//                    > const AliasedString StringAlias = "foobar"
//                    documentation
//                    > ```go
//                    > const AliasedString StringAlias = "foobar"
//                    > ```
//                    ^^^^^^^^^^^ reference 0.1.test `sg/testdata`/StringAlias#
  
  // Doer is similar to the test interface (but not the same).
//⌄ enclosing_range_start 0.1.test `sg/testdata`/TestStruct#Doer().
  func (ts *TestStruct) Doer(ctx context.Context, data string) (score int, err error) {
//      ^^ definition local 3
//         kind Variable
//         display_name ts
//         signature_documentation
//         > var ts *TestStruct
//          ^^^^^^^^^^ reference 0.1.test `sg/testdata`/TestStruct#
//                      ^^^^ definition 0.1.test `sg/testdata`/TestStruct#Doer().
//                           kind Method
//                           display_name Doer
//                           signature_documentation
//                           > func (*TestStruct).Doer(ctx context.Context, data string) (score int, err error)
//                           documentation
//                           > ```go
//                           > func (*TestStruct).Doer(ctx context.Context, data string) (score int, err error)
//                           > ```
//                           documentation
//                           > Doer is similar to the test interface (but not the same).
//                           ^^^ definition local 4
//                               kind Variable
//                               display_name ctx
//                               signature_documentation
//                               > var ctx Context
//                               ^^^^^^^ reference github.com/golang/go/src go1.22 context/
//                                       ^^^^^^^ reference github.com/golang/go/src go1.22 context/Context#
//                                                ^^^^ definition local 5
//                                                     kind Variable
//                                                     display_name data
//                                                     signature_documentation
//                                                     > var data string
//                                                              ^^^^^ definition local 6
//                                                                    kind Variable
//                                                                    display_name score
//                                                                    signature_documentation
//                                                                    > var score int
//                                                                         ^^^ definition local 7
//                                                                             kind Variable
//                                                                             display_name err
//                                                                             signature_documentation
//                                                                             > var err error
   return Score, nil
//        ^^^^^ reference 0.1.test `sg/testdata`/Score.
  }
//⌃ enclosing_range_end 0.1.test `sg/testdata`/TestStruct#Doer().
  
  // StructTagRegression is a struct that caused panic in the wild. Added here to
  // support a regression test.
  //
  // See https://github.com/tal-tech/go-zero/blob/11dd3d75ecceaa3f5772024fb3f26dec1ada8e9c/core/mapping/unmarshaler_test.go#L2272.
  type StructTagRegression struct {
//     ^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/StructTagRegression#
//                         kind Struct
//                         display_name StructTagRegression
//                         signature_documentation
//                         > type StructTagRegression struct {
//                         >     Value int "key:\",range=[:}\""
//                         > }
//                         documentation
//                         > ```go
//                         > type StructTagRegression struct {
//                         >     Value int "key:\",range=[:}\""
//                         > }
//                         > ```
//                         documentation
//                         > StructTagRegression is a struct that caused panic in the wild. Added here to
//                         > support a regression test.
//                         > 
//                         > See https://github.com/tal-tech/go-zero/blob/11dd3d75ecceaa3f5772024fb3f26dec1ada8e9c/core/mapping/unmarshaler_test.go#L2272.
   Value int `key:",range=[:}"`
// ^^^^^ definition 0.1.test `sg/testdata`/StructTagRegression#Value.
//       kind Field
//       display_name Value
//       signature_documentation
//       > struct field Value int
//       documentation
//       > ```go
//       > struct field Value int
//       > ```
  }
  
  type TestEqualsStruct = struct {
//     ^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/TestEqualsStruct#
//                      kind TypeAlias
//                      display_name TestEqualsStruct
//                      signature_documentation
//                      > type TestEqualsStruct = struct{ Value int }
//                      documentation
//                      > ```go
//                      > type TestEqualsStruct = struct{ Value int }
//                      > ```
   Value int
// ^^^^^ definition 0.1.test `sg/testdata`/TestEqualsStruct#Value.
//       kind Field
//       display_name Value
//       signature_documentation
//       > struct field Value int
//       documentation
//       > ```go
//       > struct field Value int
//       > ```
  }
  
  type ShellStruct struct {
//     ^^^^^^^^^^^ definition 0.1.test `sg/testdata`/ShellStruct#
//                 kind Struct
//                 display_name ShellStruct
//                 signature_documentation
//                 > type ShellStruct struct{ InnerStruct }
//                 documentation
//                 > ```go
//                 > type ShellStruct struct{ InnerStruct }
//                 > ```
   // Ensure this field comes before the definition
   // so that we grab the correct one in our unit
   // tests.
   InnerStruct
// ^^^^^^^^^^^ definition 0.1.test `sg/testdata`/ShellStruct#InnerStruct.
//             kind Field
//             display_name InnerStruct
//             signature_documentation
//             > struct field InnerStruct InnerStruct
//             documentation
//             > ```go
//             > struct field InnerStruct InnerStruct
//             > ```
//             documentation
//             > Ensure this field comes before the definition
//             > so that we grab the correct one in our unit
//             > tests.
// ^^^^^^^^^^^ reference 0.1.test `sg/testdata`/InnerStruct#
  }
  
  type InnerStruct struct{}
//     ^^^^^^^^^^^ definition 0.1.test `sg/testdata`/InnerStruct#
//                 kind Struct
//                 display_name InnerStruct
//                 signature_documentation
//                 > type InnerStruct struct{}
//                 documentation
//                 > ```go
//                 > type InnerStruct struct{}
//                 > ```
  
