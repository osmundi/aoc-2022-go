package main

import (
	"aoc/util"
	"fmt"
	"os"
)

func add_pending_value(cycle int, pending_values map[int]int) int {
	rv := pending_values[cycle]
	delete(pending_values, cycle)
	return rv
}

func main() {
	data := util.Data(1, "\n")
	data = data[:len(data)-1]

	if len(os.Args) < 3 {
		panic("Please provide puzzle part index as parameter `go run main.go -- 1`")
	}

	part := os.Args[2]

	fmt.Printf("Run part: %s\n", part)

	var val string
	var command string
	var add int

	cycle := 1
	register := 1 // sprite start position
	const (
		lit           = '#'
		dark          = '.'
		sprite_length = 3
	)
	pending_values := make(map[int]int)
	sum := 0
	display := [241]rune{}

	for {
		// start of cycle
		register += add_pending_value(cycle, pending_values)

		// during cycle
		if part == "1" {
			// calculate signal strenghts
			if (cycle-20)%40 == 0 {
				sum += cycle * register
			}
		} else {
			// draw display
			if cycle%40 >= register && cycle%40 < (register+sprite_length) {
				display[cycle-1] = lit
			} else {
				display[cycle-1] = dark
			}
		}

		if len(pending_values) > 0 {
			cycle++
			continue
		}

		if len(data) == 0 {
			break
		}

		// pop row from data and scan values
		val, data = data[0], data[1:]
		fmt.Sscanf(val, "%s %d", &command, &add)

		switch command {
		case "noop":
			cycle++
			continue
		case "addx":
			pending_values[cycle+2] = add
		}
		cycle++

	}

	if part == "1" {
		fmt.Printf("--------------\nSum of signal strengts: %d\n", sum)
	} else {

		for i := 0; i < 6; i++ {
			fmt.Println(string(display[i*40 : (i+1)*40]))
		}
	}
}
