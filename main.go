package main

import (
	"fmt"
	"github.com/tastapod/advent-2024/day1"
	"github.com/tastapod/advent-2024/day2"
	"github.com/tastapod/advent-2024/day3"
	"github.com/tastapod/advent-2024/day4"
	"github.com/tastapod/advent-2024/internal/parsing"
)

func main() {
	runDay1()
	runDay2()
	runDay3()
	runDay4()
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

func runDay3() {
	input := parsing.ReadDay(3)

	// part 1
	part1 := day3.SumMuls(day3.FindMuls(input))
	fmt.Printf("Day 3 part 1: total = %d\n", part1) // 187825547

	// part 2
	part2 := day3.SumEnabledMuls(input)
	fmt.Printf("Day 3 part 2: total = %d\n", part2) // 85508223
}

func runDay4() {
	input := parsing.ReadAndSplitDay(4)

	// part 1
	part1 := day4.CountWords("XMAS", input)
	fmt.Printf("Day 4 part 1: found %d words\n", part1) // 2578

	// part 2
	part2 := day4.CountCrossMAS(input)
	fmt.Printf("Day 4 part 2: found %d X-MAS", part2) // 1972
}
