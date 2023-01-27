package main

import (
	"fmt"
	. "yandex-workshop/utils"
)

func main() {
	Duration(Path("01", "1"), functionValues)
	Duration(Path("01", "2"), functionValues)
}

func functionValues(fileName string) {
	reader, file := GetReader(fileName)
	defer file.Close()

	var a, x, b, c float64
	fmt.Fscan(reader, &a, &x, &b, &c)

	outputString := fmt.Sprint(a*x*x + b*x + c)
	fmt.Println(outputString)
}
