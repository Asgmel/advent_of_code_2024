package common

import "fmt"

type Position struct {
	X int
	Y int
}

func (position Position) ToString() string {
	return fmt.Sprintf("%v.%v", position.X, position.Y)
}

func (position Position) WithinBounds(matrix [][]string) bool {
	return position.X >= 0 && position.X < len(matrix[0]) && position.Y >= 0 && position.Y < len(matrix)
}
