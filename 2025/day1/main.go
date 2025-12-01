package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	contents, err := ioutil.ReadFile("fake.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := bytes.Split(contents, []byte("\n"))
	cur := 50
	pwd := 0
  pwd2 := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		dir := line[0]
		num, _ := strconv.Atoi(string(line[1:len(line)]))
    pwd2 += num / 100
    num = num % 100
    if num == 0 {
      continue
    }
		switch dir{
		case 'L':
			cur -= num
		case 'R':
			cur += num
		}
		if cur < 0 {
      if (cur + num) != 0 {
        pwd2 += 1
      }
			cur = 100 + cur
		} else if cur > 100 {
      pwd2 += 1
    }
    cur = cur % 100
		if cur == 0 {
			pwd += 1
      pwd2 += 1
		}
	}
	fmt.Println(pwd)
	fmt.Println(pwd2)
}
