package list

import "fmt"

type List struct {
	Data string
	Next *List
}

func (l *List) Print() {
	node := l
	for node != nil {
		fmt.Print(node.Data + " -> ")
		node = node.Next
	}
	fmt.Print("None")
	fmt.Println()
}

func (l *List) Delete(index int) *List {
	i, node := 0, l
	for node != nil {
		if index == 0 {
			return node.Next
		}
		if node.Next == nil {
			return l
		}
		if index == i+1 {
			node.Next = node.Next.Next

			return l
		}
		i, node = i+1, node.Next
	}
	return l
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
