package main

import (
	"fmt"
)

type x struct {
	s string
	a float64
}

func main() {
	a := make([]int, 12)
	//fmt.Println(unsafe.Sizeof(string("hellosss")))
	fmt.Println(len(a))
}