package main

import (
	"fmt"
	"yandex-workshop/sprint-02/queue-list"
)

func main() {
	q1 := queueList.New()
	fmt.Println("--- q1 ---")
	q1.Put(-34)
	q1.Put(-23)
	q1.Get()
	q1.Size()
	q1.Get()
	q1.Size()
	q1.Get()
	q1.Get()
	q1.Put(80)
	q1.Size()

	q2 := queueList.New()
	fmt.Println("\n--- q2 ---")
	q2.Put(-66)
	q2.Put(98)
	q2.Size()
	q2.Size()
	q2.Get()
	q2.Get()
}
