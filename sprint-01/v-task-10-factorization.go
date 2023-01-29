package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	. "yandex-workshop/utils"
)

func main() {
	Run(Path("01", "1"), factorization)
	Run(Path("01", "2"), factorization)
	Run(Path("01", "3"), factorization)
}

func factorization(fileName string) {
	reader, file := GetReader(fileName)
	writer := bufio.NewWriter(os.Stdout)
	defer file.Close()

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
