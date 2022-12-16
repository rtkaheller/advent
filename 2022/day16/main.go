package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
)

type Valve struct {
	Rate    int
	Tunnels map[string]*Valve
	Tuns    []string
	Name    string
}

type MemoKey struct {
	V       string
	Visited string
	Time    int
}

type Result struct {
	V int
	A []string
}

func Release(v *Valve, visited []string, time int) (int, []string) {
	var memoVis string
	memoVisited := make([]string, len(visited))
	copy(memoVisited, visited)
	sort.Strings(memoVisited)
	for _, k := range memoVisited {
		memoVis += k
	}
	if r, ok := memo[MemoKey{V: v.Name, Visited: memoVis, Time: time}]; ok {
		return r.V, r.A
	}
	if time <= 1 {
		return 0, []string{}
	}
	maxRelease := 0
	var acts []string
	found := false
	for _, vis := range visited {
		if vis == v.Name {
			found = true
			break
		}
	}
	if !found && v.Rate > 0 {
		var chain []string
		maxRelease, chain = Release(v, append(visited, v.Name), time-1)
		maxRelease += v.Rate * time
		acts = []string{v.Name + fmt.Sprintf(" - release: %d (t: %d)", v.Rate, time)}
		acts = append(acts, chain...)
	}

	for _, dest := range v.Tunnels {
		r, chain := Release(dest, visited, time-1)
		if r > maxRelease {
			maxRelease = r
			acts = []string{}
			acts = append(acts, v.Name+" -> "+dest.Name+fmt.Sprintf(" (t: %d)", time))
			acts = append(acts, chain...)
		}
	}
	memo[MemoKey{V: v.Name, Visited: memoVis, Time: time}] = Result{maxRelease, acts}
	return maxRelease, acts
}

var memo map[MemoKey]Result

func Release2(me, el *Valve, visited []string, time int, maxValves int) (int, []string) {
	if len(visited) == maxValves {
		return 0, []string{}
	}
	var memoVis string
	memoVisited := make([]string, len(visited))
	copy(memoVisited, visited)
	sort.Strings(memoVisited)
	for _, k := range memoVisited {
		memoVis += k
	}
	curName := ""
	if me.Name > el.Name {
		curName = me.Name + el.Name
	} else {
		curName = el.Name + me.Name
	}
	if r, ok := memo[MemoKey{V: curName, Visited: memoVis, Time: time}]; ok {
		return r.V, r.A
	}
	if time <= 1 {
		return 0, []string{}
	}

	maxRelease := 0
	var acts []string
	var foundMe, foundEl bool
	for _, vis := range visited {
		if vis == me.Name {
			foundMe = true
		}
		if vis == el.Name {
			foundEl = true
		}
	}
	var ElMoves, MeMoves []*Valve

	if !foundMe && me.Rate > 0 { // Might want to release my valve
		MeMoves = append(MeMoves, me)
	}

	if !foundEl && el.Rate > 0 && el.Name != me.Name { // Might want to release el's valve (if we aren't together)
		ElMoves = append(ElMoves, el)
	}

	for _, dest := range me.Tunnels {
		MeMoves = append(MeMoves, dest)
	}
	for _, dest := range el.Tunnels {
		ElMoves = append(ElMoves, dest)
	}
	for _, mm := range MeMoves {
		for _, em := range ElMoves {
			newVis := make([]string, len(visited))
			bonus := 0
			copy(newVis, visited)
			if mm == me { // I stay to open a valve
				newVis = append(newVis, mm.Name)
				bonus += mm.Rate * time
			}
			if em == el { // El stay to open a valve
				newVis = append(newVis, em.Name)
				bonus += em.Rate * time
			}
			r, _ := Release2(mm, em, newVis, time-1, maxValves)
			if r+bonus > maxRelease {
				maxRelease = r + bonus
			}
		}
	}
	memo[MemoKey{V: curName, Visited: memoVis, Time: time}] = Result{maxRelease, acts}
	return maxRelease, []string{}
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	valves := make(map[string]*Valve)
	for _, line := range bytes.Split(contents, []byte("\n")) {
		if len(line) == 0 {
			continue
		}
		words := bytes.Split(line, []byte(" "))
		name := string(words[1])
		valve := Valve{Name: name, Tunnels: make(map[string]*Valve)}

		rate := bytes.Split(words[4], []byte("="))
		val, _ := strconv.Atoi(string(rate[1][:len(rate[1])-1]))
		valve.Rate = val
		for _, dest := range words[9:] {
			if dest[len(dest)-1] == ',' {
				valve.Tuns = append(valve.Tuns, string(dest[:len(dest)-1]))
			} else {
				valve.Tuns = append(valve.Tuns, string(dest))
			}
		}
		valves[name] = &valve
	}
	maxValves := 0
	for _, v := range valves {
		for _, t := range v.Tuns {
			v.Tunnels[t] = valves[t]
			if v.Rate > 0 {
				maxValves += 1
			}
		}
	}
	memo = make(map[MemoKey]Result)
	p1, _ := Release(valves["AA"], []string{}, 29)
	fmt.Println(p1)

	memo = make(map[MemoKey]Result)
	p2, _ := Release2(valves["AA"], valves["AA"], []string{}, 25, maxValves)
	fmt.Println(p2)
}
