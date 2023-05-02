package main

import (
	"bufio"
	"os"
)

func main() {

}

func polishNotation(fileName string) {
	reader := bufio.NewReader(os.Stdin)

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	//stack := stacks.New()

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
	_ = operations

	//for scanner.Scan() {
	//	s := scanner.Text()
	//	if _, ok := operations[s]; ok {
	//		b := ToInt(stack.Top())
	//		stack.Pop()
	//		a := ToInt(stack.Top())
	//		stack.Pop()
	//
	//		var res = operations[s](a, b)
	//		stack.Push(res)
	//	} else {
	//		n, _ := strconv.Atoi(s)
	//		stack.Push(n)
	//	}
	//}
	//
	//fmt.Println(stack.Top())
}
