package day5

import (
	"github.com/tastapod/advent-2024/internal/parsing"
	"slices"
)

type Update []string

func NewUpdate(input string) Update {
	return parsing.PartsWithSep(input, ",")
}

func ParseInput(input string) (rules []string, updates []Update) {
	var parts = parsing.PartsWithSep(input, "\n\n")
	rules = parsing.Parts(parts[0])

	var updatesInput = parsing.Parts(parts[1])
	updates = make([]Update, len(updatesInput))
	for i, p := range updatesInput {
		updates[i] = NewUpdate(p)
	}
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

func (checker *RuleChecker) IsValidUpdate(update Update) bool {
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

// SortUpdate uses the builtin [slices.SortFunc] by checking the rules table
// for each pair we want to compare.
func (checker *RuleChecker) SortUpdate(unsorted Update) (result Update) {
	// return -1 for correct order or +1 to swap them
	cmpPages := func(p1, p2 string) int {
		// this is dirty but super terse!
		return map[bool]int{true: -1, false: +1}[checker.HasRule[rule(p1, p2)]]
	}
	result = slices.SortedFunc(slices.Values(unsorted), cmpPages)
	return
}

func SumMiddleValuesOfCorrectUpdates(rules []string, updates []Update) (total int) {
	checker := NewRuleChecker(rules)

	for _, update := range updates {
		if checker.IsValidUpdate(update) {
			total += parsing.Int(update[len(update)/2])
		}
	}
	return
}

func SumMiddleValuesOfFixedUpdates(rules []string, updates []Update) (total int) {
	checker := NewRuleChecker(rules)

	for _, update := range updates {
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
