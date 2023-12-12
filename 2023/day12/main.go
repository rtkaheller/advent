package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func MemoKey(s string, c []Spring) string {
	return s + fmt.Sprintf("%v", c)
}

var memo = map[string]int{}

func Valid(spr string, can []int) bool {
	group := []int{}
	cur := 0
	in := false
	for _, c := range spr {
		if c == '#' {
			in = true
			cur += 1
		} else if in {
			in = false
			group = append(group, cur)
			cur = 0
		}
	}
	if in {
		group = append(group, cur)
	}
	correct := 0
	for i := 0; i < len(group) && i < len(can); i++ {
		if group[i] == can[i] {
			correct += 1
		}
	}
	return len(can) == len(group) && len(can) == correct
}

type Spring struct {
	L, S, E int
}

func Build(spr string, can []Spring, pos int, count []int) int {
	if len(can) == 0 {
		if Valid(spr, count) {
			return 1
		} else {
			return 0
		}
	}
	key := MemoKey(spr[pos:], can)
	if v, ok := memo[key]; ok {
		return v
	}
	cur := can[0]
	tot := 0
	for i := cur.S; i <= cur.E-cur.L; i++ {
		if i < pos {
			continue
		}
		valid := true
		for t := i; t < i+cur.L; t++ {
			if spr[t] == '.' {
				valid = false
				break
			}
		}
		if valid {
			for t := pos; t < i; t++ {
				if spr[t] == '#' {
					valid = false
					break
				}
			}
		}
		if valid && (len(can) <= 1 || spr[i+cur.L] != '#') {
			tot += Build(spr[:i]+string(bytes.Repeat([]byte{'#'}, cur.L))+spr[i+cur.L:], can[1:], i+cur.L+1, count)
		}
	}
	memo[key] = tot
	return tot
}

func Ans(spr string, count []int) int {
	tot := 0
	starts := []int{}
	ends := []int{}
	for i, c := range count {
		starts = append(starts, tot+i)
		tot += c
	}
	back := 0
	for i := len(count) - 1; i >= 0; i-- {
		ends = append([]int{len(spr) - (len(count) - 1 - i) - back}, ends...)
		back += count[i]
	}
	c := map[rune]int{'?': 0, '#': 0, '.': 0}
	for _, r := range spr {
		c[r] += 1
	}
	springs := []Spring{}
	for i, _ := range count {
		springs = append(springs, Spring{count[i], starts[i], ends[i]})
	}
	return Build(spr, springs, 0, count)
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
		s := strings.Split(string(line), " ")
		gr := []int{}
		for _, g := range strings.Fields(strings.ReplaceAll(s[1], ",", " ")) {
			v, _ := strconv.Atoi(g)
			gr = append(gr, v)
		}

		long := []int{}
		for i := 0; i < 5; i++ {
			long = append(long, gr...)
		}

		sum += Ans(s[0], gr)
		sum2 += Ans(strings.Join([]string{s[0], s[0], s[0], s[0], s[0]}, "?"), long)
	}
	fmt.Println(sum)
	fmt.Println(sum2)
}
