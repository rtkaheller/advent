package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func next(vals []int) (int, int) {
	if len(vals) == 2 {
		return vals[0] - (vals[1] - vals[0]), vals[1] - vals[0]
	}
	d := []int{}
	for i := 0; i < len(vals)-1; i++ {
		d = append(d, vals[i+1]-vals[i])
	}
	pre, nxt := next(d)
	return vals[0] - pre, vals[len(vals)-1] + nxt
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	sum := 0
	sum2 := 0
	lines := bytes.Split(contents, []byte("\n"))
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		ns := strings.Fields(string(line))
		n := []int{}
		for _, s := range ns {
			v, _ := strconv.Atoi(s)
			n = append(n, v)
		}
		pre, nxt := next(n)
		sum += nxt
		sum2 += pre
	}
	fmt.Println(sum)
	fmt.Println(sum2)
}
