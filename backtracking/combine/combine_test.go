package combine

import (
	"reflect"
	"testing"
)

func Test_combine(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// {"t-1", args{4, 2}, [][]int{{2,4}, {3, 4}, {2, 3}, {1, 2}, {1, 3}, {1, 4}}},
		{"t-2", args{5, 4}, [][]int{{1,2,3,4},{1,2,3,5},{1,2,4,5},{1,3,4,5},{2,3,4,5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combine(tt.args.n, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combine() = %v, want %v", got, tt.want)
			}
		})
	}
}