package utils

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
