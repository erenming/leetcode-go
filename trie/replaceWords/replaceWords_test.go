package replaceWords

import "testing"

func Test_replaceWords(t *testing.T) {
	type args struct {
		dict     []string
		sentence string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t-1", args{[]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery"}, "the cat was rat by the bat"},
		{"t-2", args{[]string{"a", "aa", "aaa", "aaaa"}, "a aa a aaaa aaa aaa aaa aaaaaa bbb baba ababa"}, "a a a a a a a a bbb baba a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceWords(tt.args.dict, tt.args.sentence); got != tt.want {
				t.Errorf("replaceWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
