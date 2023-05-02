package main

import "fmt"

func main() {
	maxP := -1
	F([]int{6, 3, 3, 2}, nil, 3, &maxP)
	fmt.Println(maxP)
	maxP = -1
	F([]int{5, 3, 7, 2, 8, 3}, nil, 3, &maxP)
	fmt.Println(maxP)
}

func F(A, t []int, k int, maxP *int) {
	if k == 0 {
		iMax := 0
		for i := 1; i < 3; i++ {
			if t[iMax] < t[i] {
				iMax = i
			}
		}

		if t[iMax] < t[(iMax+1)%3]+t[(iMax+2)%3] {
			if *maxP == -1 {
				*maxP = t[0] + t[1] + t[2]
			}
			if *maxP < t[0]+t[1]+t[2] {
				*maxP = t[0] + t[1] + t[2]
			}
		}
		return
	}

	if t == nil {
		t = make([]int, 0, k)
	}

	for i, v := range A {
		if A[i] == -1 {
			continue
		}
		t = append(t, v)
		A[i] = -1
		F(A, t, k-1, maxP)
		t = t[:len(t)-1]
		A[i] = v
	}
}
