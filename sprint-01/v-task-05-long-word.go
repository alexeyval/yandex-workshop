package main

import (
	"bufio"
	"fmt"
	"strings"
	. "yandex-workshop/utils"
)

func main() {
	Run(Path("01", "1"), longWord)
	Run(Path("01", "2"), longWord)
}

func longWord(fileName string) {
	reader, file := GetReader(fileName)
	defer file.Close()

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
