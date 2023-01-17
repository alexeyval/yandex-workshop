package main

import (
	"fmt"
	"sort"
	"strings"
	. "yandex-workshop/utils"
)

func main() {
	Duration(Path("1"), neighbors)
	Duration(Path("2"), neighbors)
}

func neighbors(fileName string) {
	reader, file := GetReader(fileName)
	defer file.Close()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	array := make([][]rune, n)
	for i := 0; i < n; i++ {
		array[i] = make([]rune, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &array[i][j])
		}
	}

	var xi, yj int
	fmt.Fscan(reader, &xi, &yj)

	coords := [][2]int{
		{xi, yj + 1},
		{xi - 1, yj},
		{xi + 1, yj},
		{xi, yj - 1},
	}
	var res []string
	for _, c := range coords {
		if c[0] >= 0 && c[0] < n && c[1] >= 0 && c[1] < m {
			res = append(res, fmt.Sprint(array[c[0]][c[1]]))
		}
	}
	sort.Strings(res)
	fmt.Println(strings.Join(res, ","))
}
