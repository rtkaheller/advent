package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	winning := map[int]int{}
	instances := map[int]int{}
	sum := 0
	lines := bytes.Split(contents, []byte("\n"))
	for game, line := range lines {
		if len(line) == 0 {
			continue
		}

		wins := map[int]bool{}
		card := strings.Split(string(line), ": ")
		nums := strings.Split(card[1], " | ")
		w := 0
		for i := 0; i < len(nums[0]); i += 3 {
			v, _ := strconv.Atoi(strings.Trim(nums[0][i:i+2], " "))
			wins[v] = true
		}
		for i := 0; i < len(nums[1]); i += 3 {
			v, _ := strconv.Atoi(strings.Trim(nums[1][i:i+2], " "))
			if _, ok := wins[v]; ok {
				w += 1
			}
		}
		if w > 0 {
			sum += int(math.Pow(2, float64(w-1)))
		}
		winning[game] = w
		instances[game] = 1
	}

	for i := 0; i < len(winning); i++ {
		for t := i + 1; t <= i+winning[i]; t++ {
			instances[t] += instances[i]
		}
	}
	tot := 0
	for i := 0; i < len(winning); i++ {
		tot += instances[i]
	}

	fmt.Println(sum)
	fmt.Println(tot)
}
