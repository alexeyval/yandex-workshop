package main

import (
	"fmt"
)

type Queue struct {
	data  []interface{}
	head  int //извлекаем элемент
	tail  int //добавляем элемент
	max_n int //максимально возможное количество элементов
	size  int //размер очереди
}

func MyQueueSized(maxN int) *Queue {
	return &Queue{
		data:  make([]interface{}, maxN),
		head:  0,
		tail:  0,
		max_n: maxN,
		size:  0,
	}
}

func (q *Queue) isEmpty() bool {
	return q.size == 0
}

func (q *Queue) Push(item interface{}) {
	if q.size == q.max_n {
		fmt.Println("Error")
	} else {
		q.data[q.tail] = item
		q.tail = (q.tail + 1) % q.max_n
		q.size++
	}
}

func (q *Queue) Pop() {
	if q.isEmpty() {
		fmt.Println("None")
	} else {
		tmp := q.data[q.head]
		q.data[q.head] = nil
		q.head = (q.head + 1) % q.max_n
		q.size--
		fmt.Println(tmp)
	}
}

func (q *Queue) Peek() {
	if q.isEmpty() {
		fmt.Println("None")
	} else {
		fmt.Println(q.data[q.head])
	}
}

func (q *Queue) Size() {
	fmt.Println(q.size)
}

func main() {
	// q1 := MyQueueSized(2)
	// q1.Peek()
	// q1.Push(5)
	// q1.Push(2)
	// q1.Peek()
	// q1.Size()
	// q1.Size()
	// q1.Push(1)
	// q1.Size()
	// fmt.Println(cap(q1.data))
	q2 := MyQueueSized(1)
	q2.Push(1)
	q2.Size()
	q2.Push(3)
	q2.Size()
	q2.Push(1)
	q2.Pop()
	q2.Push(1)
	q2.Pop()
	q2.Push(3)
	q2.Push(3)
	// fmt.Println("cap", cap(q2.data))

}
