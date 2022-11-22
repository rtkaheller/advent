package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

const (
	rounds = 50
)

func pixelValue(x, y int, grid [][]rune, alg []rune, inf rune) rune {
	bitString := ""
	zeroBit := "0"
	if inf == '#' {
		zeroBit = "1"
	}
	for i := -1; i <= 1; i++ {
		for t := -1; t <= 1; t++ {
			if x+i < 0 {
				bitString += zeroBit
			} else if x+i >= len(grid) {
				bitString += zeroBit
			} else if y+t < 0 {
				bitString += zeroBit
			} else if y+t >= len(grid[x+i]) {
				bitString += zeroBit
			} else {
				if grid[x+i][y+t] == '#' {
					bitString += "1"
				} else {
					bitString += "0"
				}
			}
		}
	}
	if len(bitString) != 9 {
		fmt.Println("wtf")
	}
	index, _ := strconv.ParseInt(bitString, 2, 16)
	return alg[index]
}

func expand(grid [][]rune, alg []rune, inf rune) [][]rune {
	var newGrid [][]rune
	for x := -1; x <= len(grid); x++ {
		newGrid = append(newGrid, []rune{})
		for y := -1; y <= len(grid[0]); y++ {
			newGrid[len(newGrid)-1] = append(newGrid[len(newGrid)-1], pixelValue(x, y, grid, alg, inf))
		}
	}
	return newGrid
}

func countGrid(grid [][]rune) int {
	c := 0
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == '#' {
				c++
			}
		}
	}
	return c
}
func printGrid(grid [][]rune) {
	for x := range grid {
		for y := range grid[x] {
			fmt.Printf(string(grid[x][y]))
		}
		fmt.Println()
	}
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	//contents, err := ioutil.ReadFile("fake.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := bytes.Split(contents, []byte("\n"))
	alg := []rune(string(lines[0]))
	var grid [][]rune
	for _, line := range lines[2:] {
		if len(line) == 0 {
			continue
		}
		grid = append(grid, []rune{})
		for _, r := range string(line) {
			grid[len(grid)-1] = append(grid[len(grid)-1], r)
		}
	}
	inf := '.'
	for i := 0; i < rounds; i++ {
		grid = expand(grid, alg, inf)
		inf = pixelValue(-5, -5, grid, alg, inf)
	}
	fmt.Println(countGrid(grid))
}
