package main

import (
	"fmt"
	"unsafe"
)

type x struct {
	s string
	a float64
}

func main() {
	a := x{"hello", 1.0}
	//fmt.Println(unsafe.Sizeof(string("hellosss")))
	fmt.Println(unsafe.Alignof(a))
}