package day6

import (
	"github.com/tastapod/advent-2024/internal/grids"
	"strings"
	"sync"
)

type Step struct {
	grids.Position
	grids.Dir
}

type GuardTracker struct {
	grids.Grid
	Here     Step
	History  map[Step]bool
	Obstacle grids.Position
}

func NewGuardTracker(grid grids.Grid) (gt *GuardTracker) {
	gt = &GuardTracker{
		Grid: grid,
	}

	for row := range gt.Grid.NumRows {
		rowChars := gt.Grid.Row(row)
		if col := strings.IndexAny(string(rowChars), "<^>v"); col != -1 {
			gt.Here = Step{
				grids.Position{Row: row, Col: col},
				grids.Dir(rowChars[col]),
			}
			gt.History = map[Step]bool{gt.Here: true}
			break
		}
	}
	return
}

func NewGuardTrackerWithObstacle(grid grids.Grid, obstacle grids.Position) (gt *GuardTracker) {
	gt = NewGuardTracker(grid)
	gt.Obstacle = obstacle
	return
}

func (gt *GuardTracker) CountAllPositions() int {
	gt.MoveUntilFinished()
	allPositions := make(map[grids.Position]bool)
	for step := range gt.History {
		allPositions[step.Position] = true
	}
	return len(allPositions)
}

func (gt *GuardTracker) MoveUntilFinished() (result WhatHappened) {
	for result = gt.Move(); result == Moved; result = gt.Move() {
		// we just moved!
	}
	return
}

var TurnRight = map[grids.Dir]grids.Dir{
	grids.Up:    grids.Right,
	grids.Right: grids.Down,
	grids.Down:  grids.Left,
	grids.Left:  grids.Up,
}

type WhatHappened int

const (
	Moved WhatHappened = iota
	Looped
	Exited
)

// Move moves the guard if possible and returns true, otherwise returns false
func (gt *GuardTracker) Move() (result WhatHappened) {
	var nextStep Step
	var nextPos = gt.Here.Plus(grids.Moves[gt.Here.Dir])

	var contents rune
	if nextPos == gt.Obstacle {
		contents = '#'
	} else {
		contents = gt.Grid.At(nextPos.Row, nextPos.Col)
	}

	switch contents {
	case 0:
		// walked off the grid, so we are done
		return Exited
	case '#':
		// hit an obstacle, so turn right
		nextStep = Step{gt.Here.Position, TurnRight[gt.Here.Dir]}
	default:
		// valid move, keep in this direction
		nextStep = Step{nextPos, gt.Here.Dir}
	}

	if gt.History[nextStep] {
		// we've been here, in the same direction
		result = Looped
	} else {
		// take the step
		gt.Here = nextStep
		gt.History[gt.Here] = true
		result = Moved
	}
	return
}

func CountWaysToForceLoop(grid grids.Grid) (total int) {
	results := make(chan bool) // we will only send successes along this
	trackerGroup := sync.WaitGroup{}

	// spin up the workers
	for row := 0; row < grid.NumRows; row++ {
		for col := 0; col < grid.NumCols; col++ {
			// so we can wait for them at the end
			trackerGroup.Add(1)

			// always pass the writer channel explicitly, because reasons
			go func(out chan<- bool) {
				defer trackerGroup.Done()
				tracker := NewGuardTrackerWithObstacle(grid, grids.Position{Row: row, Col: col})

				if result := tracker.MoveUntilFinished(); result == Looped {
					out <- true
				}
			}(results)
		}
	}

	// close the results channel after all the work is done
	go func() {
		trackerGroup.Wait()
		close(results)
	}()

	// keep reading until all the results are in
	for range results {
		total++
	}

	return
}
