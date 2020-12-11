package main

import (
	"fmt"

	"github.com/SoMuchForSubtlety/aoc2020/go/pkg/input"
)

func main() {
	seats := input.ReadLines(11)

	grid := make([][]rune, len(seats))
	gridCopy := make([][]rune, len(seats))
	gridOriginal := make([][]rune, len(seats))
	for i, row := range seats {
		grid[i] = []rune(row)
		gridCopy[i] = make([]rune, len(grid[i]))
		gridOriginal[i] = make([]rune, len(grid[i]))
	}

	cloneGrid(gridOriginal, grid)

	newStateFn1 := getNewStateFunc(4)
	countNeighboursFn1 := getNeighbourCountFn(func(grid [][]rune, x, y, xx, yy int) int {
		if grid[yy][xx] == '#' {
			return 1
		} else {
			return 0
		}
	})

	for {
		if changed, occupied := cloneGrid(gridCopy, grid); !changed {
			fmt.Println("part 1:", occupied)
			break
		}
		applyForEachSeat(grid, gridCopy, func(a, b [][]rune, x, y int) {
			grid[y][x] = newStateFn1(gridCopy[y][x], countNeighboursFn1(gridCopy, x, y))
		})
	}

	cloneGrid(grid, gridOriginal) // reset seats
	newStateFn2 := getNewStateFunc(5)
	countNeighboursFn2 := getNeighbourCountFn(canSeeNeighbour2)
	for {
		if changed, occupied := cloneGrid(gridCopy, grid); !changed {
			fmt.Println("part 2:", occupied)
			break
		}
		applyForEachSeat(grid, gridCopy, func(a, b [][]rune, x, y int) {
			grid[y][x] = newStateFn2(gridCopy[y][x], countNeighboursFn2(gridCopy, x, y))
		})
	}
}

func cloneGrid(dst, src [][]rune) (bool, int) {
	var diff bool
	var occupied int
	for i := range src {
		for j := range src[i] {
			if src[i][j] != dst[i][j] {
				diff = true
			}
			if src[i][j] == '#' {
				occupied++
			}
		}
		copy(dst[i], src[i])
	}
	return diff, occupied
}

func getNewStateFunc(max int) func(input rune, neighbours int) rune {
	return func(input rune, neighbours int) rune {
		switch input {
		case 'L':
			if neighbours == 0 {
				return '#'
			} else {
				return 'L'
			}
		case '#':
			if neighbours >= max {
				return 'L'
			} else {
				return '#'
			}
		default:
			return '.'
		}
	}
}

func applyForEachSeat(grid, gridCopy [][]rune, fn func(a, b [][]rune, x, y int)) {
	for y := range grid {
		for x := range grid[y] {
			fn(grid, gridCopy, x, y)
		}
	}
}

func getNeighbourCountFn(fn func(grid [][]rune, x, y, xx, yy int) int) func([][]rune, int, int) int {
	return func(grid [][]rune, x, y int) int {
		var total int
		for xx := x - 1; xx <= x+1; xx++ { // 3x3 kernel
			for yy := y - 1; yy <= y+1; yy++ {
				if (xx != x || yy != y) && !(xx < 0 || yy < 0 || xx >= len(grid[0]) || yy >= len(grid)) { // don't check yourself
					total += fn(grid, x, y, xx, yy)
				}
			}
		}
		return total
	}
}

func canSeeNeighbour2(grid [][]rune, x, y, xx, yy int) int {
	vecX := xx - x
	vecY := yy - y
	for {
		if xx < 0 || xx >= len(grid[0]) || yy < 0 || yy >= len(grid) {
			return 0 // stop checking if we hit an edge
		}
		if grid[yy][xx] == '#' { // occupied
			return 1
		}
		if grid[yy][xx] != '.' { // can't see further
			return 0
		}

		xx += vecX
		yy += vecY
	}
}
