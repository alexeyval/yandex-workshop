package main

import (
	"fmt"
	. "yandex-workshop/sprint-02/stack"
)

func main() {
	var s1 StackMax
	fmt.Println("--- s1 ---")
	s1.GetMax()
	s1.Push(7)
	s1.Pop()
	s1.Push(-2)
	s1.Push(-1)
	s1.Pop()
	s1.GetMax()
	s1.GetMax()

	var s2 StackMax
	fmt.Println("\n--- s2 ---")
	s2.GetMax()
	s2.Pop()
	s2.Pop()
	s2.Pop()
	s2.Push(10)
	s2.GetMax()
	s2.Push(-9)
}
