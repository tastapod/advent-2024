package day6_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/tastapod/advent-2024/day6"
	"github.com/tastapod/advent-2024/internal/grids"
	"github.com/tastapod/advent-2024/internal/parsing"
	"testing"
)

var Part1Start = parsing.Lines(`
....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
`)

var Part1Grid = grids.NewGrid(Part1Start, 1)

func TestCountsSteps(t *testing.T) {
	// given
	tracker := day6.NewGuardTracker(Part1Grid)

	// then
	assert.Equal(t, 41, tracker.CountAllPositions())
}

func TestCountsWaysToForceLoop(t *testing.T) {
	assert.Equal(t, 6, day6.CountWaysToForceLoop(Part1Grid))
}
