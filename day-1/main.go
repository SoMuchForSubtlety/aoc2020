package main

import (
	"fmt"

	"github.com/SoMuchForSubtlety/aoc2020/pkg/input"
)

func main() {
	nums := input.ReadInts()

	for _, num := range nums {
		for _, num2 := range nums {
			if num+num2 < 2020 {
				for _, num3 := range nums {
					if num+num2+num3 == 2020 {
						fmt.Printf("part 2: %d\n", num*num2*num3)
					}
				}
			}
			if num+num2 == 2020 {
				fmt.Printf("part 1: %d\n", num*num2)
			}
		}
	}
}
