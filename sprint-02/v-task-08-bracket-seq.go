package main

import (
	"fmt"
	. "yandex-workshop/sprint-02/stacks"
)

func main() {
	fmt.Println(isCorrectBracketSeq("{[()]}"))
	fmt.Println(isCorrectBracketSeq("()"))
}

func isCorrectBracketSeq(bracketSeq string) bool {
	s := Stack{}
	m := map[int32]int32{
		'[': ']',
		'(': ')',
		'{': '}',
	}

	for _, b := range bracketSeq {
		switch b {
		case '[', '(', '{':
			s.Push(m[b])
		default:
			if s.Top() != b {
				return false
			}
			s.Pop()
		}
	}
	if s.Len != 0 {
		return false
	}
	return true
}
