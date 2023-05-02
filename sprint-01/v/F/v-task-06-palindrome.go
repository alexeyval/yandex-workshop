package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {

}

func palindrome(fileName string) {
	reader := bufio.NewReader(os.Stdin)

	var s string
	for n := 0; n < reader.Size(); n++ {
		r, _, _ := reader.ReadRune()
		if unicode.IsLetter(r) {
			s += strings.ToLower(string(r))
		}
	}

	isPalindrome := true
	for i, j := 0, len(s)-1; i < len(s)/2; {
		if s[i] != s[j] {
			isPalindrome = false
			break
		}
		i++
		j--
	}
	fmt.Println(isPalindrome)
}
