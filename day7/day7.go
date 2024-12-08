package day7

import (
	"github.com/tastapod/advent-2024/internal/debug"
	"github.com/tastapod/advent-2024/internal/parsing"
	"strings"
)

type Operator rune

const (
	Times Operator = '*'
	Plus  Operator = '+'
)

type Result struct {
	Target    int64
	Operators []Operator
}

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

func (p *Puzzle) Solve() (results []Result) {
	var rec func(target, total int64, ops []Operator, tail []int64)

	rec = func(target, total int64, ops []Operator, tail []int64) {
		//debug.Debug("rec", target, total, ops, tail)
		noneLeft := len(tail) == 0
		if total == target && noneLeft {
			// exact result!
			//debug.Debug("Found exact result", total)
			results = append(results, Result{Target: target, Operators: ops})
		} else if noneLeft {
			// ran out of values
			//debug.Debug("Ran out of values:", target, total, tail)
			return
		} else {
			// recurse for both operators
			for _, op := range []Operator{Plus, Times} {
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
}

func SumValidEquations(input []string) (total int64) {
	debug.Debug("Checking", len(input), "lines")
	for _, line := range input {
		puzzle := NewPuzzle(line)
		results := puzzle.Solve()
		if len(results) > 0 {
			total += results[0].Target
		}
	}
	return
}