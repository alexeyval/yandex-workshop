package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	_ = degreeFour
}

func degreeFour(fileName string) {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)

	dFour := 4
	for dFour <= n {
		if dFour == n {
			fmt.Println(true)
			return
		}
		dFour *= 4
	}
	fmt.Println(false)
}
