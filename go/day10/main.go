package main

import (
	"fmt"
	"sort"

	"github.com/SoMuchForSubtlety/aoc2020/go/pkg/input"
)

func main() {
	adapters := input.ReadInts(10)
	sort.Slice(adapters, func(i, j int) bool { return adapters[i] < adapters[j] })
	var last int
	var steps [3]int
	for _, adapter := range adapters {
		step := (adapter - last) - 1
		steps[step]++
		last = adapter
	}

	steps[2]++ // phone
	fmt.Println("part 1:", steps[0]*steps[2])

	adapters = append([]int{0}, adapters...)
	combinations := make([]int, adapters[len(adapters)-1]+1)
	for i, adapter := range adapters {
		var options int
		if i >= 1 && adapter-adapters[i-1] <= 3 {
			options += combinations[adapters[i-1]]
		}
		if i >= 2 && adapter-adapters[i-2] <= 3 {
			options += combinations[adapters[i-2]]
		}
		if i >= 3 && adapter-adapters[i-3] <= 3 {
			options += combinations[adapters[i-3]]
		}
		if i == 0 {
			options = 1
		}
		combinations[adapter] = options
	}
	fmt.Println("pat 2:", combinations[len(combinations)-1])
}
