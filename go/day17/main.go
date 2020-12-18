package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/SoMuchForSubtlety/aoc2020/go/pkg/input"
)

func main() {
	reactor := initializeRactor(input.ReadLines(17))

	for i := 0; i < 6; i++ {
		reactor = reactor.cycle()
	}
	var alive int
	for _, v := range reactor {
		if v {
			alive++
		}
	}
	fmt.Println(alive)
}

type Reactor map[string]bool

func (r Reactor) get(x, y, z, w int) bool {
	return r[toKey(x, y, z, w)]
}

func (r Reactor) set(x, y, z, w int, val bool) {
	r[toKey(x, y, z, w)] = val
	if !val {
		return
	}
	for xx := x - 1; xx <= x+1; xx++ {
		for yy := y - 1; yy <= y+1; yy++ {
			for zz := z - 1; zz <= z+1; zz++ {
				for ww := w - 1; ww <= w+1; ww++ {
					key := toKey(xx, yy, zz, ww)
					if _, ok := r[key]; !ok {
						r[key] = false
					}
				}
			}
		}
	}
}

func toKey(x, y, z, w int) string {
	return fmt.Sprintf("%d,%d,%d,%d", x, y, z, w)
}

func toCoords(k string) (x, y, z, w int) {
	vals := strings.Split(k, ",")
	x, _ = strconv.Atoi(vals[0])
	y, _ = strconv.Atoi(vals[1])
	z, _ = strconv.Atoi(vals[2])
	w, _ = strconv.Atoi(vals[3])

	return x, y, z, w
}

func initializeRactor(lines []string) Reactor {
	r := make(Reactor)
	for y, line := range lines {
		for x, col := range line {
			if col == '#' {
				r.set(x, y*-1, 0, 0, true)
			} else {
				r.set(x, y*-1, 0, 0, false)
			}
		}
	}

	return r
}

func (r Reactor) countNeighbours(x, y, z, w int) int {
	var activeNeighbours int
	for xx := x - 1; xx <= x+1; xx++ {
		for yy := y - 1; yy <= y+1; yy++ {
			for zz := z - 1; zz <= z+1; zz++ {
				for ww := w - 1; ww <= w+1; ww++ {
					if r.get(xx, yy, zz, ww) {
						activeNeighbours++
					}
				}
			}
		}
	}
	if r.get(x, y, z, w) { // don't count yourself
		activeNeighbours--
	}
	return activeNeighbours
}

func (r Reactor) cycle() Reactor {
	copy := make(Reactor)
	for k, v := range r {
		copy[k] = v
	}

	for k, v := range r {
		x, y, z, w := toCoords(k)
		neighbours := r.countNeighbours(x, y, z, w)
		switch v {
		case true:
			if neighbours < 2 || neighbours > 3 {
				copy.set(x, y, z, w, false)
			}
		case false:
			if neighbours == 3 {
				copy.set(x, y, z, w, true)
			}
		}
	}
	return copy
}
