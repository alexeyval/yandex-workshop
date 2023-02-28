package main

import "fmt"

func main() {
	for _, test := range []struct {
		in   int
		want string
	}{
		{3, "((()))\n(()())\n(())()\n()(())\n()()()\n"},
		{2, "(())\n()()\n"},
	} {
		var got string
		A(test.in*2, "", 0, &got)
		if got != test.want {
			fmt.Printf("%v: got %q; want %q\n", test.in, got, test.want)
		}
	}
}

func A(n int, prefix string, count int, res *string) string {
	if n == 0 {
		*res += prefix + "\n"
		return *res
	}

	if count <= n/2 {
		A(n-1, prefix+"(", count+1, res)
	}
	if count > 0 {
		A(n-1, prefix+")", count-1, res)
	}
	return *res
}
