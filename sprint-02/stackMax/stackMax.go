package stackMax

import "fmt"

type list struct {
	n    int
	max  int
	prev *list
}

type StackMax struct {
	top *list
	len int
}

func (s *StackMax) Push(n int) {
	elem := list{n: n, max: n, prev: s.top}
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
