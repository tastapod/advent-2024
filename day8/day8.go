package day8

import (
	"github.com/tastapod/advent-2024/grids"
	"sync"
)

type Size struct {
	NumRows int
	NumCols int
}

type Pair[T comparable] struct {
	L, R T
}

func EmitNearestAntinodes(antennae Pair[grids.Position], size Size, out chan<- grids.Position) {
	offset := grids.OffsetFrom(antennae.L, antennae.R)
	for _, antinode := range []grids.Position{antennae.L.Minus(offset), antennae.L.Plus(offset.Times(2))} {
		if isInGrid(antinode, size) {
			out <- antinode
		}
	}
}

func isInGrid(pos grids.Position, size Size) bool {
	return isInRange(pos.Row, 0, size.NumRows) && isInRange(pos.Col, 0, size.NumCols)
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

func CountNearestAntinodes(input []string) int {
	return countAntinodes(input, EmitNearestAntinodes)

}

func countAntinodes(input []string, emitAntinodes func(antennae Pair[grids.Position], size Size, out chan<- grids.Position)) int {
	antennaMap := CollectAntennae(input)
	size := Size{NumRows: len(input), NumCols: len(input[0])}
	antinodes := make(map[grids.Position]bool)

	positions := make(chan grids.Position)
	waitGroup := sync.WaitGroup{}

	for _, antennae := range antennaMap {
		combinations := Combinations(antennae)
		for _, pair := range combinations {
			waitGroup.Add(1)
			go func(out chan<- grids.Position) {
				defer waitGroup.Done()
				emitAntinodes(pair, size, positions)
			}(positions)
		}
	}

	go func(out chan<- grids.Position) {
		waitGroup.Wait()
		close(positions)
	}(positions)

	for antinode := range positions {
		antinodes[antinode] = true
	}

	return len(antinodes)
}

func CountAllAntinodes(input []string) int {
	return countAntinodes(input, EmitAllAntinodes)
}

func EmitAllAntinodes(antennae Pair[grids.Position], size Size, out chan<- grids.Position) {
	offset := grids.OffsetFrom(antennae.L, antennae.R)

	for antinode := antennae.L; isInGrid(antinode, size); antinode = antinode.Plus(offset) {
		out <- antinode
	}

	for antinode := antennae.L.Minus(offset); isInGrid(antinode, size); antinode = antinode.Minus(offset) {
		out <- antinode
	}
}
