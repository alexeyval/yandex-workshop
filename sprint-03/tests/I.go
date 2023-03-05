package tests

import (
	"sort"
)

func I(A []int, k int, firstK int) []int {
	B := make([][2]int, k)
	for _, v := range A {
		B[v][0] = v
		B[v][1]++
	}

	sort.Slice(B, func(i int, j int) bool { return B[i][1] > B[j][1] })

	res := make([]int, firstK)
	for i := 0; i < firstK; i++ {
		res[i] = B[i][0]
	}

	return res
}
