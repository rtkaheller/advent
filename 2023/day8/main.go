package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

type Node struct {
	L, R *Node
	Name string
}

type Ghost struct {
	Start, Cur *Node
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	nodes := map[string]*Node{}
	lines := bytes.Split(contents, []byte("\n"))
	instr := string(lines[0])
	ghosts := []*Ghost{}
	for _, line := range lines[1:] {
		if len(line) == 0 {
			continue
		}
		n := string(line[0:3])
		nodes[n] = &Node{Name: n}
		if n[2] == 'A' {
			ghosts = append(ghosts, &Ghost{Start: nil, Cur: nodes[n]})
		}
	}

	var cur *Node
	for _, line := range lines[1:] {
		if len(line) == 0 {
			continue
		}
		if cur == nil {
			cur = nodes[string(line[0:3])]
		}
		if string(line[7:10])[2] == 'A' {
			fmt.Println(line)
		}
		if string(line[12:15])[2] == 'A' {
			fmt.Println(line)
		}
		nodes[string(line[0:3])].L = nodes[string(line[7:10])]
		nodes[string(line[0:3])].R = nodes[string(line[12:15])]
	}
	i := 0
	for i = 0; cur.Name[2] != 'Z'; i++ {
		switch instr[i%len(instr)] {
		case 'L':
			cur = cur.L
		case 'R':
			cur = cur.R
		}
	}
	fmt.Println(i)

	ends := map[*Ghost]map[int]int{}
	looped := 0
	for _, g := range ghosts {
		ends[g] = map[int]int{}
	}
	ans := 1
	for i = 1; looped != len(ghosts); i++ {
		step := instr[(i-1)%len(instr)]
		for _, g := range ghosts {
			switch step {
			case 'L':
				g.Cur = g.Cur.L
			case 'R':
				g.Cur = g.Cur.R
			}
		}
		for _, g := range ghosts {
			if g.Cur.Name[2] == 'Z' {
				if _, ok := ends[g][i%len(instr)]; !ok {
					ends[g][i%len(instr)] = i / len(instr)
					ans *= i / len(instr)
					looped += 1
				}
			}
		}
	}
	fmt.Println(ans * len(instr))
}
