package main

import (
	"fmt"
	"math"
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
	fmt.Println(math.MinInt64)
	fmt.Println(math.MinInt64 / 10)
}
