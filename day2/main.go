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

	cmds := bytes.Split(contents, []byte("\n"))
	pos := 0
	depth := 0
	aim := 0

	for _, cmd := range cmds {
		vals := bytes.Split(cmd, []byte(" "))
		if len(vals) != 2 {
			break
		}
		fmt.Println(vals)
		dir := vals[0]
		i_mag, _ := strconv.Atoi(string(vals[1]))
		switch string(dir) {
		case "forward":
			pos += i_mag
			depth += aim * i_mag
		case "down":
			aim += i_mag
		case "up":
			aim -= i_mag
		}
	}
	fmt.Println(pos * depth)
}
