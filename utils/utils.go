package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
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
		fmt.Printf("------------- %v -------------\nВвод:\n%v\n\nВывод:\n",
			path.Base(filepath.ToSlash(fileName)), strings.TrimSpace(readFileContents(fileName)))
		f(fileName)
	}
}

// TestPerformance - выводит информацию о запуске функции + время
func TestPerformance(fileName string, f func(string)) {
	fmt.Println("\n----- тест производительности -----")
	switch fileName {
	case "":
		f(fileName)
		fmt.Println()
	default:

		start := time.Now()
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
			fmt.Println(err)
			os.Exit(3)
		}
		return string(bytes)
	}
	return ""
}

// Path - находит полный путь к файлу
func Path(sprint, fileName string) string {
	if fileName == "" {
		return ""
	}
	pathProject, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	pathProject = filepath.ToSlash(pathProject)
	filePathRun, err := os.Executable()
	fileRun := path.Base(filepath.ToSlash(filePathRun))
	findTask := TaskParse(fileRun)

	mask := "*" + findTask + "*"
	pathFind := path.Join(pathProject, "*"+sprint+"*", "*", mask, "*")
	matches := pathFind
	files, _ := filepath.Glob(matches)
	for _, file := range files {
		findFileName := path.Base(filepath.ToSlash(file))
		if strings.Contains(findFileName, fileName) {
			return file
		}
	}

	fmt.Println(fmt.Errorf(`Не нашёл файл включащий "%v", он должен находиться в папке c названием `+
		"по маске %v.\nПуть для поиска такой:\n%v", fileName, mask, path.Join(pathProject, "*", "*", mask, "*")))
	os.Exit(1)
	return ""
}

// TaskParse по маске ищет task и номер задачи
func TaskParse(fileName string) string {
	wordName := ""
	switch {
	case strings.Contains(fileName, "task"):
		wordName = "task"
	case strings.Contains(fileName, "final"):
		wordName = "final"
	}
	if wordName == "" {
		fmt.Println(fmt.Errorf(`в названии запускаемого файла ` +
			`должно присутвовать task или final`))
		os.Exit(1)
		return ""
	}
	dataRegexp := regexp.MustCompile(wordName + `([\-\.\_])?(\d)*`)

	first := dataRegexp.FindString(fileName)
	first = strings.ReplaceAll(first, "_", "-")
	if first == "" {
		fmt.Println(fmt.Errorf(`в названии запускаемого файла `+
			`должно присутвовать %v с номером, примеры:
task-08
task10
task_10
task.10`, wordName))
		os.Exit(1)
		return ""
	}

	return first
}

func ToInt(n interface{}) int {
	//switch n.(type) {
	//case int:
	//	return n.(int)
	//case string:
	//	s := n.(string)
	//	i, _ := strconv.Atoi(s)
	//	return i
	//}

	return n.(int)
}
