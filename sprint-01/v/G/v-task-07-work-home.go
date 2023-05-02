package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

}

func workHome(fileName string) {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)

	res := ""
	for n > 0 {
		res = strconv.Itoa(n%2) + res
		n /= 2
	}
	fmt.Println(res)
}
