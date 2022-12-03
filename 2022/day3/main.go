package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	sum := 0
	for _, line := range bytes.Split(contents, []byte("\n")) {
		if len(line) == 0 {
			break
		}
		comp := make(map[byte]int)
		for i, item := range line {
			if i < len(line)/2 {
				comp[item] += 1
			} else if _, ok := comp[item]; ok {
				if item >= 'a' && item <= 'z' {
					sum += int(item) - 96
				} else {
					sum += int(item) - 64 + 26
				}
				break
			}
		}
	}
	fmt.Println(sum)
	first := make(map[byte]bool)
	sec := make(map[byte]bool)
	cur := 0
	sum2 := 0
	for _, line := range bytes.Split(contents, []byte("\n")) {
		if len(line) == 0 {
			break
		}
		for _, item := range line {
			switch cur {
			case 0:
				first[item] = true
			case 1:
				sec[item] = true
			case 2:
				_, in_first := first[item]
				_, in_sec := sec[item]
				if in_first && in_sec {
					if item >= 'a' && item <= 'z' {
						sum2 += int(item) - 96
					} else {
						sum2 += int(item) - 64 + 26
					}
					cur = -1
					first = make(map[byte]bool)
					sec = make(map[byte]bool)
					break
				}
			}
		}
		cur += 1
	}
	fmt.Println(sum2)
}
