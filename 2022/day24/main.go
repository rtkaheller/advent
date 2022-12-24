package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Pos struct {
	X, Y int
}

type Vec struct {
	P   Pos
	Dir byte
}

func (v Vec) Move(size Pos) Vec {
	newV := Vec{P: v.P, Dir: v.Dir}
	d := Delta(v.Dir)
	newV.P.X += d.X
	newV.P.Y += d.Y
	if newV.P.X >= size.X-1 {
		newV.P.X = 0
	}
	if newV.P.X < 0 {
		newV.P.X = size.X - 2
	}
	if newV.P.Y >= size.Y-1 {
		newV.P.Y = 0
	}
	if newV.P.Y < 0 {
		newV.P.Y = size.Y - 2
	}
	return newV
}

func Delta(d byte) Pos {
	switch d {
	case '>':
		return Pos{X: 1, Y: 0}
	case 'v':
		return Pos{X: 0, Y: 1}
	case '<':
		return Pos{X: -1, Y: 0}
	case '^':
		return Pos{X: 0, Y: -1}
	}
	return Pos{}
}

var dirs = []byte{'^', '>', 'v', '<'}

func PrintBliz(cur Pos, bliz map[Vec]bool, size Pos) {
	for y := -1; y < size.Y; y++ {
		for x := -1; x < size.X; x++ {
			if x == cur.X && y == cur.Y {
				fmt.Printf("E")
				continue
			}
			if y == -1 || x == -1 || y == size.Y-1 || x == size.X-1 {
				fmt.Printf("#")
				continue
			}
			count := 0
			var last byte
			for _, dir := range dirs {
				if _, ok := bliz[Vec{P: Pos{x, y}, Dir: dir}]; ok {
					count += 1
					last = dir
				}
			}
			if count == 0 {
				fmt.Printf(".")
			} else if count == 1 {
				fmt.Printf("%c", last)
			} else {
				fmt.Printf("%d", count)
			}
		}
		fmt.Println()
	}
}

type PathPair struct {
	P Pos
	R int
}

func (p *PathPair) H(size Pos) int {
	return p.R*Abs(size.X+size.Y) + Abs(size.X-p.P.X+size.Y-p.P.Y)
}

type Stack struct {
	Data []PathPair
	Dest Pos
}

func (s *Stack) Push(v PathPair) {
	s.Data = append(s.Data, v)
}

func (s *Stack) Pop() PathPair {
	val := s.Data[0]
	s.Data = s.Data[1:]
	return val
}

func (s Stack) Len() int {
	return len(s.Data)
}

func (s Stack) Less(i, j int) bool {
	a := s.Data[i]
	b := s.Data[j]
	return a.H(s.Dest) < b.H(s.Dest)
}

func (s Stack) Swap(i, j int) {
	s.Data[i], s.Data[j] = s.Data[j], s.Data[i]
}

var callCount = 0

func BlizPos(bliz []Vec) map[Pos]bool {
	callCount += 1
	pos := make(map[Pos]bool)
	for _, b := range bliz {
		pos[b.P] = true
	}
	return pos
}

func Path(start Pos, end Pos, bliz []Vec, size Pos, round int) (int, int) {
	var part1 int
	var there, back bool
	rounds := make(map[int][]Vec)
	rounds_pos := make(map[int]map[Pos]bool)
	rounds[round] = bliz
	rounds_pos[round] = BlizPos(bliz)
	v := make(map[PathPair]bool) //visited
	var can Stack                //candidates
	can.Dest = end

	can.Push(PathPair{P: start, R: round})
	for can.Len() > 0 {
		sort.Sort(can)
		cur := can.Pop()
		if _, ok := v[cur]; ok {
			// We've been here
			continue
		}
		v[cur] = true
		if cur.P == end {
			if !there {
				part1 = cur.R
				can = Stack{}
				can.Dest = start
				can.Push(PathPair{P: end, R: cur.R})
				v = make(map[PathPair]bool)
				there = true
				continue
			}
			if there && back {
				return part1, cur.R
			}
		}
		if cur.P == start && there && !back {
			can = Stack{}
			can.Dest = end
			can.Push(PathPair{P: start, R: cur.R})
			v = make(map[PathPair]bool)
			back = true
			continue
		}
		for _, d := range []Pos{Pos{-1, 0}, Pos{1, 0}, Pos{0, 0}, Pos{0, -1}, Pos{0, 1}} {
			check := PathPair{P: Pos{cur.P.X + d.X, cur.P.Y + d.Y}, R: cur.R + 1}

			// First time in this round means calculate bliz positions and cache
			if _, ok := rounds[check.R%((size.X-1)*(size.Y-1))]; !ok {
				var newBliz []Vec
				for _, b := range rounds[(check.R-1)%((size.X-1)*(size.Y-1))] {
					nb := b.Move(size)
					newBliz = append(newBliz, nb)
				}
				rounds[check.R%((size.X-1)*(size.Y-1))] = newBliz
				rounds_pos[check.R%((size.X-1)*(size.Y-1))] = BlizPos(newBliz)
			}

			// the dumb way I implemented this means I need to check start and end explicitly
			if !(check.P == start || check.P == end) {
				if check.P.X < 0 || check.P.X >= (size.X-1) || check.P.Y < 0 || check.P.Y >= (size.Y-1) {
					continue
				}
			}
			if _, ok := v[check]; ok {
				continue // already visited
			}
			if _, ok := rounds_pos[check.R%((size.X-1)*(size.Y-1))][check.P]; ok {
				// blizzard occupies our slot
				continue
			}
			can.Push(check)
		}
	}
	return part1, -1
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	var bliz []Vec
	size := Pos{Y: 0, X: 0}
	for y, line := range bytes.Split(contents, []byte("\n")) {
		if len(line) == 0 {
			continue
		}
		for x, c := range line {
			if c != '.' && c != '#' {
				bliz = append(bliz, Vec{P: Pos{X: x - 1, Y: y - 1}, Dir: c})
			}
			if x > size.X {
				size.X = x
			}
			if y > size.Y {
				size.Y = y
			}
		}
	}
	start := Pos{0, -1}
	end := Pos{size.X - 2, size.Y - 1}
	fmt.Println(Path(start, end, bliz, size, 0))
}
