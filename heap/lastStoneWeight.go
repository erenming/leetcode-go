package heap

import (
	"github.com/Workiva/go-datastructures/queue"
	"reflect"
)

func lastStoneWeight(stones []int) int {
	q := queue.NewPriorityQueue(0, true)
	a, _ := q.Get(1)
	return int(reflect.ValueOf(a[0]).Int())
}
