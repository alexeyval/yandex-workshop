package main

import (
	"fmt"
)

type ListNode1 struct {
	data string
	next *ListNode
}

func solutions(head *ListNode) {
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

func revers(n *ListNode) *ListNode {
	//два указателя на текущий и предыдущий
	var cur = n
	var prev *ListNode

	for cur != nil {
		//сохраним ссылку на следующий элемент, чтобы заменить ее на предыдущий
		next := cur.next
		cur.next = prev
		prev = cur
		cur = next
	}
	return prev
}

func main() {
	//инициализация листов
	n1 := &ListNode{data: "1"}
	n2 := &ListNode{data: "2"}
	n1.next = n2
	n3 := &ListNode{data: "3"}
	n2.next = n3
	n4 := &ListNode{data: "4"}
	n3.next = n4

	nodeExsempl := n1
	solutions(nodeExsempl)
	n := revers(nodeExsempl)
	solutions(n)
}
