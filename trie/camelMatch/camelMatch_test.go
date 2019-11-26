package camelMatch

import (
	"reflect"
	"testing"
)

func Test_camelMatch(t *testing.T) {
	type args struct {
		queries []string
		pattern string
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{"t-1", args{queries: []string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}, pattern: "FB"}, []bool{true, false, true, true, false}},
		{"t-2", args{queries: []string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}, pattern: "FoBa"}, []bool{true, false, true, false, false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := camelMatch(tt.args.queries, tt.args.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("camelMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
