package skiplist

import (
	"fmt"
	"math"
	"math/rand"
)

const MAX_LEVEL int = 16

type skipListNode struct {
	v        interface{}
	score    int // 分值，用于排序
	forwards []*skipListNode
	level    int
}

//type skipListLevel struct {
//	forward *skipListNode
//span    uint  // 用以计算元素的rank
//}

type SkipList struct {
	header *skipListNode // 辅助作用
	level  int
	length int
}

func newSkipListNode(v interface{}, score, level int) *skipListNode {
	return &skipListNode{v, score, make([]*skipListNode, level, level), level}
}

func NewSkipList() *SkipList {
	header := newSkipListNode(0, math.MinInt32, MAX_LEVEL)
	return &SkipList{header: header, level: 1, length: 0}
}

func (sl *SkipList) Find(v interface{}, score int) *skipListNode {
	if v == nil || sl.length == 0 {
		return nil
	}
	cur := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for nil != cur.forwards[i] {
			if cur.forwards[i].v == v && cur.forwards[i].score == score {
				return cur.forwards[i]
			} else if score < cur.forwards[i].score { // between cur -> forward (cur.forwards[i])
				break
			} else {
				cur = cur.forwards[i]
			}
		}

	}
	return cur
}

func (sl *SkipList) Insert(v interface{}, score int) bool {
	if v == nil {
		return false
	}
	// find insert position
	cur := sl.header

	update := make([]*skipListNode, MAX_LEVEL, MAX_LEVEL)

	for i := MAX_LEVEL - 1; i >= 0; i-- {
		for nil != cur.forwards[i] {
			if score < cur.forwards[i].score { // down to next level
				update[i] = cur
				break
			} else if cur.forwards[i].v == v {
				cur.forwards[i].v = v
				return true
			}
			cur = cur.forwards[i] // skip cur
		}
		if nil == cur.forwards[i] { // end point
			update[i] = cur
		}
	}

	// get the level of node randomly
	level := 1
	for i := 1; i < MAX_LEVEL; i++ {
		if rand.Int31()%7 == 1 {
			level++
		}
	}

	// splice the node
	node := &skipListNode{v: v, score: score, forwards: make([]*skipListNode, level, level), level: level}

	for i := 0; i < level; i++ {
		//if update[i] != nil {
		node.forwards[i] = update[i].forwards[i]
		update[i].forwards[i] = node
		//}
		//next := update[i].forwards[i]
		//update[i].forwards[i] = node
		//node.forwards[i] = next

	}

	if level > sl.level {
		sl.level = level
	}

	sl.length++
	return true
}

func (sl *SkipList) Delete(v interface{}, score int) bool {
	if v == nil {
		return false
	}
	// find insert position
	cur := sl.header

	update := make([]*skipListNode, MAX_LEVEL, MAX_LEVEL)
	for i := sl.level - 1; i >= 0; i-- {
		for nil != cur.forwards[i] {
			if score == cur.forwards[i].score && v == cur.forwards[i].v {
				update[i] = cur
				break
			}
			cur = cur.forwards[i] // skip cur
		}
	}
	for i, item := range update {
		if item != nil {
			for j := 0; j < item.forwards[i].level; j++ {
				item.forwards[j] = item.forwards[j].forwards[j]
			}
			return true
		}
	}
	return false
}

func (sl *SkipList) String() string {
	str := fmt.Sprintf("\nlevel:%+v, length:%+v\n", sl.level, sl.length)
	for i := sl.level - 1; i >= 0; i-- {
		line := "header"
		cur := sl.header.forwards[i]
		for cur != nil {
			line += fmt.Sprintf(" -> %v", cur.v)
			cur = cur.forwards[i]
		}
		str += fmt.Sprintf("%s\n", line)
	}
	return str
}
