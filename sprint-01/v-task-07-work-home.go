package main

import (
	"fmt"
	"strconv"
	. "yandex-workshop/utils"
)

func main() {
	Duration(Path("01", "1"), workHome)
	Duration(Path("01", "2"), workHome)
}

func workHome(fileName string) {
	reader, file := GetReader(fileName)
	defer file.Close()

	var n int
	fmt.Fscan(reader, &n)

	res := ""
	for n > 0 {
		res = strconv.Itoa(n%2) + res
		n /= 2
	}
	fmt.Println(res)
}
