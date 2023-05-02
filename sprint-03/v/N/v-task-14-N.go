package main

import (
	"fmt"
	"math"
)

type Pair [2]int
type Pairs []Pair

func (p Pair) String() string {
	return fmt.Sprint(p[0], p[1])
}

func (p Pair) Equal(p2 Pair) bool {
	return p[0] == p2[0] && p[1] == p2[1]
}
func (p Pairs) Equal(p2 []Pair) bool {
	if len(p) != len(p2) {
		return false
	}
	for i := range p {
		if !p[i].Equal(p2[i]) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(N([]Pair{
		{7, 8},
		{7, 8},
		{2, 3},
		{6, 10},
	}).Equal([]Pair{
		{2, 3},
		{6, 10},
	}))

	fmt.Println(N([]Pair{
		{2, 3},
		{5, 6},
		{3, 4},
		{3, 4},
	}).Equal([]Pair{
		{2, 4},
		{5, 6},
	}))

	fmt.Println(N([]Pair{
		{1, 3},
		{3, 5},
		{4, 6},
		{5, 6},
		{2, 4},
		{7, 10},
	}).Equal([]Pair{
		{1, 6},
		{7, 10},
	}))
}

func N(pairs []Pair) Pairs {
	if len(pairs) == 1 {
		return pairs
	}

	left := N(pairs[:len(pairs)/2])
	right := N(pairs[len(pairs)/2:])

	i, j := 0, 0

	add := func(l, r Pair, res []Pair) []Pair {
		if r[0] < l[0] {
			l, r = r, l
		}

		if l[1] >= r[0] {
			res = append(res, Pair{l[0], int(math.Max(float64(l[1]), float64(r[1])))})
		} else {
			res = append(res, l)
			res = append(res, r)
		}
		return res
	}

	var res []Pair

	for i < len(left) && j < len(right) {
		l := left[i]
		r := right[j]
		res = add(l, r, res)
		i++
		j++
	}

	for i < len(left) {
		l := left[i]
		r := res[len(res)-1]
		res = res[:len(res)-1]
		res = add(l, r, res)
		i++
	}

	for j < len(right) {
		l := right[j]
		r := res[len(res)-1]
		res = res[:len(res)-1]
		res = add(l, r, res)
		j++
	}

	return res
}
