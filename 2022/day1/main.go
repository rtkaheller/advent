package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
)

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	var elves []int

	cur_elf := 0

	calories := bytes.Split(contents, []byte("\n"))
	for _, calorie := range calories {
		if len(calorie) == 0 {
			elves = append(elves, cur_elf)
			cur_elf = 0
			continue
		}
		cal, _ := strconv.Atoi(string(calorie))
		cur_elf += cal
	}
	sort.Ints(elves)
	fmt.Println(elves[len(elves)-1])
	fmt.Println(elves[len(elves)-1] + elves[len(elves)-2] + elves[len(elves)-3])
}
