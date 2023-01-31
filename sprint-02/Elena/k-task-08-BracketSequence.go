package main

import (
	"fmt"
)

func isCorrectBracketSeq(str string) bool {
	bracketSet := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	bracketSetBool := make(map[rune]bool)

	for _, v := range bracketSet {
		bracketSetBool[v] = true
	}

	var bracket []rune

	for _, v := range str {
		if _, ok := bracketSet[v]; ok {
			bracket = append(bracket, bracketSet[v])
		} else if len(bracket) != 0 && v == bracket[len(bracket)-1] {
			bracket = bracket[:len(bracket)-1]
		} else if !bracketSetBool[v] {
			continue
		} else {
			return false
		}
	}
	return len(bracket) == 0
}

func main() {
	fmt.Println(isCorrectBracketSeq("{dgdg}"))
}
