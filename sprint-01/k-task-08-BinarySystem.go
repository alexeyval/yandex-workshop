package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main(){
	var number1, number2 string
	var rank int //старший разряд
	var res []string

	fmt.Scan(&number1, &number2)

	diff := int(math.Abs(float64(len(number1) - len(number2))))
	if diff > 0 {
		if len(number1) > len(number2) {
			number2 = strings.Repeat("0", diff) + number2
		} else {
			number1 = strings.Repeat("0", diff) + number1
		}
	}
	for i := len(number1) - 1; i >= 0; i-- {
		n1, _ := strconv.Atoi(string(number1[i]))
		n2, _ := strconv.Atoi(string(number2[i]))

		if n1 + n2 + rank >= 2 {
			res = append(res, "0")
			rank = 1
		} else {
			res = append(res, "1")
			rank = 0
		}
	}
	if rank == 1 {
		res = append(res, "1")
	}
	for i := len(res) - 1; i >= 0; i-- {
		fmt.Print(res[i])
	}
	fmt.Println()
}