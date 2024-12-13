package day4_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/tastapod/advent-2024/day4"
	"github.com/tastapod/advent-2024/internal/grids"
	"strings"
	"testing"
)

func TestFindsWordForwards(t *testing.T) {
	// given
	grid := grids.PadGrid([]string{"  XMAS  "}, 0)

	// then
	assert.Equal(t, true, day4.HasWord(
		"XMAS",
		grid,
		grids.Offset{DRow: 0, DCol: +1},
		grids.Position{Row: 0, Col: 2},
	))
}

func TestFindsWordBackwards(t *testing.T) {
	// given
	grid := grids.PadGrid([]string{"  SAMX  "}, 0)

	// then
	assert.Equal(t, true, day4.HasWord(
		"XMAS",
		grid,
		grids.Offset{DRow: 0, DCol: -1},
		grids.Position{Row: 0, Col: 5},
	))
}

func TestFindsWordDown(t *testing.T) {
	// given
	grid := grids.PadGrid([]string{
		"    X  ",
		"    M  ",
		"    A  ",
		"    S  "},
		0)

	// then
	assert.Equal(t, true, day4.HasWord(
		"XMAS",
		grid,
		grids.Offset{DRow: +1, DCol: 0},
		grids.Position{Row: 0, Col: 4},
	))
}

func TestFindsWordUpRight(t *testing.T) {
	// given
	grid := grids.PadGrid([]string{
		"     S ",
		"    A  ",
		"   M   ",
		"  X    ",
	}, 1)

	// then
	assert.Equal(t, true, day4.HasWord(
		"XMAS",
		grid,
		grids.Offset{DRow: -1, DCol: +1},
		grids.Position{Row: 3, Col: 2},
	))
}

var Part1Grid = strings.Split(strings.TrimSpace(`
MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
`), "\n")

func TestCountsAllWords(t *testing.T) {
	// given
	assert.Equal(t, 18, day4.CountWords("XMAS", Part1Grid))
}

func TestFindsCrossMAS(t *testing.T) {
	// given
	grid := grids.PadGrid([]string{
		".M.S.",
		"..A..",
		".M.S."},
		0)

	// then
	assert.True(t, day4.HasCrossMAS(grid, grids.Position{Row: 1, Col: 2}))
}

func TestCountsAllCrossMASs(t *testing.T) {
	assert.Equal(t, 9, day4.CountCrossMAS(Part1Grid))
}
