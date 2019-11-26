package topKFrequent

var _count = 0

func topKFrequent(words []string, k int) []string {
	frequency := make(map[string]int)
	for _, w := range words {
		frequency[w] += 1
	}

	// bucket to store same frequency word
	buckets := make([]*Trie, len(words)+1)
	for k, v := range frequency {
		if buckets[v] == nil {
			t := Constructor()
			buckets[v] = &t
		}
		buckets[v].Insert(k)
	}

	res := make([]string, k)
	for i := len(words) - 1; i >= 0 && !fullSlice(res); i-- {
		trie := buckets[i]
		if trie != nil {
			updateWord(trie.root, res)
		}
	}
	_count = 0
	return res
}

func fullSlice(s []string) bool {
	for _, v := range s {
		if v == "" {
			return false
		}
	}
	return true
}

func updateWord(node *Node, res []string) {
	if node == nil {
		return
	}
	if _count > len(res) - 1 {
		return
	}
	if node.isEnd {
		res[_count] = node.word
		_count += 1
	}
	for _, child := range node.children {
		if child != nil {
			updateWord(child, res)
		}
	}
}

type Node struct {
	isEnd    bool
	char     rune
	word     string
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
	node.word = word
}
