package day7

import (
	"fmt"
	"github.com/tastapod/advent-2024/internal/parsing"
	"strings"
)

type Operator string

const (
	Times  Operator = "*"
	Plus   Operator = "+"
	Concat Operator = "||"
)

type Puzzle struct {
	Target int64
	Values []int64
}

func NewPuzzle(input string) (p Puzzle) {
	parts := strings.Split(input, ": ")
	p = Puzzle{
		Target: parsing.Int64(parts[0]),
		Values: parsing.Int64s(parts[1]),
	}
	return
}

func (p *Puzzle) Solve(operators ...Operator) (result int64) {
	var rec func(target, total int64, ops []Operator, tail []int64)

	rec = func(target, total int64, ops []Operator, tail []int64) {
		switch {
		case result > 0:
			// another branch succeeded so short-circuit
		case total > target:
			// overshot so fail
		case len(tail) == 0:
			if total == target {
				// exact result!
				result = target
			}
		default:
			// try all operators with remaining values
			for _, op := range operators {
				rec(target, Apply[op](total, tail[0]), append(ops, op), tail[1:])
			}
		}
	}

	rec(p.Target, p.Values[0], nil, p.Values[1:])

	return
}

var Apply = map[Operator]func(int64, int64) int64{
	Plus: func(l, r int64) int64 {
		return l + r
	},
	Times: func(l, r int64) int64 {
		return l * r
	},
	Concat: func(l, r int64) int64 {
		return parsing.Int64(fmt.Sprintf("%d%d", l, r))
	},
}

func sumValidEquations(input []string, operators ...Operator) (total int64) {
	//debug.Debug("Checking", len(input), "lines")
	for _, line := range input {
		puzzle := NewPuzzle(line)

		if result := puzzle.Solve(operators...); result > 0 {
			total += result
		}
	}
	return
}

func SumValidEquationsPart1(input []string) int64 {
	return sumValidEquations(input, Plus, Times)
}

func SumValidEquationsPart2(input []string) (total int64) {
	return sumValidEquations(input, Plus, Times, Concat)
}
