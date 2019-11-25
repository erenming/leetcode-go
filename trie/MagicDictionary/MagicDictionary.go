package MagicDictionary

import "strings"

type MagicDictionary struct {
	trie *Trie
}

/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	t := ConstructTrie()
	return MagicDictionary{
		trie: &t,
	}
}

/** Build a dictionary through a list of words */
func (this *MagicDictionary) BuildDict(dict []string) {
	for _, d := range dict {
		this.trie.Insert(d)
	}
}

/** Returns if there is any word in the trie that equals to the given word after modifying exactly one character */
func (this *MagicDictionary) Search(word string) bool {
	words := strings.Split(word, "")
	for i, w := range strings.Split(word, "") {
		for c := 'a'; c <= 'z'; c++ {
			if w == string(c) {
				continue
			}
			words[i] = string(c)
			if this.trie.Search(strings.Join(words, "")) {
				return true
			}
			words[i] = w
		}
	}
	return false
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
func ConstructTrie() Trie {
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
			node.children[index] = &Node{isEnd: false, char: char, children: [26]*Node{}}
		}
		node = node.children[index]
	}
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	var node = this.root
	for _, char := range word {
		index := char - 'a'
		node = node.children[index]
		if node != nil && node.char == char {
			continue
		} else {
			return false
		}
	}
	return node.isEnd
}