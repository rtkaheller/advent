package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func solve(bats int, bank []int) int {
	max := 0
	best := 0
	last := -1
	for t := range bats {
		max = 0
		for i := last + 1; i < len(bank)-(bats-t)+1; i++ {
			v := bank[i]
			if v > max {
				max = v
				last = i
			}
		}
		best = best * 10
		best += max
	}
	return best
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
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		bank := []int{}
		for _, b := range line {
			v, _ := strconv.Atoi(string(b))
			bank = append(bank, v)
		}
		p1 += solve(2, bank)
		p2 += solve(12, bank)
	}
	fmt.Println(p1)
	fmt.Println(p2)
}
