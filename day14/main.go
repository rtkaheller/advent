package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
)

const (
	rounds = 40
)

type Results map[rune]int

func (r *Results) Update(char rune, inc int) {
	if _, ok := (*r)[char]; !ok {
		(*r)[char] = inc
	} else {
		(*r)[char] += inc
	}
}

var memo map[string]Results

type Insertion struct {
	First, Second, Insert rune
}

func depthInsert(template string, rules map[string]*Insertion, depth int) Results {
	if depth == 0 {
		results := make(Results)
		for _, r := range template {
			results.Update(r, 1)
		}
		return results
	}
	if r, ok := memo[fmt.Sprintf("%v, %v", string(template), depth)]; ok {
		return r
	}
	result := make(Results)
	var last rune
	for _, r := range template {
		if last != 0 {
			if insert, ok := rules[string(last)+string(r)]; ok {
				head := depthInsert(string(last)+string(insert.Insert), rules, depth-1)
				tail := depthInsert(string(insert.Insert)+string(r), rules, depth-1)
				for char, val := range head {
					result.Update(char, val)
				}
				for char, val := range tail {
					result.Update(char, val)
				}
				result.Update(last, -1)
				result.Update(insert.Insert, -1)
				result.Update(r, -1)
			}
		}
		result.Update(r, 1)
		last = r
	}
	memo[fmt.Sprintf("%v, %v", string(template), depth)] = result
	return result
}

func doInsert(template string, rules map[string]*Insertion) string {
	result := ""
	var last rune
	for _, r := range template {
		if last != 0 {
			if insert, ok := rules[string(last)+string(r)]; ok {
				result += string(insert.Insert)
			}
		}
		result += string(r)
		last = r
	}
	return result
}

func main() {
	memo = make(map[string]Results)
	contents, err := ioutil.ReadFile("input.txt")
	//contents, err := ioutil.ReadFile("fake.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	lines := bytes.Split(contents, []byte("\n"))
	template := string(lines[0])
	insertions := make(map[string]*Insertion)
	for i := 2; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			break
		}
		pairs := bytes.Split(lines[i], []byte(" -> "))
		pair := []rune(string(pairs[0]))
		insert := []rune(string(pairs[1]))
		insertions[string(pairs[0])] = &Insertion{First: pair[0], Second: pair[1], Insert: insert[0]}
	}
	counts := depthInsert(template, insertions, 10)
	min := math.MaxInt64
	max := 0
	for _, c := range counts {
		if c > max {
			max = c
		}
		if c < min {
			min = c
		}
	}
	fmt.Println(max - min)
	counts = depthInsert(template, insertions, 40)
	min = math.MaxInt64
	max = 0
	for _, c := range counts {
		if c > max {
			max = c
		}
		if c < min {
			min = c
		}
	}
	fmt.Println(max - min)
}
