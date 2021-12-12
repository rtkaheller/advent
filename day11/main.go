package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

const (
	steps = 100
)

type Point struct {
	Val     int
	Flashed int
}

type Grid struct {
	Points [][]*Point
}

func (g *Grid) Print() {
	for y := range g.Points {
		for x := range g.Points[y] {
			fmt.Printf("%v", g.Points[y][x].Flashed)
		}
		fmt.Println()
	}
	for y := range g.Points {
		for x := range g.Points[y] {
			if g.Points[y][x].Val > 9 {
				fmt.Printf("F")
			} else {
				fmt.Printf("%v", g.Points[y][x].Val)
			}
		}
		fmt.Println()
	}
}

func (g *Grid) FlashPoint(x, y, round int) {
	for i := -1; i < 2; i++ {
		for t := -1; t < 2; t++ {
			if y+i < 0 || y+i >= len(g.Points) || x+t < 0 || x+t >= len(g.Points[y+i]) {
				continue
			}
			g.Points[y+i][x+t].Val += 1
			if g.Points[y+i][x+t].Val > 9 && g.Points[y+i][x+t].Flashed < round {
				g.Points[y+i][x+t].Flashed = round
				g.FlashPoint(x+t, y+i, round)
			}
		}
	}
}

func (g *Grid) Flash(round int) int {
	for y := range g.Points {
		for x := range g.Points[y] {
			g.Points[y][x].Val += 1
		}
	}
	for y := range g.Points {
		for x := range g.Points[y] {
			if g.Points[y][x].Val > 9 && g.Points[y][x].Flashed < round {
				g.Points[y][x].Flashed = round
				g.FlashPoint(x, y, round)
			}
		}
	}
	count := 0
	for y := range g.Points {
		for x := range g.Points[y] {
			if g.Points[y][x].Val > 9 && g.Points[y][x].Flashed == round {
				g.Points[y][x].Val = 0
				count += 1
			}
		}
	}
	return count
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	//contents, err := ioutil.ReadFile("fake.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	var g Grid
	for y, line := range bytes.Split(contents, []byte("\n")) {
		if len(line) == 0 {
			continue
		}
		g.Points = append(g.Points, []*Point{})
		for _, r := range string(line) {
			val, _ := strconv.Atoi(string(r))
			g.Points[y] = append(g.Points[y], &Point{Val: val})
		}
	}
	sum := 0
	for i := 0; ; i++ {
		flashing := g.Flash(i + 1)
		if i < steps {
			sum += flashing
		}
		if flashing == len(g.Points)*len(g.Points[0]) {
			fmt.Println(sum)
			fmt.Println(i + 1)
			break
		}
	}
}
