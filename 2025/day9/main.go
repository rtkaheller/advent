package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

var minX, minY = 1000000000, 10000000000

func Abs(v int) int {
	if v < 0 {
		return -1 * v
	}
	return v
}

func Line(p1, p2 Pos, grid *map[Pos]byte) {
	if p1.X != p2.X {
		for x := min(p1.X, p2.X) + 1; x < max(p1.X, p2.X); x++ {
			(*grid)[Pos{x, p1.Y}] = 'X'
		}
	}
	if p1.Y != p2.Y {
		for y := min(p1.Y, p2.Y) + 1; y < max(p1.Y, p2.Y); y++ {
			(*grid)[Pos{p1.X, y}] = 'X'
		}
	}
}

type Pos struct {
	X, Y int
}

func Dir(p1, p2 Pos) string {
	if p1.X == p2.X {
		// Vertical line
		if p1.Y < p2.Y {
			return "Down"
		} else {
			return "Up"
		}
	} else {
		// Horizontal
		if p1.X < p2.X {
			return "Right"
		} else {
			return "Left"
		}
	}
}

func Corner(o, l, n Pos, grid *map[Pos]byte) {
	old := Dir(o, l)
	next := Dir(l, n)
	if old == next {
		(*grid)[l] = 'X'
	} else if (old == "Right" && next == "Up") || (old == "Down" && next == "Left") {
		(*grid)[l] = 'J'
	} else if (old == "Down" && next == "Right") || (old == "Left" && next == "Up") {
		(*grid)[l] = 'L'
	} else if (old == "Up" && next == "Right") || (old == "Left" && next == "Down") {
		(*grid)[l] = 'r'
	} else if (old == "Right" && next == "Down") || (old == "Up" && next == "Left") {
		(*grid)[l] = '7'
	}
}
func ScanLineY(y, eY, x int, grid *map[Pos]byte) bool {
	var start Pos
	start = Pos{x, minY - 1}

	in := false
	line := false
	var last byte
	win := false
	for c := start; c.Y <= eY; c.Y += 1 {
		v, ok := (*grid)[c]
		if !ok {
			if !in && c.Y > y {
				return false
			}
			continue
		}
		switch v {
		case 'J':
			line = false
			if last == '7' {
				if win {
					// Still in
				} else {
					in = false
				}
			} else if last == 'r' {
				if !win {
					// We weren't in, but we are now
				} else {
					// Started in, we're out now
					in = false
				}
			}
		case 'L':
			line = false
			if last == '7' {
				if !win {
					// We weren't in, but we are now
				} else {
					// Started in, we're out now
					in = false
				}
			} else if last == 'r' {
				if win {
					// Still in
				} else {
					// Started out, we're out again
					in = false
				}
			}
		case 'r':
			line = true
			last = 'r'
			win = in
			if !in {
				in = true
			}
		case '7':
			line = true
			last = '7'
			win = in
			if !in {
				in = true
			}
		case 'X':
			if line {
				continue
			}
			if in {
				in = false
			} else if !line {
				in = true
			}
		}
	}
	return true
}

func ScanLineX(x, eX, y int, grid *map[Pos]byte) bool {
	var start Pos
	start = Pos{minX - 1, y}

	in := false
	line := false
	var last byte
	win := false
	for c := start; c.X <= eX; c.X += 1 {
		v, ok := (*grid)[c]
		if !ok {
			if !in && c.X > x {
				return false
			}
			continue
		}
		switch v {
		case 'r':
			line = false
			if last == 'L' {
				if win {
					// Still in
				} else {
					in = false
				}
			} else if last == 'J' {
				if !win {
					// We weren't in, but we are now
				} else {
					// Started in, we're out now
					in = false
				}
			}
		case '7':
			line = false
			if last == 'L' {
				if !win {
					// We weren't in, but we are now
				} else {
					// Started in, we're out now
					in = false
				}
			} else if last == 'J' {
				if win {
					// Still in
				} else {
					// Started out, we're out again
					in = false
				}
			}
		case 'J', 'L':
			line = true
			last = v
			win = in
			if !in {
				in = true
			}
		case 'X':
			if line {
				continue
			}
			if in {
				in = false
			} else if !line {
				in = true
			}
		}
	}
	return true
}

func Contains(p1, p2 Pos, grid *map[Pos]byte) bool {
	iX := min(p1.X, p2.X)
	iY := min(p1.Y, p2.Y)
	aX := max(p1.X, p2.X)
	aY := max(p1.Y, p2.Y)
	return ScanLineX(iX, aX, iY, grid) &&
		ScanLineX(iX, aX, aY, grid) &&
		ScanLineY(iY, aY, iX, grid) &&
		ScanLineY(iY, aY, aX, grid)
}

func main() {
	contents, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	p1 := 0
	p2 := 0

	corn := []Pos{}
	lines := bytes.Split(contents, []byte("\n"))
	grid := map[Pos]byte{}
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		s := bytes.Split(line, []byte(","))
		x, _ := strconv.Atoi(string(s[0]))
		y, _ := strconv.Atoi(string(s[1]))
		if x < minX {
			minX = x
		}
		if y < minY {
			minY = y
		}
		corn = append(corn, Pos{x, y})
		if len(corn) > 2 {
			Corner(corn[len(corn)-3], corn[len(corn)-2], corn[len(corn)-1], &grid)
		}
	}
	Corner(corn[len(corn)-1], corn[0], corn[1], &grid)
	Corner(corn[len(corn)-2], corn[len(corn)-1], corn[0], &grid)

	Line(corn[0], corn[len(corn)-1], &grid)
	for i := 1; i < len(corn); i++ {
		Line(corn[i], corn[i-1], &grid)
	}

	for i := range corn {
		c1 := corn[i]
		for t := i; t < len(corn); t++ {
			c2 := corn[t]
			a := (Abs(c1.X-c2.X) + 1) * (Abs(c1.Y-c2.Y) + 1)
			if a > p1 {
				p1 = a
			}
			if a > p2 && Contains(c1, c2, &grid) {
				p2 = a
			}
		}
	}
	fmt.Println(p1)
	fmt.Println(p2)
}
