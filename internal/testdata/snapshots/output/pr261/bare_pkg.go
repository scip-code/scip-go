  package pr261
//        ^^^^^ definition 0.1.test `sg/pr261`/
//              kind Package
//              display_name pr261
//              signature_documentation
//              > package pr261
  
  import "sort"
//        ^^^^ reference github.com/golang/go/src go1.22 sort/
  
  // Legitimate use so goimports doesn't strip the import.
  var _ = sort.Slice
//        ^^^^ reference github.com/golang/go/src go1.22 sort/
//             ^^^^^ reference github.com/golang/go/src go1.22 sort/Slice().
  
  // Bare PkgName reference that triggered the panic.
  var _ = sort
  
