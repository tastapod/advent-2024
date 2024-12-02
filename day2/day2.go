package day2

import (
	"github.com/tastapod/advent-2024/util"
	"strings"
)

type Direction int

const (
	Decreasing Direction = -1
	Increasing Direction = +1
)

type Pair struct {
	L, R int
}

func (p Pair) Direction() Direction {
	if p.L < p.R {
		return Increasing
	} else {
		return Decreasing
	}
}

type Report struct {
	Pairs     []Pair
	Direction Direction
}

func (r *Report) IsSafe() bool {
	for _, pair := range r.Pairs {
		if pair.Direction() != r.Direction {
			return false
		}

		if abs := util.AbsInt(pair.L - pair.R); abs < 1 || abs > 3 {
			return false
		}
	}
	return true
}

func NewReport(input string) (result Report) {
	values := strings.Fields(input)
	ls := values[:len(values)-1]
	rs := values[1:]

	pairs := make([]Pair, len(ls))

	for i := 0; i < len(ls); i++ {
		pairs[i] = Pair{
			L: util.ToInt(ls[i]),
			R: util.ToInt(rs[i]),
		}
	}

	return Report{
		Pairs:     pairs,
		Direction: pairs[0].Direction(),
	}
}
