package main

import (
	"fmt"
	"strconv"
)

func main() {
	var N, K, sum int
	var s, tmp string
	var res []int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&tmp)
		s += tmp
	}
	fmt.Scan(&K)
	X, _ := strconv.Atoi(s)
	sum = K + X

	for sum != 0 {
		res = append(res, sum%10)
		sum = sum / 10
	}

	for i := 0; i <= len(res)/2-1; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}

	fmt.Println(res)
}
