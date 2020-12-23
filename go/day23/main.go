package main

import (
	"fmt"
	"strconv"
)

type Cup struct {
	label int
	next  *Cup
}

func (c *Cup) getDestinationLabel(three *Cup, max int) int {
	i := c.label - 1
	if i < 1 {
		i = max
	}
	for i == three.label || i == three.next.label || i == three.next.next.label {
		i--
		if i < 1 {
			i = max
		}
	}
	return i
}

func (c *Cup) PickUpThree() *Cup {
	start := c.next
	c.next = start.next.next.next
	return start
}

func (c *Cup) PlaceThree(start *Cup) {
	start.next.next.next = c.next
	c.next = start
}

func (c *Cup) Part1() string {
	if c.label != 1 {
		panic("wrong cup")
	}

	c = c.next
	var out string
	for c.label != 1 {
		out += strconv.Itoa(c.label)
		c = c.next
	}
	return out
}

func (c *Cup) Part2() int {
	if c.label != 1 {
		panic("wrong cup")
	}

	return c.next.label * c.next.next.label
}

func CreateCups(input string, part2 bool) (map[int]*Cup, *Cup) {
	cups := make(map[int]*Cup)
	var start, prev *Cup
	var max int
	for i, num := range input {
		label, _ := strconv.Atoi(string(num))
		if label > max {
			max = label
		}
		cup := &Cup{label: label}
		cups[cup.label] = cup
		if i == 0 {
			start = cup
		} else {
			prev.next = cup
		}

		prev = cup
	}

	if part2 {
		for i := max + 1; i <= 1_000_000; i++ {
			cup := &Cup{label: i}
			cups[cup.label] = cup
			prev.next = cup
			prev = cup
		}
	}
	prev.next = start

	return cups, start
}

func PlayGame(input string, part2 bool) map[int]*Cup {
	index, current := CreateCups(input, part2)
	rounds, maxCup := 100, 9
	if part2 {
		rounds, maxCup = 10_000_000, 1_000_000
	}

	for i := 0; i < rounds; i++ {
		three := current.PickUpThree()
		destination := index[current.getDestinationLabel(three, maxCup)]
		destination.PlaceThree(three)
		current = current.next
	}

	return index
}

func main() {
	cups := PlayGame("974618352", false)
	fmt.Println("part 1:", cups[1].Part1())

	cups = PlayGame("974618352", true)
	fmt.Println("part 2:", cups[1].Part2())
}
