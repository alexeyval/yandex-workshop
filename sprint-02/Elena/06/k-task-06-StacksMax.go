package main

import (
	"fmt"
)

type stack struct {
	data []int
}

func (s *stack) push(x int) {
	s.data = append(s.data, x)
}

func (s *stack) pop() {
	if len(s.data) == 0 {
		fmt.Println("error")
	} else {
		s.data = s.data[:len(s.data)-1]
	}
}

func (s *stack) get_max() {
	if len(s.data) == 0 {
		fmt.Println("None")
	} else {
		max := s.data[0]
		for _, v := range s.data {
			if v > max {
				max = v
			}
		}
		fmt.Println(max)
	}
}

func main() {

	var tmp stack

	tmp.get_max()
	tmp.push(7)
	tmp.pop()
	tmp.push(-2)
	tmp.push(-1)
	tmp.pop()
	tmp.get_max()
	tmp.get_max()
}
