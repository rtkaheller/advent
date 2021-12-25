package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

type Point struct {
	X, Y int
}

func Move(g [][]rune) ([][]rune, int) {
	var changed int
	var moved []Point
	for x := range g {
		for y := range g[x] {
			if g[x][y] == '>' {
				if g[x][(y+1)%len(g[x])] == '.' {
					moved = append(moved, Point{x, y})
				}
			}
		}
	}
	for _, p := range moved {
		g[p.X][(p.Y+1)%len(g[p.X])] = '>'
		g[p.X][p.Y] = '.'
		changed += 1
	}
	moved = make([]Point, 0)
	for x := range g {
		for y := range g[x] {
			if g[x][y] == 'v' {
				if g[(x+1)%len(g)][y] == '.' {
					moved = append(moved, Point{x, y})
				}
			}
		}
	}
	for _, p := range moved {
		g[(p.X+1)%len(g)][p.Y] = 'v'
		g[p.X][p.Y] = '.'
		changed += 1
	}
	return g, changed
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	//contents, err := ioutil.ReadFile("fake.txt")
	//contents, err := ioutil.ReadFile("small.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	var grid [][]rune
	for _, line := range bytes.Split(contents, []byte("\n")) {
		if len(line) == 0 {
			break
		}
		grid = append(grid, []rune{})
		for _, r := range string(line) {
			grid[len(grid)-1] = append(grid[len(grid)-1], r)
		}
	}
	c := 0
	for {
		newGrid, changed := Move(grid)
		c += 1
		if changed == 0 {
			fmt.Println(c)
			break
		}
		grid = newGrid
	}
}
