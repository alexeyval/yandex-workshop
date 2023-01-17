package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	. "yandex-workshop/utils"
)

func main() {
	Duration(Path("1"), factorization)
	Duration(Path("2"), factorization)
	Duration(Path("3"), factorization)
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
