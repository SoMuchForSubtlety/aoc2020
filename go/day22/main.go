package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/SoMuchForSubtlety/aoc2020/go/pkg/input"
)

type Deck []int

func main() {
	content := input.Read(22)
	players := strings.Split(content, "\n\n")
	player1 := toDeck(players[0])
	player2 := toDeck(players[1])
	for len(player1) > 0 && len(player2) > 0 {
		var v1, v2 int
		v1, player1 = player1.pop()
		v2, player2 = player2.pop()
		if v1 > v2 {
			player1 = player1.push(v1, v2)
		} else {
			player2 = player2.push(v2, v1)
		}
	}
	if len(player1) == 0 {
		fmt.Println("part 1:", player2.score())
	} else {
		fmt.Println("part 1:", player1.score())
	}

	_, score := recRound(toDeck(players[0]), toDeck(players[1]))
	fmt.Println("part 2:", score)
}

// returns true if player 1 wins, false if player 2 wins
func recRound(d1, d2 Deck) (bool, int) {
	var states [][2]int
	for len(d1) > 0 && len(d2) > 0 {
		state := [2]int{d1.hash(), d2.hash()}
		for _, s := range states {
			if s == state {
				return true, d1.score()
			}
		}
		states = append(states, state)

		var v1, v2 int
		v1, d1 = d1.pop()
		v2, d2 = d2.pop()
		var player1Wins bool
		if len(d1) >= v1 && len(d2) >= v2 {
			d12 := make(Deck, v1)
			copy(d12, d1)
			d22 := make(Deck, v2)
			copy(d22, d2)
			player1Wins, _ = recRound(d12, d22)
		} else {
			player1Wins = v1 > v2
		}
		if player1Wins {
			d1 = d1.push(v1, v2)
		} else {
			d2 = d2.push(v2, v1)
		}
	}
	if len(d1) > 0 {
		return true, d1.score()
	} else {
		return false, d2.score()
	}
}

func (d Deck) score() int {
	var total int
	for i := len(d) - 1; i >= 0; i-- {
		total += (d[i] * (len(d) - i))
	}
	return total
}

func (d Deck) hash() int {
	var total int
	var k int
	for i := 0; i < len(d); i++ {
		total += d[i] * k
		k += 523
	}
	return total
}

func (d Deck) pop() (int, Deck) {
	return d[0], d[1:]
}

func (d Deck) push(vals ...int) Deck {
	return append(d, vals...)
}

func toDeck(player string) Deck {
	var deck Deck
	for i, line := range strings.Split(player, "\n") {
		if i != 0 {
			val, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			deck = append(deck, val)
		}
	}
	return deck
}
