package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

const (
	diceSize = 100
)

func incDice(d int) int {
	d++
	if d > diceSize {
		d = 1
	}
	return d
}

type Winner struct {
	Player1, Player2 int64
}

type Params struct {
	p1, p2, s1, s2, left int64
	turn                 bool
}

var memo map[Params]Winner

func turn(pos, score, d int) (int, int, int) {
	pos += d
	d = incDice(d)
	pos += d
	d = incDice(d)
	pos += d
	d = incDice(d)
	pos = pos % 10
	score += pos + 1
	return pos, score, d
}

func winner(p1, p2, s1, s2, left int64, turn bool) (int64, int64) {
	params := Params{p1, p2, s1, s2, left, turn}
	if winner, ok := memo[params]; ok {
		return winner.Player1, winner.Player2
	}
	p1 = p1 % 10
	p2 = p2 % 10
	if left == 0 {
		if turn {
			s1 += p1 + 1
			if s1 >= 21 {
				memo[params] = Winner{1, 0}
				return 1, 0
			}
		} else {
			s2 += p2 + 1
			if s2 >= 21 {
				memo[params] = Winner{0, 1}
				return 0, 1
			}
		}
		turn = !turn
		left = 3
	}
	var tot1, tot2 int64
	if turn {
		w1, w2 := winner(p1+1, p2, s1, s2, left-1, turn)
		tot1 += w1
		tot2 += w2
		w1, w2 = winner(p1+2, p2, s1, s2, left-1, turn)
		tot1 += w1
		tot2 += w2
		w1, w2 = winner(p1+3, p2, s1, s2, left-1, turn)
		tot1 += w1
		tot2 += w2
	} else {
		w1, w2 := winner(p1, p2+1, s1, s2, left-1, turn)
		tot1 += w1
		tot2 += w2
		w1, w2 = winner(p1, p2+2, s1, s2, left-1, turn)
		tot1 += w1
		tot2 += w2
		w1, w2 = winner(p1, p2+3, s1, s2, left-1, turn)
		tot1 += w1
		tot2 += w2
	}
	memo[params] = Winner{tot1, tot2}
	return tot1, tot2
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	//contents, err := ioutil.ReadFile("fake.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	memo = make(map[Params]Winner)
	lines := bytes.Split(contents, []byte("\n"))
	p1 := bytes.Split(lines[0], []byte(" "))
	player1, _ := strconv.Atoi(string(p1[4]))
	p2 := bytes.Split(lines[1], []byte(" "))
	player2, _ := strconv.Atoi(string(p2[4]))

	player1 -= 1
	player2 -= 1

	p1Score := 0
	p2Score := 0
	dice := 1
	rolls := 0
	for {
		player1, p1Score, dice = turn(player1, p1Score, dice)
		rolls += 3
		if p1Score >= 1000 {
			fmt.Println(p2Score * rolls)
			break
		}
		player2, p2Score, dice = turn(player2, p2Score, dice)
		rolls += 3
		if p2Score >= 1000 {
			fmt.Println(p1Score * rolls)
			break
		}
		fmt.Println(p1Score, p2Score)
	}
	player1, _ = strconv.Atoi(string(p1[4]))
	player2, _ = strconv.Atoi(string(p2[4]))
	player1 -= 1
	player2 -= 1
	fmt.Println(winner(int64(player1), int64(player2), 0, 0, 3, true))

}
