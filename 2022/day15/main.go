package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

type Pos struct {
	X, Y int
}

type Report struct {
	Sensor, Beacon Pos
	Dist           int
}

// Go doesn't have int Abs in "math" :(
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (r *Report) Candidates(space Pos) []Pos {
	var points []Pos
	cur := Pos{r.Sensor.X - r.Dist - 1, r.Sensor.Y}
	if cur.X < 0 {
		r.Sensor.Y += 0 - cur.X
		r.Sensor.X = 0
	}
	for cur.X <= r.Sensor.X {
		points = append(points, cur)
		cur.X += 1
		cur.Y += 1
		if cur.X > space.X || cur.Y > space.Y {
			break
		}
	}

	cur = Pos{r.Sensor.X + r.Dist + 1, r.Sensor.Y}
	if cur.X > space.X {
		r.Sensor.Y += cur.X - space.X
		r.Sensor.X = space.X
	}
	for cur.X != r.Sensor.X {
		points = append(points, cur)
		cur.X -= 1
		cur.Y += 1
		if cur.X < 0 || cur.Y > space.Y {
			break
		}
	}

	cur = Pos{r.Sensor.X, r.Sensor.Y - r.Dist - 1}
	if cur.Y < 0 {
		r.Sensor.X += 0 - cur.Y
		r.Sensor.Y = 0
	}
	for cur.Y <= r.Sensor.Y {
		points = append(points, cur)
		cur.X += 1
		cur.Y += 1
		if cur.X > space.X || cur.Y > space.Y {
			break
		}
	}

	cur = Pos{r.Sensor.X, r.Sensor.Y + r.Dist + 1}
	if cur.Y > space.Y {
		r.Sensor.X += cur.Y - space.Y
		r.Sensor.Y = 0
	}
	for cur.Y <= r.Sensor.Y {
		points = append(points, cur)
		cur.X += 1
		cur.Y -= 1
		if cur.X > space.X || cur.Y < 0 {
			break
		}
	}
	return points
}

func SearchGrid(grid map[Pos]bool, space Pos) int {
	for x := 0; x < space.X; x++ {
		for y := 0; y < space.Y; y++ {
			if !grid[Pos{x, y}] {
				fmt.Println(x, y)
				return x*4000000 + y
			}
		}
		fmt.Println(x)
	}
	return -1
}

func part1(file string, row int, space Pos) (int, int) {
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return -1, -1
	}

	var reports []Report
	minX, maxX, maxDist := 0, 0, 0
	for _, line := range bytes.Split(contents, []byte("\n")) {
		if len(line) == 0 {
			continue
		}
		r := Report{}

		fmt.Sscanf(string(line), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &r.Sensor.X, &r.Sensor.Y, &r.Beacon.X, &r.Beacon.Y)
		r.Dist = Abs(r.Sensor.X-r.Beacon.X) + Abs(r.Sensor.Y-r.Beacon.Y)
		if r.Sensor.X < minX {
			minX = r.Sensor.X
		}
		if r.Sensor.X > maxX {
			maxX = r.Sensor.X
		}
		if r.Dist > maxDist {
			maxDist = r.Dist
		}
		reports = append(reports, r)
	}
	c := 0
	for x := minX - maxDist; x < maxX+maxDist; x++ {
		for _, r := range reports {
			if r.Beacon.X == x && r.Beacon.Y == row {
				break
			}
			if r.Dist >= Abs(r.Sensor.X-x)+Abs(r.Sensor.Y-row) {
				c += 1
				break
			}
		}
	}
	points := make(map[Pos]bool)
	for _, r := range reports {
		for _, p := range r.Candidates(space) {
			points[p] = true
		}
	}

	for p, _ := range points {
		safe := true
		if p.X < 0 || p.X > space.X || p.Y < 0 || p.Y > space.Y {
			continue
		}
		for _, r := range reports {
			if r.Beacon.X == p.Y && r.Beacon.Y == p.Y {
				safe = false
				break
			}
			if r.Dist >= Abs(r.Sensor.X-p.X)+Abs(r.Sensor.Y-p.Y) {
				safe = false
				break
			}
		}
		if safe {
			return c, p.X*4000000 + p.Y
		}
	}
	return c, -1
}

func main() {
	fmt.Println(part1("fake.txt", 10, Pos{20, 20}))
	fmt.Println(part1("input.txt", 2000000, Pos{4000000, 4000000}))
}
