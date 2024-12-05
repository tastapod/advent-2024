package day5

import (
	"github.com/tastapod/advent-2024/internal/parsing"
	"slices"
	"strings"
)

func ParseInput(input string) (rules, updates []string) {
	parts := parsing.PartsWithSep(input, "\n\n")
	rules = parsing.Parts(parts[0])
	updates = parsing.Parts(parts[1])
	return
}

type RuleChecker struct {
	HasRule map[string]bool // default value will be false
}

func NewRuleChecker(rules []string) (r RuleChecker) {
	r = RuleChecker{
		HasRule: map[string]bool{},
	}
	for _, rule := range rules {
		r.HasRule[rule] = true
	}
	return
}

func (checker *RuleChecker) IsValidUpdate(update []string) bool {
	for i, page := range update[1:] {
		// ensure all pages before this one should be before it
		for _, before := range update[:i+1] {
			if !checker.HasRule[rule(before, page)] {
				return false
			}
		}
	}
	return true
}

func ParseUpdate(line string) []string {
	return strings.Split(line, ",")
}

// SortUpdate uses the builtin [slices.SortFunc] by checking the rules table
// for each pair we want to compare.
func (checker *RuleChecker) SortUpdate(unsorted []string) (result []string) {
	// return -1 for correct order or +1 to swap them
	cmpPages := func(p1, p2 string) int {
		// this is dirty but super terse!
		return map[bool]int{true: -1, false: +1}[checker.HasRule[rule(p1, p2)]]
	}
	result = slices.SortedFunc(slices.Values(unsorted), cmpPages)
	return
}

func SumMiddleValuesOfCorrectUpdates(ruleLines []string, updateLines []string) (total int) {
	checker := NewRuleChecker(ruleLines)

	for _, updateLine := range updateLines {
		update := ParseUpdate(updateLine)
		if checker.IsValidUpdate(update) {
			total += parsing.Int(update[len(update)/2])
		}
	}
	return
}

func SumMiddleValuesOfFixedUpdates(ruleLines []string, updateLines []string) (total int) {
	checker := NewRuleChecker(ruleLines)

	for _, updateLine := range updateLines {
		update := ParseUpdate(updateLine)
		if !checker.IsValidUpdate(update) {
			fixed := checker.SortUpdate(update)
			total += parsing.Int(fixed[len(fixed)/2])
		}
	}
	return
}

func rule(before string, after string) string {
	return before + "|" + after
}
