package grids

import "strconv"

type Grid struct {
	grid    [][]rune
	PadSize int
	NumRows int
	NumCols int
}

// NewGrid pads a source grid of a list of strings with empty lines
// above and below, and padding on either side of each line
func NewGrid(lines []string, padSize int) (result Grid) {
	numRows := len(lines)
	numCols := len(lines[0])
	totalRows := numRows + 2*padSize
	totalCols := numCols + 2*padSize

	result = Grid{
		grid:    make([][]rune, totalRows),
		PadSize: padSize,
		NumRows: numRows,
		NumCols: numCols,
	}

	// blank the grid
	for i := range result.grid {
		result.grid[i] = make([]rune, totalCols)
	}

	// copy in the values
	for i := range numRows {
		copy(result.grid[i+padSize][padSize:], []rune(lines[i]))
	}
	return
}

func (g *Grid) At(row, col int) rune {
	return (*g).grid[row+g.PadSize][col+g.PadSize]
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
	return g.grid[row+g.PadSize][g.PadSize : g.PadSize+g.NumCols]
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

type Position struct {
	Row int
	Col int
}

func (p Position) Plus(offset Offset) Position {
	return Position{Row: p.Row + offset.DRow, Col: p.Col + offset.DCol}
}

func (p Position) Minus(offset Offset) Position {
	return Position{Row: p.Row - offset.DRow, Col: p.Col - offset.DCol}
}

func (p Position) Move(dir Dir) Position {
	return p.Plus(dir.Offset())
}

type Dir string

const (
	Up    Dir = "^"
	Down  Dir = "v"
	Left  Dir = "<"
	Right Dir = ">"
)

// Right turn
func (dir Dir) Right() Dir {
	return map[Dir]Dir{
		Up:    Right,
		Right: Down,
		Down:  Left,
		Left:  Up,
	}[dir]
}

// Left turn
func (dir Dir) Left() Dir {
	return map[Dir]Dir{
		Up:    Left,
		Left:  Down,
		Down:  Right,
		Right: Up,
	}[dir]
}

func (dir Dir) Offset() Offset {
	return map[Dir]Offset{
		Up:    {DRow: -1, DCol: +0},
		Down:  {DRow: +1, DCol: +0},
		Left:  {DRow: +0, DCol: -1},
		Right: {DRow: +0, DCol: +1},
	}[dir]
}
