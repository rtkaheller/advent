package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
)

const (
	scanRange       = 1000
	requiredOverlap = 12
)

var Rots = []Axis{
	Axis{Dir: 'X', Sign: -1},
	Axis{Dir: 'X', Sign: 1},
	Axis{Dir: 'Y', Sign: -1},
	Axis{Dir: 'Y', Sign: 1},
	Axis{Dir: 'Z', Sign: -1},
	Axis{Dir: 'Z', Sign: 1},
}

type Axis struct {
	Dir  rune
	Sign int
}

type Point struct {
	X, Y, Z int
}

func (p Point) RRotate(face, up Axis) Point {
	var newPoint Point
	if face.Dir == 'X' && up.Dir == 'Y' {
		newPoint.X = p.X * face.Sign
		newPoint.Y = p.Y * up.Sign
		newPoint.Z = p.Z * up.Sign * face.Sign
	} else if face.Dir == 'X' && up.Dir == 'Z' {
		newPoint.X = p.X * face.Sign
		newPoint.Y = p.Z * up.Sign * face.Sign * -1
		newPoint.Z = p.Y * up.Sign
	} else if face.Dir == 'Y' && up.Dir == 'X' {
		newPoint.X = p.Y * up.Sign
		newPoint.Y = p.X * face.Sign
		newPoint.Z = p.Z * up.Sign * face.Sign * -1
	} else if face.Dir == 'Y' && up.Dir == 'Z' {
		newPoint.X = p.Z * face.Sign * up.Sign
		newPoint.Y = p.X * face.Sign
		newPoint.Z = p.Y * up.Sign
	} else if face.Dir == 'Z' && up.Dir == 'X' {
		newPoint.X = p.Y * up.Sign
		newPoint.Y = p.Z * face.Sign * up.Sign
		newPoint.Z = p.X * face.Sign
	} else if face.Dir == 'Z' && up.Dir == 'Y' {
		newPoint.X = p.Z * face.Sign * up.Sign * -1
		newPoint.Y = p.Y * up.Sign
		newPoint.Z = p.X * face.Sign
	} else {
		fmt.Println("Wtf", face, up)
	}
	return newPoint
}

func (p Point) Rotate(face, up Axis) Point {
	var newPoint Point
	if face.Dir == 'X' && up.Dir == 'Y' {
		newPoint.X = p.X * face.Sign
		newPoint.Y = p.Y * up.Sign
		newPoint.Z = p.Z * up.Sign * face.Sign
	} else if face.Dir == 'X' && up.Dir == 'Z' {
		newPoint.X = p.X * face.Sign
		newPoint.Y = p.Z * up.Sign
		newPoint.Z = p.Y * up.Sign * face.Sign * -1
	} else if face.Dir == 'Y' && up.Dir == 'X' {
		newPoint.X = p.Y * face.Sign
		newPoint.Y = p.X * up.Sign
		newPoint.Z = p.Z * up.Sign * face.Sign * -1
	} else if face.Dir == 'Y' && up.Dir == 'Z' {
		newPoint.X = p.Y * face.Sign
		newPoint.Y = p.Z * up.Sign
		newPoint.Z = p.X * up.Sign * face.Sign
	} else if face.Dir == 'Z' && up.Dir == 'X' {
		newPoint.X = p.Z * face.Sign
		newPoint.Y = p.X * up.Sign
		newPoint.Z = p.Y * up.Sign * face.Sign
	} else if face.Dir == 'Z' && up.Dir == 'Y' {
		newPoint.X = p.Z * face.Sign
		newPoint.Y = p.Y * up.Sign
		newPoint.Z = p.X * up.Sign * face.Sign * -1
	} else {
		fmt.Println("Wtf", face, up)
	}
	return newPoint
}

func (p Point) Equal(p2 Point) bool {
	return p.X == p2.X && p.Y == p2.Y && p.Z == p2.Z
}

func (p Point) Less(p2 Point) bool {
	if p.X != p2.X {
		return p.X < p2.X
	}
	if p.Y != p2.Y {
		return p.Y < p2.Y
	}
	return p.Z < p2.Z
}

type Cluster []Point

func (c Cluster) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c Cluster) Less(i, j int) bool {
	return c[i].Less(c[j])
}

func (c Cluster) Len() int {
	return len(c)
}

func (c *Cluster) Print() {
	for _, p := range *c {
		fmt.Printf("%v,%v,%v\n", p.X, p.Y, p.Z)
	}
}

func (c *Cluster) Rotate(face, up Axis, reverse bool) *Cluster {
	var newCluster Cluster
	for _, p := range *c {
		if reverse {
			newCluster = append(newCluster, p.RRotate(face, up))
		} else {
			newCluster = append(newCluster, p.Rotate(face, up))
		}
	}
	return &newCluster
}

func (s *Scanner) Rotations() []Scanner {
	var scanners []Scanner
	for _, face := range Rots {
		for _, up := range Rots {
			if face.Dir != up.Dir {
				scanners = append(scanners, Scanner{Count: s.Count, Points: *s.Points.Rotate(face, up, false), Face: face, Up: up})
			}
		}
	}
	return scanners
}

func (c *Cluster) Rebase(min Point) *Cluster {
	var newCluster Cluster
	for _, p := range *c {
		newCluster = append(newCluster, Point{X: p.X - min.X, Y: p.Y - min.Y, Z: p.Z - min.Z})
	}
	return &newCluster
}

func Subsets(size int, c Cluster) []Cluster {
	var newClusters []Cluster
	if size <= 0 {
		return []Cluster{}
	}
	if size == len(c) {
		return []Cluster{c}
	}
	with := Subsets(size-1, c[1:])
	without := Subsets(size, c[1:])
	for _, w := range with {
		newClusters = append(newClusters, append(w, c[0]))
	}
	for _, w := range without {
		newClusters = append(newClusters, w)
	}
	return newClusters
}

type Scanner struct {
	RelativeTo *Scanner
	Count      int
	Points     Cluster
	Position   Point
	Face       Axis
	Up         Axis
}

func (s *Scanner) Compare(s2 Scanner) bool {
	count := 0
	for _, p := range s.Points {
		for _, p2 := range s2.Points {
			if p.Equal(p2) {
				count += 1
				break
			}
		}
	}
	return count >= requiredOverlap
}

type ScannerPair struct {
	Scanner1, Scanner2 Scanner
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	//contents, err := ioutil.ReadFile("fake.txt")
	//contents, err := ioutil.ReadFile("small.txt")
	//contents, err := ioutil.ReadFile("rot.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	var scans []Scanner
	for _, line := range bytes.Split(contents, []byte("\n")) {
		if len(line) == 0 {
			continue
		}
		if string(line[:3]) == "---" {
			scans = append(scans, Scanner{})
			scans[len(scans)-1].Count = len(scans) - 1
			continue
		}
		vals := bytes.Split(line, []byte(","))
		point := Point{}
		point.X, _ = strconv.Atoi(string(vals[0]))
		point.Y, _ = strconv.Atoi(string(vals[1]))
		if len(vals) == 3 {
			point.Z, _ = strconv.Atoi(string(vals[2]))
		}
		scans[len(scans)-1].Points = append(scans[len(scans)-1].Points, point)
	}
	//scans[0].Points.Print()
	//scans[1].Points.Print()
	//sort.Sort(scans[0].Points)
	//sort.Sort(scans[1].Points)
	//scans[0].Points.Print()
	//scans[1].Points.Print()
	//scans[0].Points.Rebase().Print()
	//scans[1].Points.Rebase().Print()
	pathToZero := make(map[int]ScannerPair)
	var fullScans [][]Scanner
	for _, scan := range scans {
		fullScans = append(fullScans, make([]Scanner, 0, 24*26))
		rots := scan.Rotations()
		for _, rot := range rots {
			for _, p := range rot.Points {
				fullScans[rot.Count] = append(fullScans[rot.Count], Scanner{Count: rot.Count, Face: rot.Face, Up: rot.Up, Points: *rot.Points.Rebase(p), Position: p})
			}
		}
	}
	var pairs []ScannerPair
	pathToZero[0] = ScannerPair{scans[0], scans[0]}
	for i, scanList := range fullScans {
		for t, fScanList := range fullScans {
			if t == i {
				continue
			}
			for _, scan := range scanList {
				found := false
				for _, fScan := range fScanList {
					if scan.Compare(fScan) {
						found = true
						pairs = append(pairs, ScannerPair{scan, fScan})
						if _, ok := pathToZero[t]; ok {
							pathToZero[i] = pairs[len(pairs)-1]
						}
					}
				}
				if found {
					break
				}
			}
			if _, ok := pathToZero[i]; ok {
				break
			}
		}
	}
	beacons := make(map[Point]bool)
	for len(pathToZero) < len(scans) {
		for _, pair := range pairs {
			if _, ok := pathToZero[pair.Scanner1.Count]; !ok {
				if _, ok := pathToZero[pair.Scanner2.Count]; ok {
					pathToZero[pair.Scanner1.Count] = pair
				}
			}
		}
	}

	var scannerPos []Point
	for i, scan := range scans {
		if _, ok := pathToZero[i]; !ok {
			continue
		}
		adjusted := scan.Points
		var scannerPosition Point
		for l := i; l != 0; {
			if path, ok := pathToZero[l]; ok {
				adjusted = *adjusted.Rotate(path.Scanner1.Face, path.Scanner1.Up, false)
				scannerPosition = scannerPosition.Rotate(path.Scanner1.Face, path.Scanner1.Up)
				adjusted = *adjusted.Rebase(path.Scanner1.Position)
				scannerPosition.X -= path.Scanner1.Position.X
				scannerPosition.Y -= path.Scanner1.Position.Y
				scannerPosition.Z -= path.Scanner1.Position.Z
				adjusted = *adjusted.Rebase(Point{
					X: -1 * path.Scanner2.Position.X,
					Y: -1 * path.Scanner2.Position.Y,
					Z: -1 * path.Scanner2.Position.Z,
				})
				scannerPosition.X += path.Scanner2.Position.X
				scannerPosition.Y += path.Scanner2.Position.Y
				scannerPosition.Z += path.Scanner2.Position.Z
				adjusted = *adjusted.Rotate(path.Scanner2.Face, path.Scanner2.Up, true)
				scannerPosition = scannerPosition.RRotate(path.Scanner2.Face, path.Scanner2.Up)
				l = path.Scanner2.Count
			} else {
				break
			}
		}
		for _, p := range adjusted {
			beacons[p] = true
		}
		scannerPos = append(scannerPos, scannerPosition)
	}
	fmt.Println(len(beacons))
	max := 0
	for _, sP := range scannerPos {
		for _, sP2 := range scannerPos {
			dist := int(math.Abs(float64(sP.X-sP2.X))) + int(math.Abs(float64(sP.Y-sP2.Y))) + int(math.Abs(float64(sP.Z-sP2.Z)))
			if dist > max {
				max = dist
			}
		}
	}
	fmt.Println(max)
}
