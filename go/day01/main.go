package main

import (
	"fmt"
	"sort"

	"github.com/SoMuchForSubtlety/aoc2020/go/pkg/input"
)

func main() {
	nums := input.ReadInts(1)

	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	i, j := sum2(nums, 2020)
	fmt.Println("part 1:", nums[i]*nums[j])

	for k, num := range nums {
		i, j = sum2(nums, 2020-num)
		if i != j && k != i && k != j {
			fmt.Println("part 2:", nums[i]*nums[j]*nums[k])
			return
		}
	}
}

func sum2(nums []int, target int) (i, j int) {
	lo := 0
	hi := len(nums) - 1

	for lo < hi && nums[lo]+nums[hi] != target {
		if nums[hi]+nums[lo] > target {
			hi--
		} else {
			lo++
		}
	}
	return lo, hi
}
