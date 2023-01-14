package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// GetReader - для чтения с консоли передаётся пустой файл
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

// Duration - выводит информацию о запуске функции
func Duration(fileName string, f func(string)) {
	start := time.Now()
	fmt.Printf("------------- %v -------------"+
		"\nВвод:\n%v\n\nВывод:\n", fileName, strings.TrimSpace(readFileContents(fileName)))
	f(fileName)
	d := time.Since(start)
	fmt.Printf("\nВремя выполнения = %v\n\n", d)
}

func readFileContents(fileName string) string {
	if fileName != "" {
		bytes, err := os.ReadFile(fileName)
		if err != nil {
			log.Fatal(err)
		}
		return string(bytes)
	}
	return ""
}
