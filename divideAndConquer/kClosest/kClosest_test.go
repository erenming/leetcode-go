package kClosest

import (
	"reflect"
	"testing"
)

func Test_kClosest(t *testing.T) {
	type args struct {
		points [][]int
		K      int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		//{"t-1", args{[][]int{{3, 3}, {5, -1}, {-2, 4}}, 2}, [][]int{{3, 3}, {-2, 4}}},
		//{"t-2", args{[][]int{{0, 1}, {1, 0}}, 2}, [][]int{{0, 1}, {1, 0}}},
		{"t-3", args{[][]int{{1, 3}, {-2, 2}, {2, -2}}, 2}, [][]int{{-2, 2}, {2, -2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kClosest(tt.args.points, tt.args.K); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
