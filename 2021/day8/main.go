package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"sort"
)

var regularSegments = map[string]int{
	"abcefg":  0,
	"cf":      1,
	"acdeg":   2,
	"acdfg":   3,
	"bcdf":    4,
	"abdfg":   5,
	"abdefg":  6,
	"acf":     7,
	"abcdefg": 8,
	"abcdfg":  9,
}

type Display struct {
	Set           string
	Segments      map[rune]bool
	ResolvedValue int
}

func (d *Display) Init(in []byte) {
	d.Set = string(in)
	d.Segments = make(map[rune]bool)
	for _, segment := range d.Set {
		d.Segments[segment] = true
	}
	switch len(d.Set) {
	case 2:
		d.ResolvedValue = 1
	case 3:
		d.ResolvedValue = 7
	case 4:
		d.ResolvedValue = 4
	case 7:
		d.ResolvedValue = 8
	}
}

func (d *Display) Resolve(results map[rune]rune) {
	var resolved []rune
	for seg, _ := range d.Segments {
		resolved = append(resolved, results[seg])
	}
	sort.Slice(resolved, func(i, j int) bool { return resolved[i] < resolved[j] })
	var resolvedStr string
	for _, r := range resolved {
		resolvedStr += string(r)
	}
	d.ResolvedValue = regularSegments[resolvedStr]
}

type Row struct {
	Signals [10]*Display
	Outputs [4]*Display
}

func (r *Row) Decode() int {
	var one, four, seven *Display
	results := make(map[rune]rune)
	var letterCounts = map[rune]int{'a': 0, 'b': 0, 'c': 0, 'd': 0, 'e': 0, 'f': 0, 'g': 0}
	for _, signal := range r.Signals {
		for seg, _ := range signal.Segments {
			letterCounts[seg] += 1
		}
		switch signal.ResolvedValue {
		case 1:
			one = signal
		case 7:
			seven = signal
		case 4:
			four = signal
		}
	}
	for char, count := range letterCounts {
		switch count {
		case 6:
			results[char] = 'b'
		case 4:
			results[char] = 'e'
		case 9:
			results[char] = 'f'
		}
	}
	for seg, _ := range seven.Segments {
		if _, ok := one.Segments[seg]; !ok {
			results[seg] = 'a'
		} else if letterCounts[seg] == 8 {
			results[seg] = 'c'
		}
	}
	for seg, _ := range four.Segments {
		if _, ok := results[seg]; !ok {
			results[seg] = 'd'
		}
	}
	for _, char := range []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g'} {
		if _, ok := results[char]; !ok {
			results[char] = 'g'
		}
	}
	val := 0
	for i, out := range r.Outputs {
		out.Resolve(results)
		val += int(math.Pow(10, float64(3-i))) * out.ResolvedValue
	}
	return val
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	var rows []*Row
	for _, line := range bytes.Split(contents, []byte("\n")) {
		if len(line) == 0 {
			continue
		}
		datas := bytes.Split(line, []byte(" | "))

		r := new(Row)
		for i, signal := range bytes.Split(datas[0], []byte(" ")) {
			disp := new(Display)
			disp.Init(signal)
			r.Signals[i] = disp
		}
		for i, output := range bytes.Split(datas[1], []byte(" ")) {
			disp := new(Display)
			disp.Init(output)
			r.Outputs[i] = disp
		}
		rows = append(rows, r)
	}
	sum := 0
	for _, row := range rows {
		for _, output := range row.Outputs {
			if output.ResolvedValue != 0 {
				sum += 1
			}
		}
	}
	fmt.Println(sum)
	sum = 0
	for _, row := range rows {
		sum += row.Decode()
	}
	fmt.Println(sum)
}
