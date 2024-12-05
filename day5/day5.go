package day5

import (
	"github.com/tastapod/advent-2024/internal/parsing"
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

func (checker *RuleChecker) IsCorrectOrder(before, after int) bool {
	_, ok := checker.Rules[Rule{before, after}]
	return ok
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
