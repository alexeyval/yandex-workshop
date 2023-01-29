package main

import (
	"bufio"
	"fmt"
	"strconv"
	"yandex-workshop/sprint-02/stacks"
	. "yandex-workshop/utils"
)

func main() {
	Duration(Path("02", "1"), polishNotation)
	Duration(Path("02", "2"), polishNotation)

	TestPerformance(Path("02", "performance-01"), polishNotation)
}

func polishNotation(fileName string) {
	reader, file := GetReader(fileName)
	defer file.Close()

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	stack := stacks.New()

	type op func(a, b int) int
	operations := map[string]op{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int {
			if b != 0 {
				return a / b
			} else {
				return 0
			}
		},
	}

	for scanner.Scan() {
		s := scanner.Text()
		if _, ok := operations[s]; ok {
			b := ToInt(stack.Top())
			stack.Pop()
			a := ToInt(stack.Top())
			stack.Pop()

			var res = operations[s](a, b)
			stack.Push(res)
		} else {
			n, _ := strconv.Atoi(s)
			stack.Push(n)
		}
	}

	fmt.Println(stack.Top())
}
