package main

import (
	"aoc/util"
	"container/heap"
	"fmt"
	"math"
	"strings"
)

var lower_case_letters = string("abcdefghijklmnopqrstuvwxyz")

type location struct {
	x int
	y int
}

type Node struct {
	location   location
	height     string
	neighbours []*Node
	came_from  *Node
	priority   int // f(n) = g(n) + h(n)
	index      int
}

func (node *Node) AddNeighbour(neighbour *Node) {
	node.neighbours = append(node.neighbours, neighbour)
}

func (node *Node) ManhattanDistance(goal *Node) int {
	return int(math.Abs(float64(node.location.x)-float64(goal.location.x)) + math.Abs(float64(node.location.y)-float64(goal.location.y)))
}

func (node *Node) CameFrom(parent *Node) {
	node.came_from = parent
}

func (node *Node) CanMove(neighbour *Node) bool {
	move_from := strings.Index(lower_case_letters, node.height)
	move_to := strings.Index(lower_case_letters, neighbour.height)
	if move_from == len(lower_case_letters)+1 {
		// moving from z
		return true
	}
	if move_from+1 >= move_to {
		// hiker can move only one step up (e.g. c->d)
		return true
	}
	return false
}

func NodeInOpenSet(node *Node, list []*Node) bool {
	for _, n := range list {
		if n == node {
			return true
		}
	}
	return false
}

// Copied from: https://pkg.go.dev/github.com/jupp0r/go-priority-queue

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest priority so we use less than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Node)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Node, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}

// https://en.wikipedia.org/wiki/A*_search_algorithm#Pseudocode
func reconstruct_path(current *Node, total_path []*Node) []*Node {
	if current.came_from == nil {
		return total_path
	} else {
		total_path = append(total_path, current)
		return reconstruct_path(current.came_from, total_path)
	}
}

// A* -search algorithm build on top of the pseudocode in wikipedia
func AStar(start *Node, goal *Node) []*Node {
	// The set of discovered nodes that may need to be (re-)expanded.
	// Initially, only the start node is known.
	openSet := make(PriorityQueue, 1)
	openSet[0] = start
	heap.Init(&openSet)

	// For node n, gScore[n] is the cost of the cheapest path from start to n currently known.
	gScore := make(map[*Node]int)
	gScore[start] = 0

	// For node n, fScore[n] := gScore[n] + h(n). fScore[n] represents our current best guess as to
	// how cheap a path could be from start to finish if it goes through n.
	// fScore == Node.priority
	fScore := make(map[*Node]int)
	fScore[start] = start.ManhattanDistance(goal)

	for {
		if len(openSet) == 0 {
			break
		}

		// current := the node in openSet having the lowest fScore[] value
		current := heap.Pop(&openSet).(*Node)

		if current == goal {
			return reconstruct_path(current, nil)
		}

		for _, neighbour := range current.neighbours {
			// d(current,neighbor) is the weight of the edge from current to neighbor
			// tentative_gScore is the distance from start to the neighbor through current
			tentative_gScore := gScore[current] + 1 // d(current, neighbour) == 1

			val, ok := gScore[neighbour]

			// This path to neighbor is better than any previous one. Record it!
			if !ok || tentative_gScore < val {
				// For node n, cameFrom[n] is the node immediately preceding it on the cheapest path from start to n currently known.
				//cameFrom := make(map[*Node]Node)
				neighbour.CameFrom(current)
				gScore[neighbour] = tentative_gScore
				fScore[neighbour] = tentative_gScore + neighbour.ManhattanDistance(goal)

				if !NodeInOpenSet(neighbour, openSet) {
					heap.Push(&openSet, neighbour)
					openSet.update(neighbour, fScore[neighbour])
				}
			}
		}
	}
	return nil
}

func draw_map(height_map []string, nodes []*Node) {
	const colorRed = "\033[0;31m"
	const colorNone = "\033[0m"
	i := 0
	color_path := false
	for y, row := range height_map {
		for x, value := range row {
			for _, node := range nodes {
				if node.location.x == x && node.location.y == y {
					color_path = true
				}
			}
			if color_path {
				fmt.Printf("%s%s%s", colorRed, string(value), colorNone)
			} else {
				fmt.Printf("%s", string(value))
			}
			i++
			if i%66 == 0 {
				fmt.Printf("\n")
			}
			color_path = false
		}
	}
}

func main() {
	data := util.Data(1, "\n")
	data = data[:len(data)-1]

	height_map := make(map[location]*Node)
	start_node := location{}
	end_node := location{}

	// init all nodes
	for y, data_row := range data {
		data_row_characters := []rune(data_row)
		for x, height := range data_row_characters {
			if string(height) == "S" {
				start_node = location{x, y}
				height_map[location{x, y}] = &(Node{location: location{x, y}, height: "a"})
			} else if string(height) == "E" {
				height_map[location{x, y}] = &(Node{location: location{x, y}, height: "z"})
				end_node = location{x, y}
			} else {
				height_map[location{x, y}] = &(Node{location: location{x, y}, height: string(height)})
			}
		}
	}

	// set neighbours for all nodes
	for y, data_row := range data {
		data_row_characters := []rune(data_row)
		for x, _ := range data_row_characters {

			val, ok := height_map[location{x + 1, y}]
			if ok {
				if height_map[location{x, y}].CanMove(val) {
					height_map[location{x, y}].AddNeighbour(val)
				}
			}
			val, ok = height_map[location{x, y + 1}]
			if ok {
				if height_map[location{x, y}].CanMove(val) {
					height_map[location{x, y}].AddNeighbour(val)
				}
			}
			val, ok = height_map[location{x - 1, y}]
			if ok {
				if height_map[location{x, y}].CanMove(val) {
					height_map[location{x, y}].AddNeighbour(val)
				}
			}
			val, ok = height_map[location{x, y - 1}]
			if ok {
				if height_map[location{x, y}].CanMove(val) {
					height_map[location{x, y}].AddNeighbour(val)
				}
			}
		}
	}

	path := AStar(height_map[start_node], height_map[end_node])

	fmt.Println("Length of the whole path:")
	fmt.Println(len(path))

	draw_map(data, path)

}
