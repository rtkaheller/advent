package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"sort"
)

type Pos struct {
	X, Y int
}

func PrintDist(grid [][]int) {
	fmt.Printf("   ")
	for x := 0; x < len(grid[0]); x++ {
		fmt.Printf("%3d", x)
	}
	fmt.Println()
	for y := 0; y < len(grid); y++ {
		fmt.Printf("%2d ", y)
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] < 0 || grid[y][x] > 10000 {
				fmt.Printf(" no")
			} else {
				fmt.Printf("%3v", grid[y][x])
			}
		}
		fmt.Println()
	}
}

func PrintGrid(grid [][]byte) {
	fmt.Printf(" ")
	for x := 0; x < len(grid[0]); x++ {
		fmt.Printf("%3d", x)
	}
	fmt.Println()
	for y := 0; y < len(grid); y++ {
		fmt.Printf("%2d ", y)
		for x := 0; x < len(grid[y]); x++ {
			fmt.Printf("%c  ", grid[y][x])
		}
		fmt.Println()
	}
}

type Stack struct {
	Data []Pos
	Dist [][]int
}

func (s *Stack) Push(v Pos) {
	s.Data = append(s.Data, v)
}

func (s *Stack) Pop() Pos {
	val := s.Data[0]
	s.Data = s.Data[1:]
	return val
}

func (s Stack) Len() int {
	return len(s.Data)
}

func (s Stack) Less(i, j int) bool {
	a := s.Data[i]
	b := s.Data[j]
	return s.Dist[a.Y][a.X] < s.Dist[b.Y][b.X]
}

func (s Stack) Swap(i, j int) {
	s.Data[i], s.Data[j] = s.Data[j], s.Data[i]
}

func FindLowGround(end Pos, grid [][]byte) int {
	v := make(map[Pos]bool) //visited
	var can Stack           //candidates

	for y := 0; y < len(grid); y++ {
		can.Dist = append(can.Dist, []int{})
		for x := 0; x < len(grid[y]); x++ {
			can.Dist[y] = append(can.Dist[y], math.MaxInt)
			can.Push(Pos{X: x, Y: y})
		}
	}
	can.Dist[end.Y][end.X] = 0

	for can.Len() > 0 {
		sort.Sort(can)
		cur := can.Pop()
		v[cur] = true
		if can.Dist[cur.Y][cur.X] == math.MaxInt {
			// If we are checking a node that isn't a neighbour to one we visited, we have no path :(
			break
		}
		for _, d := range []Pos{Pos{-1, 0}, Pos{1, 0}, Pos{0, -1}, Pos{0, 1}} {
			check := Pos{X: d.X + cur.X, Y: d.Y + cur.Y}
			if check.X < 0 || check.X >= len(grid[0]) || check.Y < 0 || check.Y >= len(grid) {
				continue //out of bounds
			}
			if _, ok := v[check]; ok {
				continue // already visited
			}
			if int(grid[cur.Y][cur.X])-int(grid[check.Y][check.X]) > 1 {
				continue // more than one step down
			}
			newDist := can.Dist[cur.Y][cur.X] + 1
			if newDist < can.Dist[check.Y][check.X] {
				can.Dist[check.Y][check.X] = newDist
			}
		}
	}
	min := math.MaxInt
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 'a' {
				if can.Dist[y][x] < min {
					min = can.Dist[y][x]
				}
			}
		}
	}
	return min
}

func FindPath(start, end Pos, grid [][]byte) int {
	v := make(map[Pos]bool) //visited
	var can Stack           //candidates

	for y := 0; y < len(grid); y++ {
		can.Dist = append(can.Dist, []int{})
		for x := 0; x < len(grid[y]); x++ {
			can.Dist[y] = append(can.Dist[y], math.MaxInt)
			can.Push(Pos{X: x, Y: y})
		}
	}
	can.Dist[start.Y][start.X] = 0

	for can.Len() > 0 {
		sort.Sort(can)
		cur := can.Pop()
		v[cur] = true
		if can.Dist[cur.Y][cur.X] == math.MaxInt {
			// If we are checking a node that isn't a neighbour to one we visited, we have no path :(
			return -1
		}
		if cur == end {
			break
		}
		for _, d := range []Pos{Pos{-1, 0}, Pos{1, 0}, Pos{0, -1}, Pos{0, 1}} {
			check := Pos{X: d.X + cur.X, Y: d.Y + cur.Y}
			if check.X < 0 || check.X >= len(grid[0]) || check.Y < 0 || check.Y >= len(grid) {
				continue //out of bounds
			}
			if _, ok := v[check]; ok {
				continue // already visited
			}
			if int(grid[check.Y][check.X])-int(grid[cur.Y][cur.X]) > 1 {
				continue // more than one step up
			}
			newDist := can.Dist[cur.Y][cur.X] + 1
			if newDist < can.Dist[check.Y][check.X] {
				can.Dist[check.Y][check.X] = newDist
			}
		}
	}
	return can.Dist[end.Y][end.X]
}

func main() {
	contents, err := ioutil.ReadFile("fake.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	var grid [][]byte
	var start, end Pos
	for y, line := range bytes.Split(contents, []byte("\n")) {
		if len(line) == 0 {
			continue
		}
		grid = append(grid, []byte{})
		for x, c := range line {
			switch c {
			case 'S':
				start.X = x
				start.Y = y
				grid[y] = append(grid[y], 'a')
			case 'E':
				end.X = x
				end.Y = y
				grid[y] = append(grid[y], 'z')
			default:
				grid[y] = append(grid[y], c)
			}
		}
	}
	fmt.Println(FindPath(start, end, grid))
	fmt.Println(FindLowGround(end, grid))
}
