package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

type Pos struct {
	X, Y int
}

type Path struct {
	Points []Pos
}

func (p *Pos) Next(dest Pos) {
	if p.X == dest.X {
		if p.Y > dest.Y {
			p.Y -= 1
		} else {
			p.Y += 1
		}
	} else {
		if p.X > dest.X {
			p.X -= 1
		} else {
			p.X += 1
		}
	}
}

func (p *Path) Expand() []Pos {
	var ret []Pos
	for i := 1; i < len(p.Points); i++ {
		for r := p.Points[i-1]; r != p.Points[i]; r.Next(p.Points[i]) {
			ret = append(ret, r)
		}
		ret = append(ret, p.Points[i])
	}
	return ret
}

func Sand(grid map[Pos]byte, start Pos, maxY int) int {
	c := 0
	sand := start
	for {
		if sand.Y > maxY {
			break
		}
		_, ok := grid[sand]
		if !ok {
			// not obstructed
			sand.Y += 1
			continue
		}
		if sand == start {
			break
		}
		if _, ok := grid[Pos{sand.X - 1, sand.Y}]; !ok {
			// down and to the left is free
			sand.X -= 1
			continue
		}
		if _, ok := grid[Pos{sand.X + 1, sand.Y}]; !ok {
			sand.X += 1
			continue
		}
		grid[Pos{X: sand.X, Y: sand.Y - 1}] = 'o'
		c += 1
		sand = start
	}
	return c
}

func main() {
	contents, err := ioutil.ReadFile("fake.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	var paths []Path

	for _, line := range bytes.Split(contents, []byte("\n")) {
		if len(line) == 0 {
			continue
		}

		paths = append(paths, Path{})
		for _, pair := range bytes.Split(line, []byte(" -> ")) {
			coord := bytes.Split(pair, []byte(","))
			x, _ := strconv.Atoi(string(coord[0]))
			y, _ := strconv.Atoi(string(coord[1]))
			paths[len(paths)-1].Points = append(paths[len(paths)-1].Points, Pos{X: x, Y: y})
		}
	}

	grid := make(map[Pos]byte)

	maxY := 0
	for _, p := range paths {
		for _, exp := range p.Expand() {
			if exp.Y > maxY {
				maxY = exp.Y
			}
			grid[exp] = '#'
		}
	}

	fmt.Println(Sand(grid, Pos{X: 500, Y: 0}, maxY))

	grid = make(map[Pos]byte)

	for _, p := range paths {
		for _, exp := range p.Expand() {
			grid[exp] = '#'
		}
	}

	for x := -5000; x < 5000; x++ {
		grid[Pos{X: x, Y: maxY + 2}] = '#'
	}
	fmt.Println(Sand(grid, Pos{X: 500, Y: 0}, maxY+2))
}
