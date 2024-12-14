package day12_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/tastapod/advent-2024/day12"
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
