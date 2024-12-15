package day12_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/tastapod/advent-2024/day12"
	"github.com/tastapod/advent-2024/internal/grids"
	"github.com/tastapod/advent-2024/internal/parsing"
	"testing"
)

var Map1 = parsing.Lines(`
AAAA
BBCD
BBCC
EEEC`)

func TestFindsRegions(t *testing.T) {
	// given
	plotMap := day12.NewPlotMap(Map1)

	// when
	regions := plotMap.FindRegions()

	// then
	assert.Equal(t, 5, len(regions))
}

func TestCalculatesPerimiterAndPrice(t *testing.T) {
	// given
	plotMap := day12.NewPlotMap(Map1)

	// when
	regions := plotMap.FindRegions()

	// then
	expectedPerimeter := []int{10, 8, 10, 4, 8}
	expectedPrice := []int{40, 32, 40, 4, 24}

	for i := range len(expectedPerimeter) {
		assert.Equal(t, expectedPerimeter[i], regions[i].Perimeter())
		assert.Equal(t, expectedPrice[i], regions[i].Price())
	}
}

var Map2 = parsing.Lines(`
RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`)

func TestCalculatesTotalPrice(t *testing.T) {
	// given
	plotMap := day12.NewPlotMap(Map2)

	// then
	assert.Equal(t, 1930, plotMap.TotalPrice())
}

func TestCalculatesTotalForPart1(t *testing.T) {
	// given
	plotMap := day12.NewPlotMap(parsing.FileLines("input.txt"))

	// then
	assert.Equal(t, 1446042, plotMap.TotalPrice())
}

func TestCountsSides(t *testing.T) {
	// given
	region := &day12.Region{
		Plots: []grids.Position{
			{0, 0},
			{0, 1},
		},
	}

	// then
	assert.Equal(t, 4, region.NumSides())
}

var Map3 = parsing.Lines(`
AAAA
BBCD
BBCC
EEEC`)

func TestCountsSidesFromExample(t *testing.T) {
	// given
	plotMap := day12.NewPlotMap(Map3)
	region := plotMap.FindRegions()[2]

	// then
	assert.Equal(t, "C", region.Crop)
	assert.Equal(t, 8, region.NumSides())
}

var Map4 = parsing.Lines(`
EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`)

func TestCalculatesDiscountedPrice(t *testing.T) {
	assert.Equal(t, 80, day12.NewPlotMap(Map3).TotalDiscountedPrice())
	assert.Equal(t, 236, day12.NewPlotMap(Map4).TotalDiscountedPrice())
}

var Map5 = parsing.Lines(`
OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`)

func TestCalculatesDiscountedPriceForRegionWithHoles(t *testing.T) {
	assert.Equal(t, 436, day12.NewPlotMap(Map5).TotalDiscountedPrice())
}

func xTestCalculatesPart2(t *testing.T) {
	input := parsing.FileLines("input.txt")
	assert.Equal(t, 0, day12.NewPlotMap(input).TotalDiscountedPrice())
}
