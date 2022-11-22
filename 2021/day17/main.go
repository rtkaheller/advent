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

func SimY(c1, c2 Coord, xVel, yVel int) (int, bool) {
	x, y := 0, 0
	maxY := 0
	for i := 0; i <= xVel; i++ {
		if x >= c1.X && x <= c2.X && y <= c1.Y && y >= c2.Y {
			return maxY, true
		}
		if x > c2.X {
			return 0, false
		}
		if y > maxY {
			maxY = y
		}
		x += xVel - i
		y += yVel
		yVel -= 1
	}
	if x < c1.X {
		return 0, false
	}
	for {
		if y > maxY {
			maxY = y
		}
		if x >= c1.X && x <= c2.X && y <= c1.Y && y >= c2.Y {
			return maxY, true
		} else if y < c2.Y {
			return 0, false
		}
		y += yVel
		yVel -= 1
	}
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	//contents, err := ioutil.ReadFile("fake.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	chunks := bytes.Split(contents, []byte(" "))
	x := bytes.Split(chunks[2][2:len(chunks[2])-1], []byte(".."))
	y := bytes.Split(chunks[3][2:len(chunks[3])-1], []byte(".."))
	var c1, c2 Coord
	c1.X, _ = strconv.Atoi(string(x[0]))
	c2.X, _ = strconv.Atoi(string(x[1]))
	c1.Y, _ = strconv.Atoi(string(y[1]))
	c2.Y, _ = strconv.Atoi(string(y[0]))

	xVelMin := 0
	for i := 0; ; i++ {
		if (i*(i+1))/2 >= c1.X {
			xVelMin = i
			break
		}
	}
	xVelMin = 0
	xVelMax := c2.X + 1
	bestY := 0
	count := 0
	for i := xVelMax; i >= xVelMin; i-- {
		for t := (-1 * c2.Y) + 1; t >= c2.Y-1; t-- {
			y, safe := SimY(c1, c2, i, -1*t)
			if safe {
				count += 1
				if y > bestY {
					bestY = y
				}
			}
		}
	}
	fmt.Println(bestY)
	fmt.Println(count)
}
