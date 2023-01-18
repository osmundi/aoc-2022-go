package main

import (
	"aoc/util"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	id              int
	items           []int
	operand         string
	operator        string
	divisible_by    int
	throw_to_true   int
	throw_to_false  int
	inspected_items []int
}

func (monkey *Monkey) Throw(inspect_item int) (int, int) {
	worry_level := monkey.SetWorryLevel(inspect_item)
	worry_level = worry_level / 3
	if worry_level%monkey.divisible_by == 0 {
		return monkey.throw_to_true, worry_level
	}
	return monkey.throw_to_false, worry_level
}

func (monkey *Monkey) Catch(item int) {
	monkey.items = append(monkey.items, item)
}

func (monkey Monkey) SetWorryLevel(worry_level int) int {
	if monkey.operand == "old" {
		if monkey.operator == "+" {
			return 2 * worry_level
		} else if monkey.operator == "*" {
			return worry_level * worry_level
		}
	} else {
		val, _ := strconv.Atoi(monkey.operand)
		if monkey.operator == "+" {
			return worry_level + val
		} else if monkey.operator == "*" {
			return worry_level * val
		}
	}
	return worry_level
}

func newMonkey(id int, items []int, operand string, operator string, divisible_by int, throw_to_true int, throw_to_false int) *Monkey {
	monkey := Monkey{id: id, items: items, operand: operand, operator: operator, divisible_by: divisible_by, throw_to_true: throw_to_true, throw_to_false: throw_to_false}
	return &monkey
}

func initMonkey(data []string) *Monkey {
	var id int
	var starting_items []int
	var operand string
	var operator string
	var divisible_by int
	var throw_to_true int
	var throw_to_false int

	_, err := fmt.Sscanf(data[0], "Monkey %d:", &id)
	if err != nil {
		panic(err)
	}

	items_split := strings.Split(data[1], ":")
	starting_items, _ = util.ConvertArrayOfStringsToInts(strings.Split(items_split[1], ","))

	_, err_scan_operations := fmt.Sscanf(strings.Trim(data[2], " "), "Operation: new = old %v %v", &operator, &operand)
	if err_scan_operations != nil {
		panic(err)
	}

	_, err_scan_test := fmt.Sscanf(strings.Trim(data[3], " "), "Test: divisible by %v", &divisible_by)
	if err_scan_test != nil {
		panic(err)
	}

	_, err_scan_throw_to_true := fmt.Sscanf(strings.Trim(data[4], " "), "If true: throw to monkey %d", &throw_to_true)
	if err_scan_throw_to_true != nil {
		panic(err)
	}
	_, err_scan_throw_to_false := fmt.Sscanf(strings.Trim(data[5], " "), "If false: throw to monkey %d", &throw_to_false)
	if err_scan_throw_to_false != nil {
		panic(err)
	}

	return newMonkey(id, starting_items, operand, operator, divisible_by, throw_to_true, throw_to_false)
}

func main() {
	data := util.Data(1, "\n")
	data = data[:len(data)-1]

	if len(os.Args) < 3 {
		panic("Please provide puzzle part index as parameter `go run main.go -- 1`")
	}

	monkeys := map[int]*Monkey{}

	for i, j := 0, 0; i < len(data); i, j = i+7, j+1 {
		monkeys[j] = initMonkey(data[i : i+6])
	}

	var inspect_item int

	for round := 1; round < 10001; round++ {
		for i := 0; i < len(monkeys); i++ {
			for {
				if len(monkeys[i].items) == 0 {
					break
				}
				inspect_item, monkeys[i].items = monkeys[i].items[0], monkeys[i].items[1:]
				monkeys[i].inspected_items = append(monkeys[i].inspected_items, inspect_item)
				throw_to, item := monkeys[i].Throw(inspect_item)
				monkeys[throw_to].Catch(item)
			}
		}
	}

	inspected_length := []int{}
	for _, value := range monkeys {
		inspected_length = append(inspected_length, len(value.inspected_items))
	}

	sort.Ints(inspected_length)

	fmt.Printf("Multiple of two most active monkeys: %d", inspected_length[len(inspected_length)-1]*inspected_length[len(inspected_length)-2])

}
