package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(E(300, []int{999, 999, 999}))
	fmt.Println(E(1000, []int{350, 999, 200}))
}

func E(k int, prices []int) int {
	sort.Ints(prices)
	count := 0
	sum := 0
	for _, v := range prices {
		if sum+v > k {
			return count
		}
		sum += v
		count++
	}
	return count
}
