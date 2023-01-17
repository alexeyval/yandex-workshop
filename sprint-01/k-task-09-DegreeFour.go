package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64

	fmt.Scan(&num)

	i := math.Log(num) / math.Log(4)

	_, frac := math.Modf(i) // определяем, есть ли дробная часть

	if frac == 0 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
