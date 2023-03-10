package main

import (
	"fmt"
	"strconv"
	"strings"
	. "yandex-workshop/utils"
)

func main() {
	Run(Path("01", "1"), listForm)
	Run(Path("01", "2"), listForm)
}

func listForm(fileName string) {
	reader, file := GetReader(fileName)
	defer file.Close()

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
