package main

import (
	"bytes"
	"fmt"
	"os"
)

func checkAdj(x, y int, grid [][]bool) int {
	count := 0
	for dY := -1; dY <= 1; dY++ {
		for dX := -1; dX <= 1; dX++ {
			if !(dX == 0 && dY == 0) {
				if y+dY >= 0 && y+dY < len(grid) &&
					x+dX >= 0 && x+dX < len(grid[y]) {
					if grid[y+dY][x+dX] {
						count += 1
					}
				}
			}
		}
	}
	return count
}

func remove(grid [][]bool) ([][]bool, int) {
	newGrid := [][]bool{}
	c := 0
	for y := range grid {
		row := []bool{}
		for x := range grid[y] {
			n := grid[y][x] && checkAdj(x, y, grid) >= 4
			if n != grid[y][x] {
				c += 1
			}
			row = append(row, n)
		}
		newGrid = append(newGrid, row)
	}
	return newGrid, c
}

func repr(grid [][]bool) string {
	s := ""
	for y := range grid {
		for x := range grid[y] {
			if !grid[y][x] {
				s += "."
			} else {
				s += "@"
			}
		}
	}
	return s
}

func render(grid [][]bool) {
	for y := range grid {
		for x := range grid[y] {
			if !grid[y][x] {
				fmt.Print(".")
			} else if checkAdj(x, y, grid) < 4 {
				fmt.Print("x")
			} else {
				fmt.Print("@")
			}
		}
		fmt.Println()
	}
}

func main() {
	contents, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	grid := [][]bool{}

	p1 := 0
	p2 := 0
	lines := bytes.Split(contents, []byte("\n"))
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		row := []bool{}
		for _, c := range line {
			row = append(row, c == '@')
		}
		grid = append(grid, row)
	}

	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] && checkAdj(x, y, grid) < 4 {
				p1 += 1
			}
		}
	}
	fmt.Println(p1)

	c := 0
	for {
		cur := repr(grid)
		grid, c = remove(grid)
		p2 += c
		if repr(grid) == cur {
			break
		}
	}
	fmt.Println(p2)
}
