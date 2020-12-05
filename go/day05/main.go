package main

import (
	"fmt"
	"sort"

	"github.com/SoMuchForSubtlety/aoc2020/go/pkg/input"
)

func main() {
	max := -1
	var ids []int
	for _, pass := range input.ReadLines(5) {
		id := search(pass, 127, 7, 0, 'B', 'F')*8 + search(pass, 7, 3, 7, 'R', 'L')
		if id > max {
			max = id
		}
		ids = append(ids, id)
	}

	sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })

	fmt.Println("part 1:", max)
	for i, id := range ids {
		if ids[i+1] == id+2 {
			fmt.Println("part 2:", id+1)
			return
		}
	}

}

func search(pass string, max, iterations, offset int, up, down byte) int {
	min, max := 0, max
	for i := 0; i < iterations; i++ {
		switch pass[i+offset] {
		case down:
			max -= (max-min)/2 + 1
		case up:
			min += (max-min)/2 + 1
		}
	}
	return min
}
