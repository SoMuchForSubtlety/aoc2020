package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/SoMuchForSubtlety/aoc2020/go/pkg/input"
)

func main() {
	tape := loadTape(input.ReadLines(8))

	c := Console{tape: tape}

	c.run(loopBreakpoint(len(tape)))
	fmt.Println("part 1:", c.acc)
	c.reset()

	graph, final := buildGraph(tape)
	graph.patchInstruction(final)
	c.run()
	fmt.Println("Part 2:", c.acc)
}

func loadTape(lines []string) []*Instruction {
	var tape []*Instruction
	for _, line := range lines {
		components := strings.Fields(line)
		arg, err := strconv.Atoi(components[1])
		if err != nil {
			panic(err)
		}
		tape = append(tape, &Instruction{op: components[0], arg: arg})
	}
	return tape
}

func loopBreakpoint(tapeLength int) func(c *Console) bool {
	executed := make([]bool, tapeLength)
	return func(c *Console) bool {
		if executed[c.pointer] {
			return true // break if we executed this instruction already
		}
		executed[c.pointer] = true
		return false
	}
}

func buildGraph(tape []*Instruction) (Graph, *Node) {
	graph := Graph{
		nodes: make([]*Node, len(tape)),
		tape:  tape,
	}

	var final *Node
	for i := range tape {
		node := graph.get(i)
		next := graph.get(i + node.op.next())
		if next == nil {
			// we move out of the tape
			if i+node.op.next() == len(tape) {
				final = node
			}
			continue
		}
		next.prev = append(next.prev, node)
		node.next = next
	}

	return graph, final
}

type Node struct {
	location int
	op       *Instruction
	prev     []*Node
	next     *Node
}

func (n *Node) startingPoints() []*Node {
	if len(n.prev) == 0 {
		return []*Node{n}
	}
	var starts []*Node
	for _, parent := range n.prev {
		starts = append(starts, parent.startingPoints()...)
	}

	return starts
}

func (n *Node) folowedBy(n2 *Node, visited []bool) bool {
	if visited[n.location] {
		return false
	}
	visited[n.location] = true
	if n.next == n2 {
		return true
	} else if n.next == nil {
		return false
	}
	return n.next.folowedBy(n2, visited)
}

type Instruction struct {
	op  string
	arg int
}

func (i Instruction) next() int {
	switch i.op {
	case "acc", "nop":
		return 1
	case "jmp":
		return i.arg
	}
	panic("not reached")
}

type Graph struct {
	nodes []*Node
	tape  []*Instruction
}

func (g Graph) get(i int) *Node {
	if i >= len(g.tape) || i < 0 {
		return nil
	}
	node := g.nodes[i]
	if node == nil {
		node = &Node{op: g.tape[i], location: i}
		g.nodes[i] = node
	}
	return node
}

func (g Graph) patchInstruction(final *Node) {
	starts := final.startingPoints()
	for _, node := range g.nodes {
		var potentialNext int
		switch node.op.op {
		case "jmp":
			potentialNext = node.location + 1
		case "nop":
			potentialNext = node.location + node.op.arg
		}
		for _, start := range starts {
			if start.location == potentialNext && g.nodes[0].folowedBy(node, make([]bool, len(g.tape))) {
				if node.op.op == "jmp" {
					node.op.op = "nop"
				} else {
					node.op.op = "jmp"
				}
				return
			}
		}
	}
}

type Console struct {
	pointer int
	tape    []*Instruction
	acc     int
}

func (c *Console) run(breakpoints ...func(c *Console) bool) {
	for c.pointer < len(c.tape) && c.pointer >= 0 {
		for _, b := range breakpoints {
			if b(c) {
				return
			}
		}
		switch c.tape[c.pointer].op {
		case "acc":
			c.acc += c.tape[c.pointer].arg
			c.pointer++
		case "jmp":
			c.pointer += c.tape[c.pointer].arg
		case "nop":
			c.pointer++
		}
	}
	if !(c.pointer == len(c.tape)) {
		panic("unexpected final pointer position")
	}
}

func (c *Console) reset() {
	c.pointer = 0
	c.acc = 0
}
