package main

import (
	"aoc/util"
	"fmt"
	"os"
)

const (
	north      = 'U'
	south      = 'D'
	west       = 'L'
	east       = 'R'
	north_west = '1'
	north_east = '2'
	south_west = '3'
	south_east = '4'
)

type coordinate struct {
	x int
	y int
}

type Rope struct {
	name             string
	location         [2]int
	visited_location [][2]int
}

func newRope(name string) *Rope {
	origin := [2]int{0, 0}
	rope := Rope{name: name, location: origin, visited_location: [][2]int{origin}}
	return &rope
}

func (rope *Rope) Move(direction rune) {
	if direction == north {
		rope.location[1]++
	} else if direction == south {
		rope.location[1]--
	} else if direction == west {
		rope.location[0]--
	} else if direction == east {
		rope.location[0]++
	} else if direction == south_east {
		rope.location[0]++
		rope.location[1]--
	} else if direction == north_east {
		rope.location[0]++
		rope.location[1]++
	} else if direction == south_west {
		rope.location[0]--
		rope.location[1]--
	} else if direction == north_west {
		rope.location[0]--
		rope.location[1]++
	}
	rope.visited_location = append(rope.visited_location, rope.location)
}

func (rope *Rope) Follow(location [2]int) {
	if location[0]-rope.location[0] == 2 {
		if location[1]-rope.location[1] == 2 {
			rope.Move(north_east)
		} else if rope.location[1]-location[1] == 2 {
			rope.Move(south_east)
		} else if location[1]-rope.location[1] == 1 {
			rope.Move(north_east)
		} else if rope.location[1]-location[1] == 1 {
			rope.Move(south_east)
		} else {
			rope.Move(east)
		}
	} else if rope.location[0]-location[0] == 2 {
		if location[1]-rope.location[1] == 2 {
			rope.Move(north_west)
		} else if rope.location[1]-location[1] == 2 {
			rope.Move(south_west)
		} else if location[1]-rope.location[1] == 1 {
			rope.Move(north_west)
		} else if rope.location[1]-location[1] == 1 {
			rope.Move(south_west)
		} else {
			rope.Move(west)
		}
	} else if rope.location[1]-location[1] == 2 {
		if location[0]-rope.location[0] == 2 {
			rope.Move(south_east)
		} else if rope.location[0]-location[0] == 2 {
			rope.Move(south_west)
		} else if location[0]-rope.location[0] == 1 {
			rope.Move(south_east)
		} else if rope.location[0]-location[0] == 1 {
			rope.Move(south_west)
		} else {
			rope.Move(south)
		}
	} else if location[1]-rope.location[1] == 2 {
		if location[0]-rope.location[0] == 2 {
			rope.Move(north_east)
		} else if rope.location[0]-location[0] == 2 {
			rope.Move(north_west)
		} else if location[0]-rope.location[0] == 1 {
			rope.Move(north_east)
		} else if rope.location[0]-location[0] == 1 {
			rope.Move(north_west)
		} else {
			rope.Move(north)
		}
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

    var knots int

	if part == "1" {
		knots = 2
	} else {
		knots = 10
	}

	ropes := map[int]*Rope{}
	for i := 0; i < knots; i++ {
		ropes[i] = newRope(fmt.Sprintf("rope_%d", i))
	}

	var dir string
	var n int

	for _, val := range data {
		_, err := fmt.Sscanf(val, "%s %d", &dir, &n)

		if err != nil {
			panic(err)
		}

		for i := 0; i < n; i++ {
			// Move head
			ropes[0].Move(util.RuneElement(dir, 0))

			// And the tail follows
			for j := 1; j < len(ropes); j++ {
				ropes[j].Follow(ropes[j-1].location)
			}
		}
	}

	unique := map[[2]int]bool{}

	for _, v := range ropes[len(ropes)-1].visited_location {
		unique[v] = true
	}

	fmt.Println("Tail has visited:")
	fmt.Println(len(unique))

}
