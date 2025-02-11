package dayEight

import (
	"strconv"

	"github.com/Asgmel/advent_of_code_2024/internal/common"
	"github.com/Asgmel/advent_of_code_2024/internal/input"
)

func getDistanceBetweenNodes(n1, n2 common.Position) common.Position {
	return common.Position{X: n2.X - n1.X, Y: n2.Y - n1.Y}
}

func getAntinodes(n1, n2 common.Position) (common.Position, common.Position) {
	distance := getDistanceBetweenNodes(n1, n2)
	return common.Position{X: n1.X - distance.X, Y: n1.Y - distance.Y}, common.Position{X: n2.X + distance.X, Y: n2.Y + distance.Y}
}

func formatInput(puzzleInput [][]string) map[string][]common.Position {
	letterMap := map[string][]common.Position{}

	for y, line := range puzzleInput {
		for x, letter := range line {
			if letter == "." {
				continue
			}
			if _, exists := letterMap[letter]; exists {
				letterMap[letter] = append(letterMap[letter], common.Position{X: x, Y: y})
			} else {
				letterMap[letter] = []common.Position{{X: x, Y: y}}
			}
		}
	}
	return letterMap
}

func checkWithinBounds(position common.Position, width, height int) bool {
	return position.X >= 0 && position.X < width && position.Y >= 0 && position.Y < height
}

func getAntinodesFromInput(formattedInput map[string][]common.Position) (antennas []common.Position, antinodes []common.Position) {
	for _, positions := range formattedInput {
		antennas = append(antennas, positions...)
		for i := 0; i < len(positions)-1; i++ {
			for j := i; j < len(positions); j++ {
				if i == j {
					continue
				}
				a1, a2 := getAntinodes(positions[i], positions[j])
				antinodes = append(antinodes, a1, a2)
			}
		}
	}
	return
}

func countValidResults(antinodes []common.Position, antennas []common.Position, maxWidth, maxHeight int) int {
	validAntinodes := []common.Position{}
antinodeLoop:
	for _, antinode := range antinodes {
		if checkWithinBounds(antinode, maxWidth, maxHeight) {
			for _, antenna := range antennas {
				if antinode == antenna {
					continue antinodeLoop
				}
			}
			validAntinodes = append(validAntinodes, antinode)
		}
	}
	return len(validAntinodes)
}

func Run() (func() string, func() string) {
	return taskOne, taskTwo
}

func taskOne() string {
	puzzleInput := input.ReadInputLetters(8, false)
	formattedInput := formatInput(puzzleInput)
	antennas, antinodes := getAntinodesFromInput(formattedInput)
	validAntinodesCount := countValidResults(antinodes, antennas, len(puzzleInput[0]), len(puzzleInput))
	return strconv.Itoa(validAntinodesCount)
}

func taskTwo() string {
	// return strconv.Itoa(sum)
	return "not implemented"
}
