package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(moduleFibonacci(4, 1))
	fmt.Println(moduleFibonacci(11, 1))
}

func moduleFibonacci(n, k int) int {
	if n <= 1 {
		return n
	}
	n = (moduleFibonacci(n-1, k) + moduleFibonacci(n-2, k)) % int(math.Pow(10, float64(k)))
	return n
}
