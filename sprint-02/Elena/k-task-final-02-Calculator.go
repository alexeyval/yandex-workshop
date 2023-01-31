package main

import (
	"fmt"
	"strconv"
	"strings"
)

type calc struct {
	data []int
}

func (s *calc) push(x int) {
	s.data = append(s.data, x)
}

func (s *calc) pop() {
	if len(s.data) == 0 {
		fmt.Println("error")
	} else {
		s.data = s.data[:len(s.data)-1]
	}
}

func operations(a, b int, op string) int {
	switch op {
	case "+": return (a + b)
	case "-": return (a - b)
	case "*": return (a * b)
	case "/":
		if b != 0 {
			return a / b
		} else {
			return 0
		}
	}
	return 0
}

func Calc(str string) int {
	var calc calc
	strSplit := strings.Split(str, " ")

	for _, v := range strSplit {
		if v == "*" || v == "/" || v == "+" || v == "-"{
			a := calc.data[len(calc.data) - 1]
			calc.pop()
			b := calc.data[len(calc.data) - 1]
			calc.pop()
			calc.push(operations(a, b, v))
		} else {
			vInt, _ := strconv.Atoi(v)
			calc.push(vInt)
		}
	}
	return calc.data[0]
}

func main(){
	str := "2 1 + 3 *"
	str2 := "7 2 + 4 * 2 +"

	fmt.Println(Calc(str))
	fmt.Println(Calc(str2))
}