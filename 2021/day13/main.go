package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

type Coord struct {
	X, Y int
}

type Fold struct {
	Up   bool
	Axis int
}

func fold(dots []*Coord, f *Fold) {
	for _, dot := range dots {
		if f.Up {
			if dot.Y > f.Axis {
				dot.Y = f.Axis - (dot.Y - f.Axis)
			}
		} else {
			if dot.X > f.Axis {
				dot.X = f.Axis - (dot.X - f.Axis)
			}
		}
	}
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	//contents, err := ioutil.ReadFile("fake.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	var dots []*Coord
	var start int
	for resume, line := range bytes.Split(contents, []byte("\n")) {
		if len(line) == 0 {
			start = resume
			break
		}
		coords := bytes.Split(line, []byte(","))

		coord := new(Coord)

		coord.X, _ = strconv.Atoi(string(coords[0]))
		coord.Y, _ = strconv.Atoi(string(coords[1]))
		dots = append(dots, coord)
	}
	var folds []*Fold
	for i, line := range bytes.Split(contents, []byte("\n")) {
		if i <= start {
			continue
		} else if len(line) == 0 {
			continue
		}
		words := bytes.Split(line, []byte(" "))
		f := bytes.Split(words[2], []byte("="))
		fold := new(Fold)
		fold.Up = (string(f[0]) == "y")
		fold.Axis, _ = strconv.Atoi(string(f[1]))
		folds = append(folds, fold)
	}
	for _, f := range folds {
		fold(dots, f)
	}
	paper := make(map[string]bool)
	var maxX, maxY int
	for _, dot := range dots {
		if dot.X > maxX {
			maxX = dot.X
		}
		if dot.Y > maxY {
			maxY = dot.Y
		}
		paper[fmt.Sprintf("%v,%v", dot.X, dot.Y)] = true
	}
	for y := 0; y < maxY+1; y++ {
		for x := 0; x < maxX; x++ {
			if paper[fmt.Sprintf("%v,%v", x, y)] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}
