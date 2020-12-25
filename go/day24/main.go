package main

import (
	"fmt"

	"github.com/SoMuchForSubtlety/aoc2020/go/pkg/input"
)

type Direction int

const (
	NE Direction = iota
	E
	SE
	SW
	W
	NW
)

type Tile struct {
	location [2]int
	black    bool
	black2   bool
}

type Tiles map[[2]int]*Tile

func (t Tiles) Get(location [2]int) *Tile {
	if neighbour := t[location]; neighbour != nil {
		return neighbour
	}
	neighbour := &Tile{location: location}
	t[location] = neighbour
	return neighbour
}

func offsetLocation(start [2]int, direction Direction) [2]int {
	newLocation := [2]int{start[0], start[1]}
	switch direction {
	case NE:
		newLocation[0] += 1
		newLocation[1] += 1
	case E:
		newLocation[0] += 2
	case SE:
		newLocation[0] += 1
		newLocation[1] -= 1
	case SW:
		newLocation[0] -= 1
		newLocation[1] -= 1
	case W:
		newLocation[0] -= 2
	case NW:
		newLocation[0] -= 1
		newLocation[1] += 1
	}

	return newLocation
}

func (t *Tile) neighbours(tiles Tiles) [6]*Tile {
	var out [6]*Tile
	for d := 0; d <= int(NW); d++ {
		out[d] = tiles.Get(offsetLocation(t.location, Direction(d)))
	}
	return out
}

func main() {
	start := [2]int{0, 0}
	tiles := make(Tiles)
	for _, line := range input.ReadLines(24) {
		current := start
		directions := toDirections(line)
		for _, direction := range directions {
			current = offsetLocation(current, direction)
		}
		tile := tiles.Get(current)
		tile.black = !tile.black
		tile.black2 = tile.black
	}

	var total int
	for _, tile := range tiles {
		if tile.black {
			total++
		}
	}
	fmt.Println("part 1:", total)

	for i := 0; i < 100; i++ {
		var currentTargets [][2]int
		for _, tile := range tiles {
			tile.black = tile.black2
			if tile.black {
				currentTargets = append(currentTargets, tile.location)
				for _, n := range tile.neighbours(tiles) {
					currentTargets = append(currentTargets, n.location)
				}
			}
		}
		for _, tile := range currentTargets {
			tiles[tile].change(tiles)
		}
	}
	var total2 int
	for _, tile := range tiles {
		if tile.black2 {
			total2++
		}
	}
	fmt.Println("part 2:", total2)
}

func (t *Tile) change(tiles Tiles) {
	var blackCount int
	for _, tile := range t.neighbours(tiles) {
		if tile.black {
			blackCount++
		}
	}
	if t.black && (blackCount == 0 || blackCount > 2) {
		t.black2 = false
	} else if !t.black && blackCount == 2 {
		t.black2 = true
	}
}

func toDirections(movement string) []Direction {
	var movements []Direction
	for len(movement) > 0 {
		if len(movement) > 1 {
			switch movement[:2] {
			case "ne":
				movements = append(movements, NE)
				movement = movement[2:]
				continue
			case "se":
				movements = append(movements, SE)
				movement = movement[2:]
				continue
			case "nw":
				movements = append(movements, NW)
				movement = movement[2:]
				continue
			case "sw":
				movements = append(movements, SW)
				movement = movement[2:]
				continue
			}
		}
		switch movement[0] {
		case 'e':
			movements = append(movements, E)
			movement = movement[1:]
		case 'w':
			movements = append(movements, W)
			movement = movement[1:]
		}
	}
	return movements
}
