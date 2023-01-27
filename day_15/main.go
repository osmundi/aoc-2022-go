package main

import (
	"aoc/util"
	"fmt"
	"math"
	"os"
	"sort"
)

type Coords struct {
	x int
	y int
}

func (coord *Coords) IncrementY() {
	coord.y++
}

type Beacon struct {
	location Coords
}

type Sensor struct {
	location Coords
	beacon   Beacon
	coverage int
}

func (sensor Sensor) GetManhattanDistance(location Coords) int {
	return int(math.Abs(float64(sensor.location.x)-float64(location.x)) + math.Abs(float64(sensor.location.y)-float64(location.y)))
}

func (sensor *Sensor) SetCoverage() {
	sensor.coverage = sensor.GetManhattanDistance(sensor.beacon.location)
}

func (sensor Sensor) GetSensorHorizontalCoverageAt(y int, limits ...int) []Coords {
	// fillCoords in range {8-(ManhattanDistance - ABS(y - sensorY))}...{8+9} (-1...17)
	md := sensor.GetManhattanDistance(sensor.beacon.location)
	var coords []Coords
	leftLimit := sensor.location.x - (md - int(math.Abs(float64(sensor.location.y)-float64(y))))
	rightLimit := sensor.location.x + (md - int(math.Abs(float64(sensor.location.y)-float64(y))))

	if len(limits) == 2 {
		if leftLimit < limits[0] {
			leftLimit = limits[0]
		}
		if rightLimit > limits[1] {
			rightLimit = limits[1]
		}
	}

	for x := leftLimit; x <= rightLimit; x++ {
		coords = append(coords, Coords{x: x, y: y})
	}

	return coords
}

func (sensor Sensor) FindNeighbours(sensors *[]Sensor) []Sensor {
	var distance int
	closestSensors := make(map[int]Sensor, len(*sensors))
	for _, neighbour := range *sensors {
		distance = sensor.GetManhattanDistance(neighbour.location)

		if distance != 0 {
			closestSensors[distance] = neighbour
		}
	}

	keys := make([]int, 0, len(closestSensors))
	for k := range closestSensors {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	var rv []Sensor
	for i, k := range keys {
		if i == 3 {
			break
		}
		rv = append(rv, closestSensors[k])
	}
	return rv
}

func getB(point Coords, m int) int {
	// y coordinate is upside down
	return -1 * (point.y - m*point.x)
}

func getInterception(m1 int, m2 int, b1 int, b2 int) Coords {
	x := (b1 - b2) / (m1 - m2)
	y := b1 - m1*x
	return Coords{x, y}
}

func (sensor Sensor) FindIntersections(sensors []*Sensor) []Coords {
	southCorner := Coords{sensor.location.x, -1 * (sensor.location.y + sensor.coverage)}

	b1 := getB(southCorner, 1)
	m1 := 1
	m2 := -1

	var interceptions []Coords
	var interception Coords

	for _, neighbour := range sensors {

		if sensor.location.x >= neighbour.location.x {
			continue
		}

		southCorner = Coords{neighbour.location.x, -1 * (neighbour.location.y + neighbour.coverage)}
		b2 := getB(southCorner, -1)
		interception = getInterception(m1, m2, b1, b2)
		interceptions = append(interceptions, interception)
	}

	return interceptions
}

func UniqueCoords(coords []Coords) []Coords {
	keys := make(map[Coords]bool)
	list := []Coords{}
	for _, entry := range coords {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func tuningFrequency(point Coords) int {
	return (point.x * 4000000) + point.y
}

func main() {
	data := util.Data(1, "\n")
	data = data[:len(data)-1]

	if len(os.Args) < 3 {
		panic("Please provide puzzle part index as parameter `go run main.go -- 1`")
	}

	part := os.Args[2]

	fmt.Printf("Run part: %s\n", part)

	targetRowY := 2000000
	targetRowCoords := make([]Coords, 0)
	var sensors []*Sensor
	var sensorX, sensorY, beaconX, beaconY int

	for _, val := range data {
		fmt.Sscanf(val, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensorX, &sensorY, &beaconX, &beaconY)
		sensors = append(sensors, &Sensor{location: Coords{x: sensorX, y: sensorY}, beacon: Beacon{location: Coords{x: beaconX, y: beaconY}}})
	}

	fmt.Println("PART 1")
	for _, sensor := range sensors {
		targetRowCoords = append(targetRowCoords, sensor.GetSensorHorizontalCoverageAt(targetRowY)...)
	}
	fmt.Println(len(UniqueCoords(targetRowCoords)))

	fmt.Println("PART 2")

	for _, sensor := range sensors {
		sensor.SetCoverage()
	}

	var searchPoints []Coords
	for _, sensor := range sensors {
		searchPoints = append(searchPoints, sensor.FindIntersections(sensors)...)
	}

	fmt.Println("Amount of points to be searched")
	fmt.Println(len(searchPoints))

	var found bool
	signalLimit := 4000000
	for _, point := range searchPoints {
		// look one step down the intersection
		point.IncrementY()

		if point.x < 0 || point.x > signalLimit || point.y < 0 || point.y > signalLimit {
			continue
		}

		found = false
		for _, sensor := range sensors {
			if sensor.GetManhattanDistance(Coords{point.x, point.y}) <= sensor.coverage {
				found = true
				break
			}
		}

		if found {
			continue
		} else {
			fmt.Println("-----FOUND----")
			fmt.Println(point)
			fmt.Println(tuningFrequency(point))
			break
		}
	}

	if !found {
		fmt.Println("Point not found")
	}
}
