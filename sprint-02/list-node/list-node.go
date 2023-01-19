package list_node

import "fmt"

type ListNode struct {
	Data string
	Next *ListNode
}

func Solution(head *ListNode) {
	node := head
	for node != nil {
		fmt.Print(node.Data + " -> ")
		node = node.Next
	}
	fmt.Print("None")
	fmt.Println()
}

func getNodeByIndex(node *ListNode, index int) *ListNode {
	for index > 0 {
		node = node.Next
		index--
	}

	return node
}

func InsertNode(head *ListNode, index int, value string) *ListNode {
	newNode := &ListNode{Data: value}
	if index == 0 {
		newNode.Next = head
		return newNode
	}

	previousNode := getNodeByIndex(head, index-1)
	newNode.Next = previousNode.Next
	previousNode.Next = newNode

	return head
}
