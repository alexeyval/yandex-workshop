package list

import "fmt"

type List struct {
	Data string
	Next *List
}

func (list *List) Print() {
	node := list
	for node != nil {
		fmt.Print(node.Data + " -> ")
		node = node.Next
	}
	fmt.Print("None")
	fmt.Println()
}

func getNodeByIndex(node *List, index int) *List {
	for index > 0 {
		node = node.Next
		index--
	}

	return node
}

func InsertNode(head *List, index int, value string) *List {
	newNode := &List{Data: value}
	if index == 0 {
		newNode.Next = head
		return newNode
	}

	previousNode := getNodeByIndex(head, index-1)
	newNode.Next = previousNode.Next
	previousNode.Next = newNode

	return head
}
