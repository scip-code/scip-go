  package server
//        ^^^^^^ definition 0.1.test `sg/sharedtestmodule/cmd/gitserver/server`/
  
  import "testing"
//        ^^^^^^^ reference github.com/golang/go/src go1.22 testing/
  
//⌄ enclosing_range_start 0.1.test `sg/sharedtestmodule/cmd/gitserver/server`/TestStuff().
  func TestStuff(t *testing.T) {
//     ^^^^^^^^^ definition 0.1.test `sg/sharedtestmodule/cmd/gitserver/server`/TestStuff().
//               kind Function
//               display_name TestStuff
//               signature_documentation
//               > func TestStuff(t *testing.T)
//               documentation
//               > ```go
//               > func TestStuff(t *testing.T)
//               > ```
//               ^ definition local 0
//                 kind Variable
//                 display_name t
//                 signature_documentation
//                 > var t *T
//                  ^^^^^^^ reference github.com/golang/go/src go1.22 testing/
//                          ^ reference github.com/golang/go/src go1.22 testing/T#
   wd := "hello"
// ^^ definition local 1
//    kind Variable
//    display_name wd
//    signature_documentation
//    > var wd string
   repo := "world"
// ^^^^ definition local 2
//      kind Variable
//      display_name repo
//      signature_documentation
//      > var repo string
  
   runCmd(t, wd, "git", "init", repo)
// ^^^^^^ reference 0.1.test `sg/sharedtestmodule/cmd/gitserver/server`/runCmd().
//        ^ reference local 0
//           ^^ reference local 1
//                              ^^^^ reference local 2
  }
//⌃ enclosing_range_end 0.1.test `sg/sharedtestmodule/cmd/gitserver/server`/TestStuff().
  
