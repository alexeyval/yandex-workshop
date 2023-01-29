package main

import "fmt"

func main() {
	fmt.Println(recursiveFibonacci(4))
	fmt.Println(recursiveFibonacci(1))
}

func recursiveFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return recursiveFibonacci(n-1) + recursiveFibonacci(n-2)
}
