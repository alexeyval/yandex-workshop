package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	. "yandex-workshop/utils"
)

func main() {
	Run(Path("01", "1"), evenAndOddNums)
	Run(Path("01", "2"), evenAndOddNums)
	Run(Path("01", "3"), evenAndOddNums)
}

func evenAndOddNums(fileName string) {
	reader, file := GetReader(fileName)
	defer file.Close()

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
