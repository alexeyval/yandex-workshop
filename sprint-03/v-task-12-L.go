package main

import (
	"fmt"
)

func main() {
	type Target struct {
		left  int
		right int
	}
	tests := []struct {
		n      int
		data   []int
		find   int
		target Target
	}{
		{6, []int{1, 2, 4, 4, 6, 8}, 3, Target{3, 5}},
		{6, []int{1, 2, 4, 4, 4, 4}, 3, Target{3, -1}},
		{6, []int{1, 2, 4, 4, 4, 4}, 10, Target{-1, -1}},
	}

	for _, tt := range tests {
		left := leftFindBinary(-1, tt.n, tt.data, tt.find)
		right := leftFindBinary(-1, tt.n, tt.data, tt.find*2)
		if left != tt.target.left || right != tt.target.right {
			fmt.Printf("find got (%v, %v), want (%v, %v)\n",
				left, right, tt.target.left, tt.target.right)
		}
	}
}

func leftFindBinary(l, r int, a []int, s int) int {
	for r-l > 1 {
		middle := (r + l) / 2
		if s > a[middle] {
			l = middle
		} else {
			r = middle
		}
	}
	if r == len(a) {
		return -1
	}
	return r + 1
}
