package main

import "math/rand"

const (
	ActionMove = iota
	ActionDevour
	ActionDivide
	ActionCount
)

type Team interface {
	Name() string
	Step(cell *Cell)
}

type TeamBase struct {
	name string
}

func (t *TeamBase) Name() string {
	return t.name
}

type TeamRandom struct {
	TeamBase
}

func NewTeamRandom(name string) *TeamRandom {
	t := TeamRandom{
		TeamBase{
			name: name,
		},
	}

	return &t
}

func (t *TeamRandom) Step(cell *Cell) {
	actioned := false
	for i := 0; i < 100 && !actioned; i++ {
		action := rand.Intn(ActionCount)
		switch action {
		case ActionMove:
			moves := []int{-1, 1}
			x := moves[rand.Intn(2)]
			y := moves[rand.Intn(2)]
			xOrY := rand.Intn(2)

			if xOrY == 0 {
				cell.move(cell.X+x, cell.Y)
				actioned = true
			} else {
				cell.move(cell.X, cell.Y+y)
				actioned = true
			}
		case ActionDevour:
			if c := bi.GetCell(cell.X+1, cell.Y); c != nil {
				cell.devour(c)
				actioned = true
			} else if c := bi.GetCell(cell.X-1, cell.Y); c != nil {
				cell.devour(c)
				actioned = true
			} else if c := bi.GetCell(cell.X, cell.Y+1); c != nil {
				cell.devour(c)
				actioned = true
			} else if c := bi.GetCell(cell.X, cell.Y-1); c != nil {
				cell.devour(c)
				actioned = true
			}
		case ActionDivide:
			avail := func(x, y int) bool {
				return bi.GetCell(x, y) == nil && !bi.GetObstacle(x, y)
			}

			switch {
			case avail(cell.X-1, cell.Y):
				cell.divide("left")
				actioned = true
			case avail(cell.X+1, cell.Y):
				cell.divide("right")
				actioned = true
			case avail(cell.X, cell.Y-1):
				cell.divide("top")
				actioned = true
			case avail(cell.X, cell.Y+1):
				cell.divide("bottom")
				actioned = true
			}
		}
	}
}
