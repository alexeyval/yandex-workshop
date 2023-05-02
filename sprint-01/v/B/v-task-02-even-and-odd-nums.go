package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

}

func evenAndOddNums(fileName string) {
	reader := bufio.NewReader(os.Stdin)

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	line := strings.Split(scanner.Text(), " ")

	flag := false
	outputString := "WIN"
	for i, num := range line {
		n, _ := strconv.Atoi(num)
		if i == 0 {
			flag = n%2 == 0
		} else {
			if flag != (n%2 == 0) {
				outputString = "FAIL"
				break
			}
		}
	}
	fmt.Println(outputString + "\n")
}
