package common

import "fmt"

type Position struct {
	X int
	Y int
}

func (position Position) ToString() string {
	return fmt.Sprintf("%v.%v", position.X, position.Y)
}
