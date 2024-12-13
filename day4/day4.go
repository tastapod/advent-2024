package day4

import "github.com/tastapod/advent-2024/grids"

func HasWord(word string, grid grids.Grid, offset grids.Offset, pos grids.Position) bool {
	var FindRest func(rest []rune, pos grids.Position) bool

	FindRest = func(rest []rune, pos grids.Position) bool {
		if len(rest) == 0 {
			return true
		} else if grid.At(pos.Row, pos.Col) != rest[0] {
			return false
		} else {
			return FindRest(rest[1:], pos.Plus(offset))
		}
	}

	return FindRest([]rune(word), pos)
}

func CountWords(word string, source []string) (result int) {
	runes := []rune(word)
	pad := len(runes) - 1
	grid := grids.PadGrid(source, pad)

	for row := range grid.NumRows {
		for col := range grid.NumCols {
			pos := grids.Position{Row: row, Col: col}
			result += countWordsAroundPosition(word, grid, pos)
		}
	}
	return
}

func countWordsAroundPosition(word string, grid grids.Grid, pos grids.Position) (result int) {
	// check all directions
	for _, dRow := range []int{-1, 0, 1} {
		for _, dCol := range []int{-1, 0, 1} {
			if HasWord(word, grid, grids.Offset{DRow: dRow, DCol: dCol}, pos) {
				result++
			}
		}
	}
	return
}

func HasCrossMAS(grid grids.Grid, pos grids.Position) bool {
	return isMAS(grid.At(pos.Row-1, pos.Col-1), grid.At(pos.Row, pos.Col), grid.At(pos.Row+1, pos.Col+1)) &&
		isMAS(grid.At(pos.Row-1, pos.Col+1), grid.At(pos.Row, pos.Col), grid.At(pos.Row+1, pos.Col-1))
}

func isMAS(r1, r2, r3 rune) bool {
	return r2 == 'A' && (r1 == 'M' && r3 == 'S' || r1 == 'S' && r3 == 'M')
}

func CountCrossMAS(source []string) (result int) {
	grid := grids.PadGrid(source, 1)

	for row := range len(source) {
		for col := range len(source[0]) {
			if HasCrossMAS(grid, grids.Position{Row: row, Col: col}) {
				result++
			}
		}
	}
	return
}
