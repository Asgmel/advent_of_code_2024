package dayFour

import (
	"strconv"

	"github.com/Asgmel/advent_of_code_2024/internal/input"
)

func Run() (func() string, func() string) {
	return taskOne, taskTwo
}

func taskOne() string {
	puzzleInput := input.ReadInputLetters(4, false)
	return strconv.Itoa(countWords(puzzleInput, "XMAS"))
}

func taskTwo() string {
	puzzleInput := input.ReadInputLetters(4, false)
	return strconv.Itoa(countMasX(puzzleInput))
}

func countWords(puzzleInput [][]string, word string) int {
	count := 0
	firstLetter := string(word[0])
	coordsToCheck := [][]int{
		{0, -1},
		{1, -1},
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
		{-1, -1},
	}

	for y := 0; y < len(puzzleInput); y++ {
		for x := 0; x < len(puzzleInput[0]); x++ {
			currentLetter := string(puzzleInput[y][x])
			if currentLetter == firstLetter {
				for _, coords := range coordsToCheck {
					if getWordByCoordinates(puzzleInput, x, y, coords[0], coords[1]) == word {
						count++
					}
				}
			}
		}
	}

	return count
}

func countMasX(puzzleInput [][]string) int {
	count := 0
	for y := 0; y < len(puzzleInput); y++ {
		for x := 0; x < len(puzzleInput[0]); x++ {
			if puzzleInput[y][x] == "A" && CheckMasXByCoordinates(puzzleInput, x, y) {
				count++
			}
		}
	}
	return count
}

func CheckMasXByCoordinates(puzzleInput [][]string, x, y int) bool {
	if x-1 < 0 || x+1 >= len(puzzleInput[0]) || y-1 < 0 || y+1 >= len(puzzleInput) {
		// string goes out of bound
		return false
	}

	wordOne := puzzleInput[y-1][x-1] + puzzleInput[y][x] + puzzleInput[y+1][x+1]
	wordTwo := puzzleInput[y+1][x-1] + puzzleInput[y][x] + puzzleInput[y-1][x+1]

	return (wordOne == "MAS" || wordOne == "SAM") && (wordTwo == "MAS" || wordTwo == "SAM")
}

func getWordByCoordinates(puzzleInput [][]string, x, y, dx, dy int) string {

	if x+(dx*3) >= len(puzzleInput[0]) || x+(dx*3) < 0 || y+(dy*3) >= len(puzzleInput) || y+(dy*3) < 0 {
		// string goes out of bounds
		return ""
	}

	return puzzleInput[y][x] + puzzleInput[y+(dy*1)][x+(dx*1)] + puzzleInput[y+(dy*2)][x+(dx*2)] + puzzleInput[y+(dy*3)][x+(dx*3)]
}
