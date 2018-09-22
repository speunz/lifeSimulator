package main

import (
	"math/rand"
)

type Cell struct {
	X   int
	Y   int
	Power int
	StomachFilling int
	Team int
}

func (e Cell) move(newX int, newY int) {
	// increase or decrease Cell coordinates by 1

	// INSERT YOUR CODE HERE #######
}

func (e Cell) devour(victim Cell) {
	// Consumes a closely standing victim if it's not more powerful then the current cell.
	// Stomach filling increases by the power of the victim but can't be more than current Cell power.
	// Victims must belong to a different team
	// INSERT YOUR CODE HERE #######
}

func (e Cell) divide(direction string) {
	// add a new Cell with power not more than current stomach filling
	// a current Cell power increases by a new Cell power
	// a current StomachFilling decreases by a new Cell power
	// a new cell will appear on the right, left, top or bottom side of the current Cell 
	// (it depends on a direction value)
	// ArCells = append(ArCells, newChildCell)

	// INSERT YOUR CODE HERE #######
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
}

func main() {
	var BattleField [100][100]int // global battlefield
	var ArCells []Cell  // global Cells list
	var i, j, probability int

	// filling battlefield with space (0) and obstacles (1)
	for i = 0; i < 100; i++ {
		for j = 0; j < 100; j++ {
			BattleField[i][j] = rand.Intn(1)
			probability = rand.Intn(100)
			if probability > 90 {
				// we will add here the cells of other players
				ArCells = append(ArCells, Cell{
					i,
					j,
					1,
					0,
					3})
			}
		}
	}

	// start battle !
	// INSERT YOUR CODE HERE #######
}
