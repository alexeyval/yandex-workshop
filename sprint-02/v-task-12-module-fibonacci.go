package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(moduleFibonacci(4, 1))
	fmt.Println(moduleFibonacci(11, 1))
}

func moduleFibonacci(n int, k float64) int {
	cache := [2]int{0, 1}

	for i := 2; i <= n; i++ {
		cache[i%2] = (cache[0] + cache[1]) % int(math.Pow(10, k))
	}

	return cache[n%2]
}
