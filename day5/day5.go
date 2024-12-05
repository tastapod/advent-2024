package day5

import (
	"github.com/tastapod/advent-2024/internal/parsing"
	"slices"
)

type Rule struct {
	Before, After int
}

func NewRule(input string) (r Rule) {
	parts := parsing.PartsWithSep(input, "|")
	return Rule{
		Before: parsing.Int(parts[0]),
		After:  parsing.Int(parts[1]),
	}
}

type Update []int

func NewUpdate(input string) Update {
	return parsing.IntsWithSep(input, ",")
}

func ParseInput(input string) (rules []Rule, updates []Update) {
	var parts = parsing.PartsWithSep(input, "\n\n")
	var rulesInput = parsing.Parts(parts[0])
	var updatesInput = parsing.Parts(parts[1])

	rules = make([]Rule, len(rulesInput))
	for i, r := range rulesInput {
		rules[i] = NewRule(r)
	}

	updates = make([]Update, len(updatesInput))
	for i, p := range updatesInput {
		updates[i] = NewUpdate(p)
	}
	return
}

type Set[T comparable] map[T]struct{}

type RuleChecker struct {
	Rules Set[Rule]
}

func NewRuleChecker(rules []Rule) (r RuleChecker) {
	r = RuleChecker{
		Rules: Set[Rule]{},
	}
	for _, rule := range rules {
		r.Rules[rule] = struct{}{} // just an existence checker
	}
	return
}

func (checker *RuleChecker) IsValidUpdate(update Update) bool {
	for i, page := range update[1:] {
		// ensure all pages before this one should be before it
		for _, before := range update[:i+1] {
			if !checker.IsCorrectOrder(before, page) {
				return false
			}
		}
	}
	return true
}

// IsCorrectOrder checks whether two pages have a mapping between them.
// We do not bother building a graph because the input has an entry for every pair combo.
func (checker *RuleChecker) IsCorrectOrder(before, after int) bool {
	_, ok := checker.Rules[Rule{before, after}]
	return ok
}

// SortUpdate uses the builtin [slices.SortFunc] by checking the rules table
// for each pair we want to compare.
func (checker *RuleChecker) SortUpdate(unsorted Update) (result Update) {
	// return -1 for correct order or +1 to swap them
	cmpPages := func(p1, p2 int) int {
		// this is dirty but super terse!
		return map[bool]int{true: -1, false: +1}[checker.IsCorrectOrder(p1, p2)]
	}
	result = slices.SortedFunc(slices.Values(unsorted), cmpPages)
	return
}

func SumMiddleValuesOfCorrectUpdates(rules []Rule, updates []Update) (total int) {
	checker := NewRuleChecker(rules)

	for _, update := range updates {
		if checker.IsValidUpdate(update) {
			total += update[len(update)/2]
		}
	}
	return
}

func SumMiddleValuesOfFixedUpdates(rules []Rule, updates []Update) (total int) {
	checker := NewRuleChecker(rules)

	for _, update := range updates {
		if !checker.IsValidUpdate(update) {
			fixed := checker.SortUpdate(update)
			total += fixed[len(fixed)/2]
		}
	}
	return
}
