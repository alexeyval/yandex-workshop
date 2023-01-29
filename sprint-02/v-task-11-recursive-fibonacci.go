package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(recursiveFibonacci(4, nil))
	fmt.Println(recursiveFibonacci(1, nil))

	fmt.Println("тест производительности")
	start := time.Now()
	fmt.Println(recursiveFibonacci(100, nil))
	d := time.Since(start)
	fmt.Printf("\nВремя выполнения = %v\n\n", d)
}

// https://tproger.ru/problems/finding-fibonacci/
func recursiveFibonacci(n int, cache []int) int {
	if cache == nil {
		cache = make([]int, n+1)
	}
	if cache[n] == 0 {
		if n <= 1 {
			cache[n] = 1
		} else {
			cache[n] = recursiveFibonacci(n-1, cache) + recursiveFibonacci(n-2, cache)
		}
	}

	return cache[n]
}
