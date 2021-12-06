package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
)

const (
	bitLength = 12
)

func search(vals []int, tar int) int {
	if len(vals) < 2 {
		return 0
	}
	if tar > vals[len(vals)/2] {
		return search(vals[len(vals)/2:], tar) + len(vals)/2
	} else if tar < vals[len(vals)/2] {
		return search(vals[:len(vals)/2], tar)
	} else {
		return len(vals) / 2
	}
}

func specificBitCount(vals []int, length, bit int) int {
	count := 0
	for _, v := range vals {
		count += (v >> (length - bit - 1)) % 2
	}
	return count

}
func countBits(vals []int, length int) []int {
	counts := make([]int, length)
	for i := 0; i < length; i++ {
		for _, v := range vals {
			counts[i] += (v >> (length - i - 1)) % 2
		}
	}
	return counts
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	var bitInts []int
	var eInts []int
	power := bytes.Split(contents, []byte("\n"))
	total := 0
	for _, bits := range power {
		var val int
		if len(bits) != bitLength {
			continue
		}
		for i, bit := range bits {
			if bit == byte('1') {
				val += 1 << (bitLength - i - 1)
			}
		}
		bitInts = append(bitInts, val)
		eInts = append(eInts, val)
		total += 1
	}
	counts := countBits(bitInts, bitLength)
	sort.Ints(bitInts)
	sort.Ints(eInts)
	gamma := 0
	epsilon := 0
	for i, c := range counts {
		if c >= total/2 {
			gamma += 1 << (bitLength - i - 1)
		} else {
			epsilon += 1 << (bitLength - i - 1)
		}
	}

	//fmt.Println(total)
	//fmt.Println(counts)
	//fmt.Println(gamma)
	//fmt.Println(epsilon)
	//fmt.Println(gamma*epsilon)

	ox := 0
	co := 0
	fmt.Println(bitInts)
	splitVal := 0
	for i := 0; i < bitLength; i++ {
		fmt.Println(i)
		count := specificBitCount(bitInts, bitLength, i)
		fmt.Println(count)
		fmt.Println(len(bitInts))
		if count >= len(bitInts)-count {
			fmt.Println("top")
			splitVal += 1 << (bitLength - i - 1)
			fmt.Println(splitVal)
			split := search(bitInts, splitVal)
			fmt.Println(split)
			if bitInts[split] == splitVal {
				fmt.Println("plus one")
				split -= 1
			}
			bitInts = bitInts[split+1:]
		} else {
			fmt.Println("bottom")
			fmt.Println(splitVal + 1<<(bitLength-i-1))
			split := search(bitInts, splitVal+1<<(bitLength-i-1))
			if bitInts[split] == splitVal+1<<(bitLength-i-1) {
				fmt.Println("plus one")
				split -= 1
			}
			fmt.Println(split)
			bitInts = bitInts[:split+1]
		}
		fmt.Println(bitInts)
		if len(bitInts) == 1 {
			ox = bitInts[0]
			break
		}
	}

	fmt.Println("eInts")
	splitVal = 0
	fmt.Println(eInts)
	for i := 0; i < bitLength; i++ {
		fmt.Println(i)
		count := specificBitCount(eInts, bitLength, i)

		fmt.Println(len(eInts))
		if count < len(eInts)-count {
			fmt.Println("top")
			splitVal += 1 << (bitLength - i - 1)
			fmt.Println(splitVal)
			split := search(eInts, splitVal)
			fmt.Println(split)
			if eInts[split] == splitVal {
				fmt.Println("plus one")
				split -= 1
			}
			eInts = eInts[split+1:]
		} else {
			fmt.Println("bottom")
			fmt.Println(splitVal + 1<<(bitLength-i-1))
			split := search(eInts, splitVal+1<<(bitLength-i-1))
			if eInts[split] == splitVal+1<<(bitLength-i-1) {
				fmt.Println("plus one")
				split -= 1
			}
			fmt.Println(split)
			eInts = eInts[:split+1]
		}
		fmt.Println(eInts)
		if len(eInts) == 1 {
			co = eInts[0]
			break
		}
	}
	fmt.Println(ox)
	fmt.Println(co)
	fmt.Println(co * ox)
}
