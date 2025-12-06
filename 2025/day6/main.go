package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Column struct {
	Nums   []int
	Strs   []string
	Op     string
	Length int
}

func Calc(nums []int, op string) int {
	r := 0
	if op == "*" {
		r = 1
	}
	for _, n := range nums {
		switch op {
		case "*":
			r *= n
		case "+":
			r += n
		}
	}
	return r
}

func (c *Column) Calc2() int {
	maxL := 0
	for _, s := range c.Strs {
		if len(s) > maxL {
			maxL = len(s)
		}
	}
	newNums := []int{}
	for i := 0; i < c.Length; i++ {
		nn := ""
		for _, s := range c.Strs {
			if s[i] != ' ' {
				nn += string(s[i])
			}
		}
		v, _ := strconv.Atoi(nn)
		newNums = append(newNums, v)
	}

	return Calc(newNums, c.Op)
}

func main() {
	contents, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	p1 := 0
	p2 := 0
	lines := bytes.Split(contents, []byte("\n"))

	c := []Column{}
	l := 1
	for _, b := range string(lines[len(lines)-2]) {
		if b == '*' || b == '+' {
			if len(c) > 0 {
				c[len(c)-1].Length = l - 1
			}
			c = append(c, Column{Nums: []int{}, Strs: []string{}, Op: string(b)})
			l = 1
		} else {
			l += 1
		}
	}
	c[len(c)-1].Length = l

	for _, line := range lines[:len(lines)-2] {
		if len(line) == 0 {
			continue
		}
		pos := 0
		for i, col := range c {
			n := line[pos : pos+col.Length]
			v, _ := strconv.Atoi(strings.Trim(string(n), " "))
			c[i].Nums = append(c[i].Nums, v)
			c[i].Strs = append(c[i].Strs, string(n))
			pos += col.Length + 1
		}
	}
	for _, col := range c {
		p1 += Calc(col.Nums, col.Op)
		p2 += col.Calc2()
	}

	fmt.Println(p1)
	fmt.Println(p2)
}
