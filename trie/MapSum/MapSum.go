package MapSum

type MapSum struct {
	root *Node
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	root := Node{isEnd: false, char: 0, children: make(map[int32]*Node)}
	return MapSum{
		root: &root,
	}
}

func (this *MapSum) Insert(key string, val int) {
	var node = this.root
	for _, char := range key {
		if node.children[char] == nil {
			node.children[char] = &Node{isEnd: false, char: char, children: make(map[int32]*Node), value: 0}
		}
		node = node.children[char]
	}
	node.isEnd = true
	node.value = val
}

func (this *MapSum) Sum(prefix string) int {
	res := 0
	node := this.root
	for _, c := range prefix {
		if _, ok := node.children[c]; !ok {
			return 0
		} else {
			node = node.children[c]
		}
	}

	return calRest(node, res)
}

func calRest(node *Node, res int) int {
	if node.isEnd {
		res += node.value
	}
	for _, c := range node.children {
		res = calRest(c, res)
	}
	return res
}

type Node struct {
	isEnd    bool
	value    int
	char     rune
	children map[int32]*Node
}

type Trie struct {
	root *Node
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
