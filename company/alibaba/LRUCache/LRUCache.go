package LRUCache

type Node struct {
	Key   int
	Value int
	Pre   *Node
	Next  *Node
}

type LRUCache struct {
	ht    map[int]*Node
	cap   int
	count int
	head  *Node
	tail  *Node
}

func Constructor(capacity int) LRUCache {
	rv := LRUCache{
		ht:    make(map[int]*Node),
		cap:   capacity,
		count: 0,
		head:  &Node{0, 0, nil, nil},
		tail:  &Node{0, 0, nil, nil},
	}
	rv.head.Next, rv.tail.Pre = rv.tail, rv.head
	return rv
}

func (this *LRUCache) MoveToHead(node *Node) {
	this.Remove(node)
	this.Add(node)
}

func (this *LRUCache) DeleteLast() {
	this.Remove(this.tail.Pre)
}

func (this *LRUCache) Remove(node *Node) {
	node.Pre.Next, node.Next.Pre = node.Next, node.Pre
	delete(this.ht, node.Key)
	this.count--
}

func (this *LRUCache) Add(node *Node) {
	node.Pre, node.Next = this.head, this.head.Next
	this.head.Next, this.head.Next.Pre = node, node
	this.ht[node.Key] = node
	this.count++
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.ht[key]
	if !ok {
		return -1
	}
	this.MoveToHead(node)
	return node.Value
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.ht[key]; ok {
		node.Value = value
		this.MoveToHead(node)
		return
	}

	node := &Node{
		Key:   key,
		Value: value,
		Pre:   nil,
		Next:  nil,
	}
	if this.count == this.cap {
		this.DeleteLast()
	}
	this.Add(node)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(cap);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
