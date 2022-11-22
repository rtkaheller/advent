package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
)

type Point struct {
	Vert       bool
	Horizontal bool
	Val        int
	Basin      bool
}

type Grid struct {
	Points [][]*Point
}

func (g *Grid) Basin(x, y, height, width int) int {
	if y < 0 || y >= height || x < 0 || x >= width {
		return 0
	}
	if g.Points[x][y].Val == 9 {
		return 0
	}
	if g.Points[x][y].Basin {
		return 0
	}
	g.Points[x][y].Basin = true
	return g.Basin(x-1, y, height, width) + g.Basin(x+1, y, height, width) + g.Basin(x, y-1, height, width) + g.Basin(x, y+1, height, width) + 1
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	var g Grid
	var height, width int
	for y, line := range bytes.Split(contents, []byte("\n")) {
		if len(line) == 0 {
			continue
		}
		width += 1
		g.Points = append(g.Points, []*Point{})
		height = 0
		for _, r := range string(line) {
			height += 1
			val, _ := strconv.Atoi(string(r))
			g.Points[y] = append(g.Points[y], &Point{Val: val})
		}
	}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			g.Points[x][y].Horizontal = true
			if x > 0 {
				if g.Points[x][y].Val >= g.Points[x-1][y].Val {
					g.Points[x][y].Horizontal = false
				}
			}
			if x < width-1 {
				if g.Points[x][y].Val >= g.Points[x+1][y].Val {
					g.Points[x][y].Horizontal = false
				}
			}
		}
	}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			g.Points[x][y].Vert = true
			if y > 0 {
				if g.Points[x][y].Val >= g.Points[x][y-1].Val {
					g.Points[x][y].Vert = false
				}
			}
			if y < height-1 {
				if g.Points[x][y].Val >= g.Points[x][y+1].Val {
					g.Points[x][y].Vert = false
				}
			}
		}
	}
	var basins []int
	sum := 0
	for x, row := range g.Points {
		for y, p := range row {
			if p.Horizontal && p.Vert {
				basins = append(basins, g.Basin(x, y, height, width))
				sum += g.Points[x][y].Val + 1
			}
		}
	}
	sort.Ints(basins)
	fmt.Println(sum)
	score := 1
	for i := len(basins) - 1; i >= len(basins)-3; i-- {
		score *= basins[i]
	}
	fmt.Println(score)
}
