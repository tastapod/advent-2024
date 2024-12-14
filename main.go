package main

import (
	"fmt"
	"github.com/tastapod/advent-2024/day1"
	"github.com/tastapod/advent-2024/day10"
	"github.com/tastapod/advent-2024/day2"
	"github.com/tastapod/advent-2024/day3"
	"github.com/tastapod/advent-2024/day4"
	"github.com/tastapod/advent-2024/day5"
	"github.com/tastapod/advent-2024/day6"
	"github.com/tastapod/advent-2024/day7"
	"github.com/tastapod/advent-2024/day8"
	"github.com/tastapod/advent-2024/day9"
	"github.com/tastapod/advent-2024/internal/grids"
	"github.com/tastapod/advent-2024/internal/parsing"
	"time"
)

func main() {
	stopwatch(runDay1)
	stopwatch(runDay2)
	stopwatch(runDay3)
	stopwatch(runDay4)
	stopwatch(runDay5)
	stopwatch(runDay6)
	stopwatch(runDay7)
	stopwatch(runDay8)
	stopwatch(runDay9)
	stopwatch(runDay10)
}

func stopwatch(fn func()) {
	start := time.Now()
	fn()
	fmt.Printf("%d µs\n", time.Since(start).Microseconds())
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
	fmt.Printf("Day 2 part 1: %d safe reports\n", numSafe) // 598

	// part 2
	numSafe = 0
	for _, reportLine := range reportLines {
		if day2.IsSafeWithTolerance(reportLine) {
			numSafe++
		}
	}
	fmt.Printf("Day 2 part 2: %d safe reports\n", numSafe) // 634
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
	fmt.Printf("Day 4 part 2: found %d X-MAS\n", part2) // 1972
}

func runDay5() {
	input := parsing.ReadDay(5)

	// part 1
	rules, updates := day5.ParseInput(input)
	part1 := day5.SumMiddleValuesOfCorrectUpdates(rules, updates)
	fmt.Printf("Day 5 part 1: total = %d\n", part1) // 6267

	part2 := day5.SumMiddleValuesOfFixedUpdates(rules, updates)
	fmt.Printf("Day 5 part 2: total = %d\n", part2) // 5184
}

func runDay6() {
	input := grids.PadGrid(parsing.ReadAndSplitDay(6), 1)

	// part 1
	guard := day6.NewGuardTracker(input)
	part1 := guard.CountAllPositions()
	fmt.Printf("Day 6 part 1: total = %d\n", part1) // 4826

	// part 2
	part2 := day6.CountWaysToForceLoop(input)
	fmt.Printf("Day 6 part 2: total = %d\n", part2) // 1721
}

func runDay7() {
	input := parsing.ReadAndSplitDay(7)

	// part 1
	part1 := day7.SumValidEquationsPart1(input)
	fmt.Printf("Day 7 part 1: total = %d\n", part1) // 10741443549536

	// part 2
	part2 := day7.SumValidEquationsPart2(input)
	fmt.Printf("Day 7 part 2: total = %d\n", part2) // 500335179214836
}

func runDay8() {
	input := parsing.ReadAndSplitDay(8)

	// part 1
	part1 := day8.CountNearestAntinodes(input)
	fmt.Printf("Day 8 part 1: total = %d\n", part1) // 327

	// part 1
	part2 := day8.CountAllAntinodes(input)
	fmt.Printf("Day 8 part 2: total = %d\n", part2) // 327
}

func runDay9() {
	input := parsing.ReadDay(9)
	var dm day9.DiskMap

	// part 1
	dm = day9.NewDiskMap(input)
	part1 := dm.DefragWholeDisk().Checksum()
	fmt.Printf("Day 9 part 1: checksum = %d\n", part1) // 6307275788409

	// part 1
	dm = day9.NewDiskMap(input)
	part2 := dm.DefragWholeDiskWithWholeFiles().Checksum()
	fmt.Printf("Day 9 part 2: checksum = %d\n", part2) // 6327174563252
}

func runDay10() {
	input := parsing.ReadAndSplitDay(10)
	trailFinder := day10.NewTrailFinder(input)

	// part 1
	part1 := trailFinder.SumTrailsFromAllTrailheads()
	fmt.Printf("Day 10 part 1: %d\n", part1)

	// part 2
	part2 := trailFinder.SumRatingsForAllTrailheads()
	fmt.Printf("Day 10 part 2: %d\n", part2)
}
