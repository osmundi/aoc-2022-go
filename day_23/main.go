package main

import (
	"aoc/util"
	"fmt"
	"os"
)

type Board map[Location]*Elf

var board Board

const (
	FREE = '.'
	ELF  = '#'
)

type Direction int

const (
	North Direction = iota
	South
	West
	East
)

type Location struct {
	x, y int
}

type Elf struct {
	Location
}

func (e Elf) String() string {
	return fmt.Sprintf("(%d, %d)", e.x, e.y)
}

func (e *Elf) ProposeMove(order []Direction, proposedMoves map[*Elf]Location) {

	_, N := board[Location{e.x, e.y - 1}]
	_, NE := board[Location{e.x + 1, e.y - 1}]
	_, NW := board[Location{e.x - 1, e.y - 1}]
	_, S := board[Location{e.x, e.y + 1}]
	_, SE := board[Location{e.x + 1, e.y + 1}]
	_, SW := board[Location{e.x - 1, e.y + 1}]
	_, W := board[Location{e.x - 1, e.y}]
	_, E := board[Location{e.x + 1, e.y}]

	if !(N || NE || NW || S || SE || SW || W || E) {
		return
	}

	for _, o := range order {
		if o == North {
			if !(N || NE || NW) {
				proposedMoves[e] = Location{e.x, e.y - 1}
				break
			}
		} else if o == South {
			if !(S || SE || SW) {
				proposedMoves[e] = Location{e.x, e.y + 1}
				break
			}
		} else if o == West {
			if !(W || NW || SW) {
				proposedMoves[e] = Location{e.x - 1, e.y}
				break
			}
		} else if o == East {
			if !(E || NE || SE) {
				proposedMoves[e] = Location{e.x + 1, e.y}
				break
			}
		}
	}
}

func (e *Elf) Move(location Location) {
	e.Location = location
}

func newElf(l Location) *Elf {
	return &Elf{l}
}

func (b Board) generateBoard(data []string, elves *[]*Elf) {
	board = make(map[Location]*Elf)
	var elf *Elf
	for y, val := range data {
		for x, tile := range val {
			if tile == ELF {
				elf = newElf(Location{x, y})
				*elves = append(*elves, elf)
				board[Location{x, y}] = elf
			}
		}
	}
}

func (b Board) drawMap(height, width int) {
	for row := 0; row < height; row++ {
		for column := 0; column < width; column++ {
			_, ok := board[Location{column, row}]
			if ok {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func setMoveOrder(moveOrder []Direction, round int) []Direction {
	if round%4 == 0 {
		return moveOrder
	}
	return append(moveOrder[(round%4):], moveOrder[:(round%4)]...)
}

func findDuplicateProposals(proposedMoves map[*Elf]Location) []*Elf {
	duplicates := make([]*Elf, 0, len(proposedMoves))
	visited := make(map[Location]*Elf, 0)
	for key, val := range proposedMoves {
		visitedKey, ok := visited[val]
		if ok {
			duplicates = append(duplicates, key)
			duplicates = append(duplicates, visitedKey)
		}
		visited[val] = key
	}
	return duplicates
}

func findBoundaries(elves []*Elf) (int, int, int, int) {
	top, bottom, left, right := elves[0].y, elves[0].y, elves[0].x, elves[0].x
	for _, elf := range elves {
		if elf.y > top {
			top = elf.y
		}
		if elf.y < bottom {
			bottom = elf.y
		}
		if elf.x > right {
			right = elf.x
		}
		if elf.x < left {
			left = elf.x
		}
	}
	return top, bottom, right, left
}

func main() {
	data := util.Data(1, "\n")
	data = data[:len(data)-1]

	if len(os.Args) < 3 {
		panic("Please provide puzzle part index as parameter `go run main.go -- 1`")
	}

	elves := make([]*Elf, 0)

	board.generateBoard(data, &elves)
	board.drawMap(12, 14)

	moveOrder := []Direction{North, South, West, East}
	currentOrder := moveOrder
	proposedMoves := make(map[*Elf]Location)
	round := 0

	for {
		currentOrder = setMoveOrder(moveOrder, round)

		for _, elf := range elves {
			elf.ProposeMove(currentOrder, proposedMoves)
		}

		if len(proposedMoves) == 0 {
			fmt.Println(round + 1)
			break
		}

		for _, key := range findDuplicateProposals(proposedMoves) {
			delete(proposedMoves, key)
		}

		for elf, location := range proposedMoves {
			delete(board, elf.Location)
			elf.Move(location)
			board[location] = elf
		}

		proposedMoves = make(map[*Elf]Location)

		round++

		if round == 10 {
			top, bottom, right, left := findBoundaries(elves)
			fmt.Println((top-bottom+1)*(right-left+1) - len(elves))
		}
	}
}
