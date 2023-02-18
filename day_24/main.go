package main

import (
	"aoc/util"
	"fmt"
	"golang.org/x/exp/slices"
	"hash/fnv"
	"os"
	"strconv"
	"strings"
)

// init global variables
var DP = make(map[uint32]int)
var board map[Location][]rune
var storm Storm

type Storm []*Blizzard

type Location struct {
	x, y int
}

func (l Location) Stringify() string {
	return strings.Join([]string{strconv.Itoa(l.x), strconv.Itoa(l.y)}, ",")
}

func (l Location) CheckFreeSquares(board map[Location][]rune) []Location {
	var freeSquares []Location
	CheckSquare(Location{l.x, l.y - 1}, &freeSquares, board)
	CheckSquare(Location{l.x, l.y + 1}, &freeSquares, board)
	CheckSquare(Location{l.x - 1, l.y}, &freeSquares, board)
	CheckSquare(Location{l.x + 1, l.y}, &freeSquares, board)
	return freeSquares
}

func CheckSquare(l Location, freeSquares *[]Location, board map[Location][]rune) {
	val, ok := board[l]
	if ok && len(val) == 1 && util.Contains(val, '.') {
		*freeSquares = append(*freeSquares, l)
	}
}

type Direction rune

const (
	North Direction = '^'
	South           = 'v'
	West            = '<'
	East            = '>'
)

const (
	FREE = '.'
	WALL = '#'
)

type Blizzard struct {
	Location
	Direction
	limit int
}

func RemoveIndex(s []rune, index int) []rune {
	if index == -1 {
		return s
	}
	return append(s[:index], s[index+1:]...)
}

func UpdateBoard(b *Blizzard, vector rune, amount int) {
	board[b.Location] = RemoveIndex(board[b.Location], slices.Index(board[b.Location], rune(b.Direction)))
	if len(board[b.Location]) == 0 {
		board[b.Location] = []rune{'.'}
	}

	if vector == 'y' {
		b.y = b.y + amount
		if b.y == 0 {
			b.y = b.limit
		} else if b.y == b.limit+1 {
			b.y = 1
		}
	} else if vector == 'x' {
		b.x = b.x + amount
		if b.x == 0 {
			b.x = b.limit
		} else if b.x == b.limit+1 {
			b.x = 1
		}
	}

	board[b.Location] = RemoveIndex(board[b.Location], slices.Index(board[b.Location], '.'))
	board[b.Location] = append(board[b.Location], rune(b.Direction))
}

func (b *Blizzard) Move() {
	if b.Direction == North {
		UpdateBoard(b, 'y', -1)
	} else if b.Direction == South {
		UpdateBoard(b, 'y', 1)
	} else if b.Direction == West {
		UpdateBoard(b, 'x', -1)
	} else if b.Direction == East {
		UpdateBoard(b, 'x', 1)
	}
}

func generateBoard(data []string) map[Location][]rune {
	board := make(map[Location][]rune)
	height := len(data) - 2
	width := len(data[0]) - 2
	for y, val := range data {
		for x, tile := range val {
			board[Location{x, y}] = []rune{tile}
			if !(tile == FREE || tile == WALL) {
				if Direction(tile) == North || Direction(tile) == South {
					storm = append(storm, &Blizzard{Location{x, y}, Direction(tile), height})
				} else {
					storm = append(storm, &Blizzard{Location{x, y}, Direction(tile), width})
				}
			}
		}
	}
	return board
}

func drawMap(height, width int, b map[Location][]rune) {
	for row := 0; row < height; row++ {
		for column := 0; column < width; column++ {
			val := b[Location{column, row}]
			if len(val) == 1 {
				fmt.Printf("%s", string(val[0]))
			} else {
				fmt.Printf("%d", len(val))
			}
		}
		fmt.Printf("\n")
	}
}

func hashState(player Location, round int) uint32 {
	state := []byte(player.Stringify() + strconv.Itoa(round))
	h := fnv.New32a()
	h.Write(state)
	return h.Sum32()
}

func reachGoal(player, goal Location, round int) int {

	if player == goal {
		return round
	}

	key := hashState(player, round)
	val, ok := DP[key]
	if ok {
		return val
	}

	var nextSquares []int

	// where the blizzards are in the current round
	currentState := simulation[round]

	if util.Contains(currentState[player], '.') {
		nextSquares = append(nextSquares, reachGoal(player, goal, round+1))
	}

	freeSquares := player.CheckFreeSquares(currentState)
	for _, moveTo := range freeSquares {
		nextSquares = append(nextSquares, reachGoal(moveTo, goal, round+1))
	}

	nextSquares = util.FilterValueFromArray(nextSquares, -1)

	if len(nextSquares) == 0 {
		DP[key] = -1
		return -1
	}

	min, _ := util.Min(nextSquares)
	DP[key] = min
	return min
}

var simulation map[int]map[Location][]rune

func createNewCopy(b map[Location][]rune) map[Location][]rune {
	targetMap := make(map[Location][]rune)
	for k, v := range b {
		dst := make([]rune, len(v))
		copy(dst, v)
		targetMap[k] = dst
	}
	return targetMap
}

func main() {
	data := util.Data(1, "\n")
	data = data[:len(data)-1]

	if len(os.Args) < 3 {
		panic("Please provide puzzle part index as parameter `go run main.go -- 1`")
	}

	part := os.Args[2]

	player := Location{1, 0}
	goal := Location{6, 5}
	if part == "2" {
		goal = Location{120, 26}
	}

	board = generateBoard(data)

	mapHeight := len(data)
	mapWidth := len(data[0])

	drawMap(mapHeight, mapWidth, board)

	// Simulate storm for n rounds
	rounds := 0
	simulatedRounds := 900
	simulation = make(map[int]map[Location][]rune)
	for {
		rounds++

		for _, b := range storm {
			b.Move()
		}

		simulation[rounds] = createNewCopy(board)

		if rounds == simulatedRounds {
			break
		}
	}

	toGoal := reachGoal(player, goal, 1)
	fmt.Printf("Took %d rounds to get to the goal!\n", toGoal-1)
	player = Location{6, 5}
	goal = Location{1, 0}
	if part == "2" {
		player = Location{120, 26}

	}

	backToStart := reachGoal(player, goal, toGoal)
	player = Location{1, 0}
	goal = Location{6, 5}
	if part == "2" {
		goal = Location{120, 26}
	}
	backToGoal := reachGoal(player, goal, toGoal+ backToStart)
	fmt.Printf("Took %d rounds to get to the goal!\n", backToGoal-1) // 13
}
