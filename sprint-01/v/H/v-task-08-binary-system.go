package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

}

func binarySystem(fileName string) {
	reader := bufio.NewReader(os.Stdin)

	var a, b string
	fmt.Fscan(reader, &a)
	fmt.Fscan(reader, &b)

	add0 := int(math.Abs(float64(len(a) - len(b))))
	switch {
	case len(a) > len(b):
		b = strings.Repeat("0", add0) + b
	case len(a) < len(b):
		a = strings.Repeat("0", add0) + a
	}

	system := 2 // система счисления
	var sum []string
	overflow := 0
	for i := len(a) - 1; i >= 0; i-- {
		ra, _ := strconv.Atoi(string(a[i]))
		rb, _ := strconv.Atoi(string(b[i]))

		sumR := ra + rb + overflow
		sum = append(sum, strconv.Itoa(sumR%system))
		overflow = sumR / system
	}
	if overflow != 0 {
		sum = append(sum, strconv.Itoa(overflow))
	}
	for i := 0; i < len(sum)/2; i++ {
		sum[i], sum[len(sum)-1-i] = sum[len(sum)-1-i], sum[i]
	}

	fmt.Println(sum)
}
