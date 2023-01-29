package main

import (
	"fmt"
	"yandex-workshop/sprint-02/queue-array"
)

func main() {
	q1 := queueArray.New(2)
	fmt.Println("--- q1 ---")
	q1.Peek()
	q1.Push(5)
	q1.Push(2)
	q1.Peek()
	q1.Size()
	q1.Size()
	q1.Push(1)
	q1.Size()

	q2 := queueArray.New(1)
	fmt.Println("\n--- q2 ---")
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
}
