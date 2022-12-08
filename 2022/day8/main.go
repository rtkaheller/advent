package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

func Scenic(s_x, s_y int, grid [][]int) int {
	score := 1
	c := 0
	for x := s_x + 1; x < len(grid[s_y]); x++ {
		c += 1
		if grid[x][s_y] >= grid[s_x][s_y] {
			break
		}
	}
	score *= c

	c = 0
	for x := s_x - 1; x >= 0; x-- {
		c += 1
		if grid[x][s_y] >= grid[s_x][s_y] {
			break
		}
	}
	score *= c

	c = 0
	for y := s_y - 1; y >= 0; y-- {
		c += 1
		if grid[s_x][y] >= grid[s_x][s_y] {
			break
		}
	}
	score *= c

	c = 0
	for y := s_y + 1; y < len(grid); y++ {
		c += 1
		if grid[s_x][y] >= grid[s_x][s_y] {
			break
		}
	}
	score *= c
	return score
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	var grid [][]int
	var visible [][]bool
	for _, line := range bytes.Split(contents, []byte("\n")) {
		if len(line) == 0 {
			continue
		}
		grid = append(grid, []int{})
		visible = append(visible, []bool{})
		for _, n := range line {
			val, _ := strconv.Atoi(string(n))
			grid[len(grid)-1] = append(grid[len(grid)-1], val)
			visible[len(grid)-1] = append(visible[len(grid)-1], false)
		}
	}
	for y := 0; y < len(grid); y++ {
		max := -1
		for x := 0; x < len(grid[y]); x++ {
			if grid[x][y] > max {
				visible[x][y] = true
				max = grid[x][y]
			}
		}
	}
	for y := 0; y < len(grid); y++ {
		max := -1
		for x := len(grid[y]) - 1; x >= 0; x-- {
			if grid[x][y] > max {
				visible[x][y] = true
				max = grid[x][y]
			}
		}
	}
	for x := 0; x < len(grid[0]); x++ {
		max := -1
		for y := 0; y < len(grid); y++ {
			if grid[x][y] > max {
				visible[x][y] = true
				max = grid[x][y]
			}
		}
	}
	for x := 0; x < len(grid[0]); x++ {
		max := -1
		for y := len(grid) - 1; y >= 0; y-- {
			if grid[x][y] > max {
				visible[x][y] = true
				max = grid[x][y]
			}
		}
	}
	count := 0
	max := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if visible[x][y] {
				count += 1
			}
			score := Scenic(x, y, grid)
			if max < score {
				max = score
			}
		}
	}
	fmt.Println(count)
	fmt.Println(max)
}
