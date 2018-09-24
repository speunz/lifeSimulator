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
	action := rand.Intn(ActionCount)
	switch action {
	case ActionMove:
		moves := []int{-1, 1}
		x := moves[rand.Intn(2)]
		y := moves[rand.Intn(2)]
		xOrY := rand.Intn(2)

		if xOrY == 0 {
			cell.move(cell.X + x)
		} else {
			cell.move(cell.Y + y)
		}
	case ActionDevour:
		if c := bi.Get(cell.X+1, cell.Y); c != nil {
			cell.devour(c)
		}
	}
}
