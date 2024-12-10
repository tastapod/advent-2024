package day8

import (
	"github.com/tastapod/advent-2024/grids"
)

type Size struct {
	NumRows int
	NumCols int
}

type Pair[T comparable] struct {
	L, R T
}

func CalculateAntinodes(antennae Pair[grids.Position], size Size) (result []grids.Position) {
	delta := antennae.L.Minus(antennae.R)
	for _, antinode := range []grids.Position{antennae.L.Move(delta), antennae.R.Unmove(delta)} {
		if isInRange(antinode.Row, 0, size.NumRows) && isInRange(antinode.Col, 0, size.NumCols) {
			result = append(result, antinode)
		}
	}
	return
}

func isInRange(val, low, high int) bool {
	return val >= low && val < high
}

func Combinations[T comparable](source []T) (result []Pair[T]) {
	result = make([]Pair[T], 0, len(source)^2/2)
	for i := 0; i < len(source); i++ {
		for j := i + 1; j < len(source); j++ {
			result = append(result, Pair[T]{source[i], source[j]})
		}
	}
	return
}

func CollectAntennae(input []string) (result map[rune][]grids.Position) {
	result = make(map[rune][]grids.Position)
	for row, rowLine := range input {
		for col, pos := range []rune(rowLine) {
			if pos != '.' {
				result[pos] = append(result[pos], grids.Position{Row: row, Col: col})
			}
		}
	}
	return
}

func CountAntinodes(input []string) int {
	// given
	antennae := CollectAntennae(input)
	size := Size{NumRows: len(input), NumCols: len(input[0])}
	antinodes := make(map[grids.Position]bool)

	for _, positions := range antennae {
		combinations := Combinations(positions)
		for _, pair := range combinations {
			for _, antinode := range CalculateAntinodes(pair, size) {
				if antinodes[antinode] {
				}
				antinodes[antinode] = true
			}
		}
	}

	return len(antinodes)
}
