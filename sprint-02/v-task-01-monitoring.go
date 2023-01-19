package main

import (
	"fmt"
	. "yandex-workshop/utils"
)

func main() {
	monitoring("")
}

func monitoring(fileName string) {
	reader, file := GetReader(fileName)
	defer file.Close()

	var n, m int
	fmt.Fscan(reader, &n, &m)
	fmt.Println(n, m)

	arr := make([][]int, m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 {
				arr[j] = make([]int, n)
			}
			fmt.Fscan(reader, &arr[j][i])
		}
	}

	for i, _ := range arr {
		fmt.Println(arr[i])
	}
}