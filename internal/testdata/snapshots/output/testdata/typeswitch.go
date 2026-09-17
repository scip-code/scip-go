  package testdata
//        ^^^^^^^^ definition 0.1.test `sg/testdata`/
  
//⌄ enclosing_range_start 0.1.test `sg/testdata`/Switch().
  func Switch(interfaceValue interface{}) bool {
//     ^^^^^^ definition 0.1.test `sg/testdata`/Switch().
//            kind Function
//            display_name Switch
//            signature_documentation
//            > func Switch(interfaceValue interface{}) bool
//            documentation
//            > ```go
//            > func Switch(interfaceValue interface{}) bool
//            > ```
//            ^^^^^^^^^^^^^^ definition local 0
//                           kind Variable
//                           display_name interfaceValue
//                           signature_documentation
//                           > var interfaceValue interface{}
   switch concreteValue := interfaceValue.(type) {
//        ^^^^^^^^^^^^^ definition local 1
//                      kind Variable
//                      display_name concreteValue
//                         ^^^^^^^^^^^^^^ reference local 0
   case int:
    return concreteValue*3 > 10
//         ^^^^^^^^^^^^^ reference local 1
   case bool:
    return !concreteValue
//          ^^^^^^^^^^^^^ reference local 1
   default:
    return false
   }
  }
//⌃ enclosing_range_end 0.1.test `sg/testdata`/Switch().
  
