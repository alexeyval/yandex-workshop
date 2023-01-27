package main

import (
	"fmt"
	"strings"
	"unicode"
	. "yandex-workshop/utils"
)

func main() {
	Duration(Path("01", "1"), palindrome)
	Duration(Path("01", "2"), palindrome)
}

func palindrome(fileName string) {
	reader, file := GetReader(fileName)
	defer file.Close()

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
