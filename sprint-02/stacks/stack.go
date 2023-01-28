package stacks

type list struct {
	value interface{}
	prev  *list
}

type Stack struct {
	top *list
	Len int
}

func (s *Stack) Push(value interface{}) {
	newElem := list{value: value, prev: s.top}
	s.top = &newElem
	s.Len++
}

func (s *Stack) Pop() {
	switch s.Len {
	case 0:
	//	fmt.Println("Error")
	default:
		s.top = s.top.prev
		s.Len--
	}
}

func (s *Stack) Top() (i interface{}) {
	switch s.Len {
	case 0:
	//	fmt.Println("Error")
	default:
		return s.top.value
	}
	return
}
