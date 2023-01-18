package main

func main() {
	nodes := []ListNode{
		{
			data: "1",
		},
		{
			data: "2",
		},
		{
			data: "3",
		},
		{
			data: "4",
		},
	}
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].next = &nodes[i+1]
	}

	Solution(nil)
	Solution(&nodes[0])
}
