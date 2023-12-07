package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func mapRange(c Conv, in int) int {
	if in >= c.S && in-c.S < c.R {
		return c.D + in - c.S
	} else {
		return in
	}
}

type Conv struct {
	S, D, R int
}

func parse(i int, lines [][]byte) (int, []Conv) {
	ret := []Conv{}
	for ; len(lines[i]) != 0; i++ {
		fields := strings.Fields(string(lines[i]))
		nums := []int{}
		for _, f := range fields {
			v, _ := strconv.Atoi(f)
			nums = append(nums, v)
		}
		ret = append(ret, Conv{D: nums[0], S: nums[1], R: nums[2]})
	}
	return i + 2, ret
}

func part2(lines [][]byte) {
	initSeeds := []int{}
	for _, s := range strings.Fields(string(lines[0][6:])) {
		v, _ := strconv.Atoi(s)
		initSeeds = append(initSeeds, v)
	}
	seeds := []int{}
	for i := 0; i+1 < len(initSeeds); i += 2 {
		for t := initSeeds[i]; t < initSeeds[i]+initSeeds[i+1]; t++ {
			seeds = append(seeds, t)
		}
	}
	i := 3
	var conv []Conv
	for i < len(lines) {
		i, conv = parse(i, lines)
		for t := 0; t < len(seeds); t++ {
			for _, c := range conv {
				if seeds[t] >= c.S && seeds[t]-c.S < c.R {
					seeds[t] = c.D + seeds[t] - c.S
					break
				}
			}
		}
	}

	m := -1
	for _, s := range seeds {
		if m == -1 || m > s {
			m = s
		}
	}
	fmt.Println(m)
}

func part1(lines [][]byte) {
	seeds := []int{}
	for _, s := range strings.Fields(string(lines[0][6:])) {
		v, _ := strconv.Atoi(s)
		seeds = append(seeds, v)
	}
	i := 3
	var conv []Conv
	for i < len(lines) {
		i, conv = parse(i, lines)
		for t := 0; t < len(seeds); t++ {
			for _, c := range conv {
				if seeds[t] >= c.S && seeds[t]-c.S < c.R {
					seeds[t] = c.D + seeds[t] - c.S
					break
				}
			}
		}
	}

	m := -1
	for _, s := range seeds {
		if m == -1 || m > s {
			m = s
		}
	}
	fmt.Println(m)
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := bytes.Split(contents, []byte("\n"))
	part1(lines)
	part2(lines)
}
