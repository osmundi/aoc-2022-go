package main

import (
	"aoc/util"
	"fmt"
	"math"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

var board map[Location]rune

const (
	void  rune = ' '
	open  rune = '.'
	solid rune = '#'
)

var warpMap map[Face]Face

type Face struct {
	location Location
	facing   Direction
	inverse  bool
}

func mapSides(n int) {
	var start, end []Face
	for _, f := range faces {
		start = generateStartSide(f[0].location, f[0].facing, f[0].inverse, n)
		end = generateStartSide(f[1].location, f[1].facing, f[1].inverse, n)
		generateWarpMapping(start, end, n)
	}
}

func generateStartSide(start Location, warpDir Direction, inverse bool, count int) []Face {
	var x, y int
	locations := []Face{}
	if warpDir == North || warpDir == South {
		x = start.x
		if inverse {
			for y := start.y; y > start.y-count; y-- {
				locations = append(locations, Face{Location{x, y}, warpDir, false})
			}
		} else {
			for y := start.y; y < start.y+count; y++ {
				locations = append(locations, Face{Location{x, y}, warpDir, false})
			}
		}
	} else {
		y = start.y
		if inverse {
			for x := start.x; x > start.x-count; x-- {
				locations = append(locations, Face{Location{x, y}, warpDir, false})
			}
		} else {
			for x := start.x; x < start.x+count; x++ {
				locations = append(locations, Face{Location{x, y}, warpDir, false})
			}
		}
	}
	return locations
}

func generateWarpMapping(start []Face, end []Face, n int) {
	for i := 0; i < n; i++ {
		warpMap[start[i]] = end[i]
		warpMap[end[i]] = start[i]
	}
}

var faces = [][]Face{
	{ // 1 <-> 6
		{Location{1, 51}, North, false},
		{Location{151, 1}, West, false},
	},
	{ // 1 <-> 4
		{Location{1, 51}, West, false},
		{Location{150, 1}, West, true},
	},
	{ // 2 <-> 6
		{Location{1, 101}, North, false},
		{Location{200, 1}, South, false},
	},
	{ // 2 <-> 5
		{Location{50, 150}, East, true},
		{Location{101, 100}, East, false},
	},
	{ // 2 <-> 3
		{Location{50, 101}, South, false},
		{Location{51, 100}, East, false},
	},
	{ // 3 <-> 4
		{Location{51, 51}, West, false},
		{Location{101, 1}, North, false},
	},
	{ // 5 <-> 6
		{Location{150, 51}, South, false},
		{Location{151, 50}, East, false},
	},
}
var test_faces = [][]Face{
	{
		{Location{1, 9}, North, false},
		{Location{5, 4}, North, true},
	},
	{
		{Location{1, 9}, West, false},
		{Location{5, 5}, North, false},
	},
	{
		{Location{8, 5}, East, false},
		{Location{12, 9}, West, true},
	},
	{
		{Location{5, 12}, East, false},
		{Location{9, 16}, North, true},
	},
	{
		{Location{1, 12}, East, false},
		{Location{12, 16}, East, true},
	},
	{
		{Location{12, 9}, South, false},
		{Location{8, 4}, South, true},
	},
	{
		{Location{5, 1}, West, false},
		{Location{12, 13}, South, false},
	},
}

func readData(data []string) (Location, map[Location]rune, string) {
	// read board and instructions into memory
	var board = make(map[Location]rune)
	var instructions string
	var first Location
	stopReading := false
	blank := false
	readFirst := true
	for x, val := range data {
		x++
		blank = strings.TrimSpace(val) == ""
		if blank {
			stopReading = true
			continue
		}
		if stopReading {
			instructions = val
			return first, board, instructions
		}
		for y, tile := range val {
			y++
			if tile != void {
				if readFirst {

					first = Location{x, y}
					readFirst = false
				}
				board[Location{x, y}] = tile
			}
		}
	}
	return Location{}, nil, ""
}

func locationInHistory(l Location, history []Location) bool {
	for _, b := range history {
		if b == l {
			return true
		}
	}
	return false
}

func drawMap(moveHistory []Location) {
	boardWidth := 150
	boardHeight := 200
	//boardWidth := 17
	//boardHeight := 14

	for row := 0; row < boardHeight; row++ {
		for column := 0; column < boardWidth; column++ {
			val, ok := board[Location{row, column}]
			if ok {
				if locationInHistory(Location{row, column}, moveHistory) {
					fmt.Printf("o")
				} else {
					fmt.Printf("%s", string(val))
				}
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}

type Location struct {
	x int
	y int
}

const (
	Right int = iota
	Left
)

type Direction uint8

const (
	East Direction = 1 << iota
	South
	West
	North
)

func (d Direction) String() string {
	return [...]string{"East", "South", "West", "North"}[d.Log()]
}

func (d Direction) Log() int {
	return int(math.Logb(float64(d)))
}

type Player struct {
	location *Location
	history  []Location
	facing   Direction
}

func (p Player) String() string {
	return fmt.Sprintf("Current: %v\nFacing: %v", p.location, p.facing)
}

func (p *Player) Warp(dir Direction) {
	steps := 1
	for {
		if dir == North {
			_, ok := board[Location{p.location.x - steps, p.location.y}]
			if ok {
				steps++
				continue
			} else {
				steps--
				val, _ := board[Location{p.location.x - steps, p.location.y}]
				if val == solid {
					break
				}
				p.location.x = p.location.x - steps
				p.Save()
				break
			}
		} else if dir == East {
			_, ok := board[Location{p.location.x, p.location.y + steps}]
			if ok {
				steps++
				continue
			} else {
				steps--
				val, _ := board[Location{p.location.x, p.location.y + steps}]
				if val == solid {
					break
				}
				p.location.y = p.location.y + steps
				p.Save()
				break
			}
		} else if dir == South {
			_, ok := board[Location{p.location.x + steps, p.location.y}]
			if ok {
				steps++
				continue
			} else {
				steps--
				val, _ := board[Location{p.location.x + steps, p.location.y}]
				if val == solid {
					break
				}
				p.location.x = p.location.x + steps
				p.Save()
				break
			}
		} else if dir == West {
			_, ok := board[Location{p.location.x, p.location.y - steps}]
			if ok {
				steps++
				continue
			} else {
				steps--
				val, _ := board[Location{p.location.x, p.location.y - steps}]
				if val == solid {
					break
				}
				p.location.y = p.location.y - steps
				p.Save()
				break
			}
		}
	}
}

func (p *Player) WarpCube(dir Direction) {
	warp := warpMap[Face{*p.location, dir, false}]
	val := board[warp.location]
	if val == solid {
		return
	} else {
		*p.location = warp.location
		p.WarpTurn(warp)
		p.Save()
	}
}

func (p *Player) MoveNorth() {
	val, ok := board[Location{p.location.x - 1, p.location.y}]
	if ok {
		if val == solid {
			return
		} else {
			p.location.x--
			p.Save()
		}
	} else {
		//p.Warp(South)
		p.WarpCube(North)
	}
}

func (p *Player) MoveWest() {
	val, ok := board[Location{p.location.x, p.location.y - 1}]
	if ok {
		if val == solid {
			return
		} else {
			p.location.y--
			p.Save()
		}
	} else {
		//p.Warp(East)
		p.WarpCube(West)
	}
}
func (p *Player) MoveEast() {
	val, ok := board[Location{p.location.x, p.location.y + 1}]
	if ok {
		if val == solid {
			return
		} else {
			p.location.y++
			p.Save()
		}
	} else {
		//p.Warp(West)
		p.WarpCube(East)
	}
}
func (p *Player) MoveSouth() {
	val, ok := board[Location{p.location.x + 1, p.location.y}]
	if ok {
		if val == solid {
			return
		} else {
			p.location.x++
			p.Save()
		}
	} else {
		//p.Warp(North)
		p.WarpCube(South)
	}
}

func (p *Player) Move(steps int) {
	for i := 0; i < steps; i++ {
		if p.facing == North {
			p.MoveNorth()
		} else if p.facing == East {
			p.MoveEast()
		} else if p.facing == South {
			p.MoveSouth()
		} else if p.facing == West {
			p.MoveWest()
		}
	}
}

func (p *Player) Save() {
	p.history = append(p.history, *(p).location)
}

func (p *Player) WarpTurn(warp Face) {
	if warp.facing == South && p.facing == North {
		return
	} else if p.facing == South && warp.facing == North {
		return
	} else if warp.facing == East && p.facing == North {
		p.Turn(Left)
	} else if warp.facing == North && p.facing == East {
		p.Turn(Right)
	} else if p.facing<<1 == warp.facing {
		p.Turn(Left)
	} else if p.facing>>1 == warp.facing {
		p.Turn(Right)
	} else {
		p.Turn(Right)
		p.Turn(Right)
	}
}

func (p *Player) Turn(dir int) {
	if dir == Right {
		if p.facing == North {
			p.facing = Direction(bits.RotateLeft8(uint8(p.facing), -3))
		} else {
			p.facing = Direction(bits.RotateLeft8(uint8(p.facing), 1))
		}
	} else if dir == Left {
		if p.facing == East {
			p.facing = Direction(bits.RotateLeft8(uint8(p.facing), 3))
		} else {
			p.facing = Direction(bits.RotateLeft8(uint8(p.facing), -1))
		}
	}
}

func (p *Player) Result() int {
	return 1000*p.location.x + 4*p.location.y + p.facing.Log()
}

func main() {
	data := util.Data(1, "\n")
	data = data[:len(data)-1]

	if len(os.Args) < 3 {
		panic("Please provide puzzle part index as parameter `go run main.go -- 1`")
	}

	part := os.Args[2]

	fmt.Printf("Run part: %s\n", part)

	var first Location
	var instructions string

	first, board, instructions = readData(data)

	if instructions == "" {
		panic("no instructions")
	}

	player := Player{location: &first, facing: East}
	player.Save()

	readBuffer := ""
	var steps int
	var seperatedInstructions [][]int
	for _, b := range instructions {
		if string(b) == "R" {
			steps, _ = strconv.Atoi(readBuffer)
			seperatedInstructions = append(seperatedInstructions, []int{steps, Right})
			readBuffer = ""
			continue
		} else if string(b) == "L" {
			steps, _ = strconv.Atoi(readBuffer)
			seperatedInstructions = append(seperatedInstructions, []int{steps, Left})
			readBuffer = ""
			continue
		} else {
			readBuffer += string(b)
		}
	}
	steps, _ = strconv.Atoi(readBuffer)
	seperatedInstructions = append(seperatedInstructions, []int{steps, -1})

	warpMap = make(map[Face]Face)
	mapSides(50)

	for _, i := range seperatedInstructions {
		player.Move(i[0])
		if player.location.x == 0 {
			break
		}
		if i[1] == -1 {
			break
		}
		player.Turn(i[1])
	}
	drawMap(player.history)
	fmt.Println(player.Result())

}
