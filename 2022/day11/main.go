package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
)

type Monkey struct {
	Items             []int64
	Operation         byte
	Const             int64
	Test, True, False int64
	Inspections       int64
}

func (m *Monkey) Inspect(monkeys []*Monkey, relief int64, gcd int64) {
	for _, item := range m.Items {
		m.Inspections += 1
		val := m.Const
		if val == 0 {
			val = item
		}
		//fmt.Printf("%v, %q, %v, %v\n", item, m.Operation, val, m.Const)
		switch m.Operation {
		case '+':
			item = item + val
		case '*':
			item = item * val
		}
		item = item / relief
		mod := item % m.Test
		if mod == 0 {
			monkeys[m.True].Items = append(monkeys[m.True].Items, item%gcd)
		} else {
			monkeys[m.False].Items = append(monkeys[m.False].Items, item%gcd)
		}
	}
	m.Items = []int64{}
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	lines := bytes.Split(contents, []byte("\n"))
	fmt.Println(monkeyBusiness(lines, 20, 3))
	fmt.Println(monkeyBusiness(lines, 10000, 1))
}

func monkeyBusiness(lines [][]byte, rounds, relief int) int {
	var monkeys []*Monkey
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if len(line) == 0 {
			continue
		}
		if string(line[0:6]) == "Monkey" {
			var m Monkey
			w := bytes.Split(lines[i+1][2:], []byte(" "))
			for _, item := range w[2:] {
				v, err := strconv.Atoi(string(item))
				if err != nil {
					v, _ = strconv.Atoi(string(item[:len(item)-1]))
				}
				m.Items = append(m.Items, int64(v))
			}
			w = bytes.Split(lines[i+2], []byte(":"))
			m.Operation = w[1][1:][10]
			v, _ := strconv.Atoi(string(w[1][1:][12:]))
			m.Const = int64(v)
			fmt.Sscanf(string(lines[i+3]), "  Test: divisible by %d", &m.Test)
			fmt.Sscanf(string(lines[i+4]), "    If true: throw to monkey %d", &m.True)
			fmt.Sscanf(string(lines[i+5]), "    If false: throw to monkey %d", &m.False)
			monkeys = append(monkeys, &m)
			i += 5
		}
	}
	gcd := int64(1)
	for _, m := range monkeys {
		gcd *= m.Test
	}
	for i := 0; i < rounds; i++ {
		for _, m := range monkeys {
			m.Inspect(monkeys, int64(relief), gcd)
		}
	}
	var s []int
	for _, m := range monkeys {
		s = append(s, int(m.Inspections))
	}
	sort.Ints(s)
	return s[len(s)-1] * s[len(s)-2]
}
