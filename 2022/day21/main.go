package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

type Monkey struct {
	Name   string
	Val    int
	HasVal bool
	V1, V2 string
	Op     byte
}

func FindHuman(nme string, mons map[string]*Monkey) bool {
	n := mons[nme]
	if n.Name == "humn" {
		return true
	}
	if n.HasVal {
		return false
	}
	return FindHuman(n.V1, mons) || FindHuman(n.V2, mons)
}

func Value(nme string, mons map[string]*Monkey) int {
	n := mons[nme]
	if n.HasVal {
		return n.Val
	}
	switch n.Op {
	case '+':
		return Value(n.V1, mons) + Value(n.V2, mons)
	case '/':
		return Value(n.V1, mons) / Value(n.V2, mons)
	case '-':
		return Value(n.V1, mons) - Value(n.V2, mons)
	case '*':
		return Value(n.V1, mons) * Value(n.V2, mons)
	}
	return -1
}

func Derive(nme string, mons map[string]*Monkey, look int) int {
	n := mons[nme]
	if n.Name == "humn" {
		return look
	}
	v := 0
	hName := ""
	var left bool
	if FindHuman(n.V1, mons) {
		v = Value(n.V2, mons)
		hName = n.V1
		left = true
	} else if FindHuman(n.V2, mons) {
		v = Value(n.V1, mons)
		hName = n.V2
	}
	switch n.Op {
	case '+':
		return Derive(hName, mons, look-v)
	case '/':
		if left {
			return Derive(hName, mons, look*v)
		} else {
			return Derive(hName, mons, v/look)
		}
	case '-':
		if left {
			return Derive(hName, mons, look+v)
		} else {
			return Derive(hName, mons, v-look)
		}
	case '*':
		return Derive(hName, mons, look/v)
	}
	return -1
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	mons := make(map[string]*Monkey)

	for _, line := range bytes.Split(contents, []byte("\n")) {
		if len(line) == 0 {
			continue
		}
		split := bytes.Split(line, []byte(": "))
		m := Monkey{Name: string(split[0])}
		v, err := strconv.Atoi(string(split[1]))
		if err == nil {
			m.Val = v
			m.HasVal = true
		} else {
			words := bytes.Split(split[1], []byte(" "))
			m.V1 = string(words[0])
			m.V2 = string(words[2])
			m.Op = words[1][0]
		}
		mons[m.Name] = &m
	}
	fmt.Println(Value("root", mons))

	mons = make(map[string]*Monkey)

	for _, line := range bytes.Split(contents, []byte("\n")) {
		if len(line) == 0 {
			continue
		}
		split := bytes.Split(line, []byte(": "))
		m := Monkey{Name: string(split[0])}
		v, err := strconv.Atoi(string(split[1]))
		if err == nil {
			m.Val = v
			if m.Name != "humn" {
				m.HasVal = true
			}
		} else {
			words := bytes.Split(split[1], []byte(" "))
			m.V1 = string(words[0])
			m.V2 = string(words[2])
			m.Op = words[1][0]
		}
		mons[m.Name] = &m
	}
	n := mons["root"]
	if FindHuman(n.V1, mons) {
		fmt.Println(Derive(n.V1, mons, Value(n.V2, mons)))
	}
	if FindHuman(n.V2, mons) {
		fmt.Println(Derive(n.V2, mons, Value(n.V1, mons)))
	}
}
