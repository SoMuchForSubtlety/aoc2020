package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/SoMuchForSubtlety/aoc2020/go/pkg/input"
)

const (
	nop Op = iota
	plus
	mul
)

type Op int

type Expr struct {
	val    int
	left   *Expr
	right  *Expr
	op     Op
	parens bool
	size   int
}

func main() {
	lines := input.ReadLines(18)
	var total1 int
	var total2 int
	for _, math := range lines {
		math = strings.ReplaceAll(strings.ReplaceAll(math, "(", "( "), ")", " )")
		total1 += ParseExpr(strings.Split(math, " "), false).Eval()
		total2 += ParseExpr(strings.Split(math, " "), true).Eval()
	}
	fmt.Println("part 1:", total1)
	fmt.Println("part 2:", total2)
}

func (o Op) String() string {
	switch o {
	case plus:
		return "+"
	case mul:
		return "*"
	default:
		return "?"
	}
}

func (e *Expr) String() string {
	if e.op == nop {
		return strconv.Itoa(e.val)
	}
	return (fmt.Sprintf("(%v%v%v)", e.left, e.op, e.right))
}

func (e *Expr) Eval() int {
	if e.op == nop {
		return e.val
	}
	left := e.left.Eval()
	right := e.right.Eval()

	switch e.op {
	case plus:
		return left + right
	case mul:
		return left * right
	default:
		panic(e.op)
	}
}

func ParseExpr(tokens []string, part2 bool) *Expr {
	var left *Expr
	var end int

	if tokens[0] == "(" {
		var c int
		for i, token := range tokens {
			if token == "(" {
				c++
			} else if token == ")" {
				c--
				if c == 0 {
					end = i
					break
				}
			}
		}
		left = ParseExpr(tokens[1:end], part2)
		if left.op != nop {
			left.parens = true
		}
		left.size = len(tokens[1:end]) + 2
	} else {
		v, err := strconv.Atoi(tokens[0])
		if err != nil {
			panic(err)
		}
		left = &Expr{val: v, size: 1}
	}
	end++
	if end == len(tokens) {
		return left
	}

	right := ParseExpr(tokens[end+1:], part2)
	op := toOp(tokens[end])

	for right.op != nop && !right.parens && (!part2 || op != mul) {
		end += right.left.size
		left = ParseExpr(tokens[0:end+1], part2)
		op = toOp(tokens[end+1])
		right = ParseExpr(tokens[end+2:], part2)
		end++
	}

	return &Expr{left: left, op: op, right: right, size: len(tokens)}
}

func toOp(o string) Op {
	switch o {
	case "+":
		return plus
	case "*":
		return mul
	default:
		panic(o)
	}
}
