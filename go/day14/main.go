package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/SoMuchForSubtlety/aoc2020/go/pkg/input"
)

var assignmentRegex = regexp.MustCompile(`mem\[([0-9]+)\] = ([0-9]+)`)

func main() {
	lines := input.ReadLines(14)

	memory := make(map[uint64]uint64)
	memory2 := make(map[uint64]uint64)

	var bitmask string
	for _, line := range lines {
		if strings.HasPrefix(line, "mask = ") {
			bitmask = strings.TrimPrefix(line, "mask = ")
			continue
		}

		address, value := parseAssignment(line)
		memory[address] = mask1(bitmask, value)
		for _, addr := range mask2(bitmask, address) {
			memory2[addr] = value
		}
	}

	var sum, sum2 uint64
	for _, v := range memory {
		sum += v
	}
	fmt.Println("part 1:", sum)

	for _, v := range memory2 {
		sum2 += v
	}
	fmt.Println("part 2:", sum2)
}

func mask2(bitmask string, address uint64) []uint64 {
	for i := uint64(len(bitmask) - 1); i < uint64(len(bitmask)); i-- {
		mask := uint64(1) << (uint64(len(bitmask)) - (i + 1))
		if bitmask[i] == '1' {
			address |= (mask) // set bit high
		}
	}

	return combinations(bitmask, address, 0)
}

func combinations(bitmask string, address uint64, offset int) []uint64 {
	for offset < len(bitmask) && bitmask[offset] != 'X' {
		offset++
	}
	if offset >= len(bitmask) {
		return []uint64{address}
	}

	address2 := address
	mask := uint64(1) << (uint64(len(bitmask)) - (uint64(offset) + 1))
	if address|mask == address { // bit is high
		address2 ^= mask // set bit low
	} else {
		address2 |= mask // set bit high
	}

	return append(combinations(bitmask, address, offset+1), combinations(bitmask, address2, offset+1)...)
}

func mask1(bitmask string, value uint64) uint64 {
	for i := uint64(len(bitmask) - 1); i < uint64(len(bitmask)); i-- {
		mask := uint64(1) << (uint64(len(bitmask)) - (i + 1))
		switch bitmask[i] {
		case '1':
			value |= mask // set bit high
		case '0':
			value |= mask // set bit high
			value ^= mask // set bit low
		}
	}

	return value
}

func parseAssignment(line string) (address, value uint64) {
	var err error
	address, err = strconv.ParseUint(assignmentRegex.FindStringSubmatch(line)[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	value, err = strconv.ParseUint(assignmentRegex.FindStringSubmatch(line)[2], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	return address, value
}
