/*
This example will return a struct from C and convert it to a native Go struct.
*/

package main

/*
#cgo CFLAGS: -Wall -g

#include <include/example.h>
*/
import "C"
import (
	"cgostruct/model"
	"fmt"
)

func main() {

	fmt.Println("cgo struct example")

	cFoo := C.get_foo()
	goFoo := CFooToGoFoo(cFoo)

	PrintGoFoo(goFoo)

}

func CFooToGoFoo(cFoo *C.Foo) *model.GoFoo {
	goFoo := model.GoFoo{
		Id:   int32(cFoo.id),
		Info: C.GoStringN(cFoo.info, C.int(cFoo.info_len)),
	}

	return &goFoo

}

func PrintGoFoo(goFoo *model.GoFoo) {
	fmt.Printf("id: %d, info: %s\n", goFoo.Id, goFoo.Info)
}
