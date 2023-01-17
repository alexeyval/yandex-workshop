package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	var strInput string
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	str = strings.Trim(str, "\n")
	str = strings.ToLower(str)
	
	re := regexp.MustCompile("[a-z0-9]+")
	strOutput := strings.Join(re.FindAllString(str, -1), "")
	for _, i := range strOutput{
		strInput = string(i) + strInput
	}
	if strInput == strOutput {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}