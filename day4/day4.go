package day4

type Pos struct {
	Row int
	Col int
}

func (p Pos) Move(dir Dir) Pos {
	return Pos{Row: p.Row + dir.DRow, Col: p.Col + dir.DCol}
}

type Dir struct {
	DRow int
	DCol int
}

type Grid [][]rune

func HasWord(word string, grid Grid, dir Dir, pos Pos) bool {
	var FindRest func(rest []rune, pos Pos) bool

	FindRest = func(rest []rune, pos Pos) bool {
		if len(rest) == 0 {
			return true
		} else if grid[pos.Row][pos.Col] != rest[0] {
			return false
		} else {
			return FindRest(rest[1:], pos.Move(dir))
		}
	}

	return FindRest([]rune(word), pos)
}

// PadGrid pads a source grid of a list of strings with empty lines
// above and below, and padding on either side of each line
func PadGrid(lines []string, pad int) (result Grid) {
	numRows := len(lines)
	numCols := len(lines[0])
	totalRows := numRows + 2*pad
	totalCols := numCols + 2*pad

	result = make(Grid, totalRows)

	// blank the grid
	for i := range result {
		result[i] = make([]rune, totalCols)
	}

	// copy in the values
	for i := range numRows {
		copy(result[i+pad][pad:], []rune(lines[i]))
	}
	return
}

func CountWords(word string, source []string) (result int) {
	runes := []rune(word)
	pad := len(runes) - 1
	grid := PadGrid(source, pad)

	for row := range len(source) {
		for col := range len(source[0]) {
			// check words in all directions
			for _, dRow := range []int{-1, 0, 1} {
				for _, dCol := range []int{-1, 0, 1} {
					if HasWord(word, grid, Dir{dRow, dCol}, Pos{row + pad, col + pad}) {
						result++
					}
				}
			}
		}
	}
	return
}

func HasCrossMAS(grid Grid, pos Pos) bool {
	return isMAS(grid[pos.Row-1][pos.Col-1], grid[pos.Row][pos.Col], grid[pos.Row+1][pos.Col+1]) &&
		isMAS(grid[pos.Row-1][pos.Col+1], grid[pos.Row][pos.Col], grid[pos.Row+1][pos.Col-1])
}

func isMAS(r1, r2, r3 rune) bool {
	return r2 == 'A' && (r1 == 'M' && r3 == 'S' || r1 == 'S' && r3 == 'M')
}

func CountCrossMAS(source []string) (result int) {
	pad := 1
	grid := PadGrid(source, pad)

	for row := range len(source) {
		for col := range len(source[0]) {
			if HasCrossMAS(grid, Pos{row + pad, col + pad}) {
				result++
			}
		}
	}
	return
}
