package main

import (
	"fmt"
	"sort"

	"github.com/SoMuchForSubtlety/aoc2020/go/pkg/input"
)

func main() {
	ints := input.ReadInts(9)

	previous := make([]int, 25)
	var location int
	preamble := true
	var invalid int
	for _, v := range ints {
		if location >= 25 {
			preamble = false
			location = 0
		}
		if !preamble {
			if !hasSum(v, ints) {
				fmt.Println("part 1:", v)
				invalid = v
				break
			}
		}
		previous[location] = v
		location++
	}

	var sumRange []int
	var sum int
	for _, v := range ints {
		if sum+v > invalid {
			var x int
			for sum+v > invalid {
				x, sumRange = sumRange[0], sumRange[1:]
				sum -= x
			}
		}
		sumRange = append(sumRange, v)
		sum += v
		if sum == invalid {
			sort.Slice(sumRange, func(i, j int) bool { return sumRange[i] > sumRange[j] })
			fmt.Println("part 2:", sumRange[0]+sumRange[len(sumRange)-1])
			return
		}
	}
}

func hasSum(num int, nums []int) bool {
	for _, x := range nums {
		for _, y := range nums {
			if x+y == num {
				return true
			}
		}
	}
	return false
}
