  package embedded
//        ^^^^^^^^ definition 0.1.test `sg/embedded`/
  
  import "fmt"
//        ^^^ reference github.com/golang/go/src go1.22 fmt/
  
//⌄ enclosing_range_start 0.1.test `sg/embedded`/RecentCommittersResults#String().
  func (r *RecentCommittersResults) String() string {
//      ^ definition local 0
//        kind Variable
//        display_name r
//        signature_documentation
//        > var r *RecentCommittersResults
//         ^^^^^^^^^^^^^^^^^^^^^^^ reference 0.1.test `sg/embedded`/RecentCommittersResults#
//                                  ^^^^^^ definition 0.1.test `sg/embedded`/RecentCommittersResults#String().
//                                         kind Method
//                                         display_name String
//                                         signature_documentation
//                                         > func (*RecentCommittersResults).String() string
//                                         relationship github.com/golang/go/src go1.22 fmt/Stringer#String(). implementation
   return fmt.Sprintf("RecentCommittersResults{Nodes: %d}", len(r.Nodes))
//        ^^^ reference github.com/golang/go/src go1.22 fmt/
//            ^^^^^^^ reference github.com/golang/go/src go1.22 fmt/Sprintf().
//                                                              ^ reference local 0
//                                                                ^^^^^ reference 0.1.test `sg/embedded`/RecentCommittersResults#Nodes.
  }
//⌃ enclosing_range_end 0.1.test `sg/embedded`/RecentCommittersResults#String().
  
  type RecentCommittersResults struct {
//     ^^^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/embedded`/RecentCommittersResults#
//                             kind Struct
//                             display_name RecentCommittersResults
//                             signature_documentation
//                             > type RecentCommittersResults struct {
//                             >     Nodes []struct {
//                             >         Authors struct {
//                             >             Nodes []struct {
//                             >                 Date      string
//                             >                 Email     string
//                             >                 Name      string
//                             >                 User      struct{ Login string }
//                             >                 AvatarURL string
//                             >             }
//                             >         }
//                             >     }
//                             >     PageInfo struct{ HasNextPage bool }
//                             > }
//                             relationship github.com/golang/go/src go1.22 fmt/Stringer# implementation
   Nodes []struct {
// ^^^^^ definition 0.1.test `sg/embedded`/RecentCommittersResults#Nodes.
//       kind Field
//       display_name Nodes
//       signature_documentation
//       > struct field Nodes []struct{Authors struct{Nodes []struct{Date string; Email string; Name string; User struct{Login string}; AvatarURL string}}}
    Authors struct {
//  ^^^^^^^ definition 0.1.test `sg/embedded`/RecentCommittersResults#$anon_90b32de18ec80596#Authors.
//          kind Field
//          display_name Authors
//          signature_documentation
//          > struct field Authors struct{Nodes []struct{Date string; Email string; Name string; User struct{Login string}; AvatarURL string}}
     Nodes []struct {
//   ^^^^^ definition 0.1.test `sg/embedded`/RecentCommittersResults#$anon_90b32de18ec80596#$anon_bb290b1f6ea0cf58#Nodes.
//         kind Field
//         display_name Nodes
//         signature_documentation
//         > struct field Nodes []struct{Date string; Email string; Name string; User struct{Login string}; AvatarURL string}
      Date  string
//    ^^^^ definition 0.1.test `sg/embedded`/RecentCommittersResults#$anon_90b32de18ec80596#$anon_bb290b1f6ea0cf58#$anon_b2a8a16c744b2d4b#Date.
//         kind Field
//         display_name Date
//         signature_documentation
//         > struct field Date string
      Email string
//    ^^^^^ definition 0.1.test `sg/embedded`/RecentCommittersResults#$anon_90b32de18ec80596#$anon_bb290b1f6ea0cf58#$anon_b2a8a16c744b2d4b#Email.
//          kind Field
//          display_name Email
//          signature_documentation
//          > struct field Email string
      Name  string
//    ^^^^ definition 0.1.test `sg/embedded`/RecentCommittersResults#$anon_90b32de18ec80596#$anon_bb290b1f6ea0cf58#$anon_b2a8a16c744b2d4b#Name.
//         kind Field
//         display_name Name
//         signature_documentation
//         > struct field Name string
      User  struct {
//    ^^^^ definition 0.1.test `sg/embedded`/RecentCommittersResults#$anon_90b32de18ec80596#$anon_bb290b1f6ea0cf58#$anon_b2a8a16c744b2d4b#User.
//         kind Field
//         display_name User
//         signature_documentation
//         > struct field User struct{Login string}
       Login string
//     ^^^^^ definition 0.1.test `sg/embedded`/RecentCommittersResults#$anon_90b32de18ec80596#$anon_bb290b1f6ea0cf58#$anon_b2a8a16c744b2d4b#$anon_d4bff1f61f45b2a1#Login.
//           kind Field
//           display_name Login
//           signature_documentation
//           > struct field Login string
      }
      AvatarURL string
//    ^^^^^^^^^ definition 0.1.test `sg/embedded`/RecentCommittersResults#$anon_90b32de18ec80596#$anon_bb290b1f6ea0cf58#$anon_b2a8a16c744b2d4b#AvatarURL.
//              kind Field
//              display_name AvatarURL
//              signature_documentation
//              > struct field AvatarURL string
     }
    }
   }
   PageInfo struct {
// ^^^^^^^^ definition 0.1.test `sg/embedded`/RecentCommittersResults#PageInfo.
//          kind Field
//          display_name PageInfo
//          signature_documentation
//          > struct field PageInfo struct{HasNextPage bool}
    HasNextPage bool
//  ^^^^^^^^^^^ definition 0.1.test `sg/embedded`/RecentCommittersResults#$anon_0a5c453971407ce4#HasNextPage.
//              kind Field
//              display_name HasNextPage
//              signature_documentation
//              > struct field HasNextPage bool
   }
  }
  
