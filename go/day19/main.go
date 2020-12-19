package main

import (
	"fmt"
	"strings"

	"github.com/SoMuchForSubtlety/aoc2020/go/pkg/input"
)

type Rule struct {
	sequences [][]*Rule
	content   string
}

type Rules map[string]*Rule

func main() {
	content := input.Read(19)
	rulesAndData := strings.Split(content, "\n\n")

	rules := ParseRules(rulesAndData[0])
	var totalMatches int
	for _, data := range strings.Split(rulesAndData[1], "\n") {
		if rules.matches(data) {
			totalMatches++
		}
	}
	fmt.Println("part 1:", totalMatches)

	rulesAndData[0] = strings.ReplaceAll(rulesAndData[0], "8: 42", "8: 42 | 42 8")
	rulesAndData[0] = strings.ReplaceAll(rulesAndData[0], "11: 42 31", "11: 42 31 | 42 11 31")
	rules = ParseRules(rulesAndData[0])
	totalMatches = 0
	for _, data := range strings.Split(rulesAndData[1], "\n") {
		if rules.matches(data) {
			totalMatches++
		}
	}
	fmt.Println("part 2:", totalMatches)
}

// returns an array of prefix lenghs of the input that can be matched by this rule
func (r *Rule) matches(input string) []int {
	if r.content != "" {
		if strings.HasPrefix(input, r.content) {
			return []int{len(r.content)}
		}
		return nil
	}
	var matches []int
	for _, sequence := range r.sequences {
		subMatches := []int{0}
		for _, rule := range sequence {
			var subMatchesN []int
			for _, match := range subMatches {
				if match == len(input) {
					continue
				}
				for _, matchN := range rule.matches(input[match:]) {
					subMatchesN = append(subMatchesN, match+matchN)
				}
			}
			subMatches = subMatchesN
		}
		matches = append(matches, subMatches...)
	}
	return matches
}

func (r Rules) matches(input string) bool {
	for _, fraction := range r.get("0").matches(input) {
		if len(input) == fraction {
			return true
		}
	}
	return false
}

func (r Rules) get(key string) *Rule {
	rule, ok := r[key]
	if !ok {
		rule = &Rule{}
		r[key] = rule
	}
	return rule
}

func ParseRules(content string) Rules {
	rules := make(Rules)
	for _, rule := range strings.Split(content, "\n") {
		parts := strings.Split(rule, ": ")
		rule := rules.get(parts[0])
		if strings.HasPrefix(parts[1], `"`) {
			rule.content = strings.Trim(parts[1], `"`)
		} else {
			parts = strings.Split(parts[1], " | ")
			for _, part := range parts {
				var sequence []*Rule
				for _, k := range strings.Split(part, " ") {
					sequence = append(sequence, rules.get(k))
				}
				rule.sequences = append(rule.sequences, sequence)
			}
		}
	}
	return rules
}
