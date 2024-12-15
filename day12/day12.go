package day12

import (
	"fmt"
	"github.com/tastapod/advent-2024/internal/debug"
	"github.com/tastapod/advent-2024/internal/grids"
	"slices"
)

type Vector struct {
	Pos grids.Position
	Dir grids.Dir
}

func (v Vector) PlotToLeft() grids.Position {
	return v.Pos.Plus(v.Dir.Left().Offset())
}

func (v Vector) PlotAhead() grids.Position {
	return v.Pos.Plus(v.Dir.Offset())
}

func (v Vector) TurnLeft() Vector {
	return Vector{Pos: v.Pos, Dir: v.Dir.Left()}
}

func (v Vector) TurnRight() Vector {
	return Vector{Pos: v.Pos, Dir: v.Dir.Right()}
}

func (v Vector) MoveForwards() Vector {
	return Vector{Pos: v.Pos.Move(v.Dir), Dir: v.Dir}
}

type Region struct {
	Crop  string
	Plots []grids.Position
}

func (r *Region) String() string {
	return fmt.Sprintf("%s: %v", r.Crop, r.Plots)
}

// Perimeter calculates the perimeter of a region.
// We start by assuming all plots have four fences, then remove pairs of
// fences for each plot touching another
func (r *Region) Perimeter() (result int) {
	result = 4 * len(r.Plots)
	for _, plot := range r.Plots {
		if slices.Contains(r.Plots, P(plot.Row-1, plot.Col)) {
			result -= 2
		}
		if slices.Contains(r.Plots, P(plot.Row, plot.Col-1)) {
			result -= 2
		}
	}
	return
}

func (r *Region) Price() int {
	return r.Perimeter() * r.Area()
}

func (r *Region) Area() int {
	return len(r.Plots)
}

func (r *Region) DiscountedPrice() int {
	return r.NumSides() * r.Area()
}

// NumSides counts the number of sides this region has, by counting the number
// of corners (turns) in a circuit.
//
// We start in the top-left corner (first plot) and 'keep our hand on the left wall'
// until we are back at the original position & direction
//
// This means:
// - turn left if there is a plot to our left
// - otherwise move forward if there is a plot ahead
// - otherwise turn right
func (r *Region) NumSides() (turns int) {
	// start in the top-left corner, facing right
	start := Vector{Pos: r.Plots[0], Dir: grids.Right}
	here := start

	for {
		switch {
		// If there is a plot to my left, move to it
		case slices.Contains(r.Plots, here.PlotToLeft()):
			here = here.TurnLeft().MoveForwards()
			turns++

		// otherwise try to move forwards along the fence
		case slices.Contains(r.Plots, here.PlotAhead()):
			here = here.MoveForwards()

		// otherwise turn right
		default:
			here = here.TurnRight()
			turns++
		}

		if here == start {
			return
		}
	}
}

type PlotMap struct {
	plots grids.Grid
}

func (pm *PlotMap) FindRegions() (result []*Region) {
	var thisRegion *Region
	regionsByPlot := make(map[grids.Position]*Region)

	joinRegion := func(row, col int, me grids.Position) (region *Region) {
		region = regionsByPlot[P(row, col)]
		region.Plots = append(region.Plots, me)
		regionsByPlot[me] = region
		return
	}

	for row := range pm.plots.NumRows {
		for col := range pm.plots.NumCols {
			me := P(row, col)
			crop := pm.plots.At(row, col)

			if pm.plots.At(row-1, col) == crop {
				// Part of region above
				above := regionsByPlot[P(row-1, col)]

				// Do we need to merge the regions above and left?
				if pm.plots.At(row, col-1) == crop {
					// Left region has same crop as above
					if left := regionsByPlot[P(row, col-1)]; above != left {
						// We need to merge the left mini-region with the one above
						for _, l := range left.Plots {
							joinRegion(row-1, col, l)
						}
						result = slices.DeleteFunc(result, func(r *Region) bool { return r == left })
						thisRegion = above
					}
				}

				// finally, join the above region
				joinRegion(row-1, col, me)
			} else if pm.plots.At(row, col-1) == crop {
				// join with left
				joinRegion(row, col-1, me)
			} else {
				// start a new region
				thisRegion = &Region{
					Crop:  string(crop),
					Plots: []grids.Position{me},
				}
				regionsByPlot[me] = thisRegion
				result = append(result, thisRegion)
			}
		}
	}
	return
}

func (pm *PlotMap) TotalPrice() (result int) {
	for _, region := range pm.FindRegions() {
		result += region.Price()
	}
	return
}

func (pm *PlotMap) TotalDiscountedPrice() (result int) {
	regions := pm.FindRegions()
	debug.Debug("Checking", len(regions), "regions")
	for _, region := range regions {
		result += region.DiscountedPrice()
	}
	return
}

func NewPlotMap(input []string) *PlotMap {
	return &PlotMap{
		grids.NewGrid(input, 1),
	}
}

// P is a convenience function to create a [grid.Position].
func P(row int, col int) grids.Position {
	return grids.Position{Row: row, Col: col}
}
