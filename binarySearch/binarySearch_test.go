package binarySearch

import (
	"fmt"
	"testing"
)

func TestFindFirstEqual(t *testing.T) {
	a := []int{1, 3, 5, 8, 8, 8, 9, 10}
	if FindFirstEqual(8, a) != 3 {
		t.Fail()
	}
	if FindFirstEqual(4, a) != -1 {
		t.Fail()
	}
}

func TestFindLastEqual(t *testing.T) {
	a := []int{1, 3, 5, 8, 8, 8, 9, 10}
	if FindLastEqual(8, a	) != 5 {
		t.Fail()
	}
	if FindLastEqual(4, a) != -1 {
		t.Fail()
	}
}

func TestFindFirstLargerEqual(t *testing.T) {
	a := []int{1, 3, 5, 8, 8, 8, 9, 10}
	if FindFirstLargerEqual(6, a) != 3 {
		t.Fail()
	}
	if FindFirstLargerEqual(8, a) != 3 {
		t.Fail()
	}
	if FindFirstLargerEqual(5, a) != 2 {
		t.Fail()
	}
	if FindFirstLargerEqual(12, a) != -1 {
		t.Fail()
	}
}

func TestFindLastLessEqual(t *testing.T) {
	a := []int{1, 3, 5, 8, 8, 8, 10, 11}
	if FindLastLessEqual(6, a) != 2 {
		t.Fail()
	}
	if FindLastLessEqual(8, a) != 5 {
		t.Fail()
	}
	if FindLastLessEqual(4, a) != 1 {
		t.Fail()
	}
	if FindLastLessEqual(0, a) != -1 {
		t.Fail()
	}
}

func TestSearchRange(t *testing.T) {
	a := []int{5,7,7,8,8,10}
	fmt.Println(searchRange(a, 8))
}
