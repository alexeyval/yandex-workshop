package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	_ = factorization
}

func factorization(fileName string) *bufio.Writer {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscan(reader, &n)

	d := 2
	for n > 1 {
		switch {
		case n%d == 0:
			_, err := writer.WriteString(strconv.Itoa(d) + " ")
			if err != nil {
				return nil
			}
			n /= d
		default:
			d += 1
		}
	}

	writer.Flush()
	return writer
}
