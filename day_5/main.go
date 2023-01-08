package main

import (
	"aoc/util"
	"fmt"
	"os"
	"strings"
)

func main() {
	data := util.Data(1, "\n")
	data = data[:len(data)-1]

	if len(os.Args) < 3 {
		panic("Please provide puzzle part index as parameter `go run main.go -- 1`")
	}

	part := os.Args[2]

	fmt.Printf("Run part: %s\n", part)

	stacks_len := int((len(data[0]) + 1) / 4)
	var stacks = make([][]rune, stacks_len)
	var move_n, move_from, move_to int

	for _, val := range data {
		// move containers
		if strings.HasPrefix(val, "move") {
			_, err := fmt.Sscanf(val, "move %d from %d to %d",
				&move_n, &move_from, &move_to)

			if err != nil {
				panic(err)
			}

			stacks[move_to-1] = append(
				stacks[move_to-1],
				util.ReverseArray(stacks[move_from-1][len(stacks[move_from-1])-move_n:])...,
			)
			stacks[move_from-1] = stacks[move_from-1][:len(stacks[move_from-1])-move_n]

		} else {
			// generate start position
			for stack_index, letter_index := 0, 1; letter_index < len(val); letter_index, stack_index = letter_index+4, stack_index+1 {
				if val[letter_index] != 32 { // space
					// prepend to stack
					stacks[stack_index] = append(
						[]rune{util.RuneElement(val, letter_index)},
						stacks[stack_index]...,
					)
				}
			}
		}
	}

	for _, stack := range stacks {
		fmt.Printf("%v", string(stack[len(stack)-1]))
	}
}
