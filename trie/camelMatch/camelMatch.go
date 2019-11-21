package camelMatch

import (
	"strings"
	"unicode"
)

func camelMatch(queries []string, pattern string) []bool {
	t := Constructor()
	for _, q := range queries {
		t.Insert(q)
	}
	qm := make(map[string]bool)
	patterns := strings.Split(pattern, "")

	Find(t.root, 0, patterns, make([]string, 0), qm)
	res := make([]bool, len(queries))
	for i, q := range queries {
		res[i] = qm[q]
	}
	return res
}

func Find(node *Node, pIndex int, patterns []string, curr []string, qm map[string]bool) {
	if pIndex >= len(patterns) {
		if node.isEnd == true {
			qm[strings.Join(curr, "")] = true
		} else {
			for _, char := range node.children {
				if unicode.IsLower(char.char) {
					Find(char, pIndex, patterns, append(curr, string(char.char)), qm)
				}
			}
		}
	} else {
		for _, char := range node.children {
			if patterns[pIndex] == string(char.char) {
				Find(char, pIndex+1, patterns, append(curr, string(char.char)), qm)
			} else if unicode.IsLower(char.char) {
				Find(char, pIndex, patterns, append(curr, string(char.char)), qm)
			}
		}
	}
}

type Node struct {
	isEnd    bool
	char     rune
	children map[int32]*Node
}

type Trie struct {
	root *Node
}

/** Initialize your data structure here. */
func Constructor() Trie {
	root := Node{isEnd: false, char: 0, children: make(map[int32]*Node)}
	return Trie{
		root: &root,
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	var node = this.root
	for _, char := range word {
		index := char
		if node.children[index] == nil {
			newNode := &Node{isEnd: false, char: char, children: make(map[int32]*Node)}
			node.children[index] = newNode
			node = newNode
		} else {
			node = node.children[index]
		}
	}
	node.isEnd = true
}
