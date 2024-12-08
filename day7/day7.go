package day7

import (
	"github.com/tastapod/advent-2024/internal/debug"
	"github.com/tastapod/advent-2024/internal/parsing"
	"math/big"
	"strings"
)

type Operator rune

const (
	Times Operator = '*'
	Plus  Operator = '+'
)

type Result struct {
	Target    *big.Int
	Operators []Operator
}

type Puzzle struct {
	Target *big.Int
	Values []*big.Int
}

func NewPuzzle(input string) (p Puzzle) {
	parts := strings.Split(input, ": ")
	p = Puzzle{
		Target: parsing.BigInt(parts[0]),
		Values: parsing.BigInts(parts[1]),
	}
	return
}

func (p *Puzzle) Solve() (results []Result) {
	var rec func(target, total *big.Int, ops []Operator, tail []*big.Int)

	rec = func(target, total *big.Int, ops []Operator, tail []*big.Int) {
		//debug.Debug("rec", target, total, ops, tail)
		noneLeft := len(tail) == 0
		if total.Cmp(target) == 0 && noneLeft {
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

var Apply = map[Operator]func(*big.Int, *big.Int) *big.Int{
	Plus: func(l, r *big.Int) *big.Int {
		return big.NewInt(0).Add(l, r)
	},
	Times: func(l, r *big.Int) *big.Int {
		return big.NewInt(0).Mul(l, r)
	},
}

func SumValidEquations(input []string) (total *big.Int) {
	//debug.Debug(total)
	total = big.NewInt(0)
	debug.Debug("Checking", len(input), "lines")
	for _, line := range input {
		puzzle := NewPuzzle(line)
		results := puzzle.Solve()
		if len(results) > 0 {
			total.Add(total, results[0].Target)
		}
	}
	return
}
