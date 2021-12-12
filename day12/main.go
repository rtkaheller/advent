package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
  "strings"
)

type Node struct {
  Big bool
  Name string
  Links []string
}

//type Link struct {
//  Start *Node
//  End *Node
//}

func paths (start *Node, visited *map[string]bool, nodes *map[string]*Node, double bool) int {
  if start.Name == "end" {
    return 1
  }
  localDub := double
  if _, ok := (*visited)[start.Name]; ok && !start.Big {
    if double == false && start.Name != "start"{
      localDub = true
    } else {
      return 0
    }
  }
  visits := 0
  local := make(map[string]bool)
  for c := range (*visited) {
    local[c] = true
  }

  if !start.Big {
    local[start.Name] = true
  }
  for _, child := range start.Links {
    visits += paths((*nodes)[child], &local, nodes, localDub)
  }
  return visits
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	//contents, err := ioutil.ReadFile("fake.txt")
	//contents, err := ioutil.ReadFile("med.txt")
	//contents, err := ioutil.ReadFile("small.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
  nodes := make(map[string]*Node)
	for _, line := range bytes.Split(contents, []byte("\n")) {
		if len(line) == 0 {
			continue
		}
	  path := bytes.Split(line, []byte("-"))
    a := string(path[0])
    b := string(path[1])
    if _, ok := nodes[a]; !ok {
      nodes[a] = &Node{Name: a, Links: []string{b}, Big: strings.ToUpper(a) == a }
    } else {
      nodes[a].Links = append(nodes[a].Links, b)
    }
    if _, ok := nodes[b]; !ok {
      nodes[b] = &Node{Name: b, Links: []string{a}, Big: strings.ToUpper(b) == b}
    } else {
      nodes[b].Links = append(nodes[b].Links, a)
    }
  }
  var start = nodes["start"]
  visited := make(map[string]bool)
  fmt.Println(paths(start, &visited, &nodes, true))
  visited = make(map[string]bool)
  fmt.Println(paths(start, &visited, &nodes, false))
}
