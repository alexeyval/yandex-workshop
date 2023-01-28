package main

import (
	"fmt"
	. "yandex-workshop/sprint-02/stacks"
)

func main() {
	fmt.Println(isCorrectBracketSeq("{[()]}"))
	fmt.Println(isCorrectBracketSeq("()"))
	fmt.Println(isCorrectBracketSeq("()))"))
	fmt.Println(isCorrectBracketSeq("(()"))
}

func isCorrectBracketSeq(bracketSeq string) bool {
	s := Stack{}
	m := map[int32]int32{'[': ']', '(': ')', '{': '}'}

	for _, b := range bracketSeq {
		if _, ok := m[b]; ok {
			s.Push(m[b])
			continue
		}
		if s.Top() != b {
			return false
		}
		s.Pop()
	}
	return s.Len == 0
}
