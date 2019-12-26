package nextLargerNodes

import (
	"reflect"
	"testing"
)

func Test_nextLargerNodes(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"t-1", args{&ListNode{2, &ListNode{1, &ListNode{5, nil}}}}, []int{5, 5, 0}},
		{"t-1", args{&ListNode{3, &ListNode{3, nil}}}, []int{0, 0}},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextLargerNodes(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextLargerNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
