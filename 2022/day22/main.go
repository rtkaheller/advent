package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

type Row struct {
	Data        []byte
	First, Last int
}

type Pos struct {
	X, Y int
}

type Vec struct {
	P    Pos
	Dir  byte
	Face string
}

func DirVal(d byte) int {
	switch d {
	case '>':
		return 0
	case 'v':
		return 1
	case '<':
		return 2
	case '^':
		return 3
	}
	return -1
}

func Rotate(r string, d byte) byte {
	if r == "L" {
		switch d {
		case '>':
			return '^'
		case 'v':
			return '>'
		case '<':
			return 'v'
		case '^':
			return '<'
		}
	}
	if r == "R" {
		switch d {
		case '>':
			return 'v'
		case 'v':
			return '<'
		case '<':
			return '^'
		case '^':
			return '>'
		}
	}

	return d
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

func Apply(cur Vec, delt Pos, grid [][]byte) Vec {
	prev := cur
	n := Vec{P: Pos{X: cur.P.X + delt.X, Y: cur.P.Y + delt.Y}, Dir: cur.Dir}
	for {
		if n.P.Y >= len(grid) {
			n.P.Y = 0
			continue
		}
		if n.P.Y < 0 {
			n.P.Y = len(grid) - 1
			continue
		}
		if n.P.X >= len(grid[n.P.Y]) {
			n.P.X = 0
			continue
		}
		if n.P.X < 0 {
			n.P.X = len(grid[n.P.Y]) - 1
			continue
		}
		if grid[n.P.Y][n.P.X] == ' ' {
			n.P.X += delt.X
			n.P.Y += delt.Y
			continue
		}
		if grid[n.P.Y][n.P.X] == '#' {
			return prev
		}
		if grid[n.P.Y][n.P.X] == '.' {
			return n
		}
	}
}

func Apply2(cur Vec, cubes map[string]Face) Vec {
	reset := len(cubes[cur.Face].Grid) - 1
	prev := cur
	delt := Delta(cur.Dir)
	n := Vec{P: Pos{X: cur.P.X + delt.X, Y: cur.P.Y + delt.Y}, Dir: cur.Dir, Face: cur.Face}
	if n.P.Y >= len(cubes[n.Face].Grid) {
		switch n.Face {
		case "top":
			n.Face = "front"
			n.P.Y = 0
		case "bottom":
			n.Face = "back"
			n.P.Y, n.P.X = n.P.X, reset
			n.Dir = '<'
		case "left":
			n.Face = "back"
			n.P.Y = 0
		case "right":
			n.Face = "front"
			n.P.X, n.P.Y = reset, n.P.X
			n.Dir = '<'
		case "front":
			n.Face = "bottom"
			n.P.Y = 0
		case "back":
			n.Face = "right"
			n.P.Y = 0
		}
	}
	if n.P.Y < 0 {
		switch n.Face {
		case "top":
			n.Face = "back"
			n.P.X, n.P.Y = 0, n.P.X
			n.Dir = '>'
		case "bottom":
			n.Face = "front"
			n.P.Y = reset
		case "left":
			n.Face = "front"
			n.P.X, n.P.Y = 0, n.P.X
			n.Dir = '>'
		case "right":
			n.Face = "back"
			n.P.Y = reset
		case "front":
			n.Face = "top"
			n.P.Y = reset
		case "back":
			n.Face = "left"
			n.P.Y = reset
		}
	}
	if n.P.X >= reset+1 {
		switch n.Face {
		case "top":
			n.Face = "right"
			n.P.X = 0
		case "bottom":
			n.Face = "right"
			n.P.X = reset
			n.P.Y = reset - n.P.Y
			n.Dir = '<'
		case "left":
			n.Face = "bottom"
			n.P.X = 0
		case "right":
			n.Face = "bottom"
			n.P.X = reset
			n.P.Y = reset - n.P.Y
			n.Dir = '<'
		case "front":
			n.Face = "right"
			n.P.X, n.P.Y = n.P.Y, reset
			n.Dir = '^'
		case "back":
			n.Face = "bottom"
			n.P.X, n.P.Y = n.P.Y, reset
			n.Dir = '^'
		}
	}
	if n.P.X < 0 {
		switch n.Face {
		case "top":
			n.Face = "left"
			n.P.X = 0
			n.P.Y = reset - n.P.Y
			n.Dir = '>'
		case "bottom":
			n.Face = "left"
			n.P.X = reset
		case "left":
			n.Face = "top"
			n.P.X = 0
			n.P.Y = reset - n.P.Y
			n.Dir = '>'
		case "right":
			n.Face = "top"
			n.P.X = reset
		case "front":
			n.Face = "left"
			n.P.X, n.P.Y = n.P.Y, 0
			n.Dir = 'v'
		case "back":
			n.Face = "top"
			n.P.X, n.P.Y = n.P.Y, 0
			n.Dir = 'v'
		}
	}
	if cubes[n.Face].Grid[n.P.Y][n.P.X] == '#' {
		return prev
	}
	if cubes[n.Face].Grid[n.P.Y][n.P.X] == '.' {
		return n
	}

	return n
}

func PrintGrid(cur Vec, grid [][]byte) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if cur.P.X == x && cur.P.Y == y {
				fmt.Printf("%c", cur.Dir)
			} else {
				fmt.Printf("%c", grid[y][x])
			}
		}
		fmt.Println()
	}
}

func Move(start Vec, grid [][]byte, dirs []string) Vec {
	cur := start
	for _, d := range dirs {
		v, err := strconv.Atoi(d)
		if err != nil {
			cur.Dir = Rotate(d, cur.Dir)
			continue
		}

		delt := Delta(cur.Dir)
		for i := 0; i < v; i++ {
			cur = Apply(cur, delt, grid)
		}
	}
	return cur
}

func Move2(start Vec, cubes map[string]Face, dirs []string) Vec {
	cur := start
	for _, d := range dirs {
		v, err := strconv.Atoi(d)
		if err != nil {
			cur.Dir = Rotate(d, cur.Dir)
			continue
		}

		for i := 0; i < v; i++ {
			nex := Apply2(cur, cubes)
			cur = nex
		}
	}
	return cur
}

type Face struct {
	Grid [][]byte
	S    string
}

func ParseCube(lines [][]byte) map[string]Face {
	top := Face{S: "top"}
	right := Face{S: "right"}
	left := Face{S: "left"}
	front := Face{S: "front"}
	back := Face{S: "back"}
	bottom := Face{S: "bottom"}
	for y := 0; y < 50; y++ {
		top.Grid = append(top.Grid, lines[y][50:100])
		right.Grid = append(right.Grid, lines[y][100:])
	}
	for y := 50; y < 100; y++ {
		front.Grid = append(front.Grid, lines[y][50:])
	}
	for y := 100; y < 150; y++ {
		left.Grid = append(left.Grid, lines[y][0:50])
		bottom.Grid = append(bottom.Grid, lines[y][50:])
	}
	for y := 150; y < 200; y++ {
		back.Grid = append(back.Grid, lines[y])
	}
	return map[string]Face{
		"top":    top,
		"right":  right,
		"left":   left,
		"front":  front,
		"back":   back,
		"bottom": bottom,
	}
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	var grid [][]byte
	lines := bytes.Split(contents, []byte("\n"))
	maxRow := 0
	for _, line := range lines[:len(lines)-2] {
		if len(line) == 0 {
			continue
		}
		grid = append(grid, line)
		if len(line) > maxRow {
			maxRow = len(line)
		}
	}
	for y := 0; y < len(grid); y++ {
		for x := len(grid[y]); x < maxRow; x++ {
			grid[y] = append(grid[y], ' ')
		}
	}
	directions := lines[len(lines)-2]
	start := Vec{P: Pos{Y: 0}, Dir: '>', Face: "top"}
	for i, c := range grid[0] {
		if c == '.' {
			start.P.X = i
			break
		}
	}
	var dirs []string
	for i := 0; i < len(directions); i++ {
		if directions[i] == 'L' || directions[i] == 'R' {
			_, err := strconv.Atoi(string(directions[:i]))
			if err == nil {
				dirs = append(dirs, string(directions[:i]))
			}
			dirs = append(dirs, string(directions[i]))
			directions = directions[i+1:]
			i = 0
		}
	}
	if len(directions) > 0 {
		dirs = append(dirs, string(directions))
	}
	start.P.X -= 50
	end := Move(start, grid, dirs)
	fmt.Println(1000*(end.P.Y+1) + 4*(end.P.X+1) + DirVal(end.Dir))
	cubes := ParseCube(lines[:len(lines)-2])
	end = Move2(start, cubes, dirs)
	var fin Pos
	switch end.Face {
	case "top":
		fin.X = end.P.X + 50
		fin.Y = end.P.Y
	case "bottom":
		fin.X = end.P.X + 50
		fin.Y = end.P.Y + 100
	case "left":
		fin.X = end.P.X
		fin.Y = end.P.Y + 100
	case "right":
		fin.X = end.P.X + 100
		fin.Y = end.P.Y
	case "front":
		fin.X = end.P.X + 50
		fin.Y = end.P.Y + 50
	case "back":
		fin.X = end.P.X
		fin.Y = end.P.Y + 150
	}
	fmt.Println(1000*(fin.Y+1) + 4*(fin.X+1) + DirVal(end.Dir))
}
