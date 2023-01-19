package main

import "yandex-workshop/sprint-02/list-node"

func main() {
	n3 := &list_node.ListNode{Data: "3"}
	n2 := &list_node.ListNode{Data: "2", Next: n3}
	n1 := &list_node.ListNode{Data: "1", Next: n2}

	list_node.Solution(n1)
	head := list_node.InsertNode(n1, 2, "new-node")
	list_node.Solution(head)
}
