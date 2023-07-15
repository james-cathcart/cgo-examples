/*

Reference: https://pkg.go.dev/cmd/cgo

Calling C function pointers is currently not supported, however you can declare
Go variables which hold C function pointers and pass them back and forth between
Go and C. C code may call function pointers received from Go. For example:

*/

package main

/*
#cgo CFLAGS: -I${SRCDIR}/include
#cgo LDFLAGS: ${SRCDIR}/include/helloworld.o

#include <include/helloworld.h>
*/
import "C"
import "fmt"

func main() {

	// get reference to the function pointer via intFunc typedef
	f := C.intFunc(C.fortytwo)

	// pass the intFunc typedef reference back to C for execution
	// and print the returned result in Go
	fmt.Println(int(C.bridge_int_func(f)))

}
