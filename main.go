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
	pair := day1.NewListPair(input.ReadDay(1))
	fmt.Printf("Day 1 part 1: %d\n", pair.SumDeltas()) // 2815556

	// part 2
	fmt.Printf("Day 1 part 2: %d\n", pair.SimilarityScore()) // 23927637
}
