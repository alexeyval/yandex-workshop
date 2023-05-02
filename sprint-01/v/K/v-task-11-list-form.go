package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	_ = listForm
}

func listForm(fileName string) {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)

	var firstNumS, secondNumS string
	for i := 0; i < n; i++ {
		var digit string
		fmt.Fscan(reader, &digit)
		firstNumS += digit
	}

	firstNum, _ := strconv.Atoi(firstNumS)
	fmt.Fscan(reader, &secondNumS)
	SecondNum, _ := strconv.Atoi(secondNumS)

	fmt.Println(strings.Join(strings.Split(strconv.Itoa(firstNum+SecondNum), ""), " "))
}
