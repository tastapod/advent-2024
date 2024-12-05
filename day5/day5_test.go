package day5_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/tastapod/advent-2024/day5"
	"testing"
)

var Part1Input = `
47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47
`

func TestParsesInput(t *testing.T) {
	// when
	rules, updates := day5.ParseInput(Part1Input)

	// then
	assert.Equal(t, 21, len(rules))
	assert.Equal(t, 6, len(updates))
}

func TestChecksPageOrder(t *testing.T) {
	// given
	rules, updates := day5.ParseInput(Part1Input)
	checker := day5.NewRuleChecker(rules)

	isCorrect := []bool{true, true, true, false, false, false} // based on the example

	// then
	for i, update := range updates {
		assert.Equal(t, isCorrect[i], checker.IsValidUpdate(day5.ParseUpdate(update)), update)
	}
}

func TestSumsMiddleValues(t *testing.T) {
	// given
	rules, updates := day5.ParseInput(Part1Input)

	// then
	assert.Equal(t, 143, day5.SumMiddleValuesOfCorrectUpdates(rules, updates))
}

func TestSortsIncorrectUpdate(t *testing.T) {
	// given
	rules, _ := day5.ParseInput(Part1Input)
	checker := day5.NewRuleChecker(rules)

	result := checker.SortUpdate(day5.ParseUpdate("75,97,47,61,53")) // incorrect
	expected := day5.ParseUpdate("97,75,47,61,53")

	assert.Equal(t, expected, result)
}

func TestSumsCorrectedMiddleValues(t *testing.T) {
	rules, updates := day5.ParseInput(Part1Input)
	assert.Equal(t, 123, day5.SumMiddleValuesOfFixedUpdates(rules, updates))
}
