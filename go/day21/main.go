package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/SoMuchForSubtlety/aoc2020/go/pkg/input"
)

type Mapping map[string][]string

func main() {
	lines := input.ReadLines(21)

	mappings := make(Mapping)
	for _, line := range lines {
		parts := strings.Split(line, " (contains ")
		ingredients := strings.Split(parts[0], " ")
		allergenes := strings.Split(parts[1][:len(parts[1])-1], ", ")
		for _, allergene := range allergenes {
			mappings.AddOptions(allergene, ingredients)
		}
	}
	mappings.EliminateDupes()
	reverseMappings := make(map[string]string)
	var allergenesSorted []string
	for k, v := range mappings {
		reverseMappings[v[0]] = k
		allergenesSorted = append(allergenesSorted, k)
	}
	sort.Slice(allergenesSorted, func(i, j int) bool { return allergenesSorted[i] < allergenesSorted[j] })
	var total int
	for _, line := range lines {
		for _, ingredient := range strings.Split(strings.Split(line, " (contains ")[0], " ") {
			if _, ok := reverseMappings[ingredient]; !ok {
				total++
			}
		}
	}
	fmt.Println("part 1:", total)

	var out string
	for _, allergene := range allergenesSorted {
		out += mappings[allergene][0] + ","
	}
	fmt.Println("part 2:", out[:len(out)-1])
}

func (m Mapping) AddOptions(allergene string, ingredients []string) {
	existringIngredients, ok := m[allergene]
	if !ok {
		m[allergene] = ingredients
		return
	}

	var intersection []string
	for _, new := range ingredients {
		for _, existing := range existringIngredients {
			if new == existing {
				intersection = append(intersection, new)
			}
		}
	}
	m[allergene] = intersection
}

func (m Mapping) EliminateDupes() {
	var done bool
	doneAllergenes := make(map[string]bool)
	for !done {
		done = true
		for allergene, ingredients := range m {
			if len(ingredients) == 1 && !doneAllergenes[allergene] {
				doneAllergenes[allergene] = true
				done = false
				for allergene2, ingredients2 := range m {
					if allergene2 == allergene {
						continue
					}
					for i, ingredient2 := range ingredients2 {
						if ingredient2 == ingredients[0] {
							m[allergene2] = remove(ingredients2, i)
							break
						}
					}
				}
				break
			}
		}
	}
}

func remove(s []string, i int) []string {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}
