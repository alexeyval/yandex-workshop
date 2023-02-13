package main

import (
	"bufio"
	"fmt"
	"os"
)
func isCorrectBracketSeq(str string) bool {
	bracketSet := map[rune]rune{
		'(': ')',
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

func RecursiyaBracketSeq(lenN, n int, prefix string){
	if n != 0 {
		RecursiyaBracketSeq(lenN, n - 1, prefix + "(")
		RecursiyaBracketSeq(lenN, n - 1, prefix + ")")
	}
	if isCorrectBracketSeq(prefix) && len(prefix) == lenN {
		fmt.Println(prefix)
	}
}

func main() {
	var N int
	buf := bufio.NewReader(os.Stdin)
	fmt.Fscan(buf, &N)
	RecursiyaBracketSeq(N * 2, N * 2, "")
}