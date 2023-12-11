package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

type Pos struct {
	X, Y int
}

func Expand(pos []Pos, m Pos, f int) []Pos {
	out := []Pos{}
	for _, p := range pos {
		out = append(out, p)
	}
	xs := map[int]bool{}
	ys := map[int]bool{}
	for _, p := range pos {
		xs[p.X] = true
		ys[p.Y] = true
	}
	dx := 0
	for x := 0; x < m.X; x++ {
		if _, ok := xs[x]; !ok {
			dx += f
		} else {
			for i, _ := range pos {
				if pos[i].X == x {
					out[i].X += dx
				}
			}
		}
	}
	dy := 0
	for y := 0; y < m.Y; y++ {
		if _, ok := ys[y]; !ok {
			dy += f
		} else {
			for i, _ := range pos {
				if pos[i].Y == y {
					out[i].Y += dy
				}
			}
		}
	}
	return out
}

func Abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func Paths(g []Pos) []int {
	dists := []int{}
	for i := 0; i < len(g); i++ {
		for t := i + 1; t < len(g); t++ {
			dists = append(dists, Abs(g[i].X-g[t].X), Abs(g[i].Y-g[t].Y))
		}
	}
	return dists
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	sum := 0
	lines := bytes.Split(contents, []byte("\n"))
	gals := []Pos{}
	for y := 0; y < len(lines); y++ {
		if len(lines[y]) == 0 {
			continue
		}
		for x := 0; x < len(lines[0]); x++ {
			if lines[y][x] == '#' {
				gals = append(gals, Pos{x, y})
			}
		}
	}
	gals2 := Expand(gals, Pos{len(lines[0]), len(lines)}, 1)
	for _, d := range Paths(gals2) {
		sum += d
	}
	fmt.Println(sum)
	sum = 0
	gals = Expand(gals, Pos{len(lines[0]), len(lines)}, 999999)
	for _, d := range Paths(gals) {
		sum += d
	}
	fmt.Println(sum)
}
