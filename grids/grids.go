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

type Position struct {
	Row int
	Col int
}

func (p Position) Move(delta Delta) Position {
	return Position{Row: p.Row + delta.DRow, Col: p.Col + delta.DCol}
}

func (p Position) Unmove(delta Delta) Position {
	return Position{Row: p.Row - delta.DRow, Col: p.Col - delta.DCol}
}

func (p Position) Minus(other Position) Delta {
	return Delta{
		DRow: p.Row - other.Row,
		DCol: p.Col - other.Col,
	}
}

type Delta struct {
	DRow int
	DCol int
}

type Dir rune

const (
	Up    Dir = '^'
	Down      = 'v'
	Left      = '<'
	Right     = '>'
)

var Moves = []Delta{
	Up:    {DRow: -1, DCol: +0},
	Down:  {DRow: +1, DCol: +0},
	Left:  {DRow: +0, DCol: -1},
	Right: {DRow: +0, DCol: +1},
}
