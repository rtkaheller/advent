package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

// Go doesn't have int Abs in "math" :(
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Pos struct {
	X, Y int
}

type Move struct {
	Dir  byte
	Dist int
}

func PrintGrid(tail, head, max, min Pos, tails map[Pos]bool) {
	fmt.Println(min, max, tail, head)
	for y := max.Y; y >= min.Y; y-- {
		for x := min.X; x <= max.X; x++ {
			p := Pos{x, y}
			if p == head {
				fmt.Printf("H")
			} else if p == tail {
				fmt.Printf("T")
			} else if _, ok := tails[p]; ok {
				fmt.Printf("#")
			} else if x == 0 && y == 0 {
				fmt.Printf("s")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}

func MoveRope(knots int, moves []Move) int {
	max := Pos{6, 6}
	min := Pos{0, 0}
	k := make([]Pos, knots)
	tails := make(map[Pos]bool)

	for _, move := range moves {
		var delta Pos
		switch move.Dir {
		case 'U':
			delta = Pos{0, 1}
		case 'D':
			delta = Pos{0, -1}
		case 'R':
			delta = Pos{1, 0}
		case 'L':
			delta = Pos{-1, 0}
		}
		for i := 0; i < move.Dist; i++ {
			k[0].X += delta.X
			k[0].Y += delta.Y
			for j, t := range k {
				if j == 0 {
					continue
				}
				distY := Abs(k[j-1].Y - t.Y)
				distX := Abs(k[j-1].X - t.X)
				if distX > 1 || distY > 1 {
					if distY != 0 {
						k[j].Y += (k[j-1].Y - t.Y) / distY
					}
					if distX != 0 {
						k[j].X += (k[j-1].X - t.X) / distX
					}
				}
				if j == len(k)-1 {
					tails[k[j]] = true
				}
			}
		}
		if max.X < k[0].X {
			max.X = k[0].X
		}
		if max.Y < k[0].Y {
			max.Y = k[0].Y
		}
		if min.X > k[0].X {
			min.X = k[0].X
		}
		if min.Y > k[0].Y {
			min.Y = k[0].Y
		}
	}
	return len(tails)
}

func main() {
	contents, err := ioutil.ReadFile("fake.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	var moves []Move
	for _, line := range bytes.Split(contents, []byte("\n")) {
		if len(line) == 0 {
			continue
		}
		pair := bytes.Split(line, []byte(" "))
		dist, _ := strconv.Atoi(string(pair[1]))
		moves = append(moves, Move{Dir: pair[0][0], Dist: dist})
	}
	fmt.Println(MoveRope(2, moves))
	fmt.Println(MoveRope(10, moves))
}
