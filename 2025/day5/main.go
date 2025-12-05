package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Span struct {
	Min int
	Max int
}

func main() {
	contents, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	p1 := 0
	p2 := 0
	first := true
	lines := bytes.Split(contents, []byte("\n"))
	ranges := []Span{}
	for _, line := range lines {
		if len(line) == 0 {
			if first {
				first = false
				continue
			} else {
				break
			}
		}
		if first {
			ids := bytes.Split(line, []byte("-"))
			v1, _ := strconv.Atoi(string(ids[0]))
			v2, _ := strconv.Atoi(string(ids[1]))
			ranges = append(ranges, Span{v1, v2})
		} else {
			id, _ := strconv.Atoi(string(line))
			for _, r := range ranges {
				if id >= r.Min && id <= r.Max {
					p1 += 1
					break
				}
			}
		}
	}
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Min < ranges[j].Min || (ranges[i].Min == ranges[j].Min && ranges[i].Max <= ranges[j].Max)
	})
	min := ranges[0].Min
	max := ranges[0].Max
	for _, s := range ranges {
		if s.Min <= max {
			if s.Max > max {
				max = s.Max
			}
		} else {
			p2 += (max + 1 - min)
			min = s.Min
			max = s.Max
		}
	}
	p2 += (max + 1 - min)
	fmt.Println(p1)
	fmt.Println(p2)
}
