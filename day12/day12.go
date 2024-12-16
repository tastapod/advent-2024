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
	Id         int
	Crop       string
	Plots      []grids.Position
	perimeter  int              // memoised length of perimeter
	numSides   int              // memoised number of sides of outer perimiter
	neighbours []grids.Position // all plots adjacent to the outer edge
}

func (r *Region) String() string {
	return fmt.Sprintf("%s: %v", r.Crop, r.Plots)
}

// Perimeter calculates the perimeter of a region.
// We start by assuming all plots have four fences, then remove pairs of
// fences for each plot touching another
func (r *Region) Perimeter() int {
	if r.perimeter > 0 {
		// memoised
		return r.perimeter
	}

	r.perimeter = 4 * len(r.Plots)
	for _, plot := range r.Plots {
		if slices.Contains(r.Plots, P(plot.Row-1, plot.Col)) {
			r.perimeter -= 2
		}
		if slices.Contains(r.Plots, P(plot.Row, plot.Col-1)) {
			r.perimeter -= 2
		}
	}
	return r.perimeter
}

func (r *Region) Area() int {
	return len(r.Plots)
}

func (r *Region) Price() int {
	return r.Perimeter() * r.Area()
}

func (r *Region) NumSides() int {
	if r.numSides == 0 {
		r.walkPerimeter()
	}
	return r.numSides
}

func (r *Region) Neighbours() []grids.Position {
	if r.neighbours == nil {
		r.walkPerimeter()
	}
	return r.neighbours
}

// walkPerimeter walks the outer perimeter of this region, counting the number of corners (turns)
// and collecting the location of all the neighbouring plots
//
// We start in the top-left corner (first plot) and 'keep our hand on the left wall'
// until we are back at the original position & direction. We store all the plots we
// visit on the way round, so we can check
//
// This means:
// - turn left if there is a plot to our left
// - otherwise move forward if there is a plot ahead
// - otherwise turn right
//
// This does not account for 'holes' where the region completely surrounds another region
func (r *Region) walkPerimeter() {
	// start in the top-left corner, facing right
	start := Vector{Pos: r.Plots[0], Dir: grids.Right}
	here := start

	for {
		switch {
		// If the plot to my left is part of this region, then move to it
		case slices.Contains(r.Plots, here.PlotToLeft()):
			here = here.TurnLeft().MoveForwards()
			r.numSides++

		// otherwise try to move forwards along the fence
		case slices.Contains(r.Plots, here.PlotAhead()):
			// capture neighbour to my left
			here = here.MoveForwards()

		// otherwise turn right
		default:
			here = here.TurnRight()
			r.numSides++
		}

		// Keep track of whatever we pass on our left
		r.neighbours = append(r.neighbours, here.PlotToLeft())

		if here == start {
			return
		}
	}
}

type PlotMap struct {
	plots         grids.Grid
	RegionsByPlot map[grids.Position]*Region
	RegionsById   map[int]*Region
}

func NewPlotMap(input []string) (pm *PlotMap) {
	pm = &PlotMap{
		plots:         grids.NewGrid(input, 1),
		RegionsByPlot: make(map[grids.Position]*Region),
		RegionsById:   make(map[int]*Region),
	}

	var nextRegionId int
	var thisRegion *Region

	joinRegion := func(row, col int, me grids.Position) (region *Region) {
		region = pm.RegionsByPlot[P(row, col)]
		region.Plots = append(region.Plots, me)
		pm.RegionsByPlot[me] = region
		return
	}

	for row := range pm.plots.NumRows {
		for col := range pm.plots.NumCols {
			me := P(row, col)
			crop := pm.plots.At(row, col)

			if pm.plots.At(row-1, col) == crop {
				// Part of region above
				above := pm.RegionsByPlot[P(row-1, col)]

				// Do we need to merge the Regions above and left?
				if pm.plots.At(row, col-1) == crop {
					// Left region has same crop as above
					if left := pm.RegionsByPlot[P(row, col-1)]; above != left {
						// We need to merge the left mini-region with the one above
						for _, l := range left.Plots {
							joinRegion(row-1, col, l)
						}
						delete(pm.RegionsById, left.Id)
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
					Id:    nextRegionId,
					Crop:  string(crop),
					Plots: []grids.Position{me},
				}
				nextRegionId++
				pm.RegionsByPlot[me] = thisRegion
				pm.RegionsById[thisRegion.Id] = thisRegion
			}
		}
	}
	return
}

func (pm *PlotMap) TotalPrice() (result int) {
	for _, region := range pm.RegionsById {
		result += region.Price()
	}
	return
}

// TotalDiscountedPrice calculates the total discounted price by
// adding all the discounted prices
func (pm *PlotMap) TotalDiscountedPrice() (result int) {
	debug.Debug("Checking", len(pm.RegionsById), "Regions")

	// build map of islands by 'host'
	islands := make(map[int][]int)

	for _, island := range pm.RegionsById {
		// check neighbours
		var neighbourIds []int

		for _, neighbour := range island.Neighbours() {
			if r, ok := pm.RegionsByPlot[neighbour]; !ok {
				// at least one neighbour is an edge so island can't be surrounded
				neighbourIds = nil
				break
			} else {
				neighbourIds = append(neighbourIds, r.Id)
			}
		}
		if len(slices.Compact(neighbourIds)) == 1 {
			// found an actual island
			hostId := neighbourIds[0]
			debug.Debug(island.Crop, "is surrounded by", hostId)
			islands[hostId] = append(islands[hostId], island.Id)
		}
	}

	// now we can calculate discounted price
	for _, region := range pm.RegionsById {
		debug.Debug(region.Crop, region.Id, "has", region.numSides, "sides and", region.Perimeter(), "fences")
		sides := region.NumSides()
		for _, islandId := range islands[region.Id] {
			islandSides := pm.RegionsById[islandId].NumSides()
			debug.Debug("Adding", islandSides, "sides for island", islandId)
			sides += islandSides
		}
		result += region.Area() * sides
	}
	return
}

// P is a convenience function to create a [grid.Position].
func P(row int, col int) grids.Position {
	return grids.Position{Row: row, Col: col}
}
