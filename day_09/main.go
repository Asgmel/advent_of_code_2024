package dayNine

import (
	"fmt"
	"strconv"

	"github.com/Asgmel/advent_of_code_2024/internal/input"
	"github.com/Asgmel/advent_of_code_2024/internal/utils"
)

func formatDiskMapBlocks(diskMap string) (blocks []int) {
	idCount := 0
	file := true
	for _, letter := range diskMap {
		num := int(letter - '0')
		for i := 0; i < num; i++ {
			if file {
				blocks = append(blocks, idCount)
			} else {
				blocks = append(blocks, -1)
			}
		}
		if file {
			idCount++
		}
		file = !file
	}
	return
}

func defragDrive(block []int) []int {
	freeAddresses := []int{}
	takenAddresses := []int{}
	for i, value := range block {
		if value == -1 {
			freeAddresses = append(freeAddresses, i)
		} else {
			takenAddresses = append(takenAddresses, i)
		}
	}

	index := 0
	for {
		nextFreeAddress := freeAddresses[index]
		nextTakenAddress := takenAddresses[len(takenAddresses)-1-index]
		if nextFreeAddress > nextTakenAddress {
			break
		}
		block[nextFreeAddress] = block[nextTakenAddress]
		block[nextTakenAddress] = -1
		index++
	}
	return block[:len(takenAddresses)]
}

func defragDriveWithoutSplittingBlocks(originalBlocks []int) []int {
	blocks := utils.CopySlice(originalBlocks)
	freeAddresses := [][]int{}
	takenAddresses := [][]int{}
	currentAddressSpace := []int{}
	currentValue := blocks[0]
	for i, value := range blocks {
		if value == currentValue {
			currentAddressSpace = append(currentAddressSpace, i)
		} else {
			if currentValue == -1 {
				freeAddresses = append(freeAddresses, currentAddressSpace)
			} else {
				takenAddresses = append(takenAddresses, currentAddressSpace)
			}
			currentAddressSpace = []int{i}
			currentValue = value
		}
	}
	if currentValue == -1 {
		freeAddresses = append(freeAddresses, currentAddressSpace)
	} else {
		takenAddresses = append(takenAddresses, currentAddressSpace)
	}

	for i := len(takenAddresses) - 1; i >= 0; i-- {
		for j := 0; j < len(freeAddresses); j++ {
			if len(takenAddresses[i]) <= len(freeAddresses[j]) {
				if takenAddresses[i][0] < freeAddresses[j][0] {
					continue
				}
				for k := range takenAddresses[i] {
					blocks[freeAddresses[j][k]] = blocks[takenAddresses[i][k]]
					blocks[takenAddresses[i][k]] = -1
				}
				freeAddresses[j] = freeAddresses[j][len(takenAddresses[i]):]
				break
			}
		}
	}
	return blocks
}

func calculateCheckSum(block []int) (sum int) {
	for i, num := range block {
		if num == -1 {
			continue
		}
		sum += i * num
	}
	return
}

func printBlock(blocks []int) {
	for _, block := range blocks {
		if block == -1 {
			fmt.Print(".")
		} else {
			fmt.Print(block)
		}
	}
	fmt.Printf("\nThe length of the blocks is: %v\n", len(blocks))
}

func Run() (func() string, func() string) {
	return taskOne, taskTwo
}

func taskOne() string {
	diskMap := input.ReadInputString(9, false)
	blocks := formatDiskMapBlocks(diskMap)
	defraggedBlocks := defragDrive(blocks)
	return strconv.Itoa(calculateCheckSum(defraggedBlocks))
}

func taskTwo() string {
	diskMap := input.ReadInputString(9, false)
	blocks := formatDiskMapBlocks(diskMap)
	defraggedBlocks := defragDriveWithoutSplittingBlocks(blocks)
	return strconv.Itoa(calculateCheckSum(defraggedBlocks))
}
