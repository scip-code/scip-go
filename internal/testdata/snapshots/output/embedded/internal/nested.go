  package nested_internal
//        ^^^^^^^^^^^^^^^ definition 0.1.test `sg/embedded/internal`/
//                        kind Package
//                        display_name nested_internal
//                        signature_documentation
//                        > package nested_internal
//                        documentation
//                        > ```go
//                        > package nested_internal
//                        > ```
  
  import (
   "fmt"
//  ^^^ reference github.com/golang/go/src go1.22 fmt/
   "sg/embedded"
//  ^^^^^^^^^^^ reference 0.1.test `sg/embedded`/
  )
  
//⌄ enclosing_range_start 0.1.test `sg/embedded/internal`/Something().
  func Something(recent embedded.RecentCommittersResults) {
//     ^^^^^^^^^ definition 0.1.test `sg/embedded/internal`/Something().
//               kind Function
//               display_name Something
//               signature_documentation
//               > func Something(recent embedded.RecentCommittersResults)
//               documentation
//               > ```go
//               > func Something(recent embedded.RecentCommittersResults)
//               > ```
//               ^^^^^^ definition local 0
//                      kind Variable
//                      display_name recent
//                      signature_documentation
//                      > var recent RecentCommittersResults
//                      ^^^^^^^^ reference 0.1.test `sg/embedded`/
//                               ^^^^^^^^^^^^^^^^^^^^^^^ reference 0.1.test `sg/embedded`/RecentCommittersResults#
   for _, commit := range recent.Nodes {
//        ^^^^^^ definition local 1
//               kind Variable
//               display_name commit
//               signature_documentation
//               > var commit struct{Authors struct{Nodes []struct{Date string; Email string; Name string; User struct{Login string}; AvatarURL string}}}
//                        ^^^^^^ reference local 0
//                               ^^^^^ reference 0.1.test `sg/embedded`/RecentCommittersResults#Nodes.
    for _, author := range commit.Authors.Nodes {
//         ^^^^^^ definition local 2
//                kind Variable
//                display_name author
//                signature_documentation
//                > var author struct{Date string; Email string; Name string; User struct{Login string}; AvatarURL string}
//                         ^^^^^^ reference local 1
//                                ^^^^^^^ reference 0.1.test `sg/embedded`/RecentCommittersResults#$anon_90b32de18ec80596#Authors.
//                                        ^^^^^ reference 0.1.test `sg/embedded`/RecentCommittersResults#$anon_90b32de18ec80596#$anon_bb290b1f6ea0cf58#Nodes.
     fmt.Println(author.Name)
//   ^^^ reference github.com/golang/go/src go1.22 fmt/
//       ^^^^^^^ reference github.com/golang/go/src go1.22 fmt/Println().
//               ^^^^^^ reference local 2
//                      ^^^^ reference 0.1.test `sg/embedded`/RecentCommittersResults#$anon_90b32de18ec80596#$anon_bb290b1f6ea0cf58#$anon_b2a8a16c744b2d4b#Name.
    }
   }
  }
//⌃ enclosing_range_end 0.1.test `sg/embedded/internal`/Something().
  
