package main

import "fmt"

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
