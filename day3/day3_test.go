package day3_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/tastapod/advent-2024/day3"
	"testing"
)

var part1Input = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))\n`

func TestFindsMuls(t *testing.T) {
	// given
	input := "mul(3,4) mul(5,6)"

	// when
	result := day3.FindMuls(input)

	// then
	assert.Equal(t, 2, len(result))
}

func TestSumsMuls(t *testing.T) {
	// given
	muls := day3.FindMuls("mul(3,4) mul(5,6)")

	// when
	result := day3.SumMuls(muls)

	// then
	assert.Equal(t, 12+30, result)
}

func TestSumsTestData(t *testing.T) {
	assert.Equal(t, 161, day3.SumMuls(day3.FindMuls(part1Input)))
}

var part2Input = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

func TestKeepsTrackOfDosAndDonts(t *testing.T) {
	assert.Equal(t, 48, day3.SumEnabledMuls(part2Input))
}
