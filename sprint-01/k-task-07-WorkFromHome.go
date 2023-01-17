package main

import (
	"fmt"
)

func main(){
	var N int
	var res []int

	fmt.Scan(&N)
	for N/2 != 0 {
		res = append(res, N % 2)
		N = N/2
	}
	res = append(res, N % 2)
	for i := len(res) - 1; i >= 0; i-- {
		fmt.Print(res[i])
	}
	fmt.Println()
}