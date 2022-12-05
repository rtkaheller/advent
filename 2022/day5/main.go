package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

type Instruction struct {
	Count, Source, Dest int
}

func part1() {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	crate_sec := true
	var stacks [][]byte
	var insts []Instruction
	lines := bytes.Split(contents, []byte("\n"))
	num_stacks := (len(lines[0]) + 1) / 4
	for i := 0; i < num_stacks; i++ {
		stacks = append(stacks, []byte{})
	}

	for _, line := range lines {
		if len(line) == 0 {
			crate_sec = false
			continue
		}
		if crate_sec {
			for i := 0; i < num_stacks; i++ {
				if line[i*4] == '[' {
					stacks[i] = append([]byte{line[i*4+1]}, stacks[i]...)
				}
			}
		} else {
			words := bytes.Split(line, []byte(" "))
			c, _ := strconv.Atoi(string(words[1]))
			s, _ := strconv.Atoi(string(words[3]))
			d, _ := strconv.Atoi(string(words[5]))
			insts = append(insts, Instruction{Count: c, Source: s - 1, Dest: d - 1})
		}
	}
	for _, inst := range insts {
		for i := 0; i < inst.Count; i++ {
			stacks[inst.Dest] = append(stacks[inst.Dest], stacks[inst.Source][len(stacks[inst.Source])-1])
			stacks[inst.Source] = stacks[inst.Source][:len(stacks[inst.Source])-1]
		}
	}
	for _, stack := range stacks {
		fmt.Printf("%s", string(stack[len(stack)-1]))
	}
	fmt.Println()
}

func part2() {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	crate_sec := true
	var stacks [][]byte
	var insts []Instruction
	lines := bytes.Split(contents, []byte("\n"))
	num_stacks := (len(lines[0]) + 1) / 4
	for i := 0; i < num_stacks; i++ {
		stacks = append(stacks, []byte{})
	}

	for _, line := range lines {
		if len(line) == 0 {
			crate_sec = false
			continue
		}
		if crate_sec {
			for i := 0; i < num_stacks; i++ {
				if line[i*4] == '[' {
					stacks[i] = append([]byte{line[i*4+1]}, stacks[i]...)
				}
			}
		} else {
			words := bytes.Split(line, []byte(" "))
			c, _ := strconv.Atoi(string(words[1]))
			s, _ := strconv.Atoi(string(words[3]))
			d, _ := strconv.Atoi(string(words[5]))
			insts = append(insts, Instruction{Count: c, Source: s - 1, Dest: d - 1})
		}
	}
	for _, inst := range insts {
		for i := inst.Count - 1; i >= 0; i-- {
			stacks[inst.Dest] = append(stacks[inst.Dest], stacks[inst.Source][len(stacks[inst.Source])-i-1])
		}
		stacks[inst.Source] = stacks[inst.Source][:len(stacks[inst.Source])-inst.Count]
	}
	for _, stack := range stacks {
		fmt.Printf("%s", string(stack[len(stack)-1]))
	}
	fmt.Println()
}

func main() {
	part1()
	part2()
}
