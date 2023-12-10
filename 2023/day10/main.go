package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

type Pos struct {
	X, Y int
}

func Draw(g [][]byte) {
	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[y]); x++ {
			fmt.Printf("%s", string(g[y][x]))
		}
		fmt.Printf("\n")
	}
}

func Answer(g [][]byte) int {
	count := 0
	for y := 1; y < len(g)-1; y += 2 {
		for x := 1; x < len(g[0])-1; x += 2 {
			if g[y][x] == '.' {
				count += 1
			}
		}
	}
	return count
}

func Extend(g [][]byte) [][]byte {
	n := [][]byte{}
	for y := 0; y < len(g)*2+1; y++ {
		n = append(n, []byte{})
		for x := 0; x < len(g[0])*2+1; x++ {
			n[y] = append(n[y], '.')
		}
	}

	for y := 1; y < len(g)*2; y += 2 {
		for x := 1; x < len(g[0])*2; x += 2 {
			n[y][x] = g[y/2][x/2]
			switch g[y/2][x/2] {
			case '-':
				n[y][x+1] = '-'
			case '|':
				n[y+1][x] = '|'
			case 'L':
				n[y][x+1] = '-'
			case 'F':
				n[y][x+1] = '-'
				n[y+1][x] = '|'
			case 'J':
			case '7':
				n[y+1][x] = '|'
			}
		}
	}
	return n
}

func Flood(g *[][]byte, p Pos) {
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if p.X+dx >= 0 && p.X+dx < len((*g)[0]) && p.Y+dy >= 0 && p.Y+dy < len(*g) {
				if (*g)[p.Y+dy][p.X+dx] == '.' {
					(*g)[p.Y+dy][p.X+dx] = 'O'
					Flood(g, Pos{p.X + dx, p.Y + dy})
				}
			}
		}
	}
}

func Fill(start Pos, path []Pos, g [][]byte) int {
	in := map[Pos]bool{}
	d1 := Pos{path[0].X - start.X, path[0].Y - start.Y}
	d2 := Pos{path[len(path)-1].X - start.X, path[len(path)-1].Y - start.Y}
	if d1.X != 0 && d2.X != 0 {
		g[start.X][start.Y] = '-'
	} else if d1.Y != 0 && d2.Y != 0 {
		g[start.X][start.Y] = '|'
	} else if d1.X > 0 && d2.Y > 0 {
		g[start.X][start.Y] = 'F'
	} else if d1.X < 0 && d2.Y > 0 {
		g[start.X][start.Y] = '7'
	} else if d1.X < 0 && d2.Y < 0 {
		g[start.X][start.Y] = 'J'
	} else if d1.X > 0 && d2.Y < 0 {
		g[start.X][start.Y] = 'L'
	}
	return len(in)
}

func Replace(path []Pos, g *[][]byte, full bool) {
	for y := 0; y < len(*g); y++ {
		for x := 0; x < len((*g)[y]); x++ {
			if (*g)[y][x] == 'S' {
				d1 := Pos{path[0].X - x, path[0].Y - y}
				d2 := Pos{path[len(path)-1].X - x, path[len(path)-1].Y - y}
				if d1.X != 0 && d2.X != 0 {
					(*g)[y][x] = '-'
				} else if d1.Y != 0 && d2.Y != 0 {
					(*g)[y][x] = '|'
				} else if d1.X > 0 && d2.Y > 0 {
					(*g)[y][x] = 'F'
				} else if d1.X < 0 && d2.Y > 0 {
					(*g)[y][x] = '7'
				} else if d1.X < 0 && d2.Y < 0 {
					(*g)[y][x] = 'J'
				} else if d1.X > 0 && d2.Y < 0 {
					(*g)[y][x] = 'L'
				}
				continue
			}
			pipe := false
			for _, p := range path {
				if x == p.X && y == p.Y {
					pipe = true
					break
				}
			}
			if !pipe {
				(*g)[y][x] = '.'
			} else if full {
				(*g)[y][x] = '#'
			}
		}
	}
}

func FindLoop(s Pos, g [][]byte) []Pos {
	n := s
	if s.X-1 >= 0 && strings.ContainsRune("-FL", rune(g[s.Y][s.X-1])) {
		n.X -= 1
	} else if s.X+1 < len(g[s.Y]) && strings.ContainsRune("-7J", rune(g[s.Y][s.X+1])) {
		n.X += 1
	} else if s.Y-1 >= 0 && strings.ContainsRune("-F7", rune(g[s.Y-1][s.X])) {
		n.Y -= 1
	} else if s.Y+1 < len(g) && strings.ContainsRune("-LJ", rune(g[s.Y+1][s.X])) {
		n.X += 1
	}
	var ret []Pos
	p := s
	for n != s {
		ret = append(ret, n)
		l := n
		switch g[n.Y][n.X] {
		case '|':
			n.Y += (n.Y - p.Y)
		case '-':
			n.X += (n.X - p.X)
		case 'L':
			if p.X > n.X {
				n.Y -= 1
			} else {
				n.X += 1
			}
		case 'J':
			if p.X < n.X {
				n.Y -= 1
			} else {
				n.X -= 1
			}
		case '7':
			if p.X < n.X {
				n.Y += 1
			} else {
				n.X -= 1
			}
		case 'F':
			if p.X > n.X {
				n.Y += 1
			} else {
				n.X += 1
			}
		default:
			fmt.Println("wtf")
		}
		p = l
	}
	return ret
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	var grid [][]byte
	var s Pos
	lines := bytes.Split(contents, []byte("\n"))
	for y := 0; y < len(lines); y++ {
		if len(lines[y]) == 0 {
			continue
		}
		grid = append(grid, []byte{})
		for x := 0; x < len(lines[y]); x++ {
			grid[y] = append(grid[y], lines[y][x])
			if lines[y][x] == 'S' {
				s.X = x
				s.Y = y
			}
		}
	}
	path := FindLoop(s, grid)
	fmt.Println(len(path)/2 + 1)
	Replace(path, &grid, false)
	newGrid := Extend(grid)
	Flood(&newGrid, Pos{0, 0})
	fmt.Println(Answer(newGrid))
}
