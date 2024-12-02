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
	report := day2.ParseReport(input)

	// then
	assert.Equal(4, len(report.Pairs))
	assert.Equal(day2.Decreasing, report.Direction)
}

type testCase struct {
	input    string
	expected bool
}

func TestValidatesReport(t *testing.T) {
	assert := assert.New(t)

	testCases := []testCase{
		{"7 6 4 2 1", true},
		{"1 2 7 8 9", false},
		{"9 7 6 2 1", false},
		{"1 3 2 4 5", false},
		{"8 6 4 4 1", false},
		{"1 3 6 7 9", true},
	}

	for _, tc := range testCases {
		// given
		report := day2.ParseReport(tc.input)

		// then
		assert.Equal(tc.expected, report.IsSafe(), tc.input)
	}
}

func TestValidatesReportWithTolerance(t *testing.T) {
	assert := assert.New(t)

	testCases := []testCase{
		{"7 6 4 2 1", true},
		{"1 2 7 8 9", false},
		{"9 7 6 2 1", false},
		{"1 3 2 4 5", true},
		{"8 6 4 4 1", true},
		{"1 3 6 7 9", true},
	}

	for _, tc := range testCases {
		// then
		assert.Equal(tc.expected, day2.IsSafeWithTolerance(tc.input))
	}
}
