package main

import (
	"aoc/util"
	"fmt"
	"hash/fnv"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var flows = make(map[string]int)
var tunnels = make(map[string][]string)
var DP = make(map[uint32]int)

func hash_state(v string, t int, opened []string) uint32 {
	t_str := strconv.Itoa(t)
	opened_str := strings.Join(opened, ",")
	state := []byte(v + t_str + opened_str)
	h := fnv.New32a()
	h.Write(state)
	return h.Sum32()
}

func solve(valve string, time int, opened []string) int {
	if time == 0 {
		return 0
	}

	key := hash_state(valve, time, opened)
	val, ok := DP[key]
	if ok {
		return val
	}

	answer := 0

	if flows[valve] > 0 && !util.Contains(opened, valve) {
		opened = append(opened, valve)
		answerWithOpenedValve := (time-1)*flows[valve] + solve(valve, time-1, opened)

		if answerWithOpenedValve > answer {
			answer = answerWithOpenedValve
		}
	}

	for _, v := range tunnels[valve] {
		answerWithNewRoute := solve(v, time-1, opened)

		if answerWithNewRoute > answer {
			answer = answerWithNewRoute
		}
	}

	DP[key] = answer

	return answer

}

func main() {
	data := util.Data(1, "\n")
	data = data[:len(data)-1]

	if len(os.Args) < 3 {
		panic("Please provide puzzle part index as parameter `go run main.go -- 1`")
	}

	part := os.Args[2]

	fmt.Printf("Run part: %s\n", part)

	var valve string
	var rate int
	var valveRoutes []string

	for _, val := range data {
		pattern := regexp.MustCompile("[A-Z]{2}")
		allSubstringMatches := pattern.FindAllString(val, -1)
		valve = allSubstringMatches[0]
		valveRoutes = allSubstringMatches[1:]
		fmt.Sscanf(val, "Valve %s has flow rate=%d; tunnels lead to valves ...", &valve, &rate)
		flows[valve] = rate
		tunnels[valve] = valveRoutes
	}

	fmt.Println(flows)
	fmt.Println(tunnels)

	fmt.Println("Solve part 1 in 30 minutes")
	rv := solve("AA", 30, nil)
	fmt.Println(rv)

}
