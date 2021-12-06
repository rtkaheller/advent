package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	fishRate     = 6
	newFishBonus = 2
	days         = 256
)

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	var fishies [9]int
	fishes := bytes.Split(contents, []byte(","))
	for _, fish := range fishes {
		if len(fish) == 0 {
			continue
		}
		left, _ := strconv.Atoi(strings.Trim(string(fish), "\n"))
		fishies[left] += 1
	}
	for i := 0; i < days; i++ {
		var newFishies [9]int
		newFishies[fishRate+newFishBonus] += fishies[0]
		newFishies[fishRate] += fishies[0]
		for t := 1; t < len(fishies); t++ {
			newFishies[t-1] += fishies[t]
		}
		fishies = newFishies
	}
	school := 0
	for i := 0; i < len(fishies); i++ {
		school += fishies[i]
	}
	fmt.Println(school)
}
