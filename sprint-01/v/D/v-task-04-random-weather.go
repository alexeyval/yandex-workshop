package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	_ = randomWeather
}

func randomWeather(fileName string) {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)

	chaoticCount := 0
	temps := make([]int, n)
	for i := 0; i < n+1; i++ {
		if i < n {
			_, _ = fmt.Fscan(reader, &temps[i])
		}
		switch i {
		case 0:
			if n == 1 {
				chaoticCount = 1
			}
		case 1:
			if n > 1 && temps[0] >= temps[1] {
				chaoticCount++
			}
		case n:
			if temps[n-1] >= temps[n-2] {
				chaoticCount++
			}
		default:
			if temps[i-2] <= temps[i-1] && temps[i-1] >= temps[i] {
				chaoticCount++
			}
		}
	}
	fmt.Println(chaoticCount)
}
