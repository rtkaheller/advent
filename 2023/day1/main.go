package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

var nums = map[string]int{
  "one": 1,
  "two": 2,
  "three": 3,
  "four": 4,
  "five": 5,
  "six": 6,
  "seven": 7,
  "eight": 8,
  "nine": 9,
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

  sum := 0
  sum2 := 0
	lines := bytes.Split(contents, []byte("\n"))
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
    first := -1
    last := -1
    for _, c := range line {
      val, err := strconv.Atoi(string(c))
      if err == nil {
        if first == -1 {
          first = val
        }
        last = val
      }
    }
    sum += first*10 + last

    first = -1
    last = -1
    for i := 0; i < len(line); i++ {
      val, err := strconv.Atoi(string(line[i]))
      if err == nil {
        first = val
        break
      }
      for s, n := range nums {
        if len(line) > i+len(s) {
          if string(line[i:i+len(s)]) == s {
            first = n
          }
        }
      }
      if first != -1 {
        break
      }
    }

    for i := len(line)-1; i >= 0; i-- {
      val, err := strconv.Atoi(string(line[i]))
      if err == nil {
        last = val
        break
      }
      for s, n := range nums {
        if i-len(s) >= 0 {
          if string(line[i-len(s)+1:i+1]) == s {
            last = n
          }
        }
      }
      if last != -1 {
        break
      }
    }
    sum2 += first*10 + last

	}
  fmt.Println(sum)
  fmt.Println(sum2)
}
