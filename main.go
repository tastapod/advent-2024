package main

import (
	"fmt"
	"github.com/tastapod/advent-2024/day1"
	"github.com/tastapod/advent-2024/input"
)

func main() {
	runDay1()
}

func runDay1() {
	// part 1
	l, r := day1.ParseInput(input.ReadDay(1))
	fmt.Printf("Day 1 part 1: %d\n", day1.SumDeltas(l, r)) // 2815556

	// part 2
	fmt.Printf("Day 1 part 2: %d\n", day1.SimilarityScore(l, r))
}
