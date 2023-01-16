package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main(){
	var N, counter int
	fmt.Scan(&N)
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	textTrim := strings.Trim(text, "\n")
	textSplit := strings.Split(textTrim, " ")

	for _, word := range textSplit {
		if len(word) > counter {
			counter = len(word)
		} 
	}

	fmt.Println(counter)
}