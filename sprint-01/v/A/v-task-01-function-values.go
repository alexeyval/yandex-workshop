package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	_ = functionValues
}

func functionValues(fileName string) {
	reader := bufio.NewReader(os.Stdin)

	var a, x, b, c float64
	fmt.Fscan(reader, &a, &x, &b, &c)

	outputString := fmt.Sprint(a*x*x + b*x + c)
	fmt.Println(outputString)
}
