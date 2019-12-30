package convertor

import (
	"reflect"
	"testing"
)

func TestTwoDimensionalIntArray(t *testing.T) {
	type args struct {
		source string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"t-1", args{`[[2,4],[3,4],[2,3],[1,2],[1,3],[1,4],]`}, [][]int{{2, 4}, {3, 4}, {2, 3}, {1, 2}, {1, 3}, {1, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoDimensionalIntArray(tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoDimensionalIntArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
