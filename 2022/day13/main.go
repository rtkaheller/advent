package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
)

type Pair struct {
	Left, Right Packet
}

type Packet struct {
	Data []Packet
	Val  int
	Int  bool
}

type Packets []*Packet

func (l Packets) Len() int {
	return len(l)
}

func (l Packets) Less(i, j int) bool {
	v, _ := Compare(*l[i], *l[j])
	return v == -1
}

func (l Packets) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func ParsePacket(line []byte) *Packet {
	val, err := strconv.Atoi(string(line))
	if err == nil {
		return &Packet{Val: val, Int: true}
	}
	var children [][]byte
	depth := 0
	comma := 0
	for i, c := range line {
		switch c {
		case '[':
			if depth == 0 {
				comma = i
			}
			depth += 1
		case ']':
			depth -= 1
			if depth == 0 {
				children = append(children, line[comma+1:i])
			}
		case ',':
			if depth == 1 {
				children = append(children, line[comma+1:i])
				comma = i
			}
		}
	}
	var ret Packet
	for _, child := range children {
		ret.Data = append(ret.Data, *ParsePacket(child))
	}
	return &ret
}

func Compare(left, right Packet) (int, string) {
	if left.Int && right.Int {
		if left.Val < right.Val {
			return -1, "int comp l < r"
		}
		if left.Val > right.Val {
			return 1, "int comp r > l"
		}
		return 0, "int comp r == l"
	}
	if left.Int && !right.Int {
		return Compare(Packet{Data: []Packet{left}}, right)
	}
	if !left.Int && right.Int {
		return Compare(left, Packet{Data: []Packet{right}})
	}
	for i := 0; i < len(left.Data); i++ {
		if i >= len(right.Data) {
			return 1, "list comp, right out of dat"
		}
		if ret, reason := Compare(left.Data[i], right.Data[i]); ret != 0 {
			return ret, reason + "list comp, left != right"
		}
	}
	if len(right.Data) == len(left.Data) {
		return 0, "list comp, end of lists"
	}
	return -1, "end of func"
}

func main() {
	contents, err := ioutil.ReadFile("fake.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	var pairs []Pair
	var packets Packets
	lines := bytes.Split(contents, []byte("\n"))
	for i := 0; i < len(lines); i += 3 {
		if len(lines[i]) == 0 {
			continue
		}

		pairs = append(pairs, Pair{Left: *ParsePacket(lines[i]), Right: *ParsePacket(lines[i+1])})
		packets = append(packets, &pairs[len(pairs)-1].Left)
		packets = append(packets, &pairs[len(pairs)-1].Right)
	}
	sum := 0
	for i, p := range pairs {
		v, _ := Compare(p.Left, p.Right)
		if v == -1 {
			sum += (i + 1)
		}
	}
	fmt.Println(sum)

	d1 := ParsePacket([]byte("[[2]]"))
	packets = append(packets, d1)
	d2 := ParsePacket([]byte("[[6]]"))
	packets = append(packets, d2)
	sort.Sort(packets)
	key := 1
	for i, p := range packets {
		if p == d1 || p == d2 {
			key *= (i + 1)
		}
	}
	fmt.Println(key)
}
