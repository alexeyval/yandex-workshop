package utils

import (
	"bufio"
	"fmt"
	"os"
)

func GetReader(fileName string) (reader *bufio.Reader, file *os.File) {
	if fileName != "" {
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		reader = bufio.NewReader(file)
	} else {
		reader = bufio.NewReader(os.Stdin)
	}
	return
}
