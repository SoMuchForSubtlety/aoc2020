package main

import (
	"fmt"
	"strconv"

	"github.com/SoMuchForSubtlety/aoc2020/go/pkg/input"
)

func main() {
	in := input.Read(23)
	cups := playGame(in, false)
	fmt.Println("part 1:", part1(cups))

	cups = playGame(in, true)
	fmt.Println("part 2:", cups[1]*cups[cups[1]])
}

func playGame(input string, part2 bool) []int {
	cups, current := createCups(input, part2)
	rounds, maxCup := 100, 9
	if part2 {
		rounds, maxCup = 10_000_000, 1_000_000
	}

	for i := 0; i < rounds; i++ {
		three := cups[current]
		cups[current] = cups[cups[cups[three]]]

		destination := getDestinationLabel(cups, current, three, maxCup)
		cups[cups[cups[three]]] = cups[destination]
		cups[destination] = three

		current = cups[current]
	}

	return cups
}

func getDestinationLabel(cups []int, current, next, max int) int {
	i := current - 1
	if i < 1 {
		i = max
	}
	for i == next || i == cups[next] || i == cups[cups[next]] {
		i--
		if i < 1 {
			i = max
		}
	}
	return i
}

func part1(cups []int) string {
	c := 1
	var out string
	for cups[c] != 1 {
		c = cups[c]
		out += strconv.Itoa(c)
	}
	return out
}

func createCups(input string, part2 bool) ([]int, int) {
	cups := make([]int, 9+1)
	if part2 {
		cups = make([]int, 1_000_000+1)
	}
	var start, prev int
	for i, num := range input {
		label, _ := strconv.Atoi(string(num))
		if i == 0 {
			start = label
		} else {
			cups[prev] = label
		}
		prev = label
	}

	if part2 {
		for i := 10; i <= 1_000_000; i++ {
			cups[prev] = i
			prev = i
		}
	}
	cups[prev] = start

	return cups, start
}
