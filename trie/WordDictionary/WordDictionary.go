package WordDictionary

type WordDictionary struct {
	root *Node
}

type Node struct {
	isEnd    bool
	char     rune
	children [26]*Node
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	root := Node{isEnd: false, char: 0, children: [26]*Node{}}
	return WordDictionary{
		root: &root,
	}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	node := this.root
	for _, c := range word {
		if node.children[c-'a'] == nil {
			node.children[c-'a'] = &Node{isEnd: false, char: c, children: [26]*Node{}}
		}
		node = node.children[c-'a']
	}
	node.isEnd = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return match(word, 0, this.root)
}

func match(word string, index int, node *Node) bool {
	if index == len(word) {
		return node.isEnd
	}
	if word[index] == '.' {
		for _, child := range node.children {
			if child != nil && match(word, index+1, child) {
				return true
			}
		}
		return false
	} else {
		next := node.children[word[index]-'a']
		return next != nil && match(word, index+1, next)
	}

}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
