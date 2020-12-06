package main

import (
	"fmt"
	"strings"

	"github.com/SoMuchForSubtlety/aoc2020/go/pkg/input"
)

func main() {
	groups := strings.Split(input.Read(6), "\n\n")

	var total1 int
	var total2 int

	for _, group := range groups {
		answers := make(map[rune]int)
		people := strings.Split(group, "\n")
		for _, person := range people {
			for _, answer := range person {
				answers[answer] = answers[answer] + 1
			}
		}
		total1 += len(answers)
		for _, count := range answers {
			if count == len(people) {
				total2++
			}
		}
	}

	fmt.Println("part 1:", total1)
	fmt.Println("part 2:", total2)
}
