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
	l, r := day1.ParseInput(input)

	// then
	assert.Equal([]int{3, 4, 2, 1, 3, 3}, l)
	assert.Equal([]int{4, 3, 5, 3, 9, 3}, r)
}

func TestSumsDeltas(t *testing.T) {
	// given
	l, r := day1.ParseInput(input)

	// when
	result := day1.SumDeltas(l, r)

	// then
	assert.Equal(t, 11, result)
}

func TestBuildsHistogram(t *testing.T) {
	assert := assert.New(t)

	// given
	_, r := day1.ParseInput(input)

	// when
	counts := day1.BuildHistogram(r)

	// then
	assert.Equal(3, counts[3])
	assert.Equal(1, counts[4])
	assert.Equal(0, counts[2])
	assert.Equal(1, counts[9])
}

func TestCalculatesSimilarityScore(t *testing.T) {
	// given
	l, r := day1.ParseInput(input)

	// when
	result := day1.SimilarityScore(l, r)

	// then
	assert.Equal(t, 31, result)

}
