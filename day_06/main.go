package daySix

import (
	"fmt"
	"strconv"

	"github.com/Asgmel/advent_of_code_2024/internal/common"
	"github.com/Asgmel/advent_of_code_2024/internal/input"
)

type direction int

const (
	north direction = iota
	east
	south
	west
)

type tile struct {
	position   common.Position
	obstructed bool
}

func newTile(position common.Position, letter string) tile {
	return tile{
		position:   position,
		obstructed: letter == "#",
	}
}

type guard struct {
	position      common.Position
	direction     direction
	visited       []common.Position
	turningPoints map[string]struct{}
}

func (guard *guard) step() {
	guard.position = guard.getNextPosition()
	guard.visited = append(guard.visited, guard.position)
}

func (guard guard) getUniqueVisited() []common.Position {
	seen := map[string]common.Position{}
	uniqueVisited := []common.Position{}

	for _, position := range guard.visited {
		if _, exists := seen[position.ToString()]; !exists {
			uniqueVisited = append(uniqueVisited, position)
			seen[position.ToString()] = position
		}
	}
	return uniqueVisited
}
func (guard guard) getPositionDirectionString() string {
	return fmt.Sprintf("%v-%v", guard.position.ToString(), guard.direction)
}

func (guard *guard) turn() {
	guard.turningPoints[guard.getPositionDirectionString()] = struct{}{}
	switch guard.direction {
	case north:
		guard.direction = east
	case east:
		guard.direction = south
	case south:
		guard.direction = west
	case west:
		guard.direction = north
	}
}

func (guard guard) getNextPosition() common.Position {
	position := guard.position
	switch guard.direction {
	case north:
		position.Y--
	case east:
		position.X++
	case south:
		position.Y++
	case west:
		position.X--
	}
	return position
}

func newGuard(position common.Position) *guard {
	return &guard{
		position:      position,
		direction:     north,
		visited:       []common.Position{position},
		turningPoints: map[string]struct{}{},
	}
}

type state struct {
	tiles   [][]tile
	guard   *guard
	looping bool
}

func (state state) getTile(position common.Position) tile {
	return state.tiles[position.Y][position.X]
}

func (state state) checkPositionWithinBounds(position common.Position) bool {
	return position.X >= 0 && position.X < len(state.tiles[0]) && position.Y >= 0 && position.Y < len(state.tiles)
}

func (state *state) calculateRoute() {
	for {
		nextPosition := state.guard.getNextPosition()
		if !state.checkPositionWithinBounds(nextPosition) {
			return
		}
		nextTile := state.getTile(nextPosition)
		if nextTile.obstructed {
			if _, exists := state.guard.turningPoints[state.guard.getPositionDirectionString()]; exists {
				state.looping = true
				return
			}
			state.guard.turn()
		} else {
			state.guard.step()
		}
	}
}

func newState(stringTiles [][]string) state {
	tiles := [][]tile{}
	guardPosition := common.Position{X: 0, Y: 0}
	for y, row := range stringTiles {
		tileRow := []tile{}
		for x, letter := range row {
			tileRow = append(tileRow, newTile(common.Position{X: x, Y: y}, letter))
			if letter == "^" {
				guardPosition = common.Position{X: x, Y: y}
			}
		}
		tiles = append(tiles, tileRow)
	}
	return state{
		tiles:   tiles,
		guard:   newGuard(guardPosition),
		looping: false,
	}
}

func Run() (func() string, func() string) {
	return taskOne, taskTwo
}

func taskOne() string {
	puzzleInput := input.ReadInputLetters(6, false)
	state := newState(puzzleInput)
	state.calculateRoute()
	return strconv.Itoa(len(state.guard.getUniqueVisited()))
}

func taskTwo() string {
	puzzleInput := input.ReadInputLetters(6, false)
	loopedStates := 0

	for y, row := range puzzleInput {
		for x, letter := range row {
			if letter == "." {
				state := newState(puzzleInput)
				state.tiles[y][x] = newTile(common.Position{X: x, Y: y}, "#")
				state.calculateRoute()
				if state.looping {
					loopedStates++
				}
			}
		}
	}
	return strconv.Itoa(loopedStates)

}
