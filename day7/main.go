package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func cost(goal int, crabs []int) int {
	move := 0
	for _, crab := range crabs {
		n := int(math.Abs(float64(crab - goal)))
		move += (n * (n + 1)) / 2
	}
	return move
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	crabs := bytes.Split(contents, []byte(","))

	total := 0
	var crab_pos []int
	for _, crab := range crabs {
		pos, _ := strconv.Atoi(strings.Trim(string(crab), "\n"))
		total += pos
		crab_pos = append(crab_pos, pos)
	}
	cur := cost(0, crab_pos)
	for i := 1; ; i++ {
		next := cost(i, crab_pos)
		if next > cur {
			fmt.Println(i - 1)
			fmt.Println(cur)
			break
		}
		cur = next
	}
}
