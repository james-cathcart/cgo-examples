/*

Reference: https://pkg.go.dev/cmd/cgo

Calling variadic C functions is not supported. It is possible to circumvent this
by using a C function wrapper. For example:

*/

package main

/*
#include <stdio.h>
#include <stdlib.h>

static void myprint(char *s) {
	printf("%s\n", s);
}
*/
import "C"
import "unsafe"

func main() {

	cs := C.CString("Hello from stdio")
	C.myprint(cs)
	C.free(unsafe.Pointer(cs))

}
