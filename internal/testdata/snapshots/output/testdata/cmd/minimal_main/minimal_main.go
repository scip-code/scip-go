  package main
//        ^^^^ definition 0.1.test `sg/testdata/cmd/minimal_main`/
//             kind Package
//             display_name main
//             signature_documentation
//             > package main
//             documentation
//             > ```go
//             > package main
//             > ```
  
  type User struct {
//     ^^^^ definition 0.1.test `sg/testdata/cmd/minimal_main`/User#
//          kind Struct
//          display_name User
//          signature_documentation
//          > type User struct {
//          >     Id   string
//          >     Name string
//          > }
//          documentation
//          > ```go
//          > type User struct {
//          >     Id   string
//          >     Name string
//          > }
//          > ```
   Id, Name string
// ^^ definition 0.1.test `sg/testdata/cmd/minimal_main`/User#Id.
//    kind Field
//    display_name Id
//    signature_documentation
//    > struct field Id string
//    documentation
//    > ```go
//    > struct field Id string
//    > ```
//     ^^^^ definition 0.1.test `sg/testdata/cmd/minimal_main`/User#Name.
//          kind Field
//          display_name Name
//          signature_documentation
//          > struct field Name string
//          documentation
//          > ```go
//          > struct field Name string
//          > ```
  }
  
  type UserResource struct{}
//     ^^^^^^^^^^^^ definition 0.1.test `sg/testdata/cmd/minimal_main`/UserResource#
//                  kind Struct
//                  display_name UserResource
//                  signature_documentation
//                  > type UserResource struct{}
//                  documentation
//                  > ```go
//                  > type UserResource struct{}
//                  > ```
  
//⌄ enclosing_range_start 0.1.test `sg/testdata/cmd/minimal_main`/main().
  func main() {}
//     ^^^^ definition 0.1.test `sg/testdata/cmd/minimal_main`/main().
//          kind Function
//          display_name main
//          signature_documentation
//          > func main()
//          documentation
//          > ```go
//          > func main()
//          > ```
//             ⌃ enclosing_range_end 0.1.test `sg/testdata/cmd/minimal_main`/main().
  
