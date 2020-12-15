package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/SoMuchForSubtlety/aoc2020/go/pkg/input"
)

func main() {
	occurences := make(map[int][]int)
	var last, round int

	startingNums := strings.Split(input.Read(15), ",")
	for i, num := range startingNums {
		n, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal(err)
		}
		occurences[n] = append(occurences[n], i+1)
		last = n
	}
	round = len(startingNums)

	for ; round < 30_000_000; round++ {
		if round == 2020 {
			fmt.Println("part 1:", last)
		}
		oc := occurences[last]
		if len(oc) == 1 {
			last = 0
		} else {
			last = oc[len(oc)-1] - oc[len(oc)-2]
		}
		occurences[last] = append(occurences[last], round+1)
	}
	fmt.Println("part 2:", last)
}
