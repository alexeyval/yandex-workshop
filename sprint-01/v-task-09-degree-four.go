package main

import (
	"fmt"
	. "yandex-workshop/utils"
)

func main() {
	Duration(Path("1"), degreeFour)
	Duration(Path("2"), degreeFour)
	Duration(Path("3"), degreeFour)
}

func degreeFour(fileName string) {
	reader, file := GetReader(fileName)
	defer file.Close()

	var n int
	fmt.Fscan(reader, &n)

	dFour := 4
	for dFour <= n {
		if dFour == n {
			fmt.Println(true)
			return
		}
		dFour *= 4
	}
	fmt.Println(false)
}
