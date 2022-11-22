package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	depths := bytes.Split(contents, []byte("\n"))
	c := 0
	sums := 0
	prev := 1000000000000
	prev_sum := 100000000000
	for i, depth := range depths {
		cur_sum := 0
		if i+2 < len(depths) {
			for t := 0; t < 3; t++ {
				cur, _ := strconv.Atoi(string(depths[i+t]))
				cur_sum += cur
			}
			if cur_sum > prev_sum {
				sums += 1
			}
			prev_sum = cur_sum
		}

		cur, _ := strconv.Atoi(string(depth))
		if cur > prev {
			c += 1
		}
		prev = cur
	}
	fmt.Println(c)
	fmt.Println(sums)
}
