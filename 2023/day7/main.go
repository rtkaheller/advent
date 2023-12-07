package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

var jokerStrength = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 1,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

var strength = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

type Hand struct {
	H string
	B int
}

var mult = 10000000000 // Enough spaces to pack

func high(hand string, str map[rune]int) int {
	c := mult / 100
	t := 0
	for _, r := range hand {
		t += str[r] * c
		c /= 100
	}
	return t
}

func handStrength(hand string, joker bool) int {
	var str map[rune]int
	if joker {
		str = jokerStrength
	} else {
		str = strength
	}
	c := map[rune]int{}
	m := 0
	for _, r := range hand {
		if _, ok := c[r]; !ok {
			c[r] = 1
		} else {
			c[r] += 1
		}
		if c[r] > m && (!joker || r != 'J') {
			m = c[r]
		}
	}
	if !joker || c['J'] == 0 {
		two := 0
		for _, i := range c {
			if i == 2 {
				two += 1
			}
		}
		switch m {
		case 5, 4:
			return (m+2)*mult + high(hand, str)
		case 3:
			if two > 0 {
				return 5*mult + high(hand, str)
			} else {
				return 4*mult + high(hand, str)
			}
		case 2:
			if two > 1 {
				return 3*mult + high(hand, str)
			} else {
				return 2*mult + high(hand, str)
			}
		default:
			return mult + high(hand, str)
		}
	} else {
		two := 0
		for r, i := range c {
			if i == 2 && r != 'J' {
				two += 1
			}
		}
		switch m + c['J'] {
		case 5, 4:
			return (m+c['J']+2)*mult + high(hand, str)
		case 3:
			if two > 1 {
				return 5*mult + high(hand, str)
			} else {
				return 4*mult + high(hand, str)
			}
		case 2:
			if two > 1 {
				return 3*mult + high(hand, str)
			} else {
				return 2*mult + high(hand, str)
			}
		default:
			// With J, we can't have less than 2 of a kind
			return 2*mult + high(hand, str)
		}
	}
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := bytes.Split(contents, []byte("\n"))
	hands := []Hand{}
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		spl := strings.Fields(string(line))
		hand := spl[0]
		bid, _ := strconv.Atoi(spl[1])
		hands = append(hands, Hand{hand, bid})
	}
	sort.SliceStable(hands, func(i, j int) bool {
		return handStrength(hands[i].H, false) < handStrength(hands[j].H, false)
	})
	sum := 0
	for i, h := range hands {
		sum += (i + 1) * h.B
	}
	fmt.Println(sum)
	sort.SliceStable(hands, func(i, j int) bool {
		return handStrength(hands[i].H, true) < handStrength(hands[j].H, true)
	})
	sum = 0
	for i, h := range hands {
		sum += (i + 1) * h.B
	}
	fmt.Println(sum)
}
