package main

import (
	. "yandex-workshop/sprint-02/list-node"
)

func main() {
	nodes := []ListNode{
		{
			Data: "1",
		},
		{
			Data: "2",
		},
		{
			Data: "3",
		},
		{
			Data: "4",
		},
	}
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = &nodes[i+1]
	}

	Solution(nil)
	Solution(&nodes[0])
}
