package day10

import (
	"github.com/tastapod/advent-2024/grids"
	"iter"
)

type TrailFinder struct {
	grids.Grid
}

type Set[T comparable] struct {
	values map[T]bool // values are keys of underlying set
}

func NewSet[T comparable](values ...T) (s Set[T]) {
	s = Set[T]{
		values: make(map[T]bool, len(values)),
	}
	for _, v := range values {
		s.values[v] = true
	}
	return
}

func (s *Set[T]) Len() int {
	return len(s.values)
}

func (s *Set[T]) IsEmpty() bool {
	return s.Len() == 0
}

// Add stores a value in the set
// Returns `true` if a value was added, `false` if it was already there
func (s *Set[T]) Add(value T) {
	s.values[value] = true
}

func (s *Set[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range s.values {
			if !yield(v) {
				return
			}
		}
	}
}

func (tf *TrailFinder) CountTrailsFrom(trailhead grids.Position) (result int) {
	trailends := NewSet[grids.Position](trailhead)
	tops := NewSet[grids.Position]()

	for !trailends.IsEmpty() {
		//debug.Debug(trailends)
		newTrailends := NewSet[grids.Position]()
		for trailend := range trailends.All() {
			//debug.Debug("Checking", trailend)
			// have we reached the top?
			height := tf.IntAt(trailend.Row, trailend.Col)
			if height == 9 {
				// found a top
				tops.Add(trailend)
				continue
			}

			// check around trailend
			for _, next := range []grids.Position{
				trailend.Plus(grids.Offset{DRow: 1}),
				trailend.Minus(grids.Offset{DRow: 1}),
				trailend.Plus(grids.Offset{DCol: 1}),
				trailend.Minus(grids.Offset{DCol: 1}),
			} {
				if tf.IntAt(next.Row, next.Col) == height+1 {
					//debug.Debug("Moving to", next, "at", height+1)
					newTrailends.Add(next)
				}
			}
		}
		trailends = newTrailends
	}
	return tops.Len()
}

func (tf *TrailFinder) CountTrailsFromAllTrailheads() (result int) {
	for row := range tf.NumRows {
		for col := range tf.NumCols {
			height := tf.IntAt(row, col)
			if height == 0 { // trailhead
				result += tf.CountTrailsFrom(grids.Position{Row: row, Col: col})
			}
		}
	}
	return
}

func NewTrailFinder(input []string) (tf TrailFinder) {
	tf = TrailFinder{
		Grid: grids.PadGrid(input, 1),
	}
	return
}
