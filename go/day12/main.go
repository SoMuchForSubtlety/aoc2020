package main

import (
	"fmt"
	"strconv"

	"github.com/SoMuchForSubtlety/aoc2020/go/pkg/input"
)

func main() {
	movements := input.ReadLines(12)

	var x, y int
	var x2, y2 int
	direction := [2]int{1, 0}
	direction2 := [2]int{10, 1}

	for _, movement := range movements {
		x, y, direction = execute(x, y, direction, movement, false)
		x2, y2, direction2 = execute(x2, y2, direction2, movement, true)
	}
	fmt.Println("part 1:", Abs(x)+Abs(y))
	fmt.Println("part 2:", Abs(x2)+Abs(y2))
}

func execute(x, y int, movementVector [2]int, op string, part2 bool) (int, int, [2]int) {
	instruction := op[0]
	arg, err := strconv.Atoi(op[1:])
	if err != nil {
		panic(err)
	}

	var xOffset, yOffset int
	switch instruction {
	case 'N':
		yOffset += arg
	case 'S':
		yOffset -= arg
	case 'E':
		xOffset += arg
	case 'W':
		xOffset -= arg
	}

	if part2 {
		movementVector[0] += xOffset
		movementVector[1] += yOffset
	} else {
		x += xOffset
		y += yOffset
	}

	switch instruction {
	case 'L', 'R':
		for steps := arg / 90; steps > 0; steps-- {
			dx := movementVector[1]
			movementVector[1] = movementVector[0]
			movementVector[0] = dx

			if instruction == 'L' {
				movementVector[0] *= -1
			} else {
				movementVector[1] *= -1
			}
		}
	case 'F':
		x += movementVector[0] * arg
		y += movementVector[1] * arg
	}

	return x, y, movementVector
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
