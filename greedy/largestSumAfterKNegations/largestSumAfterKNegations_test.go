package largestSumAfterKNegations

import "testing"

func Test_largestSumAfterKNegations(t *testing.T) {
	type args struct {
		A []int
		K int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t-1", args{[]int{2, -3, -1, 5, -4}, 2}, 13},
		{"t-2", args{[]int{4, 2, 3}, 1}, 5},
		{"t-3", args{[]int{1, 3, 2, 6, 7, 9}, 21}, 26},
		{"t-4", args{[]int{-2, 5, 0, 2, -2}, 3}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestSumAfterKNegations(tt.args.A, tt.args.K); got != tt.want {
				t.Errorf("largestSumAfterKNegations() = %v, want %v", got, tt.want)
			}
		})
	}
}

