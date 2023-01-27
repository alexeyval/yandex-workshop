package main

import (
	. "yandex-workshop/sprint-02/lists"
)

func main() {
	n4 := &List{Data: "4"}
	n3 := &List{Data: "3", Next: n4}
	n2 := &List{Data: "2", Next: n3}
	n1 := &List{Data: "1", Next: n2}

	n1.Print()
}
