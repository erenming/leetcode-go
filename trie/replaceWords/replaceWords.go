package replaceWords

import "strings"

func replaceWords(dict []string, sentence string) string {
	t := Constructor()
	for _, d := range dict {
		t.Insert(d)
	}
	sentences := strings.Split(sentence, " ")
	res := make([]string, len(sentences))
	for i, word := range sentences {
		newWord, ok := t.ReplaceString(word)
		if ok {
			res[i] = newWord
		} else {
			res[i] = word
		}
	}
	return strings.Join(res, " ")
}

type Node struct {
	isEnd    bool
	char     rune
	children [26]*Node
}

type Trie struct {
	root *Node
}

/** Initialize your data structure here. */
func Constructor() Trie {
	root := Node{isEnd: false, char: 0, children: [26]*Node{}}
	return Trie{
		root: &root,
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	var node = this.root
	for _, char := range word {
		index := char - 'a'
		if node.children[index] == nil {
			newNode := &Node{isEnd: false, char: char, children: [26]*Node{}}
			node.children[index] = newNode
			node = newNode
		} else {
			node = node.children[index]
		}
	}
	node.isEnd = true
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) ReplaceString(word string) (s string, matched bool) {
	res := ""
	node := this.root
	for _, c := range word {
		index := c - 'a'
		child := node.children[index]
		if child != nil {
			res += string(c)
			if child.isEnd {
				return res, true
			} else {
				node = child
			}
		} else {
			break
		}
	}
	return res, node.isEnd
}
