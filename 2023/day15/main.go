package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

func Hash(entry []byte) int {
	hash := 0
	for _, c := range entry {
		if c == '\n' {
			continue
		}
		hash += int(c)
		hash *= 17
		hash = hash % 256
	}
	return hash
}

type Lens struct {
	Focal int
	Index int
	Label string
}

type Box struct {
	Lenses map[string]int
	Order  []Lens
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	boxes := map[int]Box{}
	for i := 0; i < 256; i++ {
		boxes[i] = Box{map[string]int{}, []Lens{}}
	}
	sum := 0
	entries := bytes.Split(contents, []byte(","))
	for _, entry := range entries {
		if len(entry) == 0 {
			continue
		}
		if entry[len(entry)-1] == '\n' {
			entry = entry[:len(entry)-1]
		}
		sum += Hash(entry)
		var label []byte
		var lens int
		if entry[len(entry)-2] == '=' {
			label = entry[:len(entry)-2]
			lens, _ = strconv.Atoi(string(entry[len(entry)-1]))
			h := Hash(label)
			box := boxes[h]
			if _, ok := box.Lenses[string(label)]; !ok {
				newLens := Lens{Focal: lens, Index: len(box.Order), Label: string(label)}
				box.Order = append(boxes[h].Order, newLens)
				box.Lenses[string(label)] = newLens.Index
			} else {
				box.Order[box.Lenses[string(label)]].Focal = lens
			}
			boxes[h] = box
		} else {
			label = entry[:len(entry)-1]
			h := Hash(label)
			box := boxes[h]
			if l, ok := box.Lenses[string(label)]; ok {
				delete(box.Lenses, string(label))
				if l >= len(box.Order)-1 {
					box.Order = box.Order[:len(box.Order)-1]
				} else {
					box.Order = append(box.Order[:l], box.Order[l+1:]...)
				}
				for i := 0; i < len(box.Order); i++ {
					box.Order[i].Index = i
					box.Lenses[box.Order[i].Label] = i
				}
				boxes[h] = box
			}
		}
	}
	fmt.Println(sum)
	sum = 0
	for i, b := range boxes {
		for t, l := range b.Order {
			sum += (i + 1) * (t + 1) * l.Focal
		}
	}
	fmt.Println(sum)
}
