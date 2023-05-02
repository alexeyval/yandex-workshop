package main

import (
	"fmt"
)

type list struct {
	data interface{}
	next *list
}

type QueueList struct {
	head *list
	tail *list
	size int
}

func NewListQueue() *QueueList {
	tail := &list{}
	head := tail
	return &QueueList{
		head: head,
		tail: tail,
		size: 0,
	}
}

func (q *QueueList) Get() {
	if q.size == 0 {
		fmt.Println("error")
	} else {
		fmt.Println(q.head.data)
		q.head = q.head.next
		q.size--
	}
}

func (q *QueueList) Put(value int) {
	q.tail.data = value
	q.tail.next = &list{}
	q.tail = q.tail.next
	q.size++
}

func (q *QueueList) Size() {
	fmt.Println(q.size)
}

func main() {
	q1 := NewListQueue()
	// q1.Put(-34)
	// q1.Put(-23)
	// q1.Get()
	// q1.Size()
	// q1.Get()
	// q1.Size()
	// q1.Get()
	// q1.Get()
	// q1.Put(80)
	// q1.Size()
	q2 := q1
	q2.Put(-66)
	q2.Put(98)
	q2.Size()
	q2.Size()
	q2.Get()
	q2.Get()

}
