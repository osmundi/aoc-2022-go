package main

import (
	"aoc/util"
	"fmt"
	"hash/fnv"
	"os"
	"strconv"
	"strings"
)

var DP = make(map[uint32]int)

func hashState(res resources, robots Robots, time int) uint32 {
	state := []byte(res.Stringify() + robots.Stringify() + strconv.Itoa(time))
	h := fnv.New32a()
	h.Write(state)
	return h.Sum32()
}

type resources struct {
	ore      int
	clay     int
	obsidian int
	geode    int
}

func (r1 resources) Add(r2 resources) resources {
	return resources{
		r1.ore + r2.ore,
		r1.clay + r2.clay,
		r1.obsidian + r2.obsidian,
		r1.geode + r2.geode,
	}
}

func (res resources) Stringify() string {
	return strings.Join([]string{strconv.Itoa(res.ore), strconv.Itoa(res.clay), strconv.Itoa(res.obsidian), strconv.Itoa(res.geode)}, ",")
}

type RobotBluePrints interface {
	GenerateRobot(resources) (resources, bool)
	CollectResource(*resources)
	Dismiss()
	GetCount() int
	GetLimit() int
}

type Robots []RobotBluePrints

func (robots Robots) CollectResources() resources {
	res := resources{0, 0, 0, 0}
	for _, r := range robots {
		r.CollectResource(&res)
	}
	return res
}

func (robots Robots) Stringify() string {
	var cnt string
	for i := 0; i < len(robots); i++ {
		cnt += strconv.Itoa(robots[i].GetCount())
	}
	return cnt
}

type clayRobot struct {
	oreCost int
	count   int
	limit   int
}

func (r *clayRobot) Dismiss() {
	r.count--
}

func (r clayRobot) GetCount() int {
	return r.count
}
func (r clayRobot) GetLimit() int {
	return r.limit
}

func (robot *clayRobot) GenerateRobot(r resources) (resources, bool) {
	if r.ore >= robot.oreCost {
		r.ore -= robot.oreCost
		robot.count++
		return r, true
	}
	return r, false
}

func (r *clayRobot) String() string {
	return fmt.Sprintf("Clayrobots: %v (%v)", r.count, r.oreCost)
}

func (robot clayRobot) CollectResource(r *resources) {
	r.clay += robot.count * 1
}

type oreRobot struct {
	oreCost int
	count   int
	limit   int
}

func (r *oreRobot) Dismiss() {
	r.count--
}

func (r oreRobot) GetCount() int {
	return r.count
}

func (r oreRobot) GetLimit() int {
	return r.limit
}

func (robot *oreRobot) GenerateRobot(r resources) (resources, bool) {
	if r.ore >= robot.oreCost {
		r.ore -= robot.oreCost
		robot.count++
		return r, true
	}
	return r, false
}

func (robot oreRobot) CollectResource(r *resources) {
	r.ore += robot.count * 1
}

func (r *oreRobot) String() string {
	return fmt.Sprintf("Orerobots: %v (%v)", r.count, r.oreCost)
}

type obsidianRobot struct {
	oreCost  int
	clayCost int
	count    int
	limit    int
}

func (r *obsidianRobot) Dismiss() {
	r.count--
}

func (r obsidianRobot) GetCount() int {
	return r.count
}

func (r obsidianRobot) GetLimit() int {
	return r.limit
}

func (robot *obsidianRobot) GenerateRobot(r resources) (resources, bool) {
	if r.ore >= robot.oreCost && r.clay >= robot.clayCost {
		r.ore -= robot.oreCost
		r.clay -= robot.clayCost
		robot.count++
		return r, true
	}
	return r, false
}

func (robot obsidianRobot) CollectResource(r *resources) {
	r.obsidian += robot.count * 1
}

func (r *obsidianRobot) String() string {
	return fmt.Sprintf("Obsidianrobots: %v (%v, %v)", r.count, r.oreCost, r.clayCost)
}

type geodeRobot struct {
	oreCost      int
	obsidianCost int
	count        int
	limit        int
}

func (r *geodeRobot) Dismiss() {
	r.count--
}

func (r geodeRobot) GetCount() int {
	return r.count
}

func (r geodeRobot) GetLimit() int {
	return r.limit
}

func (robot *geodeRobot) GenerateRobot(r resources) (resources, bool) {
	if r.ore >= robot.oreCost && r.obsidian >= robot.obsidianCost {
		r.ore -= robot.oreCost
		r.obsidian -= robot.obsidianCost
		robot.count++
		return r, true
	}
	return r, false
}

func (robot geodeRobot) CollectResource(r *resources) {
	r.geode += robot.count * 1
}

func (r *geodeRobot) String() string {
	return fmt.Sprintf("GeodeRobots: %v (%v, %v)", r.count, r.oreCost, r.obsidianCost)
}

func grind(res resources, robots Robots, time int) int {
	if time == 0 {
		return res.geode
	}

	// Collect
	addResources := robots.CollectResources()

	// Memoize state
	key := hashState(res, robots, time)
	val, ok := DP[key]
	if ok {
		return val
	}

	geodes := grind(res.Add(addResources), robots, time-1)

	for _, robot := range robots {
		// no need to build unnecessary ore/clay/obsidian robots
		if robot.GetCount() >= robot.GetLimit() {
			continue
		}

		newRes, done := robot.GenerateRobot(res)
		if done {
			geodsWhenGeneratedNewBot := grind(newRes.Add(addResources), robots, time-1)
			if geodsWhenGeneratedNewBot > geodes {
				geodes = geodsWhenGeneratedNewBot
			}
			// we want to get back to the old state before new robot was generated
			robot.Dismiss()

            // break if robot == geodeRobot
            if robot.GetLimit() == 1000 {
                break
            }

		}
	}

	DP[key] = geodes

	return geodes
}

func main() {
	data := util.Data(1, "\n")
	data = data[:len(data)-1]

	if len(os.Args) < 3 {
		panic("Please provide puzzle part index as parameter `go run main.go -- 1`")
	}

	part := os.Args[2]

	fmt.Printf("Run part: %s\n", part)

	var blueprintIndex, oreRobotOreCost, clayRobotOreCost, obsidianRobotOreCost, obsidianRobotClayCost, geodeRobotOreCost, geodeRobotObsidianCost int
	grindTime := 24
	geodes := 0
	qualitySum := 0

	for i, val := range data {
		fmt.Println(val)

		if part == "2" && i == 3 {
			break
		}

		_, err := fmt.Sscanf(val, "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.", &blueprintIndex, &oreRobotOreCost, &clayRobotOreCost, &obsidianRobotOreCost, &obsidianRobotClayCost, &geodeRobotOreCost, &geodeRobotObsidianCost)

		if err != nil {
			panic(err)
		}

		// Dont create more ore than is needed
		oreLimit, _ := util.Max([]int{geodeRobotOreCost, obsidianRobotOreCost, clayRobotOreCost})

		robots := Robots{
			&geodeRobot{geodeRobotOreCost, geodeRobotObsidianCost, 0, 1000}, // cant be too many of these
			&obsidianRobot{obsidianRobotOreCost, obsidianRobotClayCost, 0, geodeRobotObsidianCost},
			&clayRobot{clayRobotOreCost, 0, obsidianRobotClayCost},
			&oreRobot{oreRobotOreCost, 1, oreLimit},
		}

		fmt.Println(robots)

		geodes = grind(resources{ore: 0, clay: 0, obsidian: 0, geode: 0}, robots, grindTime)

		fmt.Printf("Geodes: %d\n", geodes)

		qualitySum += blueprintIndex * geodes

		// clear cache
		DP = make(map[uint32]int)
	}

	if part == "1" {
		fmt.Println(qualitySum)
	}
}
