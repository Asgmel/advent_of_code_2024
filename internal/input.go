package input

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInputString(day int, use_test_input bool) string {
	if day < 1 || day > 25 {
		printInputError(errors.New(fmt.Sprintf("Day %v is invalid, must be between 1 and 25.", day)))
	}

	day_string := strconv.Itoa(day)

	if len(day_string) == 1 {
		day_string = "0" + day_string
	}

	path := fmt.Sprintf("inputs/day_%s", day_string)

	if use_test_input {
		path = path + "_test"
	}

	content, err := os.ReadFile(fmt.Sprintf("%s.txt", path))

	if err != nil {
		printInputError(err)
	}

	text := string(content)

	if strings.HasSuffix(text, "\n") {
		text = strings.TrimSuffix(text, "\n")
	}

	return text
}

func printInputError(err error) {
	fmt.Println("An error occured while trying to read input data:")
	panic(err)
}

func ReadInputLines(day int, use_test_input bool) []string {
	input_string := ReadInputString(day, use_test_input)
	return strings.Split(input_string, "\n")
}
