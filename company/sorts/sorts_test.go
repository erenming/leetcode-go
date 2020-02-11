package sorts

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSortSlice(t *testing.T) {
	// bubble sort
	// testSortSlice(t, "bubble")
	// select sort
	// testSortSlice(t, "select")
	// insert sort
	// testSortSlice(t, "insert")
	// merge sort
	// testSortSlice(t, "merge")
	// quick sort
	testSortSlice(t, "quick")
}

func testSortSlice(t *testing.T, method string) {
	fmt.Printf("test %s sort............\n", method)
	type args struct {
		nums   []int
		method string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"t-0", args{[]int{}, method}, []int{}},
		{"t-1", args{[]int{5, 4, 3, 1, 2, 6}, method}, []int{1, 2, 3, 4, 5, 6}},
		{"t-2", args{[]int{64, 34, 25, 12, 22, 11, 90}, method}, []int{11, 12, 22, 25, 34, 64, 90}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortSlice(tt.args.nums, tt.args.method); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
