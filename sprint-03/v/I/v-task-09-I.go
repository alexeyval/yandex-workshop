package main

import (
	"fmt"
	"yandex-workshop/sprint-03/tests"
)

func main() {
	fmt.Println(tests.I([]int{1, 2, 3, 1, 2, 3, 4}, 5, 3))
	fmt.Println(tests.I([]int{1, 1, 1, 2, 2, 3}, 5, 1))
}
