package queue

import (
	"fmt"
)

type Queue struct {
	queue []interface{}
	head  int
	tail  int
	maxN  int
	size  int
}

func New(maxN int) *Queue {
	return &Queue{
		queue: make([]interface{}, maxN),
		head:  0,
		tail:  0,
		maxN:  maxN,
		size:  0,
	}
}

func (q *Queue) Push(x interface{}) {
	if q.size == q.maxN {
		fmt.Println("error")
		return
	}
	q.queue[q.tail] = x
	q.tail = (q.tail + 1) % q.maxN
	q.size++
}

func (q *Queue) Pop() {
	value := q.queue[q.head]
	q.queue[q.head] = nil
	q.head = (q.head + 1) % q.maxN
	q.size--
	fmt.Println(value)
}

func (q *Queue) Peek() {
	if q.size == 0 {
		fmt.Println("None")
		return
	}
	fmt.Println(q.queue[q.head])
	return
}

func (q Queue) Size() int {
	fmt.Println(q.size)
	return q.size
}
