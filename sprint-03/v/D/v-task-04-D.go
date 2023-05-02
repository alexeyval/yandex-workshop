package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(D([]int{1, 2}, []int{2, 1, 3}))
	fmt.Println(D([]int{2, 1, 3}, []int{1, 1}))
}

func D(A, B []int) int {
	sort.Ints(A)
	sort.Ints(B)
	i, k := 0, 0

	res := 0
	for i < len(A) && k < len(B) {
		if B[k] >= A[i] {
			res++
		}
		i++
		k++
	}
	return res
}
