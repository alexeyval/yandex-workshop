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
	res := make([]int, len(A))
	for j, v := range count {
		for i := 1; i <= v; i++ {
			res[index] = j
			index++
		}
	}
	return res
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
