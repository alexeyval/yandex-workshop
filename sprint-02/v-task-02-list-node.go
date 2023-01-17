package main

import "fmt"

func main() {
	nodes := []ListNode{
		{
			data: "1",
		},
		{
			data: "2",
		},
		{
			data: "3",
		},
		{
			data: "4",
		},
	}
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].next = &nodes[i+1]
	}

	Solution(nil)
	Solution(&nodes[0])
}

type ListNode struct {
	data string
	next *ListNode
}

func Solution(head *ListNode) {
	node := head
	for node != nil {
		fmt.Print(node.data + " -> ")
		node = node.next
	}
	fmt.Print("None")
	fmt.Println()
}
