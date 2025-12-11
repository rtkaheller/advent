package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	C     []*Node
	Label string
}

var memo = map[*Node][][]*Node{}

func Paths(cur, goal *Node) [][]*Node {
	if cur == goal {
		return [][]*Node{[]*Node{cur}}
	}
	if v, ok := memo[cur]; ok {
		return v
	}

	r := [][]*Node{}
	for _, c := range cur.C {
		for _, p := range Paths(c, goal) {
			r = append(r, append(p, cur))
		}
	}
	memo[cur] = r
	return r
}

func main() {
	contents, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	lines := bytes.Split(contents, []byte("\n"))

	p1 := 0
	p2 := 0
	p := map[string]*Node{"out": &Node{Label: "out"}}
	child := map[string][]string{}
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		f := strings.Fields(string(line))
		k := f[0][:len(f[0])-1]
		child[k] = f[1:]
		n := Node{}
		n.Label = k
		p[k] = &n
	}
	for _, n := range p {
		for _, c := range child[n.Label] {
			n.C = append(n.C, p[c])
		}
	}
	you := p["you"]
	out := p["out"]
	svr := p["svr"]
	dac := p["dac"]
	fft := p["fft"]
	p1 = len(Paths(you, out))
	memo = map[*Node][][]*Node{}
	p2 = len(Paths(svr, fft))
	memo = map[*Node][][]*Node{}
	p2 *= len(Paths(fft, dac))
	memo = map[*Node][][]*Node{}
	p2 *= len(Paths(dac, out))

	fmt.Println(p1)
	fmt.Println(p2)
}
