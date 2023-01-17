package main

import (
	"fmt"
)

func main(){
	var N, tmp, counter int
	var slice []int

	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&tmp)
		slice = append(slice, tmp)
	}
	for j := 0 ; j < N; j++ {
		if j == 0 {
			if slice[j] > slice[j + 1]{
				counter++
			}
		} else if j == N - 1 {
			if slice[j] > slice[j - 1]{
				counter++
			}
		} else {
			if slice[j] > slice[j - 1] && slice[j] > slice[j + 1]{
				counter++
			}
		}
	}
	fmt.Println(counter)
}
