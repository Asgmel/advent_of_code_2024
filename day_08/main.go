package dayEight

import (
	"fmt"
	"strconv"

	"github.com/Asgmel/advent_of_code_2024/internal/common"
	"github.com/Asgmel/advent_of_code_2024/internal/input"
)

type bounds struct {
	minX int
	minY int
	maxX int
	maxY int
}

func newBounds(puzzleInput [][]string) bounds {
	return bounds{
		minX: 0,
		maxX: len(puzzleInput[0]) - 1,
		minY: 0,
		maxY: len(puzzleInput) - 1,
	}
}

func getDistanceBetweenNodes(n1, n2 common.Position) common.Position {
	return common.Position{X: n2.X - n1.X, Y: n2.Y - n1.Y}
}

func getAntinodesInBounds(p1, p2 common.Position, distance common.Position, bounds bounds) (antinodes []common.Position) {
	potentialAntinodes := []common.Position{
		{X: p1.X - distance.X, Y: p1.Y - distance.Y},
		{X: p2.X + distance.X, Y: p2.Y + distance.Y},
	}
	for _, antinode := range potentialAntinodes {
		if checkWithinBounds(antinode, bounds) {
			antinodes = append(antinodes, antinode)
		}
	}
	return

}

func getAllNodesInBounds(p1, p2 common.Position, distance common.Position, bounds bounds) (nodes []common.Position) {
	for {
		nodes = append(nodes, p1)
		p1 = common.Position{X: p1.X - distance.X, Y: p1.Y - distance.Y}
		if !checkWithinBounds(p1, bounds) {
			break
		}
	}

	for {
		nodes = append(nodes, p2)
		p2 = common.Position{X: p2.X + distance.X, Y: p2.Y + distance.Y}
		if !checkWithinBounds(p2, bounds) {
			break
		}
	}

	return
}

func formatInput(puzzleInput [][]string) map[string][]common.Position {
	letterMap := map[string][]common.Position{}

	for y, line := range puzzleInput {
		for x, letter := range line {
			if letter == "." {
				continue
			}
			if _, exists := letterMap[letter]; exists {
				letterMap[letter] = append(letterMap[letter], common.Position{X: x, Y: y})
			} else {
				letterMap[letter] = []common.Position{{X: x, Y: y}}
			}
		}
	}
	return letterMap
}

func checkWithinBounds(position common.Position, bounds bounds) bool {
	return position.X >= bounds.minX && position.X <= bounds.maxX && position.Y >= bounds.minY && position.Y <= bounds.maxY
}

func getAntennas(formattedInput map[string][]common.Position) []common.Position {
	antennas := []common.Position{}
	for _, value := range formattedInput {
		antennas = append(antennas, value...)
	}
	return antennas
}

func getAntennaPairs(formattedInput map[string][]common.Position) [][2]common.Position {
	antennaPairs := [][2]common.Position{}
	for _, positions := range formattedInput {
		for i := 0; i < len(positions)-1; i++ {
			for j := i; j < len(positions); j++ {
				if i == j {
					continue
				}
				antennaPairs = append(antennaPairs, [2]common.Position{positions[i], positions[j]})
			}
		}
	}
	return antennaPairs
}

func filterOverlappingNodes(antinodes []common.Position, antennas []common.Position) (validAntinodes []common.Position) {
antinodeLoop:
	for _, antinode := range antinodes {
		for _, antenna := range antennas {
			if antinode == antenna {
				continue antinodeLoop
			}
		}
		validAntinodes = append(validAntinodes, antinode)
	}
	return
}

func getNodesFromAntennaPairs(
	antennaPairs [][2]common.Position,
	bounds bounds,
	nodeFunc func(common.Position, common.Position, common.Position, bounds) []common.Position) (nodes []common.Position) {
	for _, antennaPair := range antennaPairs {
		distance := getDistanceBetweenNodes(antennaPair[0], antennaPair[1])
		a1 := antennaPair[0]
		a2 := antennaPair[1]
		nodes = append(nodes, nodeFunc(a1, a2, distance, bounds)...)
	}
	return
}

func printNodes(puzzleInput [][]string, nodes []common.Position) {
	for _, node := range nodes {
		puzzleInput[node.Y][node.X] = "#"
	}

	for _, row := range puzzleInput {
		fmt.Println(row)
	}
}

func getUniqueNodes(nodes []common.Position) (uniqueNodes []common.Position) {
	nodesMap := map[common.Position]struct{}{}
	for _, node := range nodes {
		nodesMap[node] = struct{}{}
	}

	for key := range nodesMap {
		uniqueNodes = append(uniqueNodes, key)
	}
	return
}

func Run() (func() string, func() string) {
	return taskOne, taskTwo
}

func taskOne() string {
	puzzleInput := input.ReadInputLetters(8, false)
	bounds := newBounds(puzzleInput)
	formattedInput := formatInput(puzzleInput)
	antennas := getAntennas(formattedInput)
	antennaPairs := getAntennaPairs(formattedInput)
	antinodes := getNodesFromAntennaPairs(antennaPairs, bounds, getAntinodesInBounds)
	return strconv.Itoa(len(filterOverlappingNodes(antinodes, antennas)))
}

func taskTwo() string {
	puzzleInput := input.ReadInputLetters(8, false)
	bounds := newBounds(puzzleInput)
	formattedInput := formatInput(puzzleInput)
	antennaPairs := getAntennaPairs(formattedInput)
	antinodes := getNodesFromAntennaPairs(antennaPairs, bounds, getAllNodesInBounds)
	return strconv.Itoa(len(getUniqueNodes(antinodes)))
}
