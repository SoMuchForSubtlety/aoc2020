package main

import (
	"fmt"

	"github.com/SoMuchForSubtlety/aoc2020/go/pkg/input"
)

func main() {
	nums := input.ReadInts(1)

	for i, num := range nums {
		for j, num2 := range nums {
			for l, num3 := range nums {
				if i != j && i != l && num+num2+num3 == 2020 {
					fmt.Printf("part 2: %d\n", num*num2*num3)
				}
			}
			if i != j && num+num2 == 2020 {
				fmt.Printf("part 1: %d\n", num*num2)
			}
		}
	}

}
