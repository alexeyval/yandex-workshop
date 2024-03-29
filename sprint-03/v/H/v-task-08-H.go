package main

import (
	"fmt"
	"sort"
	"strings"
)

type Test bool

func (f Test) String() string {
	res := "OK"
	if !f {
		res = "FAIL"
	}
	return res
}

func main() {
	fmt.Println(Test(H([]string{"15", "56", "2"}) == "56215" &&
		H([]string{"1", "783", "2"}) == "78321" &&
		H([]string{"2", "4", "5", "2", "J"}) == "542210"))
}

func H(A []string) string {
	sort.Slice(A, func(i, j int) bool { return A[i] > A[j] })
	return strings.Join(A, "")
}
