package main

type node struct {
	key  int
	val  int
	next *node
}

type MyHashMap struct {
	table []*node
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{make([]*node, 1000)}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	i := key % 1000
	if this.table[i] == nil {
		this.table[i] = &node{key, value, nil}
	} else {
		cur := this.table[i]
		for cur.next != nil && cur.key != key {
			cur = cur.next
		}
		if cur.key == key {
			cur.val = value
		} else {
			cur.next = &node{key, value, nil}
		}
	}
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	i := key % 1000
	cur := this.table[i]
	if cur == nil {
		return -1
	}
	for cur != nil && cur.key != key {
		cur = cur.next
	}
	if cur == nil {
		return -1
	}
	return cur.val
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	i := key % 1000
	cur := this.table[i]
	if cur == nil {
		return
	}
	if cur.key == key {
		this.table[i] = cur.next
	} else {
		var prev *node
		for cur != nil && cur.key != key {
			prev = cur
			cur = cur.next
		}
		if cur == nil {
			return
		}
		prev.next = cur.next
	}
}
