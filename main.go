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
	a := []int{0, 1}
	a = append(a, nil...)
	fmt.Println(a)
}
