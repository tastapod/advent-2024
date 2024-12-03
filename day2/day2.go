package day2

import (
	"github.com/tastapod/advent-2024/internal/ints"
	"github.com/tastapod/advent-2024/internal/parsing"
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

type ReportBuilder struct {
	Values []int
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

		if abs := ints.AbsInt(pair.L - pair.R); abs < 1 || abs > 3 {
			return false
		}
	}
	return true
}

func ParseReport(reportLine string) Report {
	return NewReport(parsing.ParseInts(reportLine))
}

func IsSafeWithTolerance(reportLine string) bool {
	data := parsing.ParseInts(reportLine)

	report := NewReport(data)
	if report.IsSafe() {
		return true
	}

	// remove one value at a time
	for i := 0; i < len(data); i++ {
		sublist := append(append([]int{}, data[0:i]...), data[i+1:]...)
		report = NewReport(sublist)
		if report.IsSafe() {
			return true
		}
	}

	// nothing worked
	return false
}

func NewReport(values []int) Report {
	ls := values[:len(values)-1]
	rs := values[1:]

	pairs := make([]Pair, 0, len(ls))

	for i := 0; i < len(ls); i++ {
		pairs = append(pairs, Pair{
			L: ls[i],
			R: rs[i],
		})
	}

	return Report{
		Pairs:     pairs,
		Direction: pairs[0].Direction(),
	}
}
