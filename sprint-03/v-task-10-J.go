package main

import (
	"fmt"
	"strings"
)

func main() {
	N := 5
	A := []int{4, 3, 9, 2, 1}
	sortBubble(N, A)

	fmt.Println()

	A = []int{12, 8, 9, 10, 11}
	sortBubble(N, A)
}

func sortBubble(N int, A []int) {
	for i := 1; i < N; i++ {
		replaces := false
		for j := 0; j < N-i; j++ {
			if A[j] > A[j+1] {
				A[j], A[j+1] = A[j+1], A[j]
				replaces = true
			}
		}
		if replaces {
			fmt.Println(strings.Trim(fmt.Sprint(A), "[]"))
		}
	}
}
