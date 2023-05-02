package main

import "fmt"

func main() {
	fmt.Println(C("abc", "ahbgdcu"))
	fmt.Println(!C("abcp", "ahpc"))
}

func C(s, t string) bool {
	if len(s) == 0 {
		return true
	}

	i := 0
	for j := range t {
		if s[i] == t[j] {
			i++
			if i == len(s) {
				return true
			}
		}
	}
	return false
}
