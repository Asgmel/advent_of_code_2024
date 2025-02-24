package dayOne

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/Asgmel/advent_of_code_2024/internal/input"
	"github.com/Asgmel/advent_of_code_2024/internal/utils"
)

func Run() (func() string, func() string) {
	return taskOne, taskTwo
}

func taskOne() string {
	puzzleInput := input.ReadInputLines(1, false)
	leftList, rightList := formatInput(puzzleInput)
	difference := findSliceDifference(leftList, rightList)
	sum := utils.SumIntSlice(difference)
	return strconv.Itoa(sum)
}

func taskTwo() string {
	puzzleInput := input.ReadInputLines(1, false)
	leftList, rightList := formatInput(puzzleInput)
	similarityScores := calculateSimilarityScore(leftList, rightList)
	sum := utils.SumIntSlice(similarityScores)
	return strconv.Itoa(sum)
}

func calculateSimilarityScore(leftList []int, rightList []int) []int {
	valueMap := map[int]int{}
	similarityScores := []int{}

	for _, val := range leftList {
		valueMap[val] = 0
	}

	for _, val := range rightList {
		_, exists := valueMap[val]
		if exists {
			valueMap[val] += val
		}
	}

	for _, val := range valueMap {
		similarityScores = append(similarityScores, val)
	}

	return similarityScores
}

func formatInput(puzzleInput []string) ([]int, []int) {
	leftList := []int{}
	rightList := []int{}

	for _, line := range puzzleInput {
		values := strings.Split(line, "   ")
		leftValue, err := strconv.Atoi(values[0])

		if err != nil {
			panic(fmt.Sprintf("Converting the left string value to an int failed: %v", err))
		}

		rightValue, err := strconv.Atoi(values[1])

		if err != nil {
			panic(fmt.Sprintf("Converting the right string value to an int failed: %v", err))
		}

		leftList = append(leftList, leftValue)
		rightList = append(rightList, rightValue)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	if len(leftList) != len(rightList) {
		panic(fmt.Sprintf("Invalid lists returned from formatInput, lenghts do not match. %v != %v", len(leftList), len(rightList)))
	}

	return leftList, rightList
}

func findSliceDifference(leftList []int, rightList []int) []int {
	differenceList := []int{}

	for i := 0; i < len(leftList); i++ {
		difference := utils.AbsoluteValue(leftList[i] - rightList[i])
		differenceList = append(differenceList, difference)
	}

	return differenceList
}
