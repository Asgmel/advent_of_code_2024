package main

import (
	"fmt"
	"github.com/Asgmel/advent_of_code_2024/internal"
)

func main() {
	input_slice := input.ReadInputLines(1, false)

	for x, line := range input_slice {
		fmt.Println(x, line)
	}
}
