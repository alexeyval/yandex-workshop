package main

import (
	"fmt"
	"strconv"
	"strings"
	. "yandex-workshop/utils"
)

func main() {
	Duration(Path("1"), listForm)
	Duration(Path("2"), listForm)
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
