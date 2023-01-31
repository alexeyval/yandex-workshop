package main

import (
	"fmt"
)

type ListNode struct {
	data string
	next *ListNode
}

func solutionsPrint(head *ListNode) {
	node := head
	prefix := ""
	for node != nil {
		if node != head {
			prefix = " -> "
		}
		fmt.Print(prefix, node.data)
		node = node.next
	}
	fmt.Println()
}

func getNodeByIndex(n *ListNode, index int) *ListNode {
	for index != 0 {
		n = n.next
		index--
	}
	return n
}

func solutionsDelete(n *ListNode, index int) *ListNode {
	if index <= 0 {
		return n
	}
	if index == 1 {
		return n.next
	}
	tmp := getNodeByIndex(n, index-1)
	tmp.next = tmp.next.next
	return n
}

func main() {
	n1 := &ListNode{data: "1"}
	n2 := &ListNode{data: "2"}
	n1.next = n2
	n3 := &ListNode{data: "3"}
	n2.next = n3
	n4 := &ListNode{data: "4"}
	n3.next = n4

	list := n1

	solutionsPrint(solutionsDelete(list, 4))
}
