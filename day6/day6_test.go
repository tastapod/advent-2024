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

func TestMutatesMap(t *testing.T) {
	// given
	startMap := "..#.#"
	ch := make(chan string)
	day6.StartMapMutator(startMap, ch)
	var results []string

	// when
	for newMap := range ch {
		results = append(results, newMap)
	}

	// then
	assert.Equal(t, 3, len(results))
	assert.Contains(t, results, "#.#.#")
	assert.Contains(t, results, ".##.#")
	assert.Contains(t, results, "..###")
}

func TestCountsWaysToForceLoop(t *testing.T) {
	assert.Equal(t, 6, day6.CountWaysToForceLoop(Part1Start))
}
