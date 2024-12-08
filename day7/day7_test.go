package day7_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tastapod/advent-2024/day7"
	"github.com/tastapod/advent-2024/internal/parsing"
	"testing"
)

func TestFindsValidCombination(t *testing.T) {
	assert := assert.New(t)

	// given
	puzzle := day7.NewPuzzle("190: 10 19")

	// when
	result := puzzle.Solve(day7.Plus, day7.Times)

	// then
	require.Equal(t, 1, len(result))
	assert.Equal(day7.Times, result[0].Operators[0])
	assert.Equal(int64(190), result[0].Target)
}

var Part1Input = parsing.Lines(`
190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`)

func TestAddsValidEquations(t *testing.T) {
	assert.Equal(t, int64(3749), day7.SumValidEquationsPart1(Part1Input))
}

func TestAddsValidEquationsWithConcat(t *testing.T) {
	assert.Equal(t, int64(11387), day7.SumValidEquationsPart2(Part1Input))
}
