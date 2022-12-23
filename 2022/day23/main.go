package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

var dirs = [][]Pos{
	[]Pos{
		Pos{X: -1, Y: -1},
		Pos{X: 0, Y: -1},
		Pos{X: 1, Y: -1},
	},
	[]Pos{
		Pos{X: -1, Y: 1},
		Pos{X: 0, Y: 1},
		Pos{X: 1, Y: 1},
	},
	[]Pos{
		Pos{X: -1, Y: -1},
		Pos{X: -1, Y: 0},
		Pos{X: -1, Y: 1},
	},
	[]Pos{
		Pos{X: 1, Y: -1},
		Pos{X: 1, Y: 0},
		Pos{X: 1, Y: 1},
	},
}

type Pos struct {
	X, Y int
}

func ElfEdge(elves map[Pos]bool) (Pos, Pos) {
	var min, max Pos
	for elf, _ := range elves {
		if elf.X < min.X {
			min.X = elf.X
		}
		if elf.Y > max.Y {
			max.Y = elf.Y
		}
		if elf.Y < min.Y {
			min.Y = elf.Y
		}
		if elf.X > max.X {
			max.X = elf.X
		}
	}
	return min, max
}

func PrintElves(elves map[Pos]bool) {
	min, max := ElfEdge(elves)
	for y := min.Y; y < max.Y+1; y++ {
		for x := min.X; x < max.X+1; x++ {
			if _, ok := elves[Pos{x, y}]; ok {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}

func Move(elves map[Pos]bool, round int) map[Pos]bool {
	prop := make(map[Pos]Pos)
	prop_c := make(map[Pos]int)
	for elf, _ := range elves {
		safe := true
		for x := -1; x <= 1; x++ {
			for y := -1; y <= 1; y++ {
				if _, ok := elves[Pos{elf.X + x, elf.Y + y}]; !(x == 0 && y == 0) && ok {
					safe = false
					break
				}
			}
			if !safe {
				break
			}
		}
		if !safe {
			prop[elf] = elf
			prop_c[elf] += 1
			for i := 0; i < 4; i++ {
				move := true
				for _, d := range dirs[(round+i)%4] {
					if _, ok := elves[Pos{elf.X + d.X, elf.Y + d.Y}]; ok {
						move = false
						break
					}
				}
				if move {
					prop[elf] = Pos{elf.X + dirs[(round+i)%4][1].X, elf.Y + dirs[(round+i)%4][1].Y}
					prop_c[elf] -= 1
					prop_c[prop[elf]] += 1
					break
				}
			}
		} else {
			prop[elf] = elf
			prop_c[elf] = 1
		}
	}
	ret := make(map[Pos]bool)
	for c, p := range prop {
		if prop_c[p] == 1 {
			ret[p] = true
		} else {
			ret[c] = true
		}
	}
	return ret
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	grid := make(map[Pos]bool)
	for y, line := range bytes.Split(contents, []byte("\n")) {
		if len(line) == 0 {
			continue
		}
		for x, c := range line {
			if c == '#' {
				grid[Pos{x, y}] = true
			}
		}
	}
	for i := 0; i < 10; i++ {
		grid = Move(grid, i)
	}
	min, max := ElfEdge(grid)
	fmt.Println((max.X-min.X+1)*(max.Y-min.Y+1) - len(grid))
	round := 10
	for {
		grid2 := Move(grid, round)
		all := true
		for elf, _ := range grid2 {
			if _, ok := grid[elf]; !ok {
				all = false
				break
			}
		}
		round += 1
		if all {
			fmt.Println(round)
			break
		}
		grid = grid2
	}
}
