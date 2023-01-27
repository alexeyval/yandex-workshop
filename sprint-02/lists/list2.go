package lists

import "fmt"

type List2 struct {
	Data interface{}
	Prev *List2
	Next *List2
}

func (l *List2) Reserve() *List2 {
	node := l
	for node != nil {
		node.Next, node.Prev = node.Prev, node.Next
		if node.Prev == nil {
			return node
		}
		node = node.Prev
	}
	return node
}

func (l *List2) Print() {
	node := l
	for node != nil {
		fmt.Printf("%v -> ", node.Data)
		node = node.Next
	}
	fmt.Print("None")
	fmt.Println()
}

func (l *List2) PrintReserve() {
	node := l
	for node != nil {
		fmt.Printf("%v -> ", node.Data)
		node = node.Prev
	}
	fmt.Print("None")
	fmt.Println()
}
