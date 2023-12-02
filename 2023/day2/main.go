package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var cubes = map[string]int{
	"blue":  14,
	"red":   12,
	"green": 13,
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	sum := 0
	power := 0
	lines := bytes.Split(contents, []byte("\n"))
	for id, line := range lines {
		if len(line) == 0 {
			continue
		}
		row := strings.Split(string(line), ":")
		draws := strings.Split(row[1], ";")
		work := true
		m := map[string]int{
			"blue":  0,
			"red":   0,
			"green": 0,
		}
		for _, draw := range draws {
			sets := strings.Split(draw, ",")
			for _, set := range sets {
				s := strings.Split(set, " ")
				n, _ := strconv.Atoi(s[1])
				if n > cubes[s[2]] {
					work = false
				}
				if n > m[s[2]] {
					m[s[2]] = n
				}
			}
		}
		power += m["blue"] * m["red"] * m["green"]
		if work {
			sum += id + 1
		}
	}
	fmt.Println(sum)
	fmt.Println(power)
}
