package findNumberOfLIS

import (
	"testing"
)

func Test_findNumberOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t-1", args{nums: []int{1, 3, 5, 4, 7}}, 2},
		{"t-2", args{nums: []int{1, 2, 4, 3, 5, 4, 7, 2}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumberOfLIS(tt.args.nums); got != tt.want {
				t.Errorf("findNumberOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
