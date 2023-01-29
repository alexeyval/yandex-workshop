package decArray

import "fmt"

type Dec struct {
	queue []interface{}
	head  int
	tail  int
	maxN  int
	size  int
}

func New(maxN int) *Dec {
	return &Dec{
		queue: make([]interface{}, maxN),
		maxN:  maxN,
	}
}

func (q *Dec) PushBack(x interface{}) {
	if q.size == q.maxN {
		fmt.Println("error")
		return
	}
	q.queue[q.tail] = x
	q.tail = (q.tail + 1) % q.maxN
	q.size++
}

func (q *Dec) PushFront(x interface{}) {
	if q.size == q.maxN {
		fmt.Println("error")
		return
	}
	q.head = (q.head - 1 + q.maxN) % q.maxN
	q.queue[q.head] = x
	q.size++
}

func (q *Dec) PopFront() {
	if q.size == 0 {
		fmt.Println("error")
		return
	}
	value := q.queue[q.head]
	q.queue[q.head] = nil
	q.head = (q.head + 1) % q.maxN
	q.size--
	fmt.Println(value)
}

func (q *Dec) PopBack() {
	if q.size == 0 {
		fmt.Println("error")
		return
	}

	q.tail = (q.tail - 1 + q.maxN) % q.maxN
	value := q.queue[q.tail]
	q.queue[q.tail] = nil
	q.size--
	fmt.Println(value)
}

func (q Dec) Size() int {
	fmt.Println(q.size)
	return q.size
}
