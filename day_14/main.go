package main

import (
	"aoc/util"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coords struct {
	x int
	y int
}

func (location *Coords) SetToStart(sandStartLocation Coords) {
	*location = Coords{x: sandStartLocation.x, y: sandStartLocation.y}
}

func (location *Coords) MoveDown() {
	location.y++
}

func (location *Coords) MoveDownLeft() {
	location.y++
	location.x--
}

func (location *Coords) MoveDownRight() {
	location.y++
	location.x++
}

type Element struct {
	location Coords
	value    string
}

func getLocation(locationString string) Coords {
	location := strings.Split(locationString, ",")

	x, err_x := strconv.Atoi(location[0])
	if err_x != nil {
		// handle error
		fmt.Println(err_x)
		os.Exit(2)
	}
	y, err_y := strconv.Atoi(location[1])
	if err_y != nil {
		// handle error
		fmt.Println(err_y)
		os.Exit(2)
	}

	return Coords{x: x, y: y}
}

func generateRocks(path string, rocks *[]Element) {
	var location Coords
	var lastLocation Coords
	edges := strings.Split(path, " -> ")
	for n, val := range edges {
		location = getLocation(val)
		if n != 0 {
			// interpolate inbetween rocks
			if lastLocation.x == location.x {
				if location.y > lastLocation.y {
					for y := lastLocation.y + 1; y < location.y; y++ {
						*rocks = append(*rocks, Element{value: "#", location: Coords{x: location.x, y: y}})
					}
				} else {
					for y := location.y + 1; y < lastLocation.y; y++ {
						*rocks = append(*rocks, Element{value: "#", location: Coords{x: location.x, y: y}})
					}
				}
			} else {
				if location.x > lastLocation.x {
					for x := lastLocation.x + 1; x < location.x; x++ {
						*rocks = append(*rocks, Element{value: "#", location: Coords{x: x, y: location.y}})
					}
				} else {
					for x := location.x + 1; x < lastLocation.x; x++ {
						*rocks = append(*rocks, Element{value: "#", location: Coords{x: x, y: location.y}})
					}
				}
			}
		}
		*rocks = append(*rocks, Element{value: "#", location: location})
		lastLocation = location
	}
}

func drawCave(cave map[Coords]*Element, limits ...int) {
	if len(limits) != 4 {
		return
	}
	sandStartLocation := Coords{500, 0}
	for y := limits[0]; y < limits[2]; y++ {
		for x := limits[3]; x <= limits[1]; x++ {
			element, ok := cave[Coords{x, y}]
			if ok {
				fmt.Printf("%s", element.value)
			} else if sandStartLocation.x == x && sandStartLocation.y == y {
				fmt.Printf("+")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
	// draw the floor
	for x := limits[3]; x <= limits[1]; x++ {
		fmt.Printf("#")
	}
}

func main() {
	data := util.Data(1, "\n")
	data = data[:len(data)-1]

	if len(os.Args) < 3 {
		panic("Please provide puzzle part index as parameter `go run main.go -- 1`")
	}

	part := os.Args[2]
	fmt.Printf("Run part: %s\n", part)

	cave := make(map[Coords]*Element)
	sandStartLocation := Coords{500, 0}

	var rocks []Element
	for _, val := range data {
		generateRocks(val, &rocks)
	}

	// create boundaries for the cave
	leftLimit := 500
	rightLimit := 500
	topLimit := 0
	bottomLimit := 0
	for _, rock := range rocks {
		if rock.location.x < leftLimit {
			leftLimit = rock.location.x
		} else if rock.location.x > rightLimit {
			rightLimit = rock.location.x
		}
		if rock.location.y > bottomLimit {
			bottomLimit = rock.location.y
		} else if rock.location.y < topLimit {
			topLimit = rock.location.y
		}
		cave[rock.location] = &rock
	}

	// set the floor
	if part == "2" {
		bottomLimit = bottomLimit + 2
	}

	// draw cave without sand
	//drawCave(cave, topLimit, rightLimit, bottomLimit, leftLimit)

	// generate sand
	var movingSand Coords
	movingSand.SetToStart(sandStartLocation)
	for {
		// check if sand has hit the bottom
		if part == "2" && movingSand.y+1 == bottomLimit {
			cave[movingSand] = &Element{location: movingSand, value: "o"}
			movingSand.SetToStart(sandStartLocation)
			continue
		}

		// If the location one step down is reserved
		_, downBlocked := cave[Coords{x: movingSand.x, y: movingSand.y + 1}]
		if downBlocked {
			_, downLeftBlocked := cave[Coords{x: movingSand.x - 1, y: movingSand.y + 1}]
			if downLeftBlocked {
				_, downRightBlocked := cave[Coords{x: movingSand.x + 1, y: movingSand.y + 1}]
				if downRightBlocked {
					// sand comes to rest
					cave[movingSand] = &Element{location: movingSand, value: "o"}

					// sand has reached the top
					if movingSand == sandStartLocation {
						break
					}
					movingSand.SetToStart(sandStartLocation)

				} else {
					movingSand.MoveDownRight()
				}
			} else {
				movingSand.MoveDownLeft()
			}
		} else {
			movingSand.MoveDown()
		}

		// check if sand is out of bounds
		if part == "1" {
			if movingSand.x <= leftLimit || movingSand.x >= rightLimit || movingSand.y > bottomLimit {
				break
			}
		}
	}

	// draw cave with sand
	if part == "1" {
		drawCave(cave, topLimit, rightLimit, bottomLimit, leftLimit)
	} else {
        
		//drawCave(cave, topLimit, rightLimit+200, bottomLimit, leftLimit-150)

	}

	// answer
	sum := 0
	for _, element := range cave {
		if element.value == "o" {
			sum++
		}
	}
	fmt.Println(sum)

}
