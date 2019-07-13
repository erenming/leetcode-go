package queue

import "fmt"

type ArrayQueue struct {
	q []interface{}
	capacity int
	head int
	tail int
}

func NewArrayQueue(n int) *ArrayQueue {
	return &ArrayQueue{make([]interface{}, n), n, 0, 0, }
}

func (this *ArrayQueue) EnQueue(v interface{}) bool {
	if this.tail == this.capacity {
		return false
	}
	this.q[this.tail] = v
	this.tail++
	return true
}

func (this *ArrayQueue) DeQueue() interface{} {
	if this.tail == this.head {
		return nil
	}
	item := this.q[this.head]
	this.head--
	return item
}

func (this *ArrayQueue) String() string {
	if this.head == this.tail {
		return "empty enqueue"
	}
	result := "head"
	for i := this.head; i <= this.tail - 1; i++ {
		result += fmt.Sprintf("<-%+v", this.q[i])
	}
	result += "<-tail"
	return result
}