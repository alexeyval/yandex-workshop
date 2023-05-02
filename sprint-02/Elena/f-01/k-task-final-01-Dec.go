package main

import (
	"fmt"
)

type Dec struct {
	data  []interface{}
	head  int //извлекаем элемент
	tail  int //добавляем элемент
	max_n int //максимально возможное количество элементов
	size  int //размер очереди
}

func NewDec(n int) *Dec {
	return &Dec{
		data:  make([]interface{}, n),
		head:  0,
		tail:  0,
		max_n: n,
		size:  0,
	}
}

func (d *Dec) push_back(value interface{}) {
	if d.size == d.max_n {
		fmt.Println("error")
	} else {
		d.data[d.tail] = value
		d.tail = (d.tail + 1) % d.max_n
		d.size++
	}
}

func (d *Dec) push_front(value interface{}) {
	if d.size == d.max_n {
		fmt.Println("error")
	} else {
		d.head = (d.head - 1 + d.max_n) % d.max_n
		d.data[d.head] = value
		d.size++
	}
}

func (d *Dec) pop_back() {
	if d.size == 0 {
		fmt.Println("error")
	} else {
		d.tail = (d.tail - 1 + d.max_n) % d.max_n
		fmt.Println(d.data[d.tail])
		d.data[d.tail] = nil
		d.size--
	}

}

func (d *Dec) pop_front() {
	if d.size == 0 {
		fmt.Println("error")
	} else {
		fmt.Println(d.data[d.head])
		d.data[d.head] = nil
		d.head = (d.head + 1) % d.max_n
		d.size--
	}

}

func main() {
	d1 := NewDec(4)
	fmt.Println("--- d1 ---")
	d1.push_front(861)
	d1.push_front(-819)
	d1.pop_back()
	d1.pop_back()

	d2 := NewDec(10)
	fmt.Println("\n--- d2 ---")
	d2.push_front(-855)
	d2.push_front(0)
	d2.pop_back()
	d2.pop_back()
	d2.push_back(844)
	d2.pop_back()
	d2.push_back(823)

	d3 := NewDec(6)
	fmt.Println("\n--- d3 ---")
	d3.push_front(-201)
	d3.push_back(959)
	d3.push_back(102)
	d3.push_front(20)
	d3.pop_front()
	d3.pop_back()
}
