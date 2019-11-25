package MagicDictionary

import "testing"

func createTrie(dict []string) *Trie {
	trie := ConstructTrie()
	for _, d := range dict {
		trie.Insert(d)
	}
	return &trie
}


func TestMagicDictionary_Search(t *testing.T) {
	trie1 := createTrie([]string{"hello", "leetcode"})
	trie2 := createTrie([]string{"hello", "hallo", "leetcode"})

	type fields struct {
		trie *Trie
	}
	type args struct {
		word string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"t-1", fields{trie1}, args{"hello"}, false},
		{"t-2", fields{trie1}, args{"hhllo"}, true},
		{"t-3", fields{trie2}, args{"hello"}, true},
		{"t-4", fields{trie2}, args{"hallo"}, true},
		{"t-5", fields{trie2}, args{"hell"}, false},
		{"t-6", fields{trie2}, args{"leetcoded"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MagicDictionary{
				trie: tt.fields.trie,
			}
			if got := this.Search(tt.args.word); got != tt.want {
				t.Errorf("MagicDictionary.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrie_Search(t *testing.T) {
	type fields struct {
		root *Node
	}
	type args struct {
		word string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Trie{
				root: tt.fields.root,
			}
			if got := this.Search(tt.args.word); got != tt.want {
				t.Errorf("Trie.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
