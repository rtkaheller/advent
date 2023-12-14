package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
)

func Draw(g [][]byte) {
	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[y]); x++ {
			fmt.Printf("%s", string(g[y][x]))
		}
		fmt.Printf("\n")
	}
}

type Pos struct {
	X, Y int
}

func Roll(rocks []*Pos, grid [][]byte, dir Pos) bool {
	moved := true
	loops := 0
	for ; moved; loops++ {
		moved = false
		for _, r := range rocks {
			if (*r).Y+dir.Y >= 0 && (*r).Y+dir.Y < len(grid) && (*r).X+dir.X >= 0 && (*r).X+dir.X < len(grid[(*r).Y]) && grid[(*r).Y+dir.Y][(*r).X+dir.X] == '.' {
				(*r).Y += dir.Y
				(*r).X += dir.X
				grid[(*r).Y-dir.Y][(*r).X-dir.X] = '.'
				grid[(*r).Y][(*r).X] = 'O'
				moved = true
			}
		}
	}
	return loops != 0
}

func main() {
	contents, err := ioutil.ReadFile("fake.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	rocks := []*Pos{}
	grid := [][]byte{}
	lines := bytes.Split(contents, []byte("\n"))
	var m Pos
	for y, line := range lines {
		if len(line) == 0 {
			continue
		}
		grid = append(grid, []byte{})
		for x, c := range line {
			grid[y] = append(grid[y], c)
			if c == 'O' {
				rocks = append(rocks, &Pos{x, y})
			}
			m.X = x
		}
		m.Y = y
	}
	Roll(rocks, grid, Pos{-1, 0})
	sum := 0
	for _, r := range rocks {
		sum += m.Y - (*r).Y + 1
	}
	fmt.Println(sum)
	rocks = []*Pos{}
	grid = [][]byte{}
	for y, line := range lines {
		if len(line) == 0 {
			continue
		}
		grid = append(grid, []byte{})
		for x, c := range line {
			grid[y] = append(grid[y], c)
			if c == 'O' {
				rocks = append(rocks, &Pos{x, y})
			}
			m.X = x
		}
		m.Y = y
	}
	moved := true
	memo := map[string][]Pos{}
	last := map[string]int{}
	lastMap := map[int]int{}
	prev := 0
	iters := 1000000000
	for i := 0; i < iters && moved; i++ {
		key := MemoKey(rocks)
		if cache, ok := memo[key]; ok {
			if last[key] < prev {
				// Loop detected
				check := ((iters - last[key] - 1) % (prev - last[key] + 1)) + last[key]
				fmt.Println(lastMap[check])
				break
			}
			for i, r := range cache {
				rocks[i].X = r.X
				rocks[i].Y = r.Y
			}
			prev = last[key]
			continue
		}
		moved = false
		if Roll(rocks, grid, Pos{0, -1}) {
			moved = true
		}
		if Roll(rocks, grid, Pos{-1, 0}) {
			moved = true
		}
		if Roll(rocks, grid, Pos{0, 1}) {
			moved = true
		}
		if Roll(rocks, grid, Pos{1, 0}) {
			moved = true
		}
		sum = 0
		for _, r := range rocks {
			sum += m.Y - (*r).Y + 1
		}
		toCache := []Pos{}
		for _, r := range rocks {
			toCache = append(toCache, Pos{(*r).X, (*r).Y})
		}
		memo[key] = toCache
		if _, ok := last[key]; !ok {
			last[key] = i
		}
		lastMap[i] = sum
	}
}

func MemoGrid(grid [][]byte) string {
	key := ""
	for _, row := range grid {
		key += string(row)
	}
	return key
}

func MemoKey(rocks []*Pos) string {
	dup := []Pos{}
	for _, rock := range rocks {
		dup = append(dup, *rock)
	}
	sort.Slice(dup, func(i, j int) bool {
		if dup[i].X == dup[j].X {
			return dup[i].Y <= dup[j].Y
		}
		return dup[i].X <= dup[j].X
	})
	result := ""
	for _, rock := range dup {
		result += fmt.Sprintf("[%v,%v]", rock.X, rock.Y)
	}
	return result
}
