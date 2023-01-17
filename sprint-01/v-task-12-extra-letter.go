package main

import (
	"fmt"
	"sort"
	. "yandex-workshop/utils"
)

func main() {
	Duration(Path("1"), extraLetter)
	Duration(Path("2"), extraLetter)
	Duration(Path("3"), extraLetter)
}

func extraLetter(fileName string) {
	reader, file := GetReader(fileName)
	defer file.Close()

	var s1, s2 string
	fmt.Fscan(reader, &s1)
	fmt.Fscan(reader, &s2)

	var arr []byte
	f := func(i, j int) bool { return arr[i] < arr[j] }

	arr = []byte(s1)
	sort.Slice(arr, f)
	s1Sort := string(arr)

	arr = []byte(s2)
	sort.Slice(arr, f)
	s2Sort := string(arr)

	var findSymbol uint8
OuterLoop:
	for i := 0; i != -1; i++ {
		switch {
		case i >= len(s1):
			findSymbol = s2Sort[i]
			break OuterLoop
		case i >= len(s2):
			findSymbol = s1Sort[i]
			break OuterLoop
		case s1Sort[i] != s2Sort[i]:
			findSymbol = s2Sort[i]
			break OuterLoop
		}
	}

	fmt.Println(string(findSymbol))
}
