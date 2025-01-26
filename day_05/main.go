package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/Asgmel/advent_of_code_2024/internal/input"
	"github.com/Asgmel/advent_of_code_2024/internal/utils"
)

type Rule struct {
	First int
	Last  int
}

func NewRule(rules string) *Rule {
	ruleSlice := utils.ConvertStrSliceToIntSlice(strings.Split(rules, "|"))
	return &Rule{
		First: ruleSlice[0],
		Last:  ruleSlice[1],
	}
}

func (rule Rule) Validate(first, last int) bool {
	return first != rule.Last || last != rule.First
}

type RuleSet struct {
	Rules []Rule
}

func (ruleSet RuleSet) ValidatePageNumbers(pageNumbers []int, swapIfError bool) (validated, swapped bool) {
	i := 0
pageLoop:
	for i < len(pageNumbers)-1 {
		for _, rule := range ruleSet.Rules {
			if !rule.Validate(pageNumbers[i], pageNumbers[i+1]) {
				if swapIfError {
					swapped = true
					temp := pageNumbers[i]
					pageNumbers[i] = pageNumbers[i+1]
					pageNumbers[i+1] = temp
					i = 0
					continue pageLoop
				} else {
					return
				}
			}
		}
		i++
	}
	validated = true
	return
}

func NewRuleSet(rawRuleSet []string) *RuleSet {
	rules := []Rule{}
	for _, rule := range rawRuleSet {
		rules = append(rules, *NewRule(rule))
	}
	return &RuleSet{
		Rules: rules,
	}
}

func main() {
	puzzleInput := input.ReadInputLines(5, false)
	commands, pageSets := splitCommandsFromPages(puzzleInput)
	ruleSet := NewRuleSet(commands)
	pageSetsSlice := changePagesStringToIntSlice(pageSets)
	taskOne(ruleSet, pageSetsSlice)
	taskTwo(ruleSet, pageSetsSlice)
}

func taskOne(ruleSet *RuleSet, pageSets [][]int) {
	validatedMiddlePages := []int{}
	for _, pageSet := range pageSets {
		validated, _ := ruleSet.ValidatePageNumbers(pageSet, false)
		if validated {
			validatedMiddlePages = append(validatedMiddlePages, getMiddleDigit(pageSet))
		}
	}
	fmt.Printf("The answer to task one is: %v\n", utils.SumIntSlice(validatedMiddlePages))
}

func taskTwo(ruleSet *RuleSet, pageSets [][]int) {
	validatedMiddlePages := []int{}
	for _, pageSet := range pageSets {
		validated, swapped := ruleSet.ValidatePageNumbers(pageSet, true)
		if validated && swapped {
			validatedMiddlePages = append(validatedMiddlePages, getMiddleDigit(pageSet))
		}
	}
	fmt.Printf("The answer to task two is: %v\n", utils.SumIntSlice(validatedMiddlePages))
}

func getMiddleDigit(numbers []int) int {
	return numbers[int(math.Floor(float64(len(numbers)/2)))]
}

func splitCommandsFromPages(puzzleInput []string) (commands []string, pageSets []string) {
	iteratingCommands := true
	for _, row := range puzzleInput {
		if row == "" {
			iteratingCommands = false
		} else {
			if iteratingCommands {
				commands = append(commands, row)
			} else {
				pageSets = append(pageSets, row)
			}
		}
	}
	return
}

func changePagesStringToIntSlice(pagesSlice []string) [][]int {
	intSlice := [][]int{}
	for _, pageString := range pagesSlice {
		intSlice = append(intSlice, utils.ConvertStrSliceToIntSlice(strings.Split(pageString, ",")))
	}
	return intSlice
}
