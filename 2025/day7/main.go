package main

import (
	"bytes"
	"fmt"
	"os"
)

type Pos struct {
	X, Y int
}

var memo = map[Pos]int{}

func Zap(start Pos, splitters map[Pos]bool, depth int) int {
	if v, ok := memo[start]; ok {
		return v
	}
	if depth <= 0 {
		return 1
	}
	r := 0
	if _, ok := splitters[start]; ok {
		r += Zap(Pos{start.X - 1, start.Y}, splitters, depth)
		r += Zap(Pos{start.X + 1, start.Y}, splitters, depth)
	} else {
		r = Zap(Pos{start.X, start.Y + 1}, splitters, depth-1)
	}
	memo[start] = r
	return r
}

func main() {
	contents, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	p1 := 0
	p2 := 0

	splitters := map[Pos]bool{}
	var start Pos

	lines := bytes.Split(contents, []byte("\n"))
	for y, line := range lines {
		if len(line) == 0 {
			continue
		}
		for x, c := range line {
			switch c {
			case 'S':
				start = Pos{x, y}
			case '^':
				splitters[Pos{x, y}] = true
			}
		}
	}

	beams := map[Pos]bool{start: true}

	for range len(lines) - 1 {
		newBeams := map[Pos]bool{}
		for beam := range beams {
			beam.Y += 1
			if _, ok := splitters[beam]; ok {
				p1 += 1
				newBeams[Pos{beam.X - 1, beam.Y}] = true
				newBeams[Pos{beam.X + 1, beam.Y}] = true
			} else {
				newBeams[Pos{beam.X, beam.Y}] = true
			}
		}
		beams = newBeams
	}
	p2 = Zap(start, splitters, len(lines)-1)

	fmt.Println(p1)
	fmt.Println(p2)
}
