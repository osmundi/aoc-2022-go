package main

import (
	"aoc/util"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Operation struct {
	ape1    string
	ape2    string
	operand string
}

func (op Operation) Solve(ape1, ape2 float64) float64 {
	if op.operand == "+" {
		return ape1 + ape2
	}
	if op.operand == "-" {
		return ape1 - ape2
	}
	if op.operand == "*" {
		return ape1 * ape2
	}
	if op.operand == "/" {
		return ape1 / ape2
	}
	return 0
}

type SyntaxTree struct {
	left     *SyntaxTree
	right    *SyntaxTree
	operand  string  // tree has only one of these values
	value    float64 // tree has only one of these values
	variable bool    // tree has only one of these values
}

func (tree SyntaxTree) String() string {
	if tree.value != 0 {
		return fmt.Sprintf("%f", tree.value)
	} else if tree.variable {
		return fmt.Sprintf("X")
	} else {
		return fmt.Sprintf("%s (%v) | (%v)", tree.operand, tree.left, tree.right)
	}
}

func (tree SyntaxTree) ContainsVariable() bool {
	if tree.value != 0 {
		return false
	}
	if tree.variable {
		return true
	}
	if tree.left.ContainsVariable() {
		return true
	}
	if tree.right.ContainsVariable() {
		return true
	}
	return false
}

func (tree *SyntaxTree) GenerateTree(ape string, solved map[string]float64, notSolved map[string]Operation) {
	val, ok := solved[ape]
	if ok {
		tree.value = val
		return
	}

	if ape == "humn" {
		tree.variable = true
		return
	}

	next, _ := notSolved[ape]

	var left SyntaxTree
	var right SyntaxTree
	left.GenerateTree(next.ape1, solved, notSolved)
	right.GenerateTree(next.ape2, solved, notSolved)

	tree.left = &left
	tree.operand = next.operand
	tree.right = &right
}

func (tree SyntaxTree) SolveEquation() float64 {
	// every ape has shout value
	if tree.value != 0 {
		return tree.value
	}
	if tree.operand == "+" {
		return tree.left.SolveEquation() + tree.right.SolveEquation()
	}
	if tree.operand == "-" {
		return tree.left.SolveEquation() - tree.right.SolveEquation()
	}
	if tree.operand == "*" {
		return tree.left.SolveEquation() * tree.right.SolveEquation()
	}
	if tree.operand == "/" {
		return tree.left.SolveEquation() / tree.right.SolveEquation()
	}
	return 0
}

func (tree SyntaxTree) BufferOperations(ops *[]string) {
	var operation string
	if tree.operand == "+" {
		operation = "-"
	}
	if tree.operand == "-" {
		operation = "+"
	}
	if tree.operand == "*" {
		operation = "/"
	}
	if tree.operand == "/" {
		operation = "*"
	}

	if tree.left.ContainsVariable() {
		// X - 3 = 10
		// -> X = 10 + 3
		*ops = append(*ops, operation+strconv.FormatFloat(tree.right.SolveEquation(), 'f', -1, 64))
		if !tree.left.variable {
			tree.left.BufferOperations(ops)
		}
	} else {
		// Division and substraction has a bit different logic when the variable 'X' is on the right side of the equation:
		// 3 - X = 10
		// -> X = -1 * (10 - 3)
		*ops = append(*ops, "inverse:"+operation+strconv.FormatFloat(tree.left.SolveEquation(), 'f', -1, 64))
		if !tree.right.variable {
			tree.right.BufferOperations(ops)
		}
	}
}

func applyOperations(start float64, operations []string) float64 {
	inverse := false
	for _, op := range operations {
		if strings.HasPrefix(op, "inverse:") {
			op = strings.TrimPrefix(op, "inverse:")
			inverse = true
		}
		result := []rune(op)
		val, err := strconv.ParseFloat(string(result[1:]), 64)
		if err != nil {
			panic(err)
		}

		if result[0] == '+' {
			if inverse {
				start = val - start
			} else {
				start += val
			}
		}
		if result[0] == '-' {
			start -= val
		}
		if result[0] == '*' {
			if inverse {
				start = val / start
			} else {
				start *= val
			}
		}
		if result[0] == '/' {
			start /= val
		}
		inverse = false
	}
	return start
}

func main() {
	data := util.Data(1, "\n")
	data = data[:len(data)-1]

	if len(os.Args) < 3 {
		panic("Please provide puzzle part index as parameter `go run main.go -- 1`")
	}

	part := os.Args[2]

	fmt.Printf("Run part: %s\n", part)

	var solved = make(map[string]float64)
	var notSolved = make(map[string]Operation)

	r, _ := regexp.Compile(`^(?P<ape>\w+):\s(?P<shout>\d+)`)

	var ape string
	var shouter1 string
	var shouter2 string
	var operator string

	for _, val := range data {
		m := r.FindStringSubmatch(val)
		if m == nil {
			fmt.Sscanf(val, "%s %s %s %s", &ape, &shouter1, &operator, &shouter2)

			if part == "2" && strings.HasPrefix(ape, "root") {
				operator = "="
			}

			notSolved[strings.TrimSuffix(ape, ":")] = Operation{shouter1, shouter2, operator}
		} else {
			result := make(map[string]string)
			for i, name := range r.SubexpNames() {
				if i != 0 && name != "" {
					result[name] = m[i]
				}
			}

			shoutValue, err := strconv.ParseFloat(result["shout"], 64)
			if err != nil {
				panic(err)
			}

			// humn not solved in part 2
			if part == "2" && result["ape"] == "humn" {
				continue
			}

			solved[result["ape"]] = float64(shoutValue)
		}
	}

	if part == "1" {
		for {
			if len(notSolved) == 0 {
				break
			}
			for key, val := range notSolved {
				solvedApe1, ok1 := solved[val.ape1]
				solvedApe2, ok2 := solved[val.ape2]
				if ok1 && ok2 {
					solved[key] = val.Solve(solvedApe1, solvedApe2)
					delete(notSolved, key)
				}
			}
		}
		fmt.Printf("%f\n", solved["root"])
	}

	if part == "2" {
		var eq SyntaxTree
		eq.GenerateTree("root", solved, notSolved)
		var ops []string
		eq.left.BufferOperations(&ops)
		fmt.Printf("%f", applyOperations(eq.right.SolveEquation(), ops))
	}
}
