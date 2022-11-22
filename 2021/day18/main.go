package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

type SnailFish struct {
	Val    int
	Parent *SnailFish
	Left   *SnailFish
	Right  *SnailFish
}

type Stack struct {
	Data []*SnailFish
}

func (s *Stack) Push(val *SnailFish) {
	s.Data = append(s.Data, val)
}

func (s *Stack) Pop() (*SnailFish, bool) {
	if len(s.Data) == 0 {
		return nil, true
	}
	val := s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1]
	return val, false
}

func (s *SnailFish) Print() {
	if s.Left == nil {
		fmt.Printf("%v", s.Val)
		return
	}
	fmt.Printf("[")
	s.Left.Print()
	fmt.Printf(",")
	s.Right.Print()
	fmt.Printf("]")
}

func (s *SnailFish) Traverse(depth int) int {
	if s.Left == nil {
		if depth >= 3 {
			return 1
		}
	}
	if s.Val >= 10 {
		return 2
	}
	val := s.Left.Traverse(depth + 1)
	switch val {
	case 1:
		s.Left.Explode()
		return -1
	case 2:
		s.Left.Split()
		return -1
	case -1:
		return -1
	}

	val = s.Right.Traverse(depth + 1)
	switch val {
	case 1:
		s.Right.Explode()
		return -1
	case 2:
		s.Right.Split()
		return -1
	case -1:
		return -1
	}
	return -1
}
func (s *SnailFish) Listify() []*SnailFish {
	if s.Left == nil {
		return []*SnailFish{s}
	}
	return append(s.Left.Listify(), s.Right.Listify()...)
}

func (s *SnailFish) Explode() bool {
	var stack Stack
	stack.Push(s)
	for {
		cur, done := stack.Pop()
		if done {
			return false
		}
		depth := 0
		for p := cur.Parent; p != nil; p = p.Parent {
			depth += 1
		}
		if cur.Right != nil {
			stack.Push(cur.Right)
		}
		if cur.Left != nil {
			stack.Push(cur.Left)
		} else {
			if depth > 4 {
				p := cur.Parent
				list := s.Listify()
				for i, l := range list {
					if l == p.Left {
						if i > 0 {
							list[i-1].Val += p.Left.Val
						}
						if i < len(list)-2 {
							list[i+2].Val += p.Right.Val
						}
						break
					}
				}
				p.Val = 0
				p.Left = nil
				p.Right = nil
				//if p.Parent.Left == p {
				//  p.Parent.Left = p.Parent.Right.Left
				//  if p.Parent.Left != nil {
				//    p.Parent.Left.Parent = p.Parent
				//  }
				//  p.Parent.Val = p.Parent.Right.Val
				//  p.Parent.Right = p.Parent.Right.Right
				//  if p.Parent.Right != nil {
				//    p.Parent.Right.Parent = p.Parent
				//  }
				//}
				//if p.Parent.Right == p {
				//  p.Parent.Right = p.Parent.Left.Right
				//  if p.Parent.Right != nil {
				//    p.Parent.Right.Parent = p.Parent
				//  }
				//  p.Parent.Val = p.Parent.Left.Val
				//    p.Parent.Left = p.Parent.Left.Left
				//  if p.Parent.Left != nil {
				//    p.Parent.Left.Parent = p.Parent
				//  }
				//}
				return true
			}
		}
	}
	return false
}

func (s *SnailFish) Split() bool {
	for _, f := range s.Listify() {
		if f.Val >= 10 {
			var left, right SnailFish
			left.Val = f.Val / 2
			right.Val = f.Val/2 + f.Val%2
			left.Parent = f
			right.Parent = f
			f.Val = 0
			f.Left = &left
			f.Right = &right
			return true
		}
	}
	return false
}

func (s *SnailFish) Reduce() {
	for {
		if !s.Explode() && !s.Split() {
			return
		}
	}
}

func AddSnails(v1, v2 *SnailFish) *SnailFish {
	result := SnailFish{Left: v1, Right: v2}
	v1.Parent = &result
	v2.Parent = &result
	return &result
}

func Parse(s string, p *SnailFish) *SnailFish {
	var f SnailFish
	f.Parent = p
	c := 0
	val, err := strconv.Atoi(s)
	if err == nil {
		f.Val = val
		return &f
	}
	for i := range s {
		switch s[i] {
		case '[':
			c++
		case ']':
			c--
		case ',':
			if c == 1 {
				f.Left = Parse(s[1:i], &f)
				f.Right = Parse(s[i+1:len(s)-1], &f)
				return &f
			}
		}
	}
	return nil
}

func (s *SnailFish) Magnitude() int {
	if s.Left == nil {
		return s.Val
	}
	return s.Left.Magnitude()*3 + s.Right.Magnitude()*2
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	//contents, err := ioutil.ReadFile("fake.txt")
	//contents, err := ioutil.ReadFile("small.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	lines := bytes.Split(contents, []byte("\n"))
	var answer = Parse(string(lines[0]), nil)
	for _, line := range lines[1:] {
		if len(line) == 0 {
			break
		}
		answer = AddSnails(answer, Parse(string(line), nil))
		answer.Reduce()
	}
	fmt.Println(answer.Magnitude())
	max := 0
	var x, y string
	for i, line := range lines {
		if len(line) == 0 {
			break
		}
		for _, line2 := range lines[i+1:] {
			if len(line2) == 0 {
				break
			}
			a := AddSnails(Parse(string(line), nil), Parse(string(line2), nil))
			a.Reduce()
			if a.Magnitude() > max {
				max = a.Magnitude()
				x = string(line)
				y = string(line2)
			}
			a = AddSnails(Parse(string(line2), nil), Parse(string(line), nil))
			a.Reduce()
			if a.Magnitude() > max {
				max = a.Magnitude()
				y = string(line)
				x = string(line2)
			}
		}
	}
	fmt.Println(max)
	fmt.Println(x)
	fmt.Println(y)
}
