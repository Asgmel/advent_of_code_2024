package daySeven

import (
	"strconv"
	"strings"

	"github.com/Asgmel/advent_of_code_2024/internal/input"
	"github.com/Asgmel/advent_of_code_2024/internal/utils"
)

type equation struct {
	sum              int
	numbers          []int
	operators        []string
	allowedOperators []string
	solvable         bool
}

func (eq equation) evaluateExpression(operators []string) int {
	result := eq.numbers[0]

	for i, operator := range operators {
		if operator == "+" {
			result += eq.numbers[i+1]
		} else if operator == "*" {
			result *= eq.numbers[i+1]
		} else if operator == "|" {
			var err error
			result, err = strconv.Atoi(strconv.Itoa(result) + strconv.Itoa(eq.numbers[i+1]))
			if err != nil {
				panic(err)
			}
		}
	}

	return result
}

func (eq equation) findOperations(operators []string, index int) (bool, []string) {
	if index == len(eq.numbers)-1 {
		if eq.evaluateExpression(operators) == eq.sum {
			return true, operators
		}
		return false, []string{}
	}

	for _, operator := range eq.allowedOperators {
		operators[index] = operator
		if valid, _ := eq.findOperations(operators, index+1); valid {
			return true, operators
		}
	}

	return false, []string{}
}
func (eq *equation) checkIfSolvable() {
	solvable, operators := eq.findOperations(eq.operators, 0)
	if solvable {
		eq.solvable = true
		eq.operators = operators
	}
}

func newEquation(line string, allowedOperators []string) equation {
	line = strings.TrimSpace(line)
	lineParts := strings.Split(line, ": ")
	sum, err := strconv.Atoi(lineParts[0])
	if err != nil {
		panic("Converting part one of the string to result int failed.")
	}

	numbers := utils.ConvertStrSliceToIntSlice(strings.Split(lineParts[1], " "))
	eq := equation{
		sum:              sum,
		numbers:          numbers,
		operators:        make([]string, len(numbers)-1),
		allowedOperators: allowedOperators,
		solvable:         false,
	}
	eq.checkIfSolvable()
	return eq
}

func Run() (func() string, func() string) {
	return taskOne, taskTwo
}

func taskOne() string {
	puzzleInput := input.ReadInputLines(7, false)
	sum := 0
	allowedOperators := []string{"+", "*"}
	for _, line := range puzzleInput {
		equation := newEquation(line, allowedOperators)
		if equation.solvable {
			sum += equation.sum
		}
	}
	return strconv.Itoa(sum)
}

func taskTwo() string {
	puzzleInput := input.ReadInputLines(7, false)
	sum := 0
	allowedOperators := []string{"+", "*", "|"}
	for _, line := range puzzleInput {
		equation := newEquation(line, allowedOperators)
		if equation.solvable {
			sum += equation.sum
		}
	}
	return strconv.Itoa(sum)
}
