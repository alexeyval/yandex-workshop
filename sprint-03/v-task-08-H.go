package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(H([]string{"15", "56", "2"}) == "56215")
	fmt.Println(H([]string{"1", "783", "2"}) == "78321")
	fmt.Println(H([]string{"2", "4", "5", "2", "10"}) == "542210")
}

func H(A []string) string {
	sort.Slice(A, func(i, j int) bool { return A[i][0] > A[j][0] })
	return strings.Join(A, "")
}
