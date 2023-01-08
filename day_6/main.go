package main

import (
	"aoc/util"
	"fmt"
	"os"
)

func find_distinct_characters(stream string, distinct_len int) int {
	var letter_buffer string
	for i := distinct_len; i < len(stream); i++ {
		letter_buffer = stream[i-distinct_len : i]
		if util.Unique(letter_buffer) {
			return i
		}
	}
	return -1
}

func main() {

	data := util.Data(1, "\n")
	data = data[:len(data)-1]

	if len(os.Args) < 3 {
		panic("Please provide puzzle part index as parameter `go run main.go -- 1`")
	}

	part := os.Args[2]

	fmt.Printf("Run part: %s\n", part)

	data_stream := data[0]

	start_of_distinct_series := -1

	if part == "1" {
		start_of_distinct_series = find_distinct_characters(data_stream, 4)
	} else {
		start_of_distinct_series = find_distinct_characters(data_stream, 14)
	}

	fmt.Println(start_of_distinct_series)

}
