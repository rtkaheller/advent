package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

type Num struct {
	Val        int
	Next, Prev *Num
}

func ListVals(first *Num) []int {
	var vals []int
	cur := first
	for {
		vals = append(vals, cur.Val)
		if cur.Next == first {
			break
		}
		cur = cur.Next
	}
	return vals[1:]
}

func Move(n *Num, first *Num) {
	length := len(ListVals(first))
	if n.Val > 0 {
		for i := 0; i < n.Val%(length-1); i++ {
			prev := n.Prev
			next := n.Next
			n2 := next.Next
			if next == first {
				i -= 1
			}
			n2.Prev = n
			n.Next = n2

			n.Prev = next
			next.Next = n

			prev.Next = next
			next.Prev = prev
		}
	}
	if n.Val < 0 {
		for i := 0; i < -1*n.Val%(length-1); i++ {
			prev := n.Prev
			next := n.Next
			p2 := prev.Prev
			if prev == first {
				i -= 1
			}
			p2.Next = n
			n.Prev = p2

			n.Next = prev
			prev.Prev = n

			prev.Next = next
			next.Prev = prev

		}
	}
}

func Mix(in []int, iterations int) []int {
	var nodes []*Num
	var first Num
	cur := &first
	for _, v := range in {
		n := Num{Val: v, Prev: cur}
		cur.Next = &n
		cur = &n
		nodes = append(nodes, &n)
	}
	cur.Next = &first
	first.Prev = cur
	for i := 0; i < iterations; i++ {
		for _, num := range nodes {
			Move(num, &first)
			cur = first.Next
		}
	}

	return ListVals(&first)
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	var file []int
	for _, line := range bytes.Split(contents, []byte("\n")) {
		if len(line) == 0 {
			continue
		}

		v, _ := strconv.Atoi(string(line))
		file = append(file, v)
	}
	res := Mix(file, 1)
	for i, v := range res {
		if v == 0 {
			fmt.Println(res[(i+1000)%len(res)] + res[(i+2000)%len(res)] + res[(i+3000)%len(res)])
			break
		}
	}

	var newFile []int
	for _, v := range file {
		newFile = append(newFile, v*811589153)
	}
	res = Mix(newFile, 10)
	for i, v := range res {
		if v == 0 {
			fmt.Println(res[(i+1000)%len(res)] + res[(i+2000)%len(res)] + res[(i+3000)%len(res)])
			break
		}
	}
}
