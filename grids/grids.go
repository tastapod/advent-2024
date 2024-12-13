package grids

import "strconv"

type Grid struct {
	Grid    [][]rune
	PadSize int
	NumRows int
	NumCols int
}

// PadGrid pads a source grid of a list of strings with empty lines
// above and below, and padding on either side of each line
func PadGrid(lines []string, padSize int) (result Grid) {
	numRows := len(lines)
	numCols := len(lines[0])
	totalRows := numRows + 2*padSize
	totalCols := numCols + 2*padSize

	result = Grid{
		Grid:    make([][]rune, totalRows),
		PadSize: padSize,
		NumRows: numRows,
		NumCols: numCols,
	}

	// blank the grid
	for i := range result.Grid {
		result.Grid[i] = make([]rune, totalCols)
	}

	// copy in the values
	for i := range numRows {
		copy(result.Grid[i+padSize][padSize:], []rune(lines[i]))
	}
	return
}

func (g *Grid) At(row, col int) rune {
	return (*g).Grid[row+g.PadSize][col+g.PadSize]
}

// IntAt returns the int value of the current cell, or -1 for invalid cells, e.g.
// non-numeric or padding cells
// We avoid the idiomatic (result, ok) semantics to make it easier to use in expression
func (g *Grid) IntAt(row, col int) int {
	if i, err := strconv.ParseInt(string(g.At(row, col)), 10, 0); err != nil {
		return -1
	} else {
		return int(i)
	}
}

// Row returns the row sans padding
func (g *Grid) Row(row int) []rune {
	return g.Grid[row+g.PadSize][g.PadSize : g.PadSize+g.NumCols]
}

type Position struct {
	Row int
	Col int
}

type Offset struct {
	DRow int
	DCol int
}

func (d Offset) Times(n int) Offset {
	return Offset{
		DRow: n * d.DRow,
		DCol: n * d.DCol,
	}
}

func OffsetFrom(p1, p2 Position) Offset {
	return Offset{
		DRow: p2.Row - p1.Row,
		DCol: p2.Col - p1.Col,
	}
}

func (p Position) Plus(offset Offset) Position {
	return Position{Row: p.Row + offset.DRow, Col: p.Col + offset.DCol}
}

func (p Position) Minus(offset Offset) Position {
	return Position{Row: p.Row - offset.DRow, Col: p.Col - offset.DCol}
}

type Dir rune

const (
	Up    Dir = '^'
	Down      = 'v'
	Left      = '<'
	Right     = '>'
)

var Moves = []Offset{
	Up:    {DRow: -1, DCol: +0},
	Down:  {DRow: +1, DCol: +0},
	Left:  {DRow: +0, DCol: -1},
	Right: {DRow: +0, DCol: +1},
}
