package main

import (
	"fmt"
)

type x struct {
	s string
	a float64
}

func a(val int) {
	for i := 0; i<3;i++ {
		fmt.Println(val+1)
	}
}

func main() {
	f := func() {}
	f = nil
	go f()
}
