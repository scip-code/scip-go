  package initial
//        ^^^^^^^ definition 0.1.test `sg/initial`/
  
//⌄ enclosing_range_start 0.1.test `sg/initial`/HoverTypeList#
  type (
   // HoverTypeList is a cool struct
   HoverTypeList struct{}
// ^^^^^^^^^^^^^ definition 0.1.test `sg/initial`/HoverTypeList#
//               kind Struct
//               display_name HoverTypeList
//               signature_documentation
//               > type HoverTypeList struct{}
//               documentation
//               > HoverTypeList is a cool struct
  )
//⌃ enclosing_range_end 0.1.test `sg/initial`/HoverTypeList#
  
//⌄ enclosing_range_start 0.1.test `sg/initial`/HoverType#
  // This should show up as well
  type HoverType struct{}
//     ^^^^^^^^^ definition 0.1.test `sg/initial`/HoverType#
//               kind Struct
//               display_name HoverType
//               signature_documentation
//               > type HoverType struct{}
//               documentation
//               > This should show up as well
//                      ⌃ enclosing_range_end 0.1.test `sg/initial`/HoverType#
  
