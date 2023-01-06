package main

import (
	"aoc/util"
	"fmt"
	"os"
)

type points_game struct {
	loss int
	duel int
	win  int
}

type points_selection struct {
	rock    int
	paper   int
	scissor int
}

var selection_shortcuts = make(map[string]string)

func play_rps(player string, opponent string) int {
	p := points_game{0, 3, 6}
	ps := points_selection{1, 2, 3}
	player_selection := selection_shortcuts[player]
	opponent_selection := selection_shortcuts[opponent]

	if player_selection == opponent_selection {
		if player_selection == "rock" {
			return p.duel + ps.rock
		} else if player_selection == "paper" {
			return p.duel + ps.paper
		} else if player_selection == "scissor" {
			return p.duel + ps.scissor
		}
	} else if player_selection == "rock" && opponent_selection == "scissor" {
		return p.win + ps.rock
	} else if player_selection == "paper" && opponent_selection == "rock" {
		return p.win + ps.paper
	} else if player_selection == "scissor" && opponent_selection == "paper" {
		return p.win + ps.scissor
	}
	if player_selection == "rock" {
		return ps.rock
	} else if player_selection == "paper" {
		return ps.paper
	} else if player_selection == "scissor" {
		return ps.scissor
	}
	return p.loss
}

func play_rps_inverse(res string, opponent string) int {
	p := points_game{0, 3, 6}
	ps := points_selection{1, 2, 3}
	result := selection_shortcuts[res]
	opponent_selection := selection_shortcuts[opponent]

	if result == "loss" {
		if opponent_selection == "rock" {
			return ps.scissor
		} else if opponent_selection == "paper" {
			return ps.rock
		} else if opponent_selection == "scissor" {
			return ps.paper
		}
	} else if result == "win" {
		if opponent_selection == "scissor" {
			return p.win + ps.rock
		} else if opponent_selection == "rock" {
			return p.win + ps.paper
		} else if opponent_selection == "paper" {
			return p.win + ps.scissor
		}
	} else {
		if opponent_selection == "rock" {
			return p.duel + ps.rock
		} else if opponent_selection == "paper" {
			return p.duel + ps.paper
		} else if opponent_selection == "scissor" {
			return p.duel + ps.scissor
		}
	}
	return -1
}

func main() {
	data := util.Data(1, "\n")


    if len(os.Args) < 3 {
		panic("Please provide puzzle part index as parameter `go run main.go -- 1`")
    }

	part := os.Args[2]

	if part == "1" {
		selection_shortcuts["X"] = "rock"
		selection_shortcuts["Y"] = "paper"
		selection_shortcuts["Z"] = "scissor"
	} else if part == "2" {
		selection_shortcuts["X"] = "loss"
		selection_shortcuts["Y"] = "draw"
		selection_shortcuts["Z"] = "win"
	} 

	selection_shortcuts["A"] = "rock"
	selection_shortcuts["B"] = "paper"
	selection_shortcuts["C"] = "scissor"

	points := 0

	for _, val := range data {
		if len(val) == 3 {
			first_parameter := string(val[0])
			second_parameter := string(val[len(val)-1])

			if part == "1" {
				points += play_rps(second_parameter, first_parameter)
			} else if part == "2" {
				points += play_rps_inverse(second_parameter, first_parameter)
			}
		} else {
			continue
		}
	}

	fmt.Println(points)
}
