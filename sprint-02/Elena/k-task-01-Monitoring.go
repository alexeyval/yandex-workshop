package main

import (
	// "bufio"
	"fmt"
	// "os"
	// "strings"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	array := make([][]byte, k)
	for i := 0; i < k; i++ {
		array[i] = make([]byte, n)
		fmt.Scan(&array[i])
	}

	// var n, m int
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Fscan(reader, &n, &m)
	// str, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	// matrix := make ([n][m] string, string)

	// strSlice := strings.Split(string(str), " ")

	fmt.Println(array)
}
