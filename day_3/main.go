package main

import (
	"aoc/util"
	"bytes"
	"fmt"
	"os"
	"unicode"
)

var lower_case_letters = []byte("abcdefghijklmnopqrstuvwxyz")

func get_char_value(char rune) int {

	if unicode.IsUpper(char) {
		position := bytes.IndexRune(lower_case_letters, unicode.ToLower(char))
		return position + 27
	}

	position := bytes.IndexRune(lower_case_letters, char)
	return position + 1
}

// refactor -> find_common_value_from_n_elements ?
func find_common_value(compartment_1 string, compartment_2 string) int {
	for _, char_1 := range compartment_1 {
		for _, char_2 := range compartment_2 {
			if char_1 == char_2 {
				return get_char_value(char_1)
			}
		}
	}
	return 0
}

func find_common_value_from_three_el(compartment_1 string, compartment_2 string, compartment_3 string) int {
	for _, char_1 := range compartment_1 {
		for _, char_2 := range compartment_2 {
			for _, char_3 := range compartment_3 {
				if char_1 == char_2 && char_1 == char_3 {
					return get_char_value(char_1)
				}
			}
		}
	}
	return 0
}

func main() {
	data := util.Data(1, "\n")
	data = data[:len(data)-1]

	if len(os.Args) < 3 {
		panic("Please provide puzzle part index as parameter `go run main.go -- 1`")
	}

	part := os.Args[2]

	sum := 0

	if part == "1" {
		for _, v := range data {
			compartment_1, compartment_2 := v[0:len(v)/2], v[len(v)/2:]
			sum += find_common_value(compartment_1, compartment_2)
		}
	} else {
		// all in the group of three elves has ONE common item type
		// and at most two of the Elves will be carrying any other item type

		if len(data)%3 != 0 {
			panic("Data can't be divided to three sections")
		}

		for i := 0; i < len(data); i += 3 {
			var section []string
			if i > len(data)-3 {
				section = data[i:]
			} else {
				section = data[i : i+3]
			}
			sum += find_common_value_from_three_el(section[0], section[1], section[2])
		}
	}

	fmt.Println(sum)

}
