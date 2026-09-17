  package switches
//        ^^^^^^^^ definition 0.1.test `sg/switches`/
//                 kind Package
//                 display_name switches
//                 signature_documentation
//                 > package switches
//                 documentation
//                 > ```go
//                 > package switches
//                 > ```
  
  // CustomSwitch does the things in a switch
  type CustomSwitch struct{}
//     ^^^^^^^^^^^^ definition 0.1.test `sg/switches`/CustomSwitch#
//                  kind Struct
//                  display_name CustomSwitch
//                  signature_documentation
//                  > type CustomSwitch struct{}
//                  documentation
//                  > ```go
//                  > type CustomSwitch struct{}
//                  > ```
//                  documentation
//                  > CustomSwitch does the things in a switch
  
  // Something does some things... and stuff
//⌄ enclosing_range_start 0.1.test `sg/switches`/CustomSwitch#Something().
  func (c *CustomSwitch) Something() bool { return false }
//      ^ definition local 0
//        kind Variable
//        display_name c
//        signature_documentation
//        > var c *CustomSwitch
//         ^^^^^^^^^^^^ reference 0.1.test `sg/switches`/CustomSwitch#
//                       ^^^^^^^^^ definition 0.1.test `sg/switches`/CustomSwitch#Something().
//                                 kind Method
//                                 display_name Something
//                                 signature_documentation
//                                 > func (*CustomSwitch).Something() bool
//                                 documentation
//                                 > ```go
//                                 > func (*CustomSwitch).Something() bool
//                                 > ```
//                                 documentation
//                                 > Something does some things... and stuff
//                                                       ⌃ enclosing_range_end 0.1.test `sg/switches`/CustomSwitch#Something().
  
//⌄ enclosing_range_start 0.1.test `sg/switches`/Switch().
  func Switch(interfaceValue interface{}) bool {
//     ^^^^^^ definition 0.1.test `sg/switches`/Switch().
//            kind Function
//            display_name Switch
//            signature_documentation
//            > func Switch(interfaceValue interface{}) bool
//            documentation
//            > ```go
//            > func Switch(interfaceValue interface{}) bool
//            > ```
//            ^^^^^^^^^^^^^^ definition local 1
//                           kind Variable
//                           display_name interfaceValue
//                           signature_documentation
//                           > var interfaceValue interface{}
   switch concreteValue := interfaceValue.(type) {
//        ^^^^^^^^^^^^^ definition local 2
//                      kind Variable
//                      display_name concreteValue
//                         ^^^^^^^^^^^^^^ reference local 1
   case int:
    return concreteValue*3 > 10
//         ^^^^^^^^^^^^^ reference local 2
   case bool:
    return !concreteValue
//          ^^^^^^^^^^^^^ reference local 2
   case CustomSwitch:
//      ^^^^^^^^^^^^ reference 0.1.test `sg/switches`/CustomSwitch#
    return concreteValue.Something()
//         ^^^^^^^^^^^^^ reference local 2
//                       ^^^^^^^^^ reference 0.1.test `sg/switches`/CustomSwitch#Something().
   default:
    return false
   }
  }
//⌃ enclosing_range_end 0.1.test `sg/switches`/Switch().
  
