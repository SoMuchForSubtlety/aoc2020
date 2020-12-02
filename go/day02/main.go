package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/SoMuchForSubtlety/aoc2020/go/pkg/input"
)

func main() {
	var valid int
	var valid2 int

	for _, line := range input.ReadLines(2) {
		min, max, char, char1, char2, pw := parseLine(line)

		count := strings.Count(pw, char)
		if count >= min && count <= max {
			valid++
		}

		if (char1 == char && char2 != char) || (char2 == char && char1 != char) {
			valid2++
		}
	}

	fmt.Printf("part 1: %d\n", valid)
	fmt.Printf("part 2: %d\n", valid2)
}

func parseLine(line string) (min, max int, char, char1, char2, pw string) {
	fields := strings.Fields(line)
	minmax := strings.Split(fields[0], "-")
	var err error
	min, err = strconv.Atoi(minmax[0])
	if err != nil {
		log.Fatal(err)
	}
	max, err = strconv.Atoi(minmax[1])
	if err != nil {
		log.Fatal(err)
	}

	return min, max, fields[1][0:1], string(fields[2][min-1]), string(fields[2][max-1]), fields[2]
}
