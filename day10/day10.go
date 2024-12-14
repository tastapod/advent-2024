package day10

import (
	"github.com/tastapod/advent-2024/internal/grids"
	"github.com/tastapod/advent-2024/internal/sets"
)

type TrailFinder struct {
	grids.Grid
}

type P = grids.Position
type D = grids.Offset

func NewTrailFinder(input []string) (tf TrailFinder) {
	tf = TrailFinder{Grid: grids.PadGrid(input, 1)}
	return
}

func (tf *TrailFinder) CountTrailsFrom(trailhead P) (result int) {
	trailends := sets.NewSet[P](trailhead)
	tops := sets.NewSet[P]()

	for !trailends.IsEmpty() {
		newTrailends := sets.NewSet[P]()
		for trailend := range trailends.All() {
			// have we reached the top?
			height := tf.IntAt(trailend.Row, trailend.Col)
			if height == 9 {
				// found a top
				tops.Add(trailend)
				continue
			}

			// check around trailend
			for _, next := range []P{
				trailend.Plus(D{DRow: 1}),
				trailend.Minus(D{DRow: 1}),
				trailend.Plus(D{DCol: 1}),
				trailend.Minus(D{DCol: 1}),
			} {
				if tf.IntAt(next.Row, next.Col) == height+1 {
					newTrailends.Add(next)
				}
			}
		}
		trailends = newTrailends
	}
	return tops.Len()
}

func (tf *TrailFinder) SumTrailsFromAllTrailheads() (total int) {
	for row := range tf.NumRows {
		for col := range tf.NumCols {
			height := tf.IntAt(row, col)
			if height == 0 { // trailhead
				total += tf.CountTrailsFrom(P{Row: row, Col: col})
			}
		}
	}
	return
}

func (tf *TrailFinder) RatingFor(trailhead P) (rating int) {
	trailends := []P{trailhead}

	for len(trailends) > 0 {
		newTrailends := make([]P, 0)
		for _, trailend := range trailends {
			// have we reached the top?
			height := tf.IntAt(trailend.Row, trailend.Col)
			if height == 9 {
				// found a top
				rating++
				continue
			}

			// check around trailend
			for _, next := range []P{
				trailend.Plus(D{DRow: 1}),
				trailend.Minus(D{DRow: 1}),
				trailend.Plus(D{DCol: 1}),
				trailend.Minus(D{DCol: 1}),
			} {
				if tf.IntAt(next.Row, next.Col) == height+1 {
					newTrailends = append(newTrailends, next)
				}
			}
		}
		trailends = newTrailends
	}
	return
}

func (tf *TrailFinder) SumRatingsForAllTrailheads() (total int) {
	for row := range tf.NumRows {
		for col := range tf.NumCols {
			if tf.IntAt(row, col) == 0 {
				total += tf.RatingFor(P{Row: row, Col: col})
			}
		}
	}
	return
}
