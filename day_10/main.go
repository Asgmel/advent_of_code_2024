package dayTen

import (
	"fmt"
	"strconv"

	"github.com/Asgmel/advent_of_code_2024/internal/common"
	"github.com/Asgmel/advent_of_code_2024/internal/input"
	"github.com/Asgmel/advent_of_code_2024/internal/utils"
)

func countValidTrails(puzzleInput [][]string, currentPosition common.Position) []common.Position {
	currentValue, err := strconv.Atoi(puzzleInput[currentPosition.Y][currentPosition.X])
	if err != nil {
		panic(err)
	}
	if currentValue == 9 {
		return []common.Position{currentPosition}
	}

	neighbours := []common.Position{
		{X: currentPosition.X, Y: currentPosition.Y - 1},
		{X: currentPosition.X + 1, Y: currentPosition.Y},
		{X: currentPosition.X, Y: currentPosition.Y + 1},
		{X: currentPosition.X - 1, Y: currentPosition.Y},
	}
	tops := []common.Position{}

	for _, neighbour := range neighbours {
		if !neighbour.WithinBounds(puzzleInput) {
			continue
		}
		neighbourValue, err := strconv.Atoi(puzzleInput[neighbour.Y][neighbour.X])
		if err != nil {
			panic(err)
		}
		if neighbourValue == currentValue+1 {
			tops = append(tops, countValidTrails(puzzleInput, neighbour)...)
		}
	}
	return tops
}

func Run() (func() string, func() string) {
	return taskOne, taskTwo
}

func taskOne() string {
	puzzleInput := input.ReadInputLetters(10, false)
	sum := 0
	for y := range puzzleInput {
		for x := range puzzleInput[y] {
			if puzzleInput[y][x] == "0" {
				sum += len(utils.FilterDuplicates(countValidTrails(puzzleInput, common.Position{X: x, Y: y})))
			}
		}
	}
	return strconv.Itoa(sum)
}

func taskTwo() string {
	puzzleInput := input.ReadInputLetters(10, false)
	sum := 0
	for y := range puzzleInput {
		for x := range puzzleInput[y] {
			if puzzleInput[y][x] == "0" {
				sum += len(countValidTrails(puzzleInput, common.Position{X: x, Y: y}))
			}
		}
	}
	return strconv.Itoa(sum)
}
