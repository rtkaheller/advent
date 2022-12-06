package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func startOfPacket(chars []byte, size int) int {
	for i, _ := range chars {
		k := make(map[byte]bool)
		for t := i; t >= 0 && t >= i-size; t-- {
			if _, ok := k[chars[t]]; ok {
				break
			} else {
				k[chars[t]] = true
			}
		}
		if len(k) == size {
			return i + 1
		}
	}
	return -1
}

func main() {
	contents, err := ioutil.ReadFile("fake.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, chars := range bytes.Split(contents, []byte("\n")) {
		if len(chars) == 0 {
			continue
		}
		fmt.Println(startOfPacket(chars, 4))
		fmt.Println(startOfPacket(chars, 14))
	}
}
