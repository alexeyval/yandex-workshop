package main

import (
	"fmt"
)

type ListNode2 struct {
	data string
	next *ListNode2
	prev *ListNode2
}

func pprint(head *ListNode2) {
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

func Revert2(head *ListNode2) *ListNode2 {
	cur := head
	for cur != nil {
		cur.next, cur.prev = cur.prev, cur.next
		if cur.prev == nil {
			return cur
		}
		cur = cur.prev
	}
	return cur
}

func main() {
	n1 := &ListNode2{data: "1"}
	n2 := &ListNode2{data: "2"}
	n1.next = n2
	n3 := &ListNode2{data: "3"}
	n2.next = n3
	n2.prev = n1
	n4 := &ListNode2{data: "4"}
	n3.next = n4
	n3.prev = n2
	n4.prev = n3

	list := n1
	pprint(Revert2(list))
}
