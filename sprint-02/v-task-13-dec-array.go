package main

import (
	"fmt"
	decArray "yandex-workshop/sprint-02/dec-array"
)

func main() {
	d1 := decArray.New(4)
	fmt.Println("--- d1 ---")
	d1.PushFront(861)
	d1.PushFront(-819)
	d1.PopBack()
	d1.PopBack()

	d2 := decArray.New(10)
	fmt.Println("\n--- d2 ---")
	d2.PushFront(-855)
	d2.PushFront(0)
	d2.PopBack()
	d2.PopBack()
	d2.PushBack(844)
	d2.PopBack()
	d2.PushBack(823)

	d3 := decArray.New(6)
	fmt.Println("\n--- d3 ---")
	d3.PushFront(-201)
	d3.PushBack(959)
	d3.PushBack(102)
	d3.PushFront(20)
	d3.PopFront()
	d3.PopBack()
}
