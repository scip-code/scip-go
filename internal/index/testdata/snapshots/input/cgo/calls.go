package review

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

func Calls() {
	ptr := C.CString("x")
	defer C.free(unsafe.Pointer(ptr))
	buf := C.malloc(16)
	defer C.free(buf)
	var out C.int
	C.fill(
		&out,
		&out,
		C.int(1),
		C.int(2),
		(*C.char)(unsafe.Pointer(ptr)))
	C.add(C.int(1), C.int(2))
	C.
		add(1, 2)
	_, _ = C.add(C.add(1, 2), 3)
	var _, _ = (C.add(1, 2))
	_, _ = C.add(1, 2), C.add(3, 4)
	_ = C.answer
	_ = C.global
	_ = C.dynamic
	_ = C.sizeof_int
	C := struct{ add func() }{add: func() {}}
	C.add()
}
