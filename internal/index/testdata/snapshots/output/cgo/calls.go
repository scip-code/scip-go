  package review
//        ^^^^^^ definition 0.1.test `sg/cgo`/
//               kind Package
//               display_name review
//               signature_documentation
//               > package review
//               documentation
//               > 
  
  /*
  #include <stdlib.h>
  #define answer 42
  static int global = 1;
  #define dynamic (global + 1)
  static int add(int a, int b) { return a + b; }
  static int fill(int *a, int *b, int c, int d, char *e) { return *a + *b + c + d + *e; }
  */
  import "C"
  import "unsafe"
//        ^^^^^^ reference github.com/golang/go/src go1.22 unsafe/
  
//⌄ enclosing_range_start 0.1.test `sg/cgo`/Calls().
  func Calls() {
//     ^^^^^ definition 0.1.test `sg/cgo`/Calls().
//           kind Function
//           display_name Calls
//           signature_documentation
//           > func Calls()
   ptr := C.CString("x")
// ^^^ definition local 0
//     kind Variable
//     display_name ptr
//     signature_documentation
//     > var ptr *_Ctype_char
//        ^^^^^^^^^ reference 0.1.test `sg/cgo`/_Cfunc_CString().
   defer C.free(unsafe.Pointer(ptr))
//       ^^^^^^ reference 0.1.test `sg/cgo`/_Cfunc_free().
//              ^^^^^^ reference github.com/golang/go/src go1.22 unsafe/
//                             ^^^ reference local 0
   buf := C.malloc(16)
// ^^^ definition local 1
//     kind Variable
//     display_name buf
//     signature_documentation
//     > var buf Pointer
//        ^^^^^^^^ reference 0.1.test `sg/cgo`/_Cfunc__CMalloc().
   defer C.free(buf)
//       ^^^^^^ reference 0.1.test `sg/cgo`/_Cfunc_free().
//              ^^^ reference local 1
   var out C.int
//     ^^^ definition local 2
//         kind Variable
//         display_name out
//         signature_documentation
//         > var out _Ctype_int
//         ^^^^^ reference 0.1.test `sg/cgo`/_Ctype_int#
   C.fill(
// ^^^^^^ reference 0.1.test `sg/cgo`/_Cfunc_fill().
    &out,
//   ^^^ reference local 2
    &out,
//   ^^^ reference local 2
    C.int(1),
//  ^^^^^ reference 0.1.test `sg/cgo`/_Ctype_int#
    C.int(2),
//  ^^^^^ reference 0.1.test `sg/cgo`/_Ctype_int#
    (*C.char)(unsafe.Pointer(ptr)))
//    ^^^^^^ reference 0.1.test `sg/cgo`/_Ctype_char#
//            ^^^^^^ reference github.com/golang/go/src go1.22 unsafe/
//                           ^^^ reference local 0
   C.add(C.int(1), C.int(2))
// ^^^^^ reference 0.1.test `sg/cgo`/_Cfunc_add().
//       ^^^^^ reference 0.1.test `sg/cgo`/_Ctype_int#
//                 ^^^^^ reference 0.1.test `sg/cgo`/_Ctype_int#
   C.
// ^^ reference 0.1.test `sg/cgo`/_Cfunc_add(). 1:5
    add(1, 2)
   _, _ = C.add(C.add(1, 2), 3)
//        ^^^^^ reference 0.1.test `sg/cgo`/_C2func_add().
//              ^^^^^ reference 0.1.test `sg/cgo`/_Cfunc_add().
   var _, _ = (C.add(1, 2))
//             ^^^^^ reference 0.1.test `sg/cgo`/_C2func_add().
   _, _ = C.add(1, 2), C.add(3, 4)
//        ^^^^^ reference 0.1.test `sg/cgo`/_Cfunc_add().
//                     ^^^^^ reference 0.1.test `sg/cgo`/_Cfunc_add().
   _ = C.answer
//     ^^^^^^^^ reference 0.1.test `sg/cgo`/_Ciconst_answer.
   _ = C.global
//     ^^^^^^^^ reference 0.1.test `sg/cgo`/_Cvar_global.
   _ = C.dynamic
//     ^^^^^^^^^ reference 0.1.test `sg/cgo`/_Cmacro_dynamic().
   _ = C.sizeof_int
//     ^^^^^^^^^^^^ reference 0.1.test `sg/cgo`/_Ciconst_sizeof_int.
   C := struct{ add func() }{add: func() {}}
// ^ definition local 3
//   kind Variable
//   display_name C
//   signature_documentation
//   > var C struct{add func()}
//              ^^^ definition local 4
//                  kind Field
//                  display_name add
//                  signature_documentation
//                  > field add func()
//                           ^^^ reference local 4
   C.add()
  }
//⌃ enclosing_range_end 0.1.test `sg/cgo`/Calls().
  
