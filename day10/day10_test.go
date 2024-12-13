package day10_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/tastapod/advent-2024/day10"
	"github.com/tastapod/advent-2024/grids"
	"github.com/tastapod/advent-2024/internal/parsing"
	"testing"
)

var Example1 = parsing.Lines(`
0123
1234
8765
9876
`)

func TestFindsTrailsFromTrailhead(t *testing.T) {
	// given
	finder := day10.NewTrailFinder(Example1)

	// when
	trails := finder.CountTrailsFrom(grids.Position{Row: 0, Col: 0})

	// then
	assert.Equal(t, 1, trails)
}

var Example2 = parsing.Lines(`
..90..9
...1.98
...2..7
6543456
765.987
876....
987....`)

func TestFindsTrailsForExample2(t *testing.T) {
	// given
	finder := day10.NewTrailFinder(Example2)

	// when
	trails := finder.CountTrailsFrom(grids.Position{Row: 0, Col: 3})

	// then
	assert.Equal(t, 4, trails)
}

var Example3 = parsing.Lines(`
89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`)

func TestFindsRoutesForAllTrailheads(t *testing.T) {
	// given
	finder := day10.NewTrailFinder(Example3)

	// when
	trails := finder.CountTrailsFromAllTrailheads()

	// then
	assert.Equal(t, 36, trails)
}

func TestCountsRoutesForPart1(t *testing.T) {
	// given
	finder := day10.NewTrailFinder(parsing.Lines(parsing.TrimFile("input.txt")))

	// when
	trails := finder.CountTrailsFromAllTrailheads()

	// then
	assert.Equal(t, 510, trails)
}
