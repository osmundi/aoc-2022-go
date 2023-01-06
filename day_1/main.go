package main

import (
	"aoc/util"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	data := util.Data(1, "\n")

	sum_of_one_set := 0
	var calories []int

	for _, v := range data {
		if v == "" {
			calories = append(calories, sum_of_one_set)
			sum_of_one_set = 0
		} else {
			v_int, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			sum_of_one_set += v_int
		}
	}

	sort.Sort(sort.IntSlice(calories))

	// part 1
	fmt.Println(calories[len(calories)-1])

	// part 2
	result := 0
	for _, v := range calories[len(calories)-3:] {
		result += v
	}
	fmt.Println(result)
}
