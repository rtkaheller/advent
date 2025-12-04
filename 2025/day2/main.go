package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	line := bytes.Split(contents, []byte("\n"))
	groups := bytes.Split(line[0], []byte(","))
	invalid := 0
	p2 := 0
	for _, group := range groups {
		if len(group) == 0 {
			continue
		}
		ends := bytes.Split(group, []byte("-"))
		first, _ := strconv.Atoi(string(ends[0]))
		second, _ := strconv.Atoi(string(ends[1]))
		for i := first; i <= second; i++ {
			id := fmt.Sprintf("%d", i)
			if id[0:len(id)/2] == id[len(id)/2:] {
				invalid += i
				p2 += i
				continue
			}
			for t := 1; t <= len(id)/2; t++ {
				if len(id)%t != 0 {
					continue
				}
				s := id[0:t]
				bad := true
				for j := t; j+t <= len(id); j += t {
					if id[j:j+t] != s {
						bad = false
					}
				}
				if bad {
					p2 += i
					break
				}
			}
		}
	}
	fmt.Println(invalid)
	fmt.Println(p2)
}
