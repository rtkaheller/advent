package main

import (
	"fmt"
	"io/ioutil"
	"math"
)

const (
	verLength = 3
	idLength  = 3
)

var bigSum = 0

func bitVal(digits string) int {
	v := 0
	for i, digit := range digits {
		if digit == '1' {
			v += int(math.Pow(2, float64(len(digits)-i-1)))
		}
	}
	return v
}

func literalVal(digits string) (int, int) {
	nonPadded := ""
	last := false
	used := 0
	for i := 0; i < len(digits) && !last; i += 5 {
		if i+4 >= len(digits) {
			break
		}
		chunk := digits[i : i+5]
		if chunk[0] == '0' {
			last = true
		}
		used += len(chunk)
		nonPadded += chunk[1:]
	}
	return bitVal(nonPadded), used
}

func parsePacket(binPacket string) (int, int) {
	version := bitVal(binPacket[:verLength])
	binPacket = binPacket[verLength:]
	id := bitVal(binPacket[:idLength])
	binPacket = binPacket[idLength:]
	bigSum += version
	totUsed := verLength + idLength
	if id == 4 {
		val, used := literalVal(binPacket)
		return val, used + totUsed
	}
	lengthTypeID := binPacket[0]
	binPacket = binPacket[1:]
	totUsed += 1
	var vals []int
	if lengthTypeID == '0' {
		bitLength := bitVal(binPacket[:15])
		binPacket = binPacket[15:]
		totUsed += 15
		locUsed := 0
		for locUsed < bitLength {
			val, used := parsePacket(binPacket)
			binPacket = binPacket[used:]
			totUsed += used
			locUsed += used
			vals = append(vals, val)
		}
	} else {
		subLength := bitVal(binPacket[:11])
		binPacket = binPacket[11:]
		totUsed += 11
		for i := 0; i < subLength; i++ {
			val, used := parsePacket(binPacket)
			binPacket = binPacket[used:]
			totUsed += used
			vals = append(vals, val)
		}
	}
	result := 0
	switch id {
	case 0:
		for _, v := range vals {
			result += v
		}
	case 1:
		result = 1
		for _, v := range vals {
			result *= v
		}
	case 2:
		result = math.MaxInt64
		for _, v := range vals {
			if v < result {
				result = v
			}
		}
	case 3:
		for _, v := range vals {
			if v > result {
				result = v
			}
		}
	case 5:
		if vals[0] > vals[1] {
			result = 1
		}
	case 6:
		if vals[0] < vals[1] {
			result = 1
		}
	case 7:
		if vals[0] == vals[1] {
			result = 1
		}
	}

	return result, totUsed
}

var hex = map[rune]string{
	'0': "0000",
	'1': "0001",
	'2': "0010",
	'3': "0011",
	'4': "0100",
	'5': "0101",
	'6': "0110",
	'7': "0111",
	'8': "1000",
	'9': "1001",
	'A': "1010",
	'B': "1011",
	'C': "1100",
	'D': "1101",
	'E': "1110",
	'F': "1111",
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	//contents, err := ioutil.ReadFile("med.txt")
	//contents, err := ioutil.ReadFile("fake.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	var binPacket string
	for _, r := range string(contents) {
		if r == rune('\n') {
			break
		}
		binPacket += hex[r]
	}
	fmt.Println(parsePacket(binPacket))
	fmt.Println(bigSum)
}
