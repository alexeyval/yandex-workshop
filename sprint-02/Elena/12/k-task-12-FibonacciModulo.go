package main

import (
	"fmt"
	"math"
)

func FibonacciModulo(number int, k float64) int {
	if number == 0 || number == 1 {
		return 1
	} else {
		prev, next := 1, 1
		for i := 0; i < number; i++ {
			tmp := next
			next = prev + next
			prev = tmp
		}
		if prev < 10 {
			return prev
		} else {
			return prev % int(math.Pow(10, k))
		}
		return prev
	}
}

func main() {
	fmt.Println(FibonacciModulo(10, 5))
}
