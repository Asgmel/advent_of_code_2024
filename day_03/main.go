package main

import (
	"fmt"
	"strings"

	"github.com/Asgmel/advent_of_code_2024/internal/input"
	"github.com/Asgmel/advent_of_code_2024/internal/utils"
)

func main() {
	taskOne()
	taskTwo()
}

func taskOne() {
	puzzleInput := input.ReadInputRegex(`mul\((\d+),(\d+)\)`, 3, false)
	result := multiplyCommands(puzzleInput)
	fmt.Printf("The result of task one is: %v\n", result)
}

func taskTwo() {
	puzzleInput := input.ReadInputRegex(`(mul\(\d+,\d+\)|do\(\)|don't\(\))`, 3, false)
	filteredPuzzleInput := filterCommands(puzzleInput)
	result := multiplyCommands(filteredPuzzleInput)
	fmt.Printf("The result of task two is %v\n", result)
}

func multiplyCommands(commands []string) int {
	sum := 0

	for _, command := range commands {
		cutCommand, _ := strings.CutPrefix(command, "mul(")
		cutCommand, _ = strings.CutSuffix(cutCommand, ")")
		stringValues := strings.Split(cutCommand, ",")
		intValues := utils.ConvertStrSliceToIntSlice(stringValues)
		sum += intValues[0] * intValues[1]
	}

	return sum
}

func filterCommands(commands []string) []string {
	updatedCommands := []string{}
	active := true

	for _, command := range commands {
		if command == "do()" {
			active = true

		} else if command == "don't()" {
			active = false

		} else if strings.HasPrefix(command, "mul(") {
			if active {
				updatedCommands = append(updatedCommands, command)
			}
		} else {
			panic(fmt.Sprintf("Invalid command value: %v", command))
		}
	}
	return updatedCommands
}
