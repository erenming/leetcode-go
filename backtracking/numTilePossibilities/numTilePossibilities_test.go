package numTilePossibilities

import (
	"testing"
)

func Test_numTilePossibilities(t *testing.T) {
	type args struct {
		tiles string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t-1", args{tiles:"AAB"}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numTilePossibilities(tt.args.tiles); got != tt.want {
				t.Errorf("numTilePossibilities() = %v, want %v", got, tt.want)
			}
		})
	}
}