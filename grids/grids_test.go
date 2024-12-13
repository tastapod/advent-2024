package grids_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/tastapod/advent-2024/grids"
	"strings"
	"testing"
)

// borrowed from Day 4
var input = strings.Split(strings.TrimSpace(`
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
	grid := grids.PadGrid(input, 3)

	// then
	padding := make([]rune, 16) // 3 + 10 + 3

	assert.Equal(16, len(grid.Grid))
	assert.Equal(16, len(grid.Grid[0]))

	assert.Equal(padding, grid.Grid[0])
	assert.Equal(padding, grid.Grid[1])
	assert.Equal(padding, grid.Grid[2])

	assert.Equal(padding, grid.Grid[13])
	assert.Equal(padding, grid.Grid[14])
	assert.Equal(padding, grid.Grid[15])
}
