package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

}

func factorization(fileName string) {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscan(reader, &n)

	d := 2
	for n > 1 {
		switch {
		case n%d == 0:
			writer.WriteString(strconv.Itoa(d) + " ")
			n /= d
		default:
			d += 1
		}
	}

	writer.Flush()
}
