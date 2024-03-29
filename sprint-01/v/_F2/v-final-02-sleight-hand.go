package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

}

func sleightHand(fileName string) {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var k int
	_, _ = fmt.Fscanln(reader, &k)

	table := map[rune]int{}
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		for _, c := range line {
			if c != '.' {
				table[c]++
			}
		}
	}

	points := 0
	for key := range table {
		if table[key] <= k*2 {
			points++
		}
	}

	fmt.Println(points)
}
