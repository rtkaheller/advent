package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

func Abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}

func Ufans3(n int) string {
	base := 1
	for i := 1; ; i++ {
		if Snafu(strings.Repeat("2", i)) >= Abs(n) {
			break
		}
		base *= 5
	}

	var b5 []int

	cur := 0
	orig := n
	for base > 1 {
		p := Abs(n)

		r := p % base
		d := p / base

		if p == 0 {
			b5 = append(b5, 0)
		} else if d == 2 {
			b5 = append(b5, 2*n/p)
		} else if d == 1 {
			if r == 0 {
				b5 = append(b5, 1*n/p)
			} else {
				if Abs(orig-(cur+(2*n/p)*base)) < Abs(orig-(cur+(1*n/p)*base)) {
					b5 = append(b5, 2*n/p)
				} else {
					b5 = append(b5, 1*n/p)
				}
			}
		} else if d == 0 {
			if Abs(orig-(cur+(1*n/p)*base)) < Abs(orig-cur) {
				b5 = append(b5, 1*n/p)
			} else {
				b5 = append(b5, 0)
			}
		}
		cur += b5[len(b5)-1] * base
		n = orig - cur
		base /= 5
	}
	b5 = append(b5, n)
	var ret string
	for i := 0; i < len(b5); i++ {
		switch b5[i] {
		case 2:
			ret += "2"
		case 1:
			ret += "1"
		case 0:
			ret += "0"
		case -1:
			ret += "-"
		case -2:
			ret += "="
		}
	}
	return ret
}

func Snafu(n string) int {
	base := 1
	ret := 0
	for i := len(n) - 1; i >= 0; i-- {
		switch n[i] {
		case '2':
			ret += 2 * base
		case '1':
			ret += base
		case '0':
			ret += 0
		case '-':
			ret -= base
		case '=':
			ret -= 2 * base
		}
		base *= 5
	}
	return ret
}

func main() {
	contents, err := ioutil.ReadFile("fake.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	sum := 0
	for _, line := range bytes.Split(contents, []byte("\n")) {
		if len(line) == 0 {
			continue
		}
		sum += Snafu(string(line))
	}
	fmt.Println(Ufans3(sum))
}
