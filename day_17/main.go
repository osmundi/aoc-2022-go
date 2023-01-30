package main

import (
	"aoc/util"
	"fmt"
)

/*
chamber:
4 |..@@@@.|
3 |.......|
2 |.......|
1 |.......|
0 +-------+

	012345678
*/
var chamber = make(map[Location]string)

type Location struct {
	x int
	y int
}

type Rock struct {
	pattern  [][]int
	moving   bool
	step     int
	gusts    []rune   // try to find
	location Location // left corner of the pattern
}

type RockSet struct {
	flat     Rock
	plus     Rock
	l        Rock
	vertical Rock
	cube     Rock
}

func (rock *Rock) GetRockShapeLocations() []Location {
	var locations []Location
	for row := 0; row < len(rock.pattern); row++ {
		for column := 0; column < len(rock.pattern[row]); column++ {
			if rock.pattern[row][column] == 1 {
				locations = append(locations, Location{rock.location.x + column, rock.location.y + row})
			}
		}
	}

	return locations
}

func (rock *Rock) GetY() int {
	// rock height
	return rock.location.y + len(rock.pattern) - 1
}

func (rock *Rock) StoppedMoving() bool {
	if rock.location.y == 1 {
		return true
	}
	return false
}

func (rock *Rock) MoveDown() {
	// there is no one blocking in the first three moves
	for _, l := range rock.GetRockShapeLocations() {
		_, rockExists := chamber[Location{l.x, l.y - 1}]
		if rockExists {
			rock.moving = false
			return
		}
	}
	rock.location.y--
	rock.step++
}

func (rock *Rock) MoveLeft() {
	for _, l := range rock.GetRockShapeLocations() {
		if l.x-1 < 1 {
			return
		}
		_, rockExists := chamber[Location{l.x - 1, l.y}]
		if rockExists {
			return
		}
	}
	rock.location.x--
}

func (rock *Rock) MoveRight() {
	for _, l := range rock.GetRockShapeLocations() {
		if l.x+1 > 7 {
			return
		}
		_, rockExists := chamber[Location{l.x + 1, l.y}]
		if rockExists {
			return
		}
	}
	rock.location.x++
}

func (rock *Rock) SetStartingLocation(highestRock int) {
	rock.location.y = highestRock + 5
	rock.location.x = 3
}

func (rock *Rock) PushWithJet(direction rune) {
	if direction == '<' {
		rock.MoveLeft()
	} else if direction == '>' {
		rock.MoveRight()
	}

	rock.gusts = append(rock.gusts, direction)
}

func drawChamber(height int) {
	for row := height; row > 0; row-- {
		for column := 1; column <= 7; column++ {
			val, ok := chamber[Location{column, row}]
			if ok {
				fmt.Printf(val)
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func GetTopRockLocation() int {
	var mapKeys []Location
	var topRock Location
	mapKeys = util.MapKeys(chamber)
	topRock = Location{0, 0}
	for _, key := range mapKeys {
		if key.y > topRock.y {
			topRock = key
		}
	}
	return topRock.y
}

type RockList []Rock

type Comparer interface {
	Compare(b RockList) int
}

func Equal(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func (a RockList) Compare(b RockList) int {
	if len(a) != len(b) {
		return 0
	}

	for i, rocks_a := range a {
		if !Equal(rocks_a.gusts, b[i].gusts) {
			return 0
		}
		if rocks_a.step != b[i].step {
			return 0
		}
	}

	fmt.Println("Found pattern between!")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("---")

	return b[len(b)-1].location.y - a[len(a)-1].location.y
}

func findDuplicates(list map[int]RockList) (int, int) {
	var comparsion int
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			comparsion = list[i].Compare(list[j])
			if comparsion > 0 {
				return comparsion, j - i
			}
		}
	}
	return 0, 0
}

func main() {
	data := util.Data(1, "\n")
	data = data[:len(data)-1]

	jetPattern := []rune(data[0])
	jetPatternLength := len(jetPattern)
	rocksInStack := 0
	turn := 0
	topRock := 0
	//rocksInStackLimit := 2022 // PART 1
	rocksInStackLimit := 10000 // PART 2

	flat := [][]int{{1, 1, 1, 1}}
	plus := [][]int{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}}
	l := [][]int{{1, 1, 1}, {0, 0, 1}, {0, 0, 1}}
	vertical := [][]int{{1}, {1}, {1}, {1}}
	cube := [][]int{{1, 1}, {1, 1}}
	fallOrder := [][][]int{flat, plus, l, vertical, cube}

	var fallingRock Rock
	fallingRock.SetStartingLocation(0)
	fallingRock.pattern = fallOrder[rocksInStack%5]
	fallingRock.moving = true

	var rockSets = make(map[int]RockList)
	var rockSetsIndex = 0

	for {
		fallingRock.MoveDown()

		if fallingRock.moving {
			fallingRock.PushWithJet(jetPattern[turn%jetPatternLength])
			turn++
		}

		if fallingRock.StoppedMoving() || !fallingRock.moving {
			rocksInStack++

			for _, l := range fallingRock.GetRockShapeLocations() {
				chamber[l] = "#"
			}

			if topRock < fallingRock.GetY() {
				topRock = fallingRock.GetY()
			}

			/*
				            FOR FINDING PATTERN DATA
				            //test how many rocks are in stack before the pattern starts to emerge
							if fallingRock.location.x == 2 && fallingRock.location.y == 391 {
				                fmt.Println("rocksInStack:")
								fmt.Println(rocksInStack)
							}

				            //find out pattern height at nth rock of the pattern
				            if rocksInStack == 264+1606 {
				                fmt.Println(fallingRock)
				                fmt.Println(fallingRock.GetY())
				            }
			*/

			if rocksInStack%5 == 0 {
				rockSetsIndex++
				rockSets[rockSetsIndex] = RockList{fallingRock}
			} else {
				if entry, ok := rockSets[rockSetsIndex]; ok {
					entry = append(entry, fallingRock)
					rockSets[rockSetsIndex] = entry
				}
			}

			if rocksInStack == rocksInStackLimit {
				fmt.Println(topRock)
				break
			}

			// set new rock
			fallingRock.SetStartingLocation(topRock)
			fallingRock.pattern = fallOrder[rocksInStack%5]
			fallingRock.moving = true
			fallingRock.gusts = nil
			fallingRock.step = 0
		}

	}

	drawChamber(20)

	setHeight, setLength := findDuplicates(rockSets)
	fmt.Println(setHeight)
	fmt.Println(setLength)

    // did rest of the calcuting manually....

	/*

				    TEST DATA:

						7 settiä (5 kiveä) -> 35 sarjoissa
						setin korkeus: 53

				        Found pattern between!
				        [{[[1 1] [1 1]] false 10 [60 60 60 62 62 62 60 60 60 62] {2 31}} {[[1 1 1 1]] false 4 [60 60 60 62] {2 37}} {[[0 1 0] [1 1 1] [0 1 0]] false 5 [62 60 62 62 60] {5 37}} {[[1 1 1] [0 0 1] [0 0 1]] false 4 [60 62 62 62] {5 40}} {[[1] [1] [1] [1]] false 9 [62 62 60 60 62 60 62 62 60] {3 38}}]
				                --->>> 24 kiveä
				        [{[[1 1] [1 1]] false 10 [60 60 60 62 62 62 60 60 60 62] {2 84}} {[[1 1 1 1]] false 4 [60 60 60 62] {2 90}} {[[0 1 0] [1 1 1] [0 1 0]] false 5 [62 60 62 62 60] {5 90}} {[[1 1 1] [0 0 1] [0 0 1]] false 4 [60 62 62 62] {5 93}} {[[1] [1] [1] [1]] false 9 [62 62 60 60 62 60 62 62 60] {3 91}}]
				                --->>> 24 + 35 kiveä

				        -->> 24 + 26 kiveä == 78
				        eli 26:nnen kiven kohdalla setistä, korkeus on kasvanut 40 kiven verran (78-38)

				        PATTERNS: (100...000 - 24) / 35 = 28571428570
						TOTAL HEIGHT: 38 (24 kiveä) + 28571428570*53 + 40(26 kiveä) = 1514285714288


		            PROD DATA:

				        Found pattern between!
				        [{[[1 1] [1 1]] false 6 [62 62 60 62 62 62] {6 388}} {[[1 1 1 1]] false 4 [60 60 60 60] {1 390}} {[[0 1 0] [1 1 1] [0 1 0]] false 4 [62 60 60 62] {3 391}} {[[1 1 1] [0 0 1] [0 0 1]] false 4 [62 60 60 62] {3 394}} {[[1] [1] [1] [1]] false 10 [62 60 60 60 60 62 62 60 60 62] {2 391}}]
				        --->>> 264 kiveä
				        --->>> 264 + 1606 kiveä = 2846

				        [{[[1 1] [1 1]] false 6 [62 62 60 62 62 62] {6 3035}} {[[1 1 1 1]] false 4 [60 60 60 60] {1 3037}} {[[0 1 0] [1 1 1] [0 1 0]] false 4 [62 60 60 62] {3 3038}} {[[1 1 1] [0 0 1] [0 0 1]] false 4 [62 60 60 62] {3 3041}} {[[1] [1] [1] [1]] false 10 [62 60 60 60 60 62 62 60 60 62] {2 3038}}]
				        ---
				        346
				        setin korkeus: 2647

				        346*5 = 1730

				        PATTERNS: (100...000 - 264) / 1730 = 578034681


				        Kiviä sarjassa: 100...000 - 1730*578034681 = 1870
				        Kiviä lopussa: 1870 - 264 = 1606

				        Pattern height: 578034681*2647 = 1530057800607
						TOTAL HEIGHT: 391 (264 kiveä) + 1530057800607 + 2455(1606 kiveä) = 1514285714288

	*/

}
