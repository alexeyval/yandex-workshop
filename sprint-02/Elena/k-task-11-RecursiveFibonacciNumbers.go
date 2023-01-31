package main

import (
	"fmt"
	"time"
)

func Recursia(number int) int{
	if number == 0 || number == 1 {
		return (1)
	} else {
		return (Recursia(number - 1) + Recursia(number - 2))
	}

}

func main(){
	// fmt.Println(Recursia(32))

	fmt.Println("тест производительности")
	start := time.Now()
	fmt.Println(Recursia(32))
	d := time.Since(start)
	fmt.Printf("\nВремя выполнения = %v\n\n", d)
}