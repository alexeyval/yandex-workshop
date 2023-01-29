package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println(moduleFibonacci(4, 1))
	fmt.Println(moduleFibonacci(11, 1))

	fmt.Println("тест производительности")
	start := time.Now()
	fmt.Println(moduleFibonacci(int(math.Pow(10, 6)), 8))
	d := time.Since(start)
	fmt.Printf("\nВремя выполнения = %v\n\n", d)
}

// https://tproger.ru/problems/finding-fibonacci/
func moduleFibonacci(n int, k float64) int {
	cache := [2]int{0, 1}

	for i := 2; i <= n; i++ {
		cache[i%2] = (cache[0] + cache[1]) % int(math.Pow(10, k))
	}

	return cache[n%2]
}
