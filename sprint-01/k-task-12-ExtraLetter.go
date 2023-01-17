package main

import (
	"fmt"
	"sort"
)

func SortSymbol(s string) []byte{
	var arr []byte

	for i, _ := range s {
		arr = append(arr, s[i])
	}

	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	return arr
}

func main() {
	var s1, s2 string

	fmt.Scan(&s1, &s2)

	arr1 := SortSymbol(s1)
	arr2 := SortSymbol(s2)

	if len(arr1) < len(arr2) {
		arr1 = append(arr1, '0')
		for i, _ := range arr2 {
			if arr1[i] != arr2[i]{
				fmt.Println(string(arr2[i]))
				break
			} else {
				continue
			}
		}
	}
}