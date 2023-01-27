package stacks

import "fmt"

type listMax struct {
	n    int
	max  int
	prev *listMax
}

type StackMax struct {
	top *listMax
	len int
}

func (s *StackMax) Push(n int) {
	elem := listMax{n: n, max: n, prev: s.top}
	if s.len != 0 && elem.max < s.top.max {
		elem.max = s.top.max
	}
	s.top = &elem
	s.len++
}

func (s *StackMax) Pop() {
	switch s.len {
	case 0:
		fmt.Println("Error")
	default:
		s.top = s.top.prev
		s.len--
	}
}

func (s *StackMax) GetMax() {
	switch s.len {
	case 0:
		fmt.Println("None")
	default:
		fmt.Println(s.top.max)
	}
}
