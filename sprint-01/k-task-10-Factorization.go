package main

import (
	"fmt"
	"math"
)

func main(){
	var N int
	var res []int

	fmt.Scan(&N)
	tmp := math.Sqrt(float64(N))

	for i := 2; i <= int(tmp); i++{
		for int(N) % i == 0 {
			res = append(res, i);
			N /= i;
		}
	}

	if (N != 1){
		res = append(res, N)
	}

	fmt.Println(res)
}