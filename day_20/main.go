package main

import (
	"aoc/util"
	"errors"
	"fmt"
	"os"
	"strconv"
)

// handle also negative values (which go doesn't do by default)
func mod(a, b int) int {
	return (a%b + b) % b
}

// https://freshman.tech/snippets/go/concatenate-slices/
func concatMultipleSlices[T any](slices [][]T) []T {
	var totalLen int

	for _, s := range slices {
		totalLen += len(s)
	}

	result := make([]T, totalLen)

	var i int

	for _, s := range slices {
		i += copy(result[i:], s)
	}

	return result
}

type Value struct {
	index int
	value int
}

type Arrangement []Value

func (arr Arrangement) String() string {
	var values string
	for _, val := range arr {
		values += strconv.Itoa(val.value) + " "
	}

	return fmt.Sprintf("%s", values)
}

func (arr Arrangement) findCurrentIndex(val Value) (int, error) {
	for i, v := range arr {
		if val == v {
			return i, nil
		}
	}
	return 0, errors.New("Cannot find index")
}

func (arr Arrangement) findIndexOfZero() (int, error) {
	for i, v := range arr {
		if v.value == 0 {
			return i, nil
		}
	}
	return 0, errors.New("Cannot find zero")
}

func (arr Arrangement) sumThreeValues() int {
	index, err := arr.findIndexOfZero()
	if err != nil {
		panic(err)
	}
	first := arr[(1000+index)%len(arr)]
	second := arr[(2000+index)%len(arr)]
	third := arr[(3000+index)%len(arr)]
	return first.value + second.value + third.value
}

func (arr *Arrangement) moveValue(val Value) {

	// exit early
	if val.value == 0 {
		return
	}

	index, err := arr.findCurrentIndex(val)
	if err != nil {
		panic(err)
	}

	reIndex := mod(index+val.value, len(*arr)-1)
	if reIndex == index {
		return
	}

	if (index+val.value)%(len(*arr)-1) == 0 {
		// if moving right in the arrangement -> warp to the start of the list
		// and if moving left -> warp to the end of the list
		part1 := make(Arrangement, index)
		_ = copy(part1, (*arr)[:index])
		part2 := make(Arrangement, len(*arr)-(index+1))
		_ = copy(part2, (*arr)[index+1:])

		if val.value > 0 {
			// reIndex = 0
			*arr = concatMultipleSlices([][]Value{[]Value{val}, part1, part2})
		} else {
			// reIndex = len(*arr)
			*arr = concatMultipleSlices([][]Value{part1, part2, []Value{val}})
		}
	} else {

		if reIndex > index {
			start := make(Arrangement, index)
			_ = copy(start, (*arr)[:index])

			part1 := make(Arrangement, reIndex-index)
			_ = copy(part1, (*arr)[index+1:reIndex+1])
			part2 := make(Arrangement, len(*arr)-(reIndex+1))
			_ = copy(part2, (*arr)[reIndex+1:])

			*arr = concatMultipleSlices([][]Value{start, part1, []Value{val}, part2})
		} else {
			part1 := make(Arrangement, reIndex)
			_ = copy(part1, (*arr)[:reIndex])

			part2 := make(Arrangement, index-reIndex)
			_ = copy(part2, (*arr)[reIndex:index])

			end := make(Arrangement, len(*arr)-(index+1))
			_ = copy(end, (*arr)[index+1:])

			*arr = concatMultipleSlices([][]Value{part1, []Value{val}, part2, end})
		}

	}
}

func main() {
	data := util.Data(1, "\n")
	data = data[:len(data)-1]

	if len(os.Args) < 3 {
		panic("Please provide puzzle part index as parameter `go run main.go -- 1`")
	}

	part := os.Args[2]

	fmt.Printf("Run part: %s\n", part)

	initialArrangement := make(Arrangement, len(data))
	decryptKey := 811589153

	for i, val := range data {
		intVal, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		initialArrangement[i] = Value{i, intVal * decryptKey}
	}

	currentArrangement := make(Arrangement, len(data))
	_ = copy(currentArrangement, initialArrangement)

	for i := 0; i < 10; i++ {
		for _, val := range initialArrangement {
			currentArrangement.moveValue(val)
		}

		if i == 0 {
			fmt.Printf("Result part 1: %d\n", currentArrangement.sumThreeValues())
		}
	}

	fmt.Printf("Result part 2: %d\n", currentArrangement.sumThreeValues())
}
