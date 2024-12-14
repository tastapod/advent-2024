package day12

import (
	"fmt"
	"github.com/tastapod/advent-2024/internal/grids"
	"slices"
)

type PlotMap struct {
	plots grids.Grid
}

type Region struct {
	crop  string
	Plots []grids.Position
}

func (r *Region) String() string {
	return fmt.Sprintf("%s: %v", r.crop, r.Plots)
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
	return r.Perimeter() * len(r.Plots)
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
				// join with above
				above := regionsByPlot[P(row-1, col)]

				// Do we need to merge the regions above and left?
				if pm.plots.At(row, col-1) == crop {
					// left has same crop as above
					if left := regionsByPlot[P(row, col-1)]; above != left {
						// different regions, so merge them
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
					crop:  string(crop),
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

func NewPlotMap(input []string) *PlotMap {
	return &PlotMap{
		grids.NewGrid(input, 1),
	}
}

// P is a convenience function to create a [grid.Position].
func P(row int, col int) grids.Position {
	return grids.Position{Row: row, Col: col}
}
