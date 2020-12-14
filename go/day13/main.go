package main

import (
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"

	"github.com/SoMuchForSubtlety/aoc2020/go/pkg/input"
)

func main() {
	lines := input.ReadLines(13)

	minStart, err := strconv.Atoi(lines[0])
	if err != nil {
		log.Fatal(err)
	}

	var busLines []int
	for _, line := range strings.Split(lines[1], ",") {
		if line == "x" {
			busLines = append(busLines, -1)
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		busLines = append(busLines, num)
	}

	bestBus := -1
	for _, bus := range busLines {
		if bus != -1 && (bestBus == -1 || (bus-minStart%bus < bestBus-minStart%bestBus)) {
			bestBus = bus
		}
	}

	fmt.Println("part 1:", bestBus*(bestBus-minStart%bestBus))

	var remainders []*big.Int
	var values []*big.Int

	for i, bus := range busLines {
		remainders = append(remainders, big.NewInt(int64((bus-i)%bus)))
		values = append(values, big.NewInt(int64(bus)))
	}
	fmt.Println("part 2:", crt(remainders, values))
}

func crt(a, n []*big.Int) *big.Int {
	p := new(big.Int).Set(n[0])
	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}
	var x, q, s, z big.Int
	for i, n1 := range n {
		q.Div(p, n1)
		z.GCD(nil, &s, n1, &q)
		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}
	return x.Mod(&x, p)
}
