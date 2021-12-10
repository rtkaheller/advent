package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
)

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	m := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
		'>': '<',
	}
	points := map[rune]int{
		')': 3,
		'}': 1197,
		']': 57,
		'>': 25137,
	}
	subPoints := map[rune]int{
		'(': 1,
		'{': 3,
		'[': 2,
		'<': 4,
	}
	var stack []rune
	sum := 0
	var completeSums []int
	for _, line := range bytes.Split(contents, []byte("\n")) {
		if len(line) == 0 {
			continue
		}
		stack = []rune{}
		found := false
		for _, r := range string(line) {
			switch r {
			case '(', '{', '[', '<':
				stack = append(stack, r)
			case ')', '}', ']', '>':
				if stack[len(stack)-1] != m[r] {
					sum += points[r]
					found = true
				}
				if len(stack) != 0 {
					stack = stack[:len(stack)-1]
				}
			}
			if found {
				break
			}
		}
		if !found {
			completeSums = append(completeSums, 0)
			for i := len(stack) - 1; i >= 0; i-- {
				completeSums[len(completeSums)-1] = completeSums[len(completeSums)-1]*5 + subPoints[stack[i]]
			}
		}
	}
	sort.Ints(completeSums)
	fmt.Println(sum)
	fmt.Println(completeSums[len(completeSums)/2])
}
