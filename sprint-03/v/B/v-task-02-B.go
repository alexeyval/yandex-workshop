package main

import "fmt"

func main() {
	m := map[uint8]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	for _, test := range []struct {
		in, want string
	}{
		{"7", "p q r s "},
		{"23", "ad ae af bd be bf cd ce cf "},
		{"92", "wa wb wc xa xb xc ya yb yc za zb zc "},
		{"234", "adg adh adi aeg aeh aei afg afh afi " +
			"bdg bdh bdi beg beh bei bfg bfh bfi " +
			"cdg cdh cdi ceg ceh cei cfg cfh cfi "},
	} {
		var got string
		B(m, test.in, "", &got)
		if got != test.want {
			fmt.Printf("%v: got %q; want %q\n", test.in, got, test.want)
		}
	}
}

func B(m map[uint8]string, n, prefix string, c *string) {
	if len(n) == 0 {
		*c += prefix + " "
		return
	}
	if len(n) == 1 {
		for _, k := range m[n[0]] {
			B(m, "", prefix+string(k), c)
		}
		return
	}

	_ = m[n[0]] + m[n[1]]
	for _, i := range m[n[0]] {
		for _, k := range m[n[1]] {
			B(m, n[2:], prefix+string(i)+string(k), c)
		}
	}
}
