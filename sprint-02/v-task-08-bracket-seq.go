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
	parentheses := map[int32]int32{']': '[', ')': '(', '}': '{'}

	for _, b := range bracketSeq {
		if _, ok := parentheses[b]; !ok {
			s.Push(b)
			continue
		}
		if s.Top() != parentheses[b] {
			return false
		}
		s.Pop()
	}
	return s.Len == 0
}
