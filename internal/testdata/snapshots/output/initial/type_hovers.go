  package initial
//        ^^^^^^^ definition 0.1.test `sg/initial`/
  
  type (
   // HoverTypeList is a cool struct
   HoverTypeList struct{}
// ^^^^^^^^^^^^^ definition 0.1.test `sg/initial`/HoverTypeList#
//               kind Struct
//               display_name HoverTypeList
//               signature_documentation
//               > type HoverTypeList struct{}
//               documentation
//               > ```go
//               > type HoverTypeList struct{}
//               > ```
//               documentation
//               > HoverTypeList is a cool struct
  )
  
  // This should show up as well
  type HoverType struct{}
//     ^^^^^^^^^ definition 0.1.test `sg/initial`/HoverType#
//               kind Struct
//               display_name HoverType
//               signature_documentation
//               > type HoverType struct{}
//               documentation
//               > ```go
//               > type HoverType struct{}
//               > ```
//               documentation
//               > This should show up as well
  
