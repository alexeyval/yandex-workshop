package main

import (
	"bufio"
	"fmt"
	"os"
)

func TwoBikes(arr []int, price, index int) int{
	if index == len(arr) {
		return -1
	}
	if arr[index] < price {
		return TwoBikes(arr, price, index + 1)
	}
	return index + 1
}

func main(){
	var N, tmp, price int
	var bikes []int
	buf := bufio.NewReader(os.Stdin)
	fmt.Fscan(buf, &N)
	for i := 0; i < N; i++{
		fmt.Fscan(buf, &tmp)
		bikes = append(bikes, tmp)
	}
	fmt.Fscan(buf, &price)
	bikes1 := TwoBikes(bikes, price, 0)
	bikes2 := TwoBikes(bikes, price * 2, 0)
	fmt.Println(bikes1, bikes2)
}