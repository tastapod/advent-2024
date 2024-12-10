package grids

type Grid [][]rune

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

func (g *Grid) At(p Position) rune {
	return (*g)[p.Row][p.Col]
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
