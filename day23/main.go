package main

import (
	"fmt"
)

var energy = map[rune]int{
	'A': 1,
	'B': 10,
	'C': 100,
	'D': 1000,
}

var home = map[rune]int{
	'A': 3,
	'B': 5,
	'C': 7,
	'D': 9,
}

var solved = [7][13]rune{
	[13]rune{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
	[13]rune{'#', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '#'},
	[13]rune{'#', '#', '#', 'A', '#', 'B', '#', 'C', '#', 'D', '#', '#', '#'},
	[13]rune{' ', ' ', '#', 'A', '#', 'B', '#', 'C', '#', 'D', '#', ' ', ' '},
	[13]rune{' ', ' ', '#', 'A', '#', 'B', '#', 'C', '#', 'D', '#', ' ', ' '},
	[13]rune{' ', ' ', '#', 'A', '#', 'B', '#', 'C', '#', 'D', '#', ' ', ' '},
	[13]rune{' ', ' ', '#', '#', '#', '#', '#', '#', '#', '#', '#', ' ', ' '},
}

func sample() [7][13]rune {
	return [7][13]rune{
		[13]rune{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
		[13]rune{'#', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '#'},
		[13]rune{'#', '#', '#', 'B', '#', 'C', '#', 'B', '#', 'D', '#', '#', '#'},
		[13]rune{' ', ' ', '#', 'D', '#', 'C', '#', 'B', '#', 'A', '#', ' ', ' '},
		[13]rune{' ', ' ', '#', 'D', '#', 'B', '#', 'A', '#', 'C', '#', ' ', ' '},
		[13]rune{' ', ' ', '#', 'A', '#', 'D', '#', 'C', '#', 'A', '#', ' ', ' '},
		[13]rune{' ', ' ', '#', '#', '#', '#', '#', '#', '#', '#', '#', ' ', ' '},
	}
}

func input() [7][13]rune {
	return [7][13]rune{
		[13]rune{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
		[13]rune{'#', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '#'},
		[13]rune{'#', '#', '#', 'D', '#', 'C', '#', 'A', '#', 'B', '#', '#', '#'},
		[13]rune{' ', ' ', '#', 'D', '#', 'C', '#', 'B', '#', 'A', '#', ' ', ' '},
		[13]rune{' ', ' ', '#', 'D', '#', 'B', '#', 'A', '#', 'C', '#', ' ', ' '},
		[13]rune{' ', ' ', '#', 'B', '#', 'C', '#', 'D', '#', 'A', '#', ' ', ' '},
		[13]rune{' ', ' ', '#', '#', '#', '#', '#', '#', '#', '#', '#', ' ', ' '},
	}
}

type Point struct {
	X, Y int
}

type Pod struct {
	Pos  Point
	Type rune
}

type Move struct {
	Start, End Point
	Final      bool
	Type       rune
}

func PrintGrid(grid [7][13]rune) {
	for x := range grid {
		for y := range grid[x] {
			fmt.Printf("%v", string(grid[x][y]))
		}
		fmt.Println()
	}
}

func (m *Move) DoMove(grid [7][13]rune) ([7][13]rune, int, bool) {
	if grid[m.Start.X][m.Start.Y] != m.Type {
		return grid, 0, false
	}
	cur := Point{m.Start.X, m.Start.Y}
	cost := 0
	for {
		if cur.Y < m.End.Y {
			if grid[cur.X][cur.Y+1] == '.' {
				cost += 1
				cur.Y += 1
				continue
			}
		} else if cur.Y > m.End.Y {
			if grid[cur.X][cur.Y-1] == '.' {
				cost += 1
				cur.Y -= 1
				continue
			}
		}
		// Must be at the right Y
		if cur.X < m.End.X {
			if grid[cur.X+1][cur.Y] == '.' {
				cost += 1
				cur.X += 1
				continue
			}
		} else if cur.X > m.End.X {
			if grid[cur.X-1][cur.Y] == '.' {
				cost += 1
				cur.X -= 1
				continue
			}
		}
		if cur == m.End {
			grid[m.Start.X][m.Start.Y] = '.'
			grid[m.End.X][m.End.Y] = m.Type
			return grid, cost * energy[m.Type], true
		}
		return grid, 0, false
	}
}

type Result struct {
	Score int
	Legal bool
	Moves []Move
}

func CalcMoves(grid [7][13]rune, moved [7][13]bool) []Move {
	var openSpaces = FindOpenSpaces(grid)
	var pods []*Pod
	minSafe := make(map[rune]int)
	for x := range grid {
		for y := range grid[x] {
			switch grid[x][y] {
			case 'A':
				pods = append(pods, &Pod{Pos: Point{x, y}, Type: grid[x][y]})
			case 'B':
				pods = append(pods, &Pod{Pos: Point{x, y}, Type: grid[x][y]})
			case 'C':
				pods = append(pods, &Pod{Pos: Point{x, y}, Type: grid[x][y]})
			case 'D':
				pods = append(pods, &Pod{Pos: Point{x, y}, Type: grid[x][y]})
			}
			if 5-x == minSafe[grid[x][y]] {
				minSafe[grid[x][y]] = 5 - x
			}
		}
	}
	var moves []Move
	for _, pod := range pods {
		for _, space := range openSpaces {
			if space == pod.Pos {
				continue
			}
			if space.X == 1 && (space.Y == 3 || space.Y == 5 || space.Y == 7 || space.Y == 9) {
				continue
			}
			if space.X != 1 && space.Y != home[pod.Type] {
				continue
			}
			if pod.Pos.X == 1 && space.X == 1 {
				continue
			}
			if space.Y == home[pod.Type] && space.X > 5-minSafe[pod.Type] {
				// We're home, don't go anywhere
				continue
			}
			if pod.Pos.X != 1 && moved[pod.Pos.X][pod.Pos.Y] {
				continue
			}
			if (pod.Pos.X == 1) || (space.X != 1) {
				moves = append([]Move{Move{Start: pod.Pos, End: space, Type: pod.Type, Final: (pod.Pos.X == 1) || (space.X != 1)}}, moves...)
			} else {
				moves = append(moves, Move{Start: pod.Pos, End: space, Type: pod.Type, Final: (pod.Pos.X == 1) || (space.X != 1)})
			}
		}
	}
	return moves
}

type MemoKey struct {
	Grid  [7][13]rune
	Moved [7][13]bool
}

var memo map[MemoKey]Result

func FindPaths(grid [7][13]rune, moved [7][13]bool) (int, bool, []Move) {
	if val, ok := memo[MemoKey{grid, moved}]; ok {
		return val.Score, val.Legal, val.Moves
	}
	min := int(^uint(0) >> 1)
	var found bool
	var minChain []Move
	moves := CalcMoves(grid, moved)
	for _, move := range moves {
		newGrid, cost, legal := move.DoMove(grid)
		if !legal {
			continue
		}
		if newGrid == solved {
			found = true
			if cost < min {
				min = cost
				minChain = []Move{move}
			}
			break
		}
		var newMoved [7][13]bool
		for x := range moved {
			for y := range moved[x] {
				newMoved[x][y] = moved[x][y]
			}
		}
		newMoved[move.Start.X][move.Start.Y] = false
		newMoved[move.End.X][move.End.Y] = true
		best, legal, chain := FindPaths(newGrid, newMoved)
		if ((best + cost) < min) && legal {
			found = true
			minChain = make([]Move, 0)
			for _, m := range chain {
				minChain = append(minChain, m)
			}
			minChain = append(minChain, move)
			min = best + cost
		}
	}
	memo[MemoKey{grid, moved}] = Result{Score: min, Legal: found, Moves: minChain}
	return min, found, minChain
}

func FindOpenSpaces(grid [7][13]rune) []Point {
	var openSpaces []Point
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] != '#' && grid[x][y] != ' ' {
				openSpaces = append(openSpaces, Point{x, y})
			}
		}
	}
	return openSpaces
}

func main() {

	memo = make(map[MemoKey]Result)
	//grid := sample()
	grid := input()
	score, _, _ := FindPaths(grid, [7][13]bool{})
	fmt.Println(score)
}
