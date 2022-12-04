package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

type Pair struct {
	Lower, Upper int
}

func Parse(in []byte) Pair {
	l := bytes.Split(in, []byte("-"))
	lower, _ := strconv.Atoi(string(l[0]))
	upper, _ := strconv.Atoi(string(l[1]))
	return Pair{Lower: lower, Upper: upper}
}

func main() {
	contents, err := ioutil.ReadFile("fake.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	c := 0
	c2 := 0
	for _, line := range bytes.Split(contents, []byte("\n")) {
		if len(line) == 0 {
			continue
		}
		pairs := bytes.Split(line, []byte(","))
		first := Parse(pairs[0])
		sec := Parse(pairs[1])
		if (first.Lower <= sec.Lower && first.Upper >= sec.Upper) || (sec.Lower <= first.Lower && sec.Upper >= first.Upper) {
			c += 1
		}
		if (first.Lower <= sec.Lower && first.Upper >= sec.Lower) ||
			(first.Lower <= sec.Upper && first.Upper >= sec.Upper) ||
			(sec.Lower <= first.Lower && sec.Upper >= first.Lower) ||
			(sec.Lower <= first.Upper && sec.Upper >= first.Upper) {
			c2 += 1
		}
	}
	fmt.Println(c)
	fmt.Println(c2)
}
