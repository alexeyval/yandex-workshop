package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

}

func nearestZero(fileName string) {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(os.Stdout)

	paths := make([]int, n)
	zeroPoz := -1
	countIters := 0
	for i := 0; i < n; i++ {
		countIters++
		scanner.Scan()
		line := scanner.Text()

		num, _ := strconv.Atoi(line)
		s := int(math.Abs(float64(i - zeroPoz)))
		if num != 0 {
			paths[i] = s
		} else /* num == 0 */ {
			s = int(math.Abs(float64(i - zeroPoz)))

			if zeroPoz == -1 {
				for j := 0; j < i; j++ {
					paths[j] = i - j
					countIters++
				}
			} else {
				for j := i - s/2; j < i; j++ {
					paths[j] = i - j
					countIters++
				}
			}
			zeroPoz = i
		}
	}
	for _, v := range paths {
		outputString := strconv.Itoa(v)
		writer.WriteString(outputString + " ")
	}

	writer.Flush()

	fmt.Println("\nn =", n)
	fmt.Println("countIters =", countIters)
}
