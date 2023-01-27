package main

import (
	. "yandex-workshop/sprint-02/lists"
)

func main() {
	n4 := &List2{Data: "4"}
	n3 := &List2{Data: "3", Next: n4}
	n4.Prev = n3
	n2 := &List2{Data: "2", Next: n3}
	n3.Prev = n2
	n1 := &List2{Data: "1", Next: n2}
	n2.Prev = n1

	n1.Print()
	n4.PrintReserve()

	n1.Reserve().Print()
}
