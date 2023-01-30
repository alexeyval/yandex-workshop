package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(recursiveFibonacci(4))
	fmt.Println(recursiveFibonacci(1))

	fmt.Println("тест производительности")
	start := time.Now()
	fmt.Println(recursiveFibonacci(100))
	d := time.Since(start)
	fmt.Printf("\nВремя выполнения = %v\n\n", d)
}

// https://tproger.ru/problems/finding-fibonacci/
func recursiveFibonacci(n int) int {
	return recursiveFibonacciMain(n, nil)
}

func recursiveFibonacciMain(n int, cache []int) int {
	if cache == nil {
		cache = make([]int, n+1)
	}
	if cache[n] == 0 {
		if n <= 1 {
			cache[n] = n
		} else {
			cache[n] = recursiveFibonacciMain(n-1, cache) + recursiveFibonacciMain(n-2, cache)
		}
	}

	return cache[n]
}
