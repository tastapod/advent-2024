package main

import (
	"fmt"
	"github.com/tastapod/advent-2024/day1"
	"github.com/tastapod/advent-2024/day2"
	"github.com/tastapod/advent-2024/parsing"
)

func main() {
	runDay1()
	runDay2()
}

func runDay1() {
	// part 1
	pair := day1.NewListPair(parsing.ReadDay(1))
	fmt.Printf("Day 1 part 1: %d\n", pair.SumDeltas()) // 2815556

	// part 2
	fmt.Printf("Day 1 part 2: %d\n", pair.SimilarityScore()) // 23927637
}

func runDay2() {
	// part 1
	reportLines := parsing.ReadAndSplitDay(2)
	numSafe := 0
	for _, reportLine := range reportLines {
		report := day2.ParseReport(reportLine)
		if report.IsSafe() {
			numSafe++
		}
	}
	fmt.Printf("Day 2 part 1: %d safe reports\n", numSafe)

	// part 2
	numSafe = 0
	for _, reportLine := range reportLines {
		if day2.IsSafeWithTolerance(reportLine) {
			numSafe++
		}
	}
	fmt.Printf("Day 2 part 2: %d safe reports\n", numSafe)
}
