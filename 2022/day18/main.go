package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

var dirs = []Pos{
	Pos{X: 1, Y: 0, Z: 0},
	Pos{X: -1, Y: 0, Z: 0},
	Pos{X: 0, Y: 1, Z: 0},
	Pos{X: 0, Y: -1, Z: 0},
	Pos{X: 0, Y: 0, Z: 1},
	Pos{X: 0, Y: 0, Z: -1},
}

type Pos struct {
	X, Y, Z int
}

type Cube struct {
	P Pos
}

type Stack struct {
	Data []Pos
}

func (s *Stack) Pop() Pos {
	v := s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1]
	return v
}

func (s *Stack) Push(v Pos) {
	s.Data = append(s.Data, v)
}

func Fill(grid map[Pos]bool, max Pos) map[Pos]bool {
	result := make(map[Pos]bool)
	start := Pos{-1, -1, -1}
	var s Stack
	s.Push(start)
	for len(s.Data) != 0 {
		cur := s.Pop()
		if _, ok := grid[cur]; ok {
			continue
		}
		if _, ok := result[cur]; !ok {
			result[cur] = true
			for _, d := range dirs {
				can := Pos{X: d.X + cur.X, Y: d.Y + cur.Y, Z: d.Z + cur.Z}
				if can.X < -1 || can.X > max.X+1 || can.Y < -1 || can.Y > max.Y+1 || can.Z < -1 || can.Z > max.Z+1 {
					continue
				}
				s.Push(can)
			}
		}
	}
	return result
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	var cubes []Cube
	grid := make(map[Pos]bool)
	var maxP Pos
	for _, line := range bytes.Split(contents, []byte("\n")) {
		if len(line) == 0 {
			continue
		}
		c := bytes.Split(line, []byte(","))
		x, _ := strconv.Atoi(string(c[0]))
		y, _ := strconv.Atoi(string(c[1]))
		z, _ := strconv.Atoi(string(c[2]))
		if x > maxP.X {
			maxP.X = x
		}
		if y > maxP.Y {
			maxP.Y = y
		}
		if z > maxP.Z {
			maxP.Z = z
		}
		cubes = append(cubes, Cube{P: Pos{X: x, Y: y, Z: z}})
		grid[cubes[len(cubes)-1].P] = true
	}
	filled := Fill(grid, maxP)
	sum := 0
	sum2 := 0
	for _, k := range cubes {
		for _, d := range dirs {
			can := Pos{k.P.X + d.X, k.P.Y + d.Y, k.P.Z + d.Z}
			if _, ok := grid[can]; !ok {
				if _, in_fill := filled[can]; in_fill {
					sum2 += 1
				}
				sum += 1
			}
		}
	}
	fmt.Println(sum)
	fmt.Println(sum2)
}
