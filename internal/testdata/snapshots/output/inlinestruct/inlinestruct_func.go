  package inlinestruct
//        ^^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/
  
//⌄ enclosing_range_start 0.1.test `sg/inlinestruct`/InFuncSig#
  type InFuncSig struct {
//     ^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/InFuncSig#
//               kind Struct
//               display_name InFuncSig
//               signature_documentation
//               > type InFuncSig struct{ value bool }
// ⌄ enclosing_range_start 0.1.test `sg/inlinestruct`/InFuncSig#value.
   value bool
// ^^^^^ definition 0.1.test `sg/inlinestruct`/InFuncSig#value.
//       kind Field
//       display_name value
//       signature_documentation
//       > struct field value bool
//          ⌃ enclosing_range_end 0.1.test `sg/inlinestruct`/InFuncSig#value.
  }
//⌃ enclosing_range_end 0.1.test `sg/inlinestruct`/InFuncSig#
  
//⌄ enclosing_range_start 0.1.test `sg/inlinestruct`/rowsCloseHook.
  var rowsCloseHook = func() func(InFuncSig, *error) { return nil }
//    ^^^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/rowsCloseHook.
//                  kind Variable
//                  display_name rowsCloseHook
//                  signature_documentation
//                  > var rowsCloseHook func() func(InFuncSig, *error)
//                                ^^^^^^^^^ reference 0.1.test `sg/inlinestruct`/InFuncSig#
//                                                                ⌃ enclosing_range_end 0.1.test `sg/inlinestruct`/rowsCloseHook.
  
