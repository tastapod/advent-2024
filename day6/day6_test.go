package day6_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/tastapod/advent-2024/day6"
	"strings"
	"testing"
)

var Part1Start = strings.TrimSpace(`
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

func TestCountsSteps(t *testing.T) {
	// given
	guard := day6.NewGuard(Part1Start)

	// then
	assert.Equal(t, 41, guard.CountAllPositions())
}
