package dayTwo

import (
	"fmt"
	"strings"

	"github.com/Asgmel/advent_of_code_2024/internal/input"
	"github.com/Asgmel/advent_of_code_2024/internal/utils"
)

func Run() {
	puzzleInput := input.ReadInputLines(2, false)
	formattedInput := formatInput(puzzleInput)
	taskOne(formattedInput)
	taskTwo(formattedInput)
}

func taskOne(formattedInput [][]int) {
	count := 0
	for _, row := range formattedInput {
		if (checkAscending(row) || checkDescending(row)) && checkDifference(row) {
			count++
		}
	}
	fmt.Printf("The answer to task one is: %v\n", count)
}

func taskTwo(formattedInput [][]int) {
	count := 0
inputLoop:
	for _, row := range formattedInput {
		for i := range row {
			copyRow := utils.CopySlice(row)
			tempRow := append(copyRow[:i], copyRow[i+1:]...)
			if (checkAscending(tempRow) || checkDescending(tempRow)) && checkDifference(tempRow) {
				count++
				continue inputLoop
			}
		}
	}
	fmt.Printf("The answer to task two is: %v\n", count)
}

func formatInput(puzzleInput []string) [][]int {
	formattedInput := [][]int{}
	for _, row := range puzzleInput {
		strSlice := strings.Split(row, " ")
		formattedInput = append(formattedInput, utils.ConvertStrSliceToIntSlice(strSlice))
	}
	return formattedInput
}

func checkAscending(row []int) bool {
	for i := 0; i < len(row)-1; i++ {
		if row[i] > row[i+1] {
			return false
		}
	}
	return true
}

func checkDescending(row []int) bool {
	for i := 0; i < len(row)-1; i++ {
		if row[i] < row[i+1] {
			return false
		}
	}
	return true
}

func checkDifference(row []int) bool {
	for i := 0; i < len(row)-1; i++ {
		difference := utils.AbsoluteValue(row[i] - row[i+1])
		if difference < 1 || difference > 3 {
			return false
		}
	}
	return true
}
