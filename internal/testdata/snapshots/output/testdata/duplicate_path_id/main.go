  package gosrc
//        ^^^^^ definition 0.1.test `sg/testdata/duplicate_path_id`/
//              kind Package
//              display_name gosrc
//              signature_documentation
//              > package gosrc
//              documentation
//              > ```go
//              > package gosrc
//              > ```
  
  type importMeta struct{}
//     ^^^^^^^^^^ definition 0.1.test `sg/testdata/duplicate_path_id`/importMeta#
//                kind Struct
//                display_name importMeta
//                signature_documentation
//                > type importMeta struct{}
//                documentation
//                > ```go
//                > type importMeta struct{}
//                > ```
  
  type sourceMeta struct{}
//     ^^^^^^^^^^ definition 0.1.test `sg/testdata/duplicate_path_id`/sourceMeta#
//                kind Struct
//                display_name sourceMeta
//                signature_documentation
//                > type sourceMeta struct{}
//                documentation
//                > ```go
//                > type sourceMeta struct{}
//                > ```
  
//⌄ enclosing_range_start 0.1.test `sg/testdata/duplicate_path_id`/fetchMeta().
  func fetchMeta() (string, *importMeta, *sourceMeta) {
//     ^^^^^^^^^ definition 0.1.test `sg/testdata/duplicate_path_id`/fetchMeta().
//               kind Function
//               display_name fetchMeta
//               signature_documentation
//               > func fetchMeta() (string, *importMeta, *sourceMeta)
//               documentation
//               > ```go
//               > func fetchMeta() (string, *importMeta, *sourceMeta)
//               > ```
//                           ^^^^^^^^^^ reference 0.1.test `sg/testdata/duplicate_path_id`/importMeta#
//                                        ^^^^^^^^^^ reference 0.1.test `sg/testdata/duplicate_path_id`/sourceMeta#
   panic("hmm")
  }
//⌃ enclosing_range_end 0.1.test `sg/testdata/duplicate_path_id`/fetchMeta().
  
//⌄ enclosing_range_start 0.1.test `sg/testdata/duplicate_path_id`/init().
  func init() {}
//     ^^^^ definition 0.1.test `sg/testdata/duplicate_path_id`/init().
//          kind Function
//          display_name init
//          signature_documentation
//          > func init()
//          documentation
//          > ```go
//          > func init()
//          > ```
//             ⌃ enclosing_range_end 0.1.test `sg/testdata/duplicate_path_id`/init().
//⌄ enclosing_range_start 0.1.test `sg/testdata/duplicate_path_id`/init().
  func init() {}
//     ^^^^ definition 0.1.test `sg/testdata/duplicate_path_id`/init().
//          kind Function
//          display_name init
//          signature_documentation
//          > func init()
//          documentation
//          > ```go
//          > func init()
//          > ```
//             ⌃ enclosing_range_end 0.1.test `sg/testdata/duplicate_path_id`/init().
//⌄ enclosing_range_start 0.1.test `sg/testdata/duplicate_path_id`/init().
  func init() {}
//     ^^^^ definition 0.1.test `sg/testdata/duplicate_path_id`/init().
//          kind Function
//          display_name init
//          signature_documentation
//          > func init()
//          documentation
//          > ```go
//          > func init()
//          > ```
//             ⌃ enclosing_range_end 0.1.test `sg/testdata/duplicate_path_id`/init().
  
