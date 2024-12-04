package day4_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/tastapod/advent-2024/day4"
	"strings"
	"testing"
)

func TestFindsWordForwards(t *testing.T) {
	// given
	grid := [][]rune{[]rune("  XMAS  ")}

	// then
	assert.Equal(t, true, day4.HasWord(
		"XMAS",
		grid,
		day4.Dir{DRow: 0, DCol: +1},
		day4.Pos{Row: 0, Col: 2},
	))
}

func TestFindsWordBackwards(t *testing.T) {
	// given
	grid := [][]rune{[]rune("  SAMX  ")}

	// then
	assert.Equal(t, true, day4.HasWord(
		"XMAS",
		grid,
		day4.Dir{DRow: 0, DCol: -1},
		day4.Pos{Row: 0, Col: 5},
	))
}

func TestFindsWordDown(t *testing.T) {
	// given
	grid := [][]rune{
		[]rune("    X  "),
		[]rune("    M  "),
		[]rune("    A  "),
		[]rune("    S  "),
	}

	// then
	assert.Equal(t, true, day4.HasWord(
		"XMAS",
		grid,
		day4.Dir{DRow: +1, DCol: 0},
		day4.Pos{Row: 0, Col: 4},
	))
}

func TestFindsWordUpRight(t *testing.T) {
	// given
	grid := [][]rune{
		[]rune("     S "),
		[]rune("    A  "),
		[]rune("   M   "),
		[]rune("  X    "),
	}

	// then
	assert.Equal(t, true, day4.HasWord(
		"XMAS",
		grid,
		day4.Dir{DRow: -1, DCol: +1},
		day4.Pos{Row: 3, Col: 2},
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

func TestBuildsPaddedGrid(t *testing.T) {
	assert := assert.New(t)

	// given
	grid := day4.PadGrid(Part1Grid, 3)

	// then
	padding := make([]rune, 16) // 3 + 10 + 3

	assert.Equal(16, len(grid))
	assert.Equal(16, len(grid[0]))

	assert.Equal(padding, grid[0])
	assert.Equal(padding, grid[1])
	assert.Equal(padding, grid[2])

	assert.Equal(padding, grid[13])
	assert.Equal(padding, grid[14])
	assert.Equal(padding, grid[15])
}

func TestCountsAllWords(t *testing.T) {
	// given
	assert.Equal(t, 18, day4.CountWords("XMAS", Part1Grid))
}
