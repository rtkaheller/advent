package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := bytes.Split(contents, []byte("\n"))
	timesStr := strings.Fields(string(lines[0][6:]))
	distsStr := strings.Fields(string(lines[1][10:]))
	tot := 1
	for i := 0; i < len(timesStr); i++ {
		time, _ := strconv.Atoi(timesStr[i])
		dist, _ := strconv.Atoi(distsStr[i])
		count := 0
		for t := 0; t < time; t++ {
			if dist < t*time-t*t {
				count += 1
			}
		}
		tot *= count
	}
	fmt.Println(tot)

	time, _ := strconv.Atoi(strings.ReplaceAll(string(lines[0][6:]), " ", ""))
	dist, _ := strconv.Atoi(strings.ReplaceAll(string(lines[1][10:]), " ", ""))
	count := 0
	for t := 0; t < time; t++ {
		if dist < t*time-t*t {
			count += 1
		}
	}
	fmt.Println(count)
}
