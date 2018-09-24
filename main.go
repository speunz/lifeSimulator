package main

import (
	"fmt"
	"math/rand"

	"github.com/andlabs/ui"
)

var BattleField [100][100]int // global battlefield
var ArCells []Cell            // global Cells list
var Teams []Team

type BattleInfo struct {
	BattleField *[100][100]int
	Cells       *[]Cell
	Teams       []Team
}

// Size returns the size of the battlefield
func (bi *BattleInfo) Size() (int, int) {
	// hardcoded to 100, 100 according to the problem statement
	return 100, 100
}

// Get returns the cell at position x, y
func (bi *BattleInfo) GetCell(x, y int) *Cell {
	for _, c := range *bi.Cells {
		if c.X == x && c.Y == y {
			return &c
		}
	}

	return nil
}

var bi = &BattleInfo{&BattleField, &ArcCells, Teams}

// GetObstacle returns true if there's an obstacle at the position x, y
func (bi *BattleInfo) GetObstacle(x, y int) bool {
	return (bi.BattleField[x][y] != 0)
}

type Cell struct {
	X              int
	Y              int
	Power          int
	StomachFilling int
	Team           int
}

func (e Cell) move(newX int, newY int) {
	// increase or decrease Cell coordinates by 1
	e.X = newX
	e.Y = newY
}

func (e Cell) devour(victim Cell) {
	// Consumes a closely standing victim if it's not more powerful then the current cell.
	// Stomach filling increases by the power of the victim but can't be more than current Cell power.
	// Victims must belong to a different team
	// INSERT YOUR CODE HERE #######

	if victim.Team != e.Team && e.Power >= victim.Power {
		e.StomachFilling += victim.Power
		if e.StomachFilling > e.Power {
			e.StomachFilling = e.Power
		}
		victim.X = -1
		victim.Y = -1
	}
}

func (e Cell) divide(direction string) {
	// add a new Cell with power not more than current stomach filling
	// a current Cell power increases by a new Cell power
	// a current StomachFilling decreases by a new Cell power
	// a new cell will appear on the right, left, top or bottom side of the current Cell
	// (it depends on a direction value)
	// ArCells = append(ArCells, newChildCell)

	// INSERT YOUR CODE HERE #######

	newPow := e.Power / 2
	if newPow > e.StomachFilling {
		newPow = e.StomachFilling
	}

	nc := Cell{
		X:              e.X,
		Y:              e.Y,
		Power:          newPow,
		StomachFilling: 0,
		Team:           e.Team,
	}

	switch divide {
	case "left":
		nc.X--
	case "right":
		nc.X++
	case "top":
		nc.Y--
	case "bottom":
		nc.Y++
	}

	if nc.X >= 0 && nc.X < 100 && nc.Y >= 0 && nc.Y < 100 {
		e.Power += nc.Power
		e.StomachFilling -= nc.Power
		*ArcCells = append(*ArcCells, nc)
	}
}

func (e Cell) step() {
	// this method will be called for each step of the game
	// during this step current Cell has a right to move and to divide one time

	// if Cell wants to
	// e.move(newX, newY)

	// if Cell wants to
	// e.divide("right")

	// if Cell wants to and can
	// e.devour(ArCells[3])

	// INSERT YOUR CODE HERE #######

	Teams[e.Team].Step(&e)
}

func main() {
	var i, j, probability int
	const numTeams = 5

	for i := 0; i < numTeams; i++ {
		Teams = append(Teams, NewTeamRandom(fmt.Sprintf("Team %d", i)))
	}

	// filling battlefield with space (0) and obstacles (1)
	for i = 0; i < 100; i++ {
		for j = 0; j < 100; j++ {
			BattleField[i][j] = (rand.Intn(100) / 90)
			probability = rand.Intn(100)
			if probability > 90 {
				// we will add here the cells of other players
				ArCells = append(ArCells, Cell{
					i,
					j,
					1,
					0,
					rand.Intn(numTeams)})
			}
		}
	}

	// start battle !
	// INSERT YOUR CODE HERE #######

	bi := BattleInfo{
		BattleField: &BattleField,
		Cells:       &ArCells,
	}
	ui.Main(createUI(&bi))
}
