package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

}

func monitoring(fileName string) {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(reader, &n, &m)

	arr := make([][]string, m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 {
				arr[j] = make([]string, n)
			}
			fmt.Fscan(reader, &arr[j][i])
		}
	}

	for i, _ := range arr {
		fmt.Println(strings.Join(arr[i], " "))
	}
}
