package main

import (
	"aoc/util"
	"fmt"
	"math"
	"strings"
)

var multiples = []int{-2, -1, 0, 1, 2}
var limits = make([]int, 20)

func snafuToDecimal(snafu string) int {
	snafuRunes := []rune(snafu)
	decimal := 0
	for i, val := range util.ReverseArray(snafuRunes) {
		if val == '=' {
			decimal += -2 * int(math.Pow(5, float64(i)))
		} else if val == '-' {
			decimal += -1 * int(math.Pow(5, float64(i)))
		} else if val == '0' {
			continue
		} else if val == '1' {
			decimal += 1 * int(math.Pow(5, float64(i)))
		} else if val == '2' {
			decimal += 2 * int(math.Pow(5, float64(i)))
		}
	}
	return decimal
}

func decimalToSnafu(val int, history []rune) {
	if val < 0 {
		for exp, limit := range limits {

			limit = limit * -1
			currentPower := int(math.Pow(5, float64(exp))) * -1
			if val == currentPower {
				history[exp] = '-'
				break
			}

			if val == 2*currentPower {
				history[exp] = '='
				break
			}

			if val < limit {
				continue
			} else {
				if currentPower-limits[exp-1] < val {
					history[exp] = '-'
					decimalToSnafu(val-1*currentPower, history)
				} else {
					history[exp] = '='
					decimalToSnafu(val-2*currentPower, history)
				}
				break
			}
		}

	} else {

		for exp, limit := range limits {
			currentPower := int(math.Pow(5, float64(exp)))
			if val == currentPower {
				history[exp] = '1'
				break
			}

			if val == 2*currentPower {
				history[exp] = '2'
				break
			}

			if int(math.Abs(float64(val))) > limit {
				continue
			} else {
				if currentPower+limits[exp-1] < val {
					history[exp] = '2'
					decimalToSnafu(val-2*currentPower, history)
				} else {
					history[exp] = '1'
					decimalToSnafu(val-1*currentPower, history)
				}
				break
			}

		}
	}
}

func Sum(ints []int) int {
	sum := 0
	for _, val := range ints {
		sum += val
	}
	return sum
}

func main() {
	data := util.Data(1, "\n")

	fmt.Println(len(data))

	decimals := make([]int, len(data)-1)

	for i, row := range data {
		if len(row) == 0 {
			continue
		}
		decimals[i] = snafuToDecimal(row)
	}

	sum := Sum(decimals)
	fmt.Println(sum)

    // init limits for every power
	for i := 0; i < 20; i++ {
		if i == 0 {
			limits[i] = 2 * int(math.Pow(5, float64(i)))
		} else {
			limits[i] = 2*int(math.Pow(5, float64(i))) + limits[i-1]
		}
	}

	snafu := make([]rune, 20)
	for i := range snafu {
		snafu[i] = '0'
	}

	decimalToSnafu(sum, snafu)
    snafuString := strings.TrimRight(string(snafu), "0")
    fmt.Println(util.ReverseString(snafuString))
}
