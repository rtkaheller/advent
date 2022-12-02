package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func score2(opp, resp rune) int {
  switch opp {
  case 'A': // Rock
    switch resp {
    case 'X': // Lose
      return 0 + 3
    case 'Y': // Draw
      return 3 + 1
    case 'Z': // Win
      return 6 + 2
    }
  case 'B': // Paper
    switch resp {
    case 'X': // Lose
      return 0 + 1
    case 'Y': // Draw
      return 3 + 2
    case 'Z': // Win
      return 6 + 3
    }
  case 'C': // Scissor
    switch resp {
    case 'X': // Lose
      return 0 + 2
    case 'Y': // Draw
      return 3 + 3
    case 'Z': // Win
      return 6 + 1
    }
  }
  return 0
}
func score(opp, resp rune) int {
  switch opp {
  case 'A': // Rock
    switch resp {
    case 'X': // Rock
      return 1 + 3
    case 'Y': // Paper
      return 2 + 6
    case 'Z': // Scissor
      return 3 + 0
    }
  case 'B': // Paper
    switch resp {
    case 'X': // Rock
      return 1 + 0
    case 'Y': // Paper
      return 2 + 3
    case 'Z': // Scissor
      return 3 + 6
    }
  case 'C': // Scissor
    switch resp {
    case 'X': // Rock
      return 1 + 6
    case 'Y': // Paper
      return 2 + 0
    case 'Z': // Scissor
      return 3 + 3
    }
  }
  return 0
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

  games := make(map[rune]rune)
	rows := bytes.Split(contents, []byte("\n"))
  sum := 0
  sum2 := 0
	for _, row := range rows {
    if len(row) == 0 {
      break
    }
    games[rune(string(row)[0])] = rune(string(row)[2])
    sum += score(rune(string(row)[0]), rune(string(row)[2]))
    sum2 += score2(rune(string(row)[0]), rune(string(row)[2]))
	}
  fmt.Println(sum)
  fmt.Println(sum2)
}
