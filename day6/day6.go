package day6

import (
	"github.com/tastapod/advent-2024/grids"
	"github.com/tastapod/advent-2024/internal/debug"
	"github.com/tastapod/advent-2024/internal/parsing"
	"strings"
)

type Step struct {
	grids.Pos
	grids.Dir
}

type Guard struct {
	grids.Grid
	Here    Step
	History map[Step]bool
}

func NewGuard(mapInput string) (g *Guard) {
	g = &Guard{
		Grid: grids.PadGrid(parsing.Parts(mapInput), 1),
	}

	for row, rowChars := range g.Grid {
		if col := strings.IndexAny(string(rowChars), "<^>v"); col != -1 {
			g.Here = Step{
				grids.Pos{Row: row, Col: col},
				grids.Dir(rowChars[col]),
			}
			g.History = map[Step]bool{g.Here: true}

			// don't forget to reset our starting point!
			rowChars[col] = '.'
			debug.Debug("Starting at", g.Here)
			break
		}
	}

	return
}

func (g *Guard) CountAllPositions() int {
	for g.Move() {
		// we just moved!
	}
	allPositions := make(map[grids.Pos]bool)
	for step := range g.History {
		allPositions[step.Pos] = true
	}
	return len(allPositions)
}

var TurnRight = map[grids.Dir]grids.Dir{
	grids.Up:    grids.Right,
	grids.Right: grids.Down,
	grids.Down:  grids.Left,
	grids.Left:  grids.Up,
}

// Move moves the guard if possible and returns true, otherwise returns false
func (g *Guard) Move() (moved bool) {
	var nextStep Step
	var nextPos = g.Here.Move(grids.Moves[g.Here.Dir])

	contents := g.Grid[nextPos.Row][nextPos.Col]

	switch contents {
	case '.':
		// valid move, keep in this direction
		nextStep = Step{nextPos, g.Here.Dir}
	case '#':
		// hit an obstacle, so turn right
		nextStep = Step{g.Here.Pos, TurnRight[g.Here.Dir]}
	default:
		// walked off the grid, so we are done
		return false
	}

	if g.History[nextStep] {
		// we've been here, in the same direction
		//debug.Debug("Already visited", nextStep)
		moved = false
	} else {
		// take the step
		//debug.Debug("Moving to", nextStep)
		g.Here = nextStep
		g.History[g.Here] = true
		moved = true
	}
	return
}
