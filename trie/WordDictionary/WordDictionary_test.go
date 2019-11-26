package WordDictionary

import "testing"

func TestWordDictionary_Search(t *testing.T) {

	type args struct {
		word string
	}
	tests := []struct {
		name   string
		fields []string
		args   args
		want   bool
	}{
		{"t-1", []string{"bad", "dad", "mad"}, args{"pad"}, false},
		{"t-2", []string{"bad", "dad", "mad"}, args{"bad"}, true},
		{"t-3", []string{"bad", "dad", "mad"}, args{".ad"}, true},
		{"t-4", []string{"bad", "dad", "mad"}, args{"b.."}, true},
		{"t-5", []string{}, args{"a"}, false},
		{"t-6", []string{"at", "and", "an", "add", "a"}, args{".at"}, false},
		{"t-7", []string{"at", "and", "an", "add", "a"}, args{"."}, true},
		{"t-8", []string{"at", "and", "an", "add", "a"}, args{"b."}, false},
		{"t-9", []string{"at", "and", "an", "add", "a"}, args{"ad"}, false},


	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Constructor()
			for _, word := range tt.fields {
				this.AddWord(word)
			}
			if got := this.Search(tt.args.word); got != tt.want {
				t.Errorf("WordDictionary.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
