package day2_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/tastapod/advent-2024/day2"
	"testing"
)

/**
Game plan
- Parse input into lines
- Parse each line into a Report containing number pairs
*/

func TestParsesReport(t *testing.T) {
	assert := assert.New(t)

	// given
	input := "7 6 4 2 1"

	// when
	report := day2.NewReport(input)

	// then
	assert.Equal(4, len(report.Pairs))
	assert.Equal(day2.Decreasing, report.Direction)
}

func TestValidatesReport(t *testing.T) {
	assert := assert.New(t)

	type testCase struct {
		input    string
		expected bool
	}

	testCases := []testCase{
		{"7 6 4 2 1", true},
		{"1 2 7 8 9", false},
		{"9 7 6 2 1", false},
		{"1 3 2 4 5", false},
		{"1 3 6 7 9", true},
	}

	for _, tc := range testCases {
		// given
		report := day2.NewReport(tc.input)

		// then
		assert.Equal(tc.expected, report.IsSafe())
	}
}
