package main

import "fmt"

type stackEffective struct {
	data int
	max  int
	prev *stackEffective
}

type stackMax struct {
	cur *stackEffective
}

func (n *stackMax) push(x int) {
	if n.cur == nil {
		elem := &stackEffective{data: x, max: x}
		n.cur = elem
	} else {
		tmpMax := n.cur.max
		if n.cur.max < x {
			tmpMax = x
		}
		elem := &stackEffective{
			data: x,
			max:  tmpMax,
			prev: n.cur,
		}
		n.cur = elem
	}
}

func (n *stackMax) pop() {
	if n.cur == nil {
		fmt.Println("error")
	} else {
		n.cur = n.cur.prev
	}
}

func (n *stackMax) get_max() {
	if n.cur == nil {
		fmt.Println("none")
	} else {
		fmt.Println(n.cur.max)
	}

}

func main() {
	var d stackMax
	d.pop()
	d.pop()
	d.push(4)
	d.push(-5)
	d.push(7)
	d.pop()
	d.pop()
	d.get_max()
	d.pop()
	d.get_max()
}
