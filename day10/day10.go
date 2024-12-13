package day10

import (
	"github.com/tastapod/advent-2024/internal/grids"
	"github.com/tastapod/advent-2024/internal/sets"
)

type TrailFinder struct {
	grids.Grid
}

func (tf *TrailFinder) CountTrailsFrom(trailhead grids.Position) (result int) {
	trailends := sets.NewSet[grids.Position](trailhead)
	tops := sets.NewSet[grids.Position]()

	for !trailends.IsEmpty() {
		//debug.Debug(trailends)
		newTrailends := sets.NewSet[grids.Position]()
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
