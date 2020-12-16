package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/SoMuchForSubtlety/aoc2020/go/pkg/input"
)

type rule struct {
	name       string
	min1, max1 int
	min2, max2 int
}

type ticket []ticketField

type ticketField struct {
	val           int
	matchingRules []*rule
}

func main() {
	txt := input.Read(16)
	parts := strings.Split(txt, "\n\n")
	var rules []*rule

	r := regexp.MustCompile(`^([a-z ]+): (\d+)-(\d+) or (\d+)-(\d+)$`)
	for _, line := range strings.Split(parts[0], "\n") {
		parts := r.FindStringSubmatch(line)
		rules = append(rules, &rule{
			name: parts[1],
			min1: toInt(parts[2]),
			max1: toInt(parts[3]),
			min2: toInt(parts[4]),
			max2: toInt(parts[5]),
		})
	}

	myTicket := toInts(strings.Split(parts[1], "\n")[1])
	var tickets [][]int
	for i, ticket := range strings.Split(parts[2], "\n") {
		if i == 0 {
			continue
		}
		tickets = append(tickets, toInts(ticket))
	}
	tickets = append(tickets, myTicket)

	var part1Total int
	var validTickets []ticket
	for _, tt := range tickets {
		ticketValid := true
		ticket := ticket{}
		for _, f := range tt {
			var found bool
			field := ticketField{
				val: f,
			}
			for _, rule := range rules {
				if (f >= rule.min1 && f <= rule.max1) || (f >= rule.min2 && f <= rule.max2) {
					found = true
					field.matchingRules = append(field.matchingRules, rule)
				}
			}
			if !found {
				ticketValid = false
				part1Total += f
			}
			ticket = append(ticket, field)
		}
		if ticketValid {
			validTickets = append(validTickets, ticket)
		}
	}
	fmt.Println("part 1:", part1Total)

	possibleRules := make([]map[*rule]bool, len(myTicket))
	for i := 0; i < len(myTicket); i++ {
		matches := make(map[*rule]int)
		for _, ticket := range validTickets {
			for _, rule := range ticket[i].matchingRules {
				matches[rule]++
			}
		}
		matchAll := make(map[*rule]bool)
		for r, matchCount := range matches {
			if matchCount == len(validTickets) {
				matchAll[r] = true
			}
		}
		possibleRules[i] = matchAll
	}
	deleted := make([]bool, len(possibleRules))
	invalid := true
	for invalid {
		invalid = false
		for i := 0; i < len(myTicket); i++ {
			if len(possibleRules[i]) == 1 && !deleted[i] {
				var r *rule
				for k := range possibleRules[i] {
					r = k
				}
				for j := 0; j < len(possibleRules); j++ {
					if i == j {
						continue
					}
					delete(possibleRules[j], r)
				}
				deleted[i] = true
				invalid = true
				break
			}
		}
	}
	part2 := 1
	for i, value := range myTicket {
		var r *rule
		for k := range possibleRules[i] {
			r = k
		}
		if strings.HasPrefix(r.name, "departure") {
			part2 *= value
		}
	}
	fmt.Println("part 2:", part2)
}

func toInts(s string) []int {
	var ints []int
	vals := strings.Split(s, ",")
	for _, v := range vals {
		ints = append(ints, toInt(v))
	}
	return ints
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
