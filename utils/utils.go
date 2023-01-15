package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
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
	switch fileName {
	case "":
		fmt.Println("---------------------------------------")
		f(fileName)
		fmt.Println()
	default:
		start := time.Now()
		fmt.Printf("------------- %v -------------\nВвод:\n%v\n\nВывод:\n",
			path.Base(filepath.ToSlash(fileName)), strings.TrimSpace(readFileContents(fileName)))
		f(fileName)
		d := time.Since(start)
		fmt.Printf("\nВремя выполнения = %v\n\n", d)
	}
}

// readFileContents - считывает всё содержимое файла
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

// Path - находит полный путь файла
func Path(fileName string) string {
	if fileName == "" {
		return ""
	}
	pathProject, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	matches := path.Join(pathProject, "*", "*", "*", "*.txt")
	files, _ := filepath.Glob(matches)
	for _, file := range files {
		if strings.HasSuffix(file, fileName) {
			return file
		}
	}

	fmt.Println(fmt.Errorf("не нашёл файл %v", fileName))
	os.Exit(1)
	return ""
}
