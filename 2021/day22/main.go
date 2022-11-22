package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
)

const (
	minDimension = -50
	maxDimension = 50
)

type Point struct {
	X, Y, Z int
}

type Cube struct {
	Vert1, Vert2 Point
	Direction    bool
	Missing      []*Cube
}

func (c *Cube) Score() int {
	score := (c.Vert2.X - c.Vert1.X + 1) * (c.Vert2.Y - c.Vert1.Y + 1) * (c.Vert2.Z - c.Vert1.Z + 1)
	if !c.Direction {
		score *= -1
	}
	for _, cube := range c.Missing {
		score += cube.Score()
	}
	return score
}

func (c *Cube) Intersect(c2 Cube) {
	if (c.Vert2.X >= c2.Vert1.X && c.Vert1.X <= c2.Vert2.X) &&
		(c.Vert2.Y >= c2.Vert1.Y && c.Vert1.Y <= c2.Vert2.Y) &&
		(c.Vert2.Z >= c2.Vert1.Z && c.Vert1.Z <= c2.Vert2.Z) {
		x1 := int(math.Max(float64(c.Vert1.X), float64(c2.Vert1.X)))
		x2 := int(math.Min(float64(c.Vert2.X), float64(c2.Vert2.X)))
		y1 := int(math.Max(float64(c.Vert1.Y), float64(c2.Vert1.Y)))
		y2 := int(math.Min(float64(c.Vert2.Y), float64(c2.Vert2.Y)))
		z1 := int(math.Max(float64(c.Vert1.Z), float64(c2.Vert1.Z)))
		z2 := int(math.Min(float64(c.Vert2.Z), float64(c2.Vert2.Z)))
		newCube := Cube{Vert1: Point{x1, y1, z1}, Vert2: Point{x2, y2, z2}, Direction: c2.Direction}
		for _, cube := range c.Missing {
			cube.Intersect(newCube)
		}
		if newCube.Direction == c.Direction {
			newCube.Direction = !c.Direction
		}
		c.Missing = append(c.Missing, &newCube)
	}
}

type Instruction struct {
	MinX, MinY, MinZ int
	MaxX, MaxY, MaxZ int
	Direction        bool
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	//contents, err := ioutil.ReadFile("fake.txt")
	//contents, err := ioutil.ReadFile("small.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	var grid [101][101][101]bool
	var instructions []Instruction
	for _, line := range bytes.Split(contents, []byte("\n")) {
		if len(line) == 0 {
			break
		}
		step := bytes.Split(line, []byte(" "))
		coords := bytes.Split(step[1], []byte(","))
		newInstruction := Instruction{}
		newInstruction.Direction = string(step[0]) == "on"
		x := bytes.Split(coords[0][2:], []byte(".."))
		newInstruction.MinX, _ = strconv.Atoi(string(x[0]))
		newInstruction.MaxX, _ = strconv.Atoi(string(x[1]))
		y := bytes.Split(coords[1][2:], []byte(".."))
		newInstruction.MinY, _ = strconv.Atoi(string(y[0]))
		newInstruction.MaxY, _ = strconv.Atoi(string(y[1]))
		z := bytes.Split(coords[2][2:], []byte(".."))
		newInstruction.MinZ, _ = strconv.Atoi(string(z[0]))
		newInstruction.MaxZ, _ = strconv.Atoi(string(z[1]))
		instructions = append(instructions, newInstruction)
	}
	for _, instruct := range instructions {
		for x := instruct.MinX; x <= instruct.MaxX; x++ {
			if x < minDimension || x > maxDimension {
				continue
			}
			for y := instruct.MinY; y <= instruct.MaxY; y++ {
				if y < minDimension || y > maxDimension {
					continue
				}
				for z := instruct.MinZ; z <= instruct.MaxZ; z++ {
					if z < minDimension || z > maxDimension {
						continue
					}
					grid[x+50][y+50][z+50] = instruct.Direction
				}
			}
		}
	}
	c := 0
	for x := range grid {
		for y := range grid[x] {
			for z := range grid[y] {
				if grid[x][y][z] {
					c += 1
				}
			}
		}
	}
	fmt.Println(c)
	var cubes []*Cube
	for _, instruct := range instructions {
		cubes = append(cubes, &Cube{
			Vert1: Point{
				X: instruct.MinX,
				Y: instruct.MinY,
				Z: instruct.MinZ,
			},
			Vert2: Point{
				X: instruct.MaxX,
				Y: instruct.MaxY,
				Z: instruct.MaxZ,
			},
			Direction: instruct.Direction,
		})
		for _, cube := range cubes[:len(cubes)-1] {
			cube.Intersect(*cubes[len(cubes)-1])
		}
	}
	sum := 0
	for _, cube := range cubes {
		if cube.Direction {
			sum += cube.Score()
		}
	}
	fmt.Println(sum)
}
