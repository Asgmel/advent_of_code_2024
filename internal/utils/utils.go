package utils

import (
	"strconv"
)

func AbsoluteValue(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func SumIntSlice(slice []int) int {
	sum := 0
	for _, val := range slice {
		sum += val
	}
	return sum
}

func ConvertStrSliceToIntSlice(strSlice []string) []int {
	intSlice := []int{}
	for _, value := range strSlice {
		intValue, err := strconv.Atoi(value)

		if err != nil {
			panic(err)
		}

		intSlice = append(intSlice, intValue)
	}

	return intSlice
}

func CopySlice[T any](slice []T) []T {
	return append(slice[:0:0], slice...)
}

func FilterDuplicates[T comparable](slice []T) []T {
	seen := map[T]bool{}
	result := []T{}

	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}
