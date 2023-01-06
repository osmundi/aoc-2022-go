package main

import (
	"aoc/util"
	"fmt"
	"os"
	"strings"
	//"unicode"
)

func contains_one_another(ranges []string) bool {

	var a_left, b_left, a_right, b_right int

	_, err := fmt.Sscanf(ranges[0], "%d-%d", &a_left, &a_right)

	if err != nil {
		panic(err)
	}

	_, err_b := fmt.Sscanf(ranges[1], "%d-%d", &b_left, &b_right)

	if err_b != nil {
		panic(err_b)
	}

	if a_left <= b_left && a_right >= b_right {
		return true
	} else if a_left >= b_left && a_right <= b_right {
		return true
	}
	return false
}

func overlaps(ranges []string) bool {

	var a_left, b_left, a_right, b_right int

	_, err := fmt.Sscanf(ranges[0], "%d-%d", &a_left, &a_right)

	if err != nil {
		panic(err)
	}

	_, err_b := fmt.Sscanf(ranges[1], "%d-%d", &b_left, &b_right)

	if err_b != nil {
		panic(err_b)
	}

	if a_left <= b_left && a_right >= b_left {
		return true
	} else if a_left <= b_right && a_right >= b_right {
		return true
	}
	return false

    // 764 too low
}

func main() {
	data := util.Data(1, "\n")
	data = data[:len(data)-1]

	if len(os.Args) < 3 {
		panic("Please provide puzzle part index as parameter `go run main.go -- 1`")
	}

	part := os.Args[2]

	fmt.Println(part)
	fmt.Println(data[0])

	sum := 0

	for _, section_ranges := range data {

		ranges := strings.Split(section_ranges, ",")


        if part == "1" {
            if contains_one_another(ranges) {
                sum++
            }
        } else {
            if overlaps(ranges) || contains_one_another(ranges) {
                sum++
            }
        }
	}

	fmt.Println(sum)

	//    13-23,14-24
	//    13-65,13-64
}
