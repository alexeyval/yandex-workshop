package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(moduleFibonacci(4, 1))
	fmt.Println(moduleFibonacci(11, 1))
}

func moduleFibonacci(n, k int) int {
	if n <= 1 {
		return n
	}
	s := fmt.Sprint(moduleFibonacci(n-1, k) + moduleFibonacci(n-2, k))
	n, _ = strconv.Atoi(s[len(s)-k:])
	return n
}
