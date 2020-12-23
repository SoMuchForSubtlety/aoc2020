package main

import "testing"

func BenchmarkPart1(b *testing.B) {
	cups, current := createCups("974618352", false)
	maxCup := 9

	for i := 0; i < b.N; i++ {
		three := cups[current]
		cups[current] = cups[cups[cups[three]]]

		destination := getDestinationLabel(cups, current, three, maxCup)
		cups[cups[cups[three]]] = cups[destination]
		cups[destination] = three

		current = cups[current]
	}
}

func BenchmarkPart2(b *testing.B) {
	cups, current := createCups("974618352", true)
	maxCup := 1_000_000

	for i := 0; i < b.N; i++ {
		three := cups[current]
		cups[current] = cups[cups[cups[three]]]

		destination := getDestinationLabel(cups, current, three, maxCup)
		cups[cups[cups[three]]] = cups[destination]
		cups[destination] = three

		current = cups[current]
	}
}
