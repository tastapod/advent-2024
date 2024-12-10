package day8_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/tastapod/advent-2024/day8"
	"github.com/tastapod/advent-2024/grids"
	"github.com/tastapod/advent-2024/internal/parsing"
	"testing"
)

func TestCalculatesAntinodes(t *testing.T) {
	// given
	var antennae = day8.Pair[grids.Position]{
		L: grids.Position{Row: 3, Col: 4},
		R: grids.Position{Row: 5, Col: 5},
	}

	// when
	antinodes := antinodes(antennae, day8.Size{NumRows: 10, NumCols: 10})

	// then
	assert.Equal(t, 2, len(antinodes))
	assert.Contains(t, antinodes, grids.Position{Row: 1, Col: 3})
	assert.Contains(t, antinodes, grids.Position{Row: 7, Col: 6})
}

func antinodes(antennae day8.Pair[grids.Position], size day8.Size) []grids.Position {
	ch := make(chan grids.Position, 2)
	day8.EmitNearestAntinodes(antennae, size, ch)
	close(ch)
	var antinodes []grids.Position

	for antinode := range ch {
		antinodes = append(antinodes, antinode)
	}
	return antinodes
}

func TestIgnoresAntinodesOutsideOfGrid(t *testing.T) {
	// given
	var antennae = day8.Pair[grids.Position]{
		L: grids.Position{Row: 3, Col: 4},
		R: grids.Position{Row: 4, Col: 8},
	}

	// when
	antinodes := antinodes(antennae, day8.Size{NumRows: 10, NumCols: 10})

	// then
	assert.Equal(t, 1, len(antinodes))
	assert.Contains(t, antinodes, grids.Position{Row: 2, Col: 0})
}

func TestGeneratesCombinationsOfPairs(t *testing.T) {
	// given
	source := []rune{'a', 'b', 'c'}

	// when
	combinations := day8.Combinations(source)

	// then
	assert.Contains(t, combinations, day8.Pair[rune]{L: 'a', R: 'b'})
	assert.Contains(t, combinations, day8.Pair[rune]{L: 'a', R: 'c'})
	assert.Contains(t, combinations, day8.Pair[rune]{L: 'b', R: 'c'})
}

var Input = parsing.Lines(`
............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............
`)

func TestCollectsDifferentTypesOfAntenna(t *testing.T) {
	// when
	antennae := day8.CollectAntennae(Input)

	// then
	assert.Equal(t, 4, len(antennae['0']))
	assert.Contains(t, antennae['0'], grids.Position{Row: 1, Col: 8})

	assert.Equal(t, 3, len(antennae['A']))
	assert.Contains(t, antennae['A'], grids.Position{Row: 8, Col: 8})
	assert.Contains(t, antennae['A'], grids.Position{Row: 9, Col: 9})
}

func TestCountsPositionsOfNearestAntinodes(t *testing.T) {
	assert.Equal(t, 14, day8.CountNearestAntinodes(Input))

}

func TestCountsPositionsOfAllAntinodes(t *testing.T) {
	assert.Equal(t, 34, day8.CountAllAntinodes(Input))

}