package maxSubArray

import (
	"math"
	"testing"
)

func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t-1", args{[]int{-2,1,-3,4,-1,2,1,-5,4}}, 6},
		{"t-2", args{[]int{-1}}, -1},
		{"t-3", args{[]int{}}, int(math.MinInt64)},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
