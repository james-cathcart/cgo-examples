/*

Get *char from C into a Go string

*/

package main

/*
#cgo CFLAGS: -Wall -g

#include <stdlib.h>

char* GetText() {
	return "Some text from C";
}

*/
import "C"
import (
	"fmt"
)

func main() {

	cMessage := C.GetText()
	goMessage := C.GoString(cMessage)
	fmt.Printf("Message from C: %s\n", goMessage)
}
