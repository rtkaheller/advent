package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
)

type Coord struct {
	X, Y int
}

type stack struct {
	Data    []Coord
	Present map[Coord]bool
	F       [][]int
	G       [][]int
}

func h(start, end Coord) int {
	return int(math.Sqrt(math.Pow(float64(start.X-end.X), 2) + math.Pow(float64(start.Y-end.Y), 2)))
}

func (s *stack) Push(val Coord) {
	if _, ok := s.Present[val]; !ok {
		s.Data = append(s.Data, val)
		s.Present[val] = true
	}
}

func (s *stack) Pop() Coord {
	val := s.Data[0]
	s.Data = s.Data[1:]
	s.Present[val] = false
	return val
}

func (s stack) Len() int {
	return len(s.Data)
}

func (s stack) Less(i, j int) bool {
	a := s.Data[i]
	b := s.Data[j]
	return s.F[a.X][a.Y] < s.F[b.X][b.Y]
}
func (s stack) Swap(i, j int) {
	s.Data[i], s.Data[j] = s.Data[j], s.Data[i]
}

func Path(start, end Coord, grid [][]int) int {
	var data stack
	data.Present = make(map[Coord]bool)
	for x := range grid {
		data.F = append(data.F, make([]int, len(grid[x])))
		data.G = append(data.G, make([]int, len(grid[x])))
		for y := range grid[x] {
			data.F[x][y] = math.MaxInt64
			data.G[x][y] = math.MaxInt64
		}
	}
	data.G[start.X][start.Y] = 0
	data.F[start.X][start.Y] = h(start, end)
	data.Push(start)
	for data.Len() != 0 {
		sort.Sort(data)
		cur := data.Pop()

		if cur.X == end.X && cur.Y == end.Y {
			return data.G[cur.X][cur.Y]
		}

		if cur.X-1 >= 0 {
			try := Coord{X: cur.X - 1, Y: cur.Y}
			nDist := data.G[cur.X][cur.Y] + grid[try.X][try.Y]
			if nDist < data.G[try.X][try.Y] {
				data.G[try.X][try.Y] = nDist
				data.F[try.X][try.Y] = nDist + h(try, end)
				data.Push(try)
			}
		}
		if cur.Y-1 >= 0 {
			try := Coord{X: cur.X, Y: cur.Y - 1}
			nDist := data.G[cur.X][cur.Y] + grid[try.X][try.Y]
			if nDist < data.G[try.X][try.Y] {
				data.G[try.X][try.Y] = nDist
				data.F[try.X][try.Y] = nDist + h(try, end)
				data.Push(try)
			}
		}
		if cur.X+1 < len(grid) {
			try := Coord{X: cur.X + 1, Y: cur.Y}
			nDist := data.G[cur.X][cur.Y] + grid[try.X][try.Y]
			if nDist < data.G[try.X][try.Y] {
				data.G[try.X][try.Y] = nDist
				data.F[try.X][try.Y] = nDist + h(try, end)
				data.Push(try)
			}
		}
		if cur.Y+1 < len(grid[cur.X]) {
			try := Coord{X: cur.X, Y: cur.Y + 1}
			nDist := data.G[cur.X][cur.Y] + grid[try.X][try.Y]
			if nDist < data.G[try.X][try.Y] {
				data.G[try.X][try.Y] = nDist
				data.F[try.X][try.Y] = nDist + h(try, end)
				data.Push(try)
			}
		}
	}
	return data.G[end.X-1][end.Y-1]
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	//contents, err := ioutil.ReadFile("fake.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	var risk [][]int
	for x, line := range bytes.Split(contents, []byte("\n")) {
		if len(line) == 0 {
			break
		}
		risk = append(risk, make([]int, len(string(line))))
		for y, r := range string(line) {
			if len(line) == 0 {
				break
			}

			val, _ := strconv.Atoi(string(r))
			risk[x][y] = val
		}
	}
	fmt.Println(Path(Coord{X: 0, Y: 0}, Coord{X: len(risk), Y: len(risk[0])}, risk))
	var bigRisk [][]int
	for x := range risk {
		bigRisk = append(bigRisk, make([]int, len(risk[x])*5))
		for i := 0; i < 5; i++ {
			for y := range risk[x] {
				newVal := risk[x][y] + i
				if newVal > 9 {
					newVal = newVal - 9
				}
				bigRisk[x][y+len(risk[x])*i] = newVal
			}
		}
	}
	for i := 1; i < 5; i++ {
		for x := range risk {
			bigRisk = append(bigRisk, make([]int, len(bigRisk[x])))
			for y := range bigRisk[x] {
				newVal := bigRisk[x][y] + i
				if newVal > 9 {
					newVal = newVal - 9
				}
				bigRisk[x+len(risk[x])*i][y] = newVal
			}
		}
	}
	//for x := range bigRisk {
	//  for y := range bigRisk[x] {
	//    fmt.Printf("%v", bigRisk[x][y])
	//  }
	//  fmt.Println()
	//}
	fmt.Println(Path(Coord{X: 0, Y: 0}, Coord{X: len(bigRisk), Y: len(bigRisk[0])}, bigRisk))
}
