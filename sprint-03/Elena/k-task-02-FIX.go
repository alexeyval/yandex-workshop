package main

import (
	"bufio"
	"fmt"
	"os"
)

/*

Этот код просто выстрел себе в ногу
Тут используется рекурсия для того,
чтобы просто сложить две строки? Зачем?
Это как реализовать сложение двух чисел через рекурсию

Тоже самое решение я вписал просто в обычный цикл
в main:
for j := 0; j < len(dict[num2]); j++ {
	fmt.Print(dict[num1][j]+dict[num2][j], " ")
}
аналогично этому:
for i := 0; i < len(slice); i++ {
	fmt.Print(s+slice[i], " ")
}

func Combinations(n int, s string, slice []string) {
	if n != 0 {
		for i := 0; i < len(slice); i++ {
			Combinations(n-1, s+slice[i], slice)
		}
	} else {
		fmt.Print(s, " ")
	}
}*/

func main() {
	dict := make([][]string, 8)
	for i := range dict {
		dict[i] = make([]string, 4)
	}
	dict = [][]string{
		2: {"a", "b", "c"},
		3: {"d", "e", "f"},
		4: {"g", "h", "i"},
		5: {"j", "k", "l"},
		6: {"m", "n", "o"},
		7: {"p", "q", "r", "s"},
		8: {"t", "u", "v"},
		9: {"w", "x", "y", "z"},
	}
	var N int
	buf := bufio.NewReader(os.Stdin)
	fmt.Fscan(buf, &N)
	num1 := N / 10
	num2 := N % 10
	for i := 0; i < len(dict[num1]); i++ {
		for j := 0; j < len(dict[num2]); j++ {
			fmt.Print(dict[num1][j]+dict[num2][j], " ")
		}
	}
	fmt.Println()
}
