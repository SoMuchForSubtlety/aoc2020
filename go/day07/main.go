package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/SoMuchForSubtlety/aoc2020/go/pkg/input"
)

type baggage map[string]*bag

func main() {
	// colour -> bag
	var bags baggage = make(map[string]*bag)

	for _, rule := range input.ReadLines(7) {
		bags.parseRule(rule)
	}

	var count int
	for _, b := range bags {
		if b.contains("shiny gold") {
			count++
		}
	}

	fmt.Println("part 1:", count)
	fmt.Println("part 2:", bags["shiny gold"].childCount())
}

func (b baggage) parseRule(rule string) {
	ruleParts := strings.Split(rule, " bags contain ") // -> ["light red", "1 bright white bag, 2 muted yellow bags."]
	container := b.bag(ruleParts[0])

	if ruleParts[1] == "no other bags." {
		return
	}

	ruleParts[1] = strings.ReplaceAll(ruleParts[1], ".", "")     // -> 1 bright white bag, 2 muted yellow bags
	ruleParts[1] = strings.ReplaceAll(ruleParts[1], " bags", "") // -> 1 bright white bag, 2 muted yellow
	ruleParts[1] = strings.ReplaceAll(ruleParts[1], " bag", "")  // -> 1 bright white, 2 muted yellow

	for _, childString := range strings.Split(ruleParts[1], ", ") {

		childParts := strings.SplitN(childString, " ", 2) // -> ["1", "bright white"]
		amount, err := strconv.Atoi(childParts[0])
		if err != nil {
			panic(err)
		}
		container.addChild(quantity{bag: b.bag(childParts[1]), amount: amount})
	}
}

func (b baggage) bag(color string) *bag {
	container, ok := b[color]
	if !ok {
		container = &bag{colour: color}
		b[color] = container
	}

	return container
}

type bag struct {
	colour   string
	children []quantity
}

func (b *bag) addChild(b2 quantity) {
	b.children = append(b.children, b2)
}

func (b *bag) contains(colour string) bool {
	for _, c := range b.children {
		if c.bag.colour == colour || c.bag.contains(colour) {
			return true
		}
	}
	return false
}

func (b *bag) childCount() int {
	var total int
	for _, c := range b.children {
		total += (c.bag.childCount() + 1) * c.amount // (children of c + c itself) * amount of c
	}
	return total
}

type quantity struct {
	amount int
	bag    *bag
}
