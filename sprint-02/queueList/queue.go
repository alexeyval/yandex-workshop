package queueList

import (
	"fmt"
)

type list struct {
	value interface{}
	next  *list
}

type Queue struct {
	head *list
	tail *list
	size int
}

func New() *Queue {
	tail := &list{}
	head := tail
	return &Queue{
		head: head,
		tail: tail,
		size: 0,
	}
}

func (q *Queue) Put(x interface{}) {
	q.tail.value = x
	q.tail.next = &list{}
	q.tail = q.tail.next
	q.size++
}

func (q *Queue) Get() {
	if q.size == 0 {
		fmt.Println("error")
		return
	}
	head := q.head
	q.head = q.head.next
	head.next = nil
	q.size--
	fmt.Println(head.value)
}

func (q Queue) Size() int {
	fmt.Println(q.size)
	return q.size
}
