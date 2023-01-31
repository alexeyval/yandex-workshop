package main

import (
	"fmt"
)

type ListNode struct {
	data string
	next *ListNode
}

func solutionFind(head *ListNode, data string) int {
	index, node := 0, head

	for node != nil {
		if node.data == data {
			return index
		}
		index, node = index+1, node.next
	}
	return -1
}

func main() {
	n1 := &ListNode{data: "1"}
	n2 := &ListNode{data: "2"}
	n1.next = n2
	n3 := &ListNode{data: "2"}
	n2.next = n3
	n4 := &ListNode{data: "4"}
	n3.next = n4

	list := n1
	fmt.Println(solutionFind(list, "2"))
}
