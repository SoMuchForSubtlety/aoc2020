package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/SoMuchForSubtlety/aoc2020/go/pkg/input"
)

type Tile struct {
	id         string
	content    [][]rune
	neighbours map[Direction]*Neighbour
}

type Neighbour struct {
	tile     *Tile
	border   Direction
	reversed bool
}

type Direction int

const (
	top Direction = iota
	right
	bottom
	left
)

type Tiles map[string]*Tile

func directions() []Direction {
	return []Direction{top, right, bottom, left}
}

func main() {

	tiles := ParseTiles(input.Read(20))
	for k, tile1 := range tiles {
		for k2, tile2 := range tiles {
			if k == k2 {
				continue
			}
			for direction2, border := range tile2.GetBorders() {
				for _, direction := range directions() {
					border1 := tile1.Border(direction)
					if border1 == border {
						tile1.neighbours[direction] = &Neighbour{tile: tile2, reversed: false, border: Direction(direction2)}
					} else if Reverse(border1) == border {
						tile1.neighbours[direction] = &Neighbour{tile: tile2, reversed: true, border: Direction(direction2)}
					}
				}
			}
		}
	}
	corners := 1
	var bottomLeft *Tile
	for id, t := range tiles {
		if len(t.neighbours) == 2 {
			if (t.neighbours[left] == nil && t.neighbours[bottom] == nil) || bottomLeft == nil {
				bottomLeft = t
			}
			val, _ := strconv.Atoi(strings.Split(id, " ")[1])
			corners *= val
		}
	}
	fmt.Println("part 1:", corners)

	// make sure we area actually on bottom left
	for !(bottomLeft.neighbours[left] == nil && bottomLeft.neighbours[bottom] == nil) {
		bottomLeft.RotateClockwise(1)
	}

	// start from the bottom left and move upwards as follows
	// -->
	// <--
	// x->
	current := bottomLeft
	direction1, direction2 := right, left
	var lastDirection Direction
	var x, y int
	for {
		if y == 0 {
			x++
		}
		next := current.neighbours[direction1]
		lastDirection = direction1
		if next == nil {
			y++
			next = current.neighbours[top]
			lastDirection = top
			direction1, direction2 = direction2, direction1
			if next == nil {
				break
			}
		}
		// rotate the next tile in the sequence util the borders match up
		next.tile.RotateClockwise(2 + (int(lastDirection) - int(next.border)))
		if current.Border(lastDirection) != next.tile.Border(lastDirection.Opposite()) {
			next.tile.Flip(lastDirection)
		}

		current = next.tile
	}

	ocean := tiles.Stitch(x*(len(current.content[0])-2), y*(len(current.content)-2))
	var monsters int
	for {
		monsters = countMonsters(ocean)
		if monsters > 0 {
			break
		}
		Flip(ocean, top)
		monsters = countMonsters(ocean)
		if monsters > 0 {
			break
		}
		Flip(ocean, top)
		Flip(ocean, right)
		monsters = countMonsters(ocean)
		if monsters > 0 {
			break
		}
		Flip(ocean, right)

		RotateClockwise(ocean)
	}
	fmt.Println("part 2:", countWaves(ocean))
}

func countWaves(ocean [][]rune) int {
	var total int
	for y := 0; y < len(ocean); y++ {
		for x := 0; x < len(ocean[y]); x++ {
			if ocean[y][x] == '#' {
				total++
			}
		}
	}
	return total
}

func countMonsters(ocean [][]rune) int {
	monster := [][]rune{
		[]rune("                  # "),
		[]rune("#    ##    ##    ###"),
		[]rune(" #  #  #  #  #  #   "),
	}

	var monsters int
	for y := 0; y < len(ocean)-2; y++ {
		for x := 0; x < len(ocean[y])-19; x++ {
			for mY := 0; mY < len(monster); mY++ {
				for mX := 0; mX < len(monster[mY]); mX++ {
					if monster[mY][mX] != '#' {
						continue
					}
					if ocean[y+mY][x+mX] != '#' {
						goto next
					}
				}
			}
			monsters++
			for mY := 0; mY < len(monster); mY++ {
				for mX := 0; mX < len(monster[mY]); mX++ {
					if monster[mY][mX] != '#' {
						continue
					}
					ocean[y+mY][x+mX] = 'O'
				}
			}
		next:
		}
	}

	return monsters
}

func (d Direction) Opposite() Direction {
	switch d {
	case top:
		return bottom
	case bottom:
		return top
	case left:
		return right
	case right:
		return left
	default:
		panic("not reached")
	}
}

func (t *Tile) Border(d Direction) string {
	switch d {
	case top:
		return string(t.content[0])
	case bottom:
		return string(t.content[len(t.content)-1])
	}

	var border string
	for _, row := range t.content {
		if d == right {
			border += string(row[len(row)-1])
		} else {
			border += string(row[0])
		}
	}

	return border
}

func (t *Tile) GetBorders() []string {
	return []string{t.Border(top), t.Border(right), t.Border(bottom), t.Border(left)}
}

func (t *Tile) Flip(d Direction) {
	if d == left || d == right {
		tmp := t.neighbours[top]
		t.neighbours[top] = t.neighbours[bottom]
		t.neighbours[bottom] = tmp
	} else {
		tmp := t.neighbours[left]
		t.neighbours[left] = t.neighbours[right]
		t.neighbours[right] = tmp
	}

	Flip(t.content, d)
}

func (t *Tile) RotateClockwise(times int) {
	if times < 0 {
		times += 4
	}
	for times > 0 {
		times--
		RotateClockwise(t.content)
		tmp := t.neighbours[top]
		t.neighbours[top] = t.neighbours[left]
		t.neighbours[left] = t.neighbours[bottom]
		t.neighbours[bottom] = t.neighbours[right]
		t.neighbours[right] = tmp
	}
}

func (t Tiles) Stitch(xDim, yDim int) [][]rune {
	out := make([][]rune, yDim)
	for i := range out {
		out[i] = make([]rune, xDim)
	}
	var current *Tile
	for _, tile := range t {
		if tile.neighbours[top] == nil && tile.neighbours[left] == nil {
			current = tile
			break
		}
	}
	var xBase, yBase int
	direction1, direction2 := right, left
	xDirection := 1
	for {
		for y := 1; y < len(current.content)-1; y++ {
			for x := 1; x < len(current.content[y])-1; x++ {
				out[yBase+y-1][xBase+x-1] = current.content[y][x]
			}
		}
		xBase += (len(current.content[0]) - 2) * xDirection
		if xBase >= xDim || xBase < 0 {
			yBase += len(current.content) - 2
			xDirection *= -1
			xBase += (len(current.content[0]) - 2) * xDirection
		}

		next := current.neighbours[direction1]
		if next == nil {
			next = current.neighbours[bottom]
			direction1, direction2 = direction2, direction1
			if next == nil {
				break
			}
		}
		current = next.tile
	}

	return out
}

func ParseTiles(input string) Tiles {
	tiles := make(Tiles)
	for _, tile := range strings.Split(input, "\n\n") {
		parts := strings.Split(tile, ":\n")
		var content [][]rune
		for _, row := range strings.Split(parts[1], "\n") {
			content = append(content, []rune(row))
		}
		tiles[parts[0]] = &Tile{id: parts[0], content: content, neighbours: make(map[Direction]*Neighbour)}
	}
	return tiles
}

func Flip(tile [][]rune, d Direction) {
	if d == left || d == right {
		for i, j := 0, len(tile)-1; i < j; i, j = i+1, j-1 {
			tile[i], tile[j] = tile[j], tile[i]
		}
	} else {
		for y := 0; y < len(tile); y++ {
			for i, j := 0, len(tile[y])-1; i < j; i, j = i+1, j-1 {
				tile[y][i], tile[y][j] = tile[y][j], tile[y][i]
			}
		}
	}
}

func RotateClockwise(tile [][]rune) {
	n := len(tile)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			temp := tile[i][j]
			tile[i][j] = tile[n-1-j][i]
			tile[n-1-j][i] = tile[n-1-i][n-1-j]
			tile[n-1-i][n-1-j] = tile[j][n-1-i]
			tile[j][n-1-i] = temp
		}
	}
}

func Reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}
