package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	. "yandexWorkshop/sprint_01/utils"
)

func main() {
	binarySystem("10")
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
		b = strings.Repeat("0", add0)
	case len(a) < len(b):
		a = strings.Repeat("0", add0)
	}

	system := 2 // система счисления
	var sum []string
	overflow := 0
	for i := len(a) - 1; i >= 0; i-- {
		ra, _ := strconv.Atoi(string(a[i]))
		rb, _ := strconv.Atoi(string(b[i]))

		if ra+rb+overflow >= system {
			sum = append(sum, "0")
			overflow = 1
		} else {
			sum = append(sum, strconv.Itoa(ra+rb+overflow))
			overflow = 0
		}
	}
	if overflow != 0 {
		sum = append(sum, strconv.Itoa(overflow))
	}
	for i := 0; i < len(sum)/2; i++ {
		sum[i], sum[len(sum)-1-i] = sum[len(sum)-1-i], sum[i]
	}

	fmt.Println(sum)
}
