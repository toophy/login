// main.go
package main

/*
#include "hehe.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

type Mallocer struct {
}

func (this *Mallocer) Alloc(s int) {
}

func (this *Mallocer) Free(s interface{}) {

}

func (this *Mallocer) relloc(s int, a int) {
	if a < s {
		return
	}

	block := make([]byte, a)

	first := uintptr(unsafe.Pointer(&block[0]))

	count := a/s - 1

	for i := 0; i < count; i++ {
		x := first + uintptr(i*s)
		*(*uintptr)(unsafe.Pointer(x)) = x
	}

	fmt.Print(block)
}

func Tcmalloc() {

	var mc Mallocer
	mc.relloc(10, 100)

	// block_1 := make([]byte, 4096)
	// block_2 := make([]byte, 4096)
	// block_3 := make([]byte, 4096)

	// *(*uintptr)(unsafe.Pointer(&block_1[0])) = uintptr(unsafe.Pointer(&block_2[0]))
	// *(*uintptr)(unsafe.Pointer(&block_2[0])) = uintptr(unsafe.Pointer(&block_3[0]))
	// *(*uintptr)(unsafe.Pointer(&block_3[0])) = uintptr(0)

	// fmt.Print(block_1[:8])
	// fmt.Print("\n")
	// fmt.Print(block_2[:8])
	// fmt.Print("\n")
	// fmt.Print(block_3[:8])
	// fmt.Print("\n")

	// fmt.Println(&block_1[0])
	// fmt.Println(&block_2[0])
	// fmt.Println(&block_3[0])
}

func main() {

	C.ChangePrcName(1)

	Tcmalloc()
}
