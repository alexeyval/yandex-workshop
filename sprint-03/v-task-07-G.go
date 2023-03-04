package main

import "fmt"

func main() {
	fmt.Println(Equal(G([]int{0, 2, 1, 2, 0, 0, 1}, 3), []int{0, 0, 0, 1, 1, 2, 2}))
	fmt.Println(Equal(G([]int{2, 1, 2, 0, 1}, 3), []int{0, 1, 1, 2, 2}))
	fmt.Println(Equal(G([]int{2, 1, 1, 2, 0, 2}, 3), []int{0, 1, 1, 2, 2, 2}))
}

func G(A []int, k int) []int {
	count := make([]int, k)
	for _, v := range A {
		count[v]++
	}

	index := 0
	for i, v := range count {
		for amount := 1; amount <= v; amount++ {
			A[index] = i
			index++
		}
	}
	return A
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
