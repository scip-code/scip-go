  package inlinestruct
//        ^^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/
  
  type InFuncSig struct {
//     ^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/InFuncSig#
//               kind Struct
//               display_name InFuncSig
//               signature_documentation
//               > type InFuncSig struct{ value bool }
//               documentation
//               > ```go
//               > type InFuncSig struct{ value bool }
//               > ```
   value bool
// ^^^^^ definition 0.1.test `sg/inlinestruct`/InFuncSig#value.
//       kind Field
//       display_name value
//       signature_documentation
//       > struct field value bool
//       documentation
//       > ```go
//       > struct field value bool
//       > ```
  }
  
  var rowsCloseHook = func() func(InFuncSig, *error) { return nil }
//    ^^^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/rowsCloseHook.
//                  kind Variable
//                  display_name rowsCloseHook
//                  signature_documentation
//                  > var rowsCloseHook func() func(InFuncSig, *error)
//                  documentation
//                  > ```go
//                  > var rowsCloseHook func() func(InFuncSig, *error)
//                  > ```
//                                ^^^^^^^^^ reference 0.1.test `sg/inlinestruct`/InFuncSig#
  
