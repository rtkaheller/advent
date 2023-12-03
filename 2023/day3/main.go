package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

type Pos struct {
	X, Y int
}

type Number struct {
	Val int
	P   []Pos
}

type Point struct {
	N *Number
	C byte
}

func main() {
	contents, err := ioutil.ReadFile("fake.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	sum := 0
	lines := bytes.Split(contents, []byte("\n"))
	var grid [][]Point
	var nums []Number
	for x, line := range lines {
		if len(line) == 0 {
			continue
		}

		grid = append(grid, []Point{})
		cur_num := ""
		num_start := 0
		for y, c := range line {
			var p Point
			p.C = c
			_, err := strconv.Atoi(string(c))
			if err == nil {
				if cur_num == "" {
					num_start = y
				}
				cur_num += string(c)
			} else {
				if cur_num != "" {
					v, _ := strconv.Atoi(cur_num)
					var n Number
					n.Val = v
					n.P = []Pos{}
					for t := num_start; t < y; t++ {
						n.P = append(n.P, Pos{x, t})
						grid[x][t].N = &n
					}
					nums = append(nums, n)
					cur_num = ""
				}
			}
			grid[x] = append(grid[x], p)
		}
		if cur_num != "" {
			v, _ := strconv.Atoi(cur_num)
			var n Number
			n.Val = v
			n.P = []Pos{}
			for t := num_start; t < len(grid[x]); t++ {
				n.P = append(n.P, Pos{x, t})
				grid[x][t].N = &n
			}
			nums = append(nums, n)
			cur_num = ""
		}
	}
	parts := map[*Number]bool{}
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y].C != '.' {
				_, err := strconv.Atoi(string(grid[x][y].C))
				if err != nil {
					for dx := -1; dx <= 1; dx++ {
						for dy := -1; dy <= 1; dy++ {
							if grid[x+dx][y+dy].N != nil {
								parts[grid[x+dx][y+dy].N] = true
							}
						}
					}
				}
			}
		}
	}
	for p, _ := range parts {
		sum += p.Val
	}
	fmt.Println(sum)

	gear_sum := 0
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y].C != '.' {
				_, err := strconv.Atoi(string(grid[x][y].C))
				if err != nil {
					c := map[*Number]bool{}
					for dx := -1; dx <= 1; dx++ {
						for dy := -1; dy <= 1; dy++ {
							if grid[x+dx][y+dy].N != nil {
								c[grid[x+dx][y+dy].N] = true
							}
						}
					}
					if len(c) == 2 {
						s := 1
						for g, _ := range c {
							s *= g.Val
						}
						gear_sum += s
					}
				}
			}
		}
	}
	fmt.Println(gear_sum)
}
