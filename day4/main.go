package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

const (
	gridSize = 5
)

type square struct {
	Val    int
	Called bool
}

type grid struct {
	Grid [5][5]square
	Won  bool
}

func (g *grid) score() int {
	score := 0
	for x := 0; x < gridSize; x++ {
		for y := 0; y < gridSize; y++ {
			if !g.Grid[x][y].Called {
				score += g.Grid[x][y].Val
			}
		}
	}
	return score
}

func (g *grid) call(val int) {
	for x := 0; x < gridSize; x++ {
		for y := 0; y < gridSize; y++ {
			if g.Grid[x][y].Val == val {
				g.Grid[x][y].Called = true
				return
			}
		}
	}
}

func (g *grid) win() bool {
	if g.Won {
		return g.Won
	}
	for x := 0; x < gridSize; x++ {
		winner := true
		for y := 0; y < gridSize; y++ {
			if !g.Grid[x][y].Called {
				winner = false
				break
			}
		}
		if winner {
			g.Won = true
			return true
		}
	}
	for y := 0; y < gridSize; y++ {
		winner := true
		for x := 0; x < gridSize; x++ {
			if !g.Grid[x][y].Called {
				winner = false
				break
			}
		}
		if winner {
			g.Won = true
			return true
		}
	}
	return false
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := bytes.Split(contents, []byte("\n"))
	calls := bytes.Split(lines[0], []byte(","))
	var grids []*grid
	var newGrid *grid
	y := 0
	for _, line := range lines[1:] {
		if len(line) == 0 {
			if newGrid != nil {
				grids = append(grids, newGrid)
			}
			newGrid = new(grid)
			y = -1
		}
		x := 0
		for _, num := range bytes.Split(line, []byte(" ")) {
			if len(num) != 0 {
				val, _ := strconv.Atoi(string(num))
				newGrid.Grid[x][y] = square{Val: val, Called: false}
				x += 1
			}
		}
		y += 1
	}
	for _, call := range calls {
		for _, grid := range grids {
			callVal, _ := strconv.Atoi(string(call))
			grid.call(callVal)
			if !grid.Won && grid.win() {
				fmt.Println(grid.score())
				fmt.Println(callVal)
				fmt.Println(grid.score() * callVal)
			}
		}
	}
}
