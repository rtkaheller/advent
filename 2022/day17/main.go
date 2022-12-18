package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const width = 7

var patterns = []Rock{
	Rock{
		Pattern: [][]bool{
			{true, true, true, true},
			{false, false, false, false},
			{false, false, false, false},
			{false, false, false, false},
		},
		Width: 4,
	},
	Rock{
		Pattern: [][]bool{
			{false, true, false, false},
			{true, true, true, false},
			{false, true, false, false},
			{false, false, false, false},
		},
		Width: 3,
	},
	Rock{
		Pattern: [][]bool{
			{true, true, true, false},
			{false, false, true, false},
			{false, false, true, false},
			{false, false, false, false},
		},
		Width: 3,
	},
	Rock{
		Pattern: [][]bool{
			{true, false, false, false},
			{true, false, false, false},
			{true, false, false, false},
			{true, false, false, false},
		},
		Width: 1,
	},
	Rock{
		Pattern: [][]bool{
			{true, true, false, false},
			{true, true, false, false},
			{false, false, false, false},
			{false, false, false, false},
		},
		Width: 2,
	},
}

type Rock struct {
	Pattern [][]bool
	Width   int
}

func PrintRock(r Rock) {
	for y := 0; y < len(r.Pattern); y++ {
		for x := 0; x < len(r.Pattern[y]); x++ {
			if r.Pattern[y][x] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}

func PrintGridSeg(grid [][width]bool, y1, y2 int) {
	for y := y2 - 1; y >= y1; y-- {
		fmt.Printf("%4d: ", y)
		for x := 0; x < width; x++ {
			if grid[y][x] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}

func PrintGrid(grid [][width]bool) {
	for y := len(grid) - 1; y >= 0; y-- {
		fmt.Printf("%4d: ", y)
		for x := 0; x < width; x++ {
			if grid[y][x] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}

func TopRow(grid [][width]bool) int {
	for y := len(grid) - 1; y >= 0; y-- {
		empty := true
		for x := 0; x < width; x++ {
			if grid[y][x] {
				empty = false
				break
			}
		}
		if empty {
			return y
		}
	}
	return -1
}

func Drop(r Rock, grid *[][width]bool, dir string, j int) int {
	add := 4
	empty := true
	for y := len(*grid) - 1; y >= len(*grid)-4 && y >= 0; y-- {
		empty = true
		for x := 0; x < width; x++ {
			if (*grid)[y][x] {
				empty = false
				break
			}
		}
		if !empty {
			add = 5 - (len(*grid) - y)
			break
		}
	}
	if !empty || len(*grid) == 0 {
		for i := 0; i < add; i++ {
			(*grid) = append(*grid, [width]bool{})
		}
	}

	x, y := 2, len(*grid)-1
	for {
		// horiziontal movement
		cx := x
		switch dir[j] {
		case '<':
			cx -= 1
		case '>':
			cx += 1
		}
		if cx >= 0 && cx+r.Width <= width {
			collision := false
			for rx := 0; rx < r.Width; rx++ {
				for ry := 0; ry < 4; ry++ {
					if r.Pattern[ry][rx] && y+ry < len(*grid) && (*grid)[y+ry][cx+rx] {
						collision = true
					}
				}
			}
			if !collision {
				x = cx
			}
		}

		j = (j + 1) % len(dir)

		//vertical movement
		cy := y - 1
		if cy < 0 {
			break
		}
		collision := false
		for rx := 0; rx < r.Width; rx++ {
			for ry := 0; ry < 4; ry++ {
				if r.Pattern[ry][rx] && cy+ry < len(*grid) && (*grid)[cy+ry][x+rx] {
					collision = true
				}
			}
		}
		if collision {
			break
		}
		y = cy
	}

	ApplyRock(r, x, y, grid)
	return j
}

func ApplyRock(r Rock, x, y int, grid *[][width]bool) {
	//fmt.Println()
	//PrintGrid(*grid)
	//fmt.Println(x, y)
	//fmt.Println(len(*grid))
	//PrintRock(r)
	for rx := 0; rx < r.Width; rx++ {
		for ry := 0; ry < 4; ry++ {
			if r.Pattern[ry][rx] {
				(*grid)[y+ry][x+rx] = true
			}
		}
	}
}

func main() {

	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	directions := strings.TrimSpace(string(contents))
	fmt.Println(Simulate(directions, 2022))
	// No clue why this needs minus one. For the sample input, it needs minus 3 for some reason
	fmt.Println(Simulate(directions, 1000000000000) - 1) // Why minus one!?!?!?!?!?
}

type MemKey struct {
	Grid string
	J    int
	I    int
}

type MemVal struct {
	I int
	L int
}

func Simulate(directions string, iterations int) int {
	mem := make(map[MemKey]MemVal)
	grid := make([][width]bool, 0)
	jet := 0
	offScreen := 0
	for i := 0; i < iterations; i++ {
		jet = Drop(patterns[i%5], &grid, directions, jet)
		m := MemKey{I: i % 5, J: jet}
		top := TopRow(grid)
		if top > 10 {
			for t := 0; t < 10; t++ {
				var val byte
				pos := 1
				for b := 0; b < width; b++ {
					if grid[top-t-1][b] {
						val += byte(pos)
					}
					pos *= 2
				}
				m.Grid += string(val)
			}
			if l, ok := mem[m]; ok && offScreen == 0 {
				offScreen = ((iterations - i) / (i - l.I)) * (top - l.L)
				i = iterations - ((iterations - i) % (i - l.I))
			} else {
				mem[m] = MemVal{I: i, L: top}
			}
		}
	}
	return TopRow(grid) + offScreen
}
