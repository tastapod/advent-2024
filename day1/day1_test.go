package day1_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/tastapod/advent-2024/day1"
	"strings"
	"testing"
)

var input = strings.TrimSpace(`
3   4
4   3
2   5
1   3
3   9
3   3
`)

func TestParsesInput(t *testing.T) {
	assert := assert.New(t)

	// when
	pair := day1.NewListPair(input)

	// then
	assert.Equal([]int{1, 2, 3, 3, 3, 4}, pair.L)
	assert.Equal([]int{3, 3, 3, 4, 5, 9}, pair.R)
}

func TestSumsDeltas(t *testing.T) {
	// given
	pair := day1.NewListPair(input)

	// when
	result := pair.SumDeltas()

	// then
	assert.Equal(t, 11, result)
}

func TestBuildsRunLengthHistogram(t *testing.T) {
	assert := assert.New(t)

	// given
	pair := day1.NewListPair(input)

	// when
	result := day1.RunLengthHistogram(pair.R)

	// then
	assert.Equal(3, result[3])
	assert.Equal(1, result[4])
	assert.Equal(0, result[2])
	assert.Equal(1, result[9])
}

func TestCalculatesSimilarityScore(t *testing.T) {
	// given
	pair := day1.NewListPair(input)

	// when
	result := pair.SimilarityScore()

	// then
	assert.Equal(t, 31, result)
}
