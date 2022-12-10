package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

func Tick(c, v int, crt *[40][6]bool) {
	x := c % 40
	y := c / 40
	if v < -1 || v > 41 {
		return
	}
	if x >= v-1 && x <= v+1 {
		crt[x][y] = true
	}
}

func main() {
	contents, err := ioutil.ReadFile("fake.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	signals := []int{}
	v := 1
	var crt [40][6]bool
	for _, line := range bytes.Split(contents, []byte("\n")) {
		if len(line) == 0 {
			continue
		}
		inst := bytes.Split(line, []byte(" "))
		switch string(inst[0]) {
		case "noop":
			Tick(len(signals), v, &crt)
			signals = append(signals, v)
		case "addx":
			Tick(len(signals), v, &crt)
			signals = append(signals, v)
			val, _ := strconv.Atoi(string(inst[1]))
			Tick(len(signals), v, &crt)
			signals = append(signals, v)
			v += val
		}
	}
	fmt.Println(signals[19]*20 + signals[59]*60 + signals[99]*100 + signals[139]*140 + signals[179]*180 + signals[219]*220)
	for y := 0; y < 6; y++ {
		for x := 0; x < 40; x++ {
			if crt[x][y] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}
