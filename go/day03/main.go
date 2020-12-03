package main

import (
	"fmt"

	"github.com/SoMuchForSubtlety/aoc2020/go/pkg/input"
)

func main() {
	lines := input.ReadLines(3)

	fmt.Println("part 1:", slope(3, 1, lines))
	fmt.Println("part 2:", slope(1, 1, lines)*slope(3, 1, lines)*slope(5, 1, lines)*slope(7, 1, lines)*slope(1, 2, lines))
}

func slope(xOffset, yOffset int, lines []string) int {
	var x, y, trees int

	for y < len(lines)-1 {
		x = (x + xOffset) % len(lines[0])
		y += yOffset

		if lines[y][x] == '#' {
			trees++
		}
	}
	return trees
}
