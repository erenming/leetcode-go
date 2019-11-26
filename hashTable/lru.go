package hashTable

type DLinkedNode struct {
	Key   int
	Value int
	Prev  *DLinkedNode
	Next  *DLinkedNode
}

type LRUCache struct {
	Capacity int
	Count    int
	Cache    map[int]*DLinkedNode
	Head     *DLinkedNode
	Tail     *DLinkedNode
}

func Constructor(capacity int) LRUCache {
	l := new(LRUCache)
	l.Count = 0
	l.Capacity = capacity
	l.Cache = make(map[int]*DLinkedNode, capacity)
	l.Head = &DLinkedNode{0, 0, nil, nil}
	l.Tail = &DLinkedNode{0, 0, nil, nil}
	l.Head.Next = l.Tail
	l.Tail.Prev = l.Head
	return *l
}

func (this *LRUCache) AddNode(node *DLinkedNode) {
	node.Prev, node.Next = this.Head, this.Head.Next
	this.Head.Next, this.Head.Next.Prev = node, node
	this.Cache[node.Key] = node
}

func (this *LRUCache) RemoveNode(node *DLinkedNode) {
	node.Prev.Next, node.Next.Prev = node.Next, node.Prev
	delete(this.Cache, node.Key)
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.RemoveNode(node)
	this.AddNode(node)
}

func (this *LRUCache) popTail() *DLinkedNode {
	res := this.Tail.Prev
	this.RemoveNode(res)
	return res
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.Cache[key]
	if !ok {
		return -1
	}
	this.moveToHead(node)
	return node.Value
}

func (this *LRUCache) Put(key int, value int) {
	existed, ok := this.Cache[key]
	if !ok {
		node := &DLinkedNode{key, value, nil, nil}
		this.Cache[key] = node
		this.AddNode(node)

		this.Count++
		if this.Count > this.Capacity {
			this.popTail()
			this.Count--
		}
	} else {
		existed.Value = value
		this.moveToHead(existed)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
