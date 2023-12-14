package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func Abs(v int) int {
	if v < 0 {
		return v * -1
	}
	return v
}

func NewMid(d []int) (int, bool) {
	for i := 0; i < len(d)-1; i++ {
		match := true
		fixed := false
		for t := 0; i-t >= 0 && i+t < len(d)-1; t++ {
			if d[i-t] != d[i+t+1] {
				diff := Abs(d[i-t] - d[i+t+1])
				if !fixed && (diff&(diff-1)) == 0 {
					fixed = true
				} else {
					match = false
					break
				}
			}
		}
		if match && fixed {
			return i + 1, match
		}
	}
	return 0, false
}

func Mid(d []int) (int, bool) {
	for i := 0; i < len(d)-1; i++ {
		match := true
		for t := 0; i-t >= 0 && i+t < len(d)-1; t++ {
			if d[i-t] != d[i+t+1] {
				match = false
				break
			}
		}
		if match {
			return i + 1, match
		}
	}
	return 0, false
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	sum := 0
	sum2 := 0
	rows := []int{}
	cols := []string{}

	lines := bytes.Split(contents, []byte("\n"))

	for i := 0; i < len(lines[0]); i++ {
		cols = append(cols, "")
	}

	for l, line := range lines {
		if len(line) == 0 {
			col := []int{}
			for _, c := range cols {
				v, _ := strconv.ParseInt(string(c), 2, 64)
				col = append(col, int(v))
			}
			if v, ok := Mid(rows); ok {
				sum += 100 * v
			} else if v, ok := Mid(col); ok {
				sum += v
			}
			if v, ok := NewMid(rows); ok {
				sum2 += 100 * v
			} else if v, ok := NewMid(col); ok {
				sum2 += v
			}
			rows = []int{}
			if l+1 >= len(lines) {
				break
			}
			cols = []string{}
			for i := 0; i < len(lines[l+1]); i++ {
				cols = append(cols, "")
			}
			continue
		}
		v, _ := strconv.ParseInt(strings.ReplaceAll(strings.ReplaceAll(string(line), ".", "0"), "#", "1"), 2, 64)
		rows = append(rows, int(v))
		for i := 0; i < len(line); i++ {
			if line[i] == '.' {
				cols[i] += "0"
			} else {
				cols[i] += "1"
			}
		}
	}
	fmt.Println(sum)
	fmt.Println(sum2)
}
