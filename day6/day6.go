package day6

import (
	"github.com/tastapod/advent-2024/grids"
	"strings"
	"sync"
)

type Step struct {
	grids.Pos
	grids.Dir
}

type GuardTracker struct {
	grids.Grid
	Here     Step
	History  map[Step]bool
	Obstacle grids.Pos
}

func NewGuardTracker(grid grids.Grid) (g *GuardTracker) {
	g = &GuardTracker{
		Grid: grid,
	}

	for row, rowChars := range g.Grid {
		if col := strings.IndexAny(string(rowChars), "<^>v"); col != -1 {
			g.Here = Step{
				grids.Pos{Row: row, Col: col},
				grids.Dir(rowChars[col]),
			}
			g.History = map[Step]bool{g.Here: true}
			break
		}
	}
	return
}

func NewGuardTrackerWithObstacle(grid grids.Grid, obstacle grids.Pos) (gt *GuardTracker) {
	gt = NewGuardTracker(grid)
	gt.Obstacle = obstacle
	return
}

func (gt *GuardTracker) CountAllPositions() int {
	gt.MoveUntilFinished()
	allPositions := make(map[grids.Pos]bool)
	for step := range gt.History {
		allPositions[step.Pos] = true
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
	var nextPos = gt.Here.Move(grids.Moves[gt.Here.Dir])

	var contents rune
	if nextPos == gt.Obstacle {
		contents = '#'
	} else {
		contents = gt.Grid[nextPos.Row][nextPos.Col]
	}

	switch contents {
	case 0:
		// walked off the grid, so we are done
		return Exited
	case '#':
		// hit an obstacle, so turn right
		nextStep = Step{gt.Here.Pos, TurnRight[gt.Here.Dir]}
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
	results := make(chan bool)
	trackerGroup := sync.WaitGroup{}
	readerGroup := sync.WaitGroup{}

	// reader for results, counts into total
	readerGroup.Add(1)
	go func(in <-chan bool) {
		defer readerGroup.Done()
		// block until there is some data, then read until closed
		for range in {
			total++
		}
	}(results)

	// then spin up the workers
	for row := 1; row < len(grid)-1; row++ {
		for col := 1; col < len(grid[0])-1; col++ {
			trackerGroup.Add(1)
			go func(out chan<- bool) {
				defer trackerGroup.Done()
				tracker := NewGuardTrackerWithObstacle(grid, grids.Pos{Row: row, Col: col})
				if tracker.MoveUntilFinished() == Looped {
					out <- true
				}
			}(results)
		}
	}

	trackerGroup.Wait()
	close(results)
	readerGroup.Wait()
	return
}
