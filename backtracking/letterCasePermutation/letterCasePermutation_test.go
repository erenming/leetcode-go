package letterCasePermutation

import (
	"reflect"
	"testing"
)

func Test_letterCasePermutation(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"t-1", args{"a1b2"}, []string{"a1b2", "a1B2", "A1B2", "A1b2"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := letterCasePermutation(tt.args.S); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("letterCasePermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
