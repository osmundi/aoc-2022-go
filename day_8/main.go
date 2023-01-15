package main

import (
	"aoc/util"
	"fmt"
	"os"
)

func calculate_scenic_score(trees *[][]int, i int, j int) int {

    left_sum, right_sum, up_sum, down_sum := 0,0,0,0

    // view to left
	for column := j-1; column >= 0; column-- {
		if (*trees)[i][j] > (*trees)[i][column] {
            left_sum++
		} else {
            left_sum++
            break
        }
	}

    // view to right
	for column := j+1; column <= 98; column++ {
		if (*trees)[i][j] > (*trees)[i][column] {
            right_sum++
		} else {
            right_sum++
            break
        }
	}

    // view up
	for row := i-1; row >= 0; row-- {
		if (*trees)[i][j] > (*trees)[row][j] {
            up_sum++
		} else {
            up_sum++
            break
        }
	}

    // view down
	for row := i+1; row <= 98; row++ {
		if (*trees)[i][j] > (*trees)[row][j] {
            down_sum++
		} else {
            down_sum++
            break
        }
	}
    return left_sum*right_sum*up_sum*down_sum
}


func hidden_from_left(trees *[][]int, i int, j int) bool {
	for column := 0; column < j; column++ {
		if (*trees)[i][j] <= (*trees)[i][column] {
			return true
		}
	}
	return false
}

func hidden_from_right(trees *[][]int, i int, j int) bool {
	for column := 98; column > j; column-- {
		if (*trees)[i][j] <= (*trees)[i][column] {
			return true
		}
	}
	return false
}

func hidden_from_top(trees *[][]int, i int, j int) bool {
	for row := 0; row < i; row++ {
		if (*trees)[i][j] <= (*trees)[row][j] {
			return true
		}
	}
	return false
}

func hidden_from_bottom(trees *[][]int, i int, j int) bool {
	for row := 98; row > i; row-- {
		if (*trees)[i][j] <= (*trees)[row][j] {
			return true
		}
	}
	return false
}

func runes_to_ints(runes []rune) []int {
	var ints []int
	for _, r := range runes {
		ints = append(ints, int(r))
	}
	return ints
}

func main() {
	data := util.Data(1, "\n")
	data = data[:len(data)-1]

	if len(os.Args) < 3 {
		panic("Please provide puzzle part index as parameter `go run main.go -- 1`")
	}

	part := os.Args[2]

	fmt.Printf("Run part: %s\n", part)

	// init 2d matrix
	trees := make([][]int, len(data))
	for i, val := range data {
		trees[i] = runes_to_ints([]rune(val))
	}

	if part == "1" {
		sum := 0
		_ = sum

		for i := 0; i < len(trees); i++ {
			for j := 0; j < len(trees[i]); j++ {
				if hidden_from_left(&trees, i, j) && hidden_from_right(&trees, i, j) && hidden_from_bottom(&trees, i, j) && hidden_from_top(&trees, i, j) {
					sum += 1
				}
			}
		}
		fmt.Printf("Trees visible: %d\n", 99*99-sum)
	} else {
        score := 0
        top_score := 0

		for i := 0; i < len(trees); i++ {
			for j := 0; j < len(trees[i]); j++ {
                score = calculate_scenic_score(&trees, i, j)
                if score > top_score {
                    top_score = score
                }
			}
		}
        // 129 too low
        fmt.Println(top_score)
	}

}
