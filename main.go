package main

import (
	"fmt"
	"github.com/tastapod/advent-2024/day1"
	"github.com/tastapod/advent-2024/day2"
	"github.com/tastapod/advent-2024/util"
)

func main() {
	runDay1()
	runDay2()
}

func runDay1() {
	// part 1
	pair := day1.NewListPair(util.ReadDay(1))
	fmt.Printf("Day 1 part 1: %d\n", pair.SumDeltas()) // 2815556

	// part 2
	fmt.Printf("Day 1 part 2: %d\n", pair.SimilarityScore()) // 23927637
}

func runDay2() {
	// part 1
	reportLines := util.ReadAndSplitDay(2)
	numSafe := 0
	for _, reportLine := range reportLines {
		report := day2.NewReport(reportLine)
		if report.IsSafe() {
			numSafe++
		}
	}
	fmt.Printf("Day 2 part 1: %d safe reports\n", numSafe)
}
