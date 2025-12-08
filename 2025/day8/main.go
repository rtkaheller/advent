package main

import (
	"bytes"
	"flag"
	"fmt"
	"maps"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
)

type Pair struct {
	J1, J2 *Junc
}

type Pos struct {
	X, Y, Z int
}

type Junc struct {
	P       Pos
	Circuit int
}

var Circs = map[int][]*Junc{}
var cCount = 0

func Dist(p1, p2 Pos) float64 {
	r := math.Sqrt(
		math.Pow(math.Abs(float64(p1.X-p2.X)), 2) +
			math.Pow(math.Abs(float64(p1.Y-p2.Y)), 2) +
			math.Pow(math.Abs(float64(p1.Z-p2.Z)), 2),
	)
	return r
}

func Join(j1, j2 *Junc) {
	if j1.Circuit == j2.Circuit && j2.Circuit != 0 {
		return
	}
	if j1.Circuit == 0 && j2.Circuit == 0 {
		cCount += 1
		j1.Circuit = cCount
		j2.Circuit = cCount
		Circs[cCount] = []*Junc{j1, j2}
	} else if j1.Circuit == 0 {
		// j2 is the one connected
		j1.Circuit = j2.Circuit
		Circs[j1.Circuit] = append(Circs[j1.Circuit], j1)
	} else if j2.Circuit == 0 {
		// j1 is the one connected
		j2.Circuit = j1.Circuit
		Circs[j2.Circuit] = append(Circs[j2.Circuit], j2)
	} else {
		// Both are connected, merge to the lower value
		if j1.Circuit < j2.Circuit {
			old := j2.Circuit
			for _, j := range Circs[j2.Circuit] {
				j.Circuit = j1.Circuit
				Circs[j1.Circuit] = append(Circs[j1.Circuit], j)
			}
			delete(Circs, old)
		} else {
			old := j1.Circuit
			for _, j := range Circs[j1.Circuit] {
				j.Circuit = j2.Circuit
				Circs[j2.Circuit] = append(Circs[j2.Circuit], j)
			}
			delete(Circs, old)
		}
	}
}

func main() {
	flag.Parse()
	args := flag.Args()
	file := args[0]
	count, _ := strconv.Atoi(args[1])
	contents, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	p1 := 0
	p2 := 0

	junctions := []*Junc{}
	lines := bytes.Split(contents, []byte("\n"))
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		coords := bytes.Split(line, []byte(","))
		x, _ := strconv.Atoi(string(coords[0]))
		y, _ := strconv.Atoi(string(coords[1]))
		z, _ := strconv.Atoi(string(coords[2]))
		junctions = append(junctions, &Junc{P: Pos{x, y, z}, Circuit: 0})
	}

	dists := []Pair{}
	for i := range junctions {
		for t := i + 1; t < len(junctions); t++ {
			if i == t {
				continue
			}
			dists = append(dists, Pair{J1: junctions[i], J2: junctions[t]})
		}
	}
	sort.Slice(dists, func(i, j int) bool {
		return Dist(dists[i].J1.P, dists[i].J2.P) < Dist(dists[j].J1.P, dists[j].J2.P)
	})

	for i := 0; i < len(dists); i++ {
		Join(dists[i].J1, dists[i].J2)
		if i == count {
			circ := map[int]int{}
			for _, j := range junctions {
				circ[j.Circuit] += 1
			}
			delete(circ, 0)
			v := slices.Collect(maps.Values(circ))
			sort.Ints(v)

			p1 = v[len(v)-1] * v[len(v)-2] * v[len(v)-3]
		}
		s := 0
		for _, l := range Circs {
			s += len(l)
		}
		if s == len(junctions) {
			p2 = dists[i].J1.P.X * dists[i].J2.P.X
			break
		}
	}

	fmt.Println(p1)
	fmt.Println(p2)
}
