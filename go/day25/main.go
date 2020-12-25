package main

import (
	"fmt"

	"github.com/SoMuchForSubtlety/aoc2020/go/pkg/input"
)

func main() {
	pubs := input.ReadInts(25)
	v := 1
	var cardLoops int
	for v != pubs[0] {
		v = (v * 7) % 20201227
		cardLoops++
	}
	v = 1
	for i := 0; i < cardLoops; i++ {
		v = (v * pubs[1]) % 20201227
	}
	fmt.Println("part 1:", v)
}
