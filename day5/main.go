package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

type Coord struct {
	X, Y int
}

type Line struct {
	Start, End Coord
}

type Grid struct {
	Points [][]int
}

func (g *Grid) AddLine(l *Line) {
	if l.Start.X == l.End.X {
		if l.Start.Y < l.End.Y {
			for i := l.Start.Y; i <= l.End.Y; i++ {
				g.Points[l.Start.X][i] += 1
			}
		} else {
			for i := l.End.Y; i <= l.Start.Y; i++ {
				g.Points[l.Start.X][i] += 1
			}
		}
	} else if l.Start.Y == l.End.Y {
		if l.Start.X < l.End.X {
			for i := l.Start.X; i <= l.End.X; i++ {
				g.Points[i][l.Start.Y] += 1
			}
		} else {
			for i := l.End.X; i <= l.Start.X; i++ {
				g.Points[i][l.Start.Y] += 1
			}
		}
	} else {
		var dx, dy int
		if l.Start.X < l.End.X {
			dx = 1
		} else {
			dx = -1
		}
		if l.Start.Y < l.End.Y {
			dy = 1
		} else {
			dy = -1
		}
		for x, y := l.Start.X, l.Start.Y; x != l.End.X+dx && y != l.End.Y+dy; x, y = x+dx, y+dy {
			g.Points[x][y] += 1
		}
	}
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	txtlines := bytes.Split(contents, []byte("\n"))
	var maxX, maxY int
	var Lines []*Line
	for _, line := range txtlines {
		if len(line) == 0 {
			continue
		}
		newLine := new(Line)
		points := bytes.Split(line, []byte(" -> "))
		start := bytes.Split(points[0], []byte(","))
		end := bytes.Split(points[1], []byte(","))
		x, _ := strconv.Atoi(string(start[0]))
		y, _ := strconv.Atoi(string(start[1]))
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
		newLine.Start = Coord{X: x, Y: y}

		x, _ = strconv.Atoi(string(end[0]))
		y, _ = strconv.Atoi(string(end[1]))
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
		newLine.End = Coord{X: x, Y: y}
		Lines = append(Lines, newLine)
	}
	var g Grid
	g.Points = make([][]int, maxX+1)
	for i := 0; i < maxX+1; i++ {
		g.Points[i] = make([]int, maxY+1)
	}
	for _, line := range Lines {
		g.AddLine(line)
	}
	c := 0
	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			if g.Points[x][y] > 1 {
				c++
			}
		}
	}
	fmt.Println(c)
}
