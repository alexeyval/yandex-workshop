package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	_ = longWord
}

func longWord(fileName string) {
	reader := bufio.NewReader(os.Stdin)

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	scanner.Scan()
	line := scanner.Text()

	words := strings.Split(line, " ")
	maxWord := words[0]
	for _, word := range words {
		if len(word) > len(maxWord) {
			maxWord = word
		}
	}
	fmt.Printf("%v\n%v\n", maxWord, len(maxWord))
}
