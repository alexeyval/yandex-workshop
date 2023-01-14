package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	. "yandexWorkshop/utils"
)

func main() {
	Duration(Path("1.txt"), binarySystem)
	Duration(Path("2.txt"), binarySystem)
}

func binarySystem(fileName string) {
	reader, file := GetReader(fileName)
	defer file.Close()

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

	system := 10 // система счисления
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
