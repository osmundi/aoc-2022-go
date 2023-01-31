package main

import (
	"aoc/util"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Direction Vectors to generate 6 neighbors of the cell
var X = []int{1, -1, 0, 0, 0, 0}
var Y = []int{0, 0, 1, -1, 0, 0}
var Z = []int{0, 0, 0, 0, 1, -1}

type Cube struct {
	x          int
	y          int
	z          int
	neighbours []*Cube
}

type Location struct {
	x int
	y int
	z int
}

func (a *Cube) Compare(b *Cube) bool {

	if a.x == b.x && a.y == b.y && math.Abs(float64(a.z)-float64(b.z)) == 1 {
		return true
	} else if a.y == b.y && a.z == b.z && math.Abs(float64(a.x)-float64(b.x)) == 1 {
		return true
	} else if a.x == b.x && a.z == b.z && math.Abs(float64(a.y)-float64(b.y)) == 1 {
		return true
	}

	return false
}

func FindAirPocket(grid map[Location]bool, x int, y int, z int, edge [2]int) bool {

	if x <= edge[0] || y <= edge[0] || z <= edge[0] {
		return false
	}
	if x >= edge[1] || y >= edge[1] || z >= edge[1] {
		return false
	}

	for k := 0; k < 6; k++ {
		dx := x + X[k]
		dy := y + Y[k]
		dz := z + Z[k]

		if grid[Location{dx, dy, dz}] {
			continue
		} else {
			grid[Location{dx, dy, dz}] = true
		}
		return FindAirPocket(grid, dx, dy, dz, edge)
	}
	return true
}

func generateGrid(limits [2]int) map[Location]bool {
	low := limits[0]
	high := limits[1]
	var grid = make(map[Location]bool)
	for x := low; x <= high; x++ {
		for y := low; y <= high; y++ {
			for z := low; z <= high; z++ {
				grid[Location{x, y, z}] = false
			}
		}
	}
	return grid
}

func mapClone(src map[Location]bool) map[Location]bool {
	dst := make(map[Location]bool, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

func main() {
	data := util.Data(1, "\n",true)
	data = data[:len(data)-1]

	if len(os.Args) < 3 {
		panic("Please provide puzzle part index as parameter `go run main.go -- 1`")
	}

	part := os.Args[2]

	fmt.Printf("Run part: %s\n", part)

	var grid map[Location]bool

	var limits [2]int

	if part == "1" {
		limits = [2]int{0, 13}
	} else {
		limits = [2]int{-2, 21}
	}
	grid = generateGrid(limits)

	cubes := []*Cube{}

	for _, val := range data {
		coords := strings.Split(val, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])

		cubes = append(cubes, &Cube{x: x, y: y, z: z})

		grid[Location{x, y, z}] = true
	}

	// PART 1
	closedSides := 0
	for i, cube := range cubes {
		for j := i + 1; j < len(cubes); j++ {
			if cube.Compare(cubes[j]) {
				cube.neighbours = append(cube.neighbours, cubes[j])
				cubes[j].neighbours = append(cubes[j].neighbours, cube)
			}
		}
	}

	for _, cube := range cubes {
		closedSides += len(cube.neighbours)
	}
	fmt.Println(len(cubes)*6 - closedSides)

	// PART 2
	// launch a dfs to find edges outside the cube
	inPocket := 0
	for _, cube := range cubes {
		// sum all sides that might have air pockets
		for k := 0; k < 6; k++ {
			dx := cube.x + X[k]
			dy := cube.y + Y[k]
			dz := cube.z + Z[k]

			if grid[Location{dx, dy, dz}] {
				continue
			}

			if FindAirPocket(mapClone(grid), dx, dy, dz, limits) {
				inPocket++
			}

		}

	}

	fmt.Println(len(cubes)*6 - closedSides - inPocket)

}
