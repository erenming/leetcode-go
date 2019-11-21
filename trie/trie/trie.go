package main

import "fmt"

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

/** Returns if the word is in the trie. */
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

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	var node = this.root
	for _, char := range prefix {
		index := char - 'a'
		node = node.children[index]
		if node != nil && node.char == char {
			continue
		} else {
			return false
		}
	}
	return true
}

func main() {
	t := Constructor()
	t.Search("app")
	t.Insert("apple")
	fmt.Println(t.Search("apple"))
	fmt.Println(t.Search("app"))
	fmt.Println(t.StartsWith("app"))
	t.Insert("app")
	fmt.Println(t.Search("app"))
}
