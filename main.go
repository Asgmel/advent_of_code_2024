package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/Asgmel/advent_of_code_2024/day_01"
	"github.com/Asgmel/advent_of_code_2024/day_02"
	"github.com/Asgmel/advent_of_code_2024/day_03"
	"github.com/Asgmel/advent_of_code_2024/day_04"
	"github.com/Asgmel/advent_of_code_2024/day_05"
	"github.com/Asgmel/advent_of_code_2024/day_06"
	"github.com/Asgmel/advent_of_code_2024/day_07"
	"github.com/Asgmel/advent_of_code_2024/day_08"
	"github.com/Asgmel/advent_of_code_2024/day_09"
	"github.com/Asgmel/advent_of_code_2024/day_10"
)

// callerName returns the name of the function skip frames up the call stack.
func callerName(skip int) string {
	const unknown = "unknown"
	pcs := make([]uintptr, 1)
	n := runtime.Callers(skip+2, pcs)
	if n < 1 {
		return unknown
	}
	frame, _ := runtime.CallersFrames(pcs).Next()
	if frame.Function == "" {
		return unknown
	}
	return frame.Function
}

func timer() func() {
	name := callerName(1)
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func main() {
	// add new days here as they are finished
	dayMap := map[string]func() (func() string, func() string){
		"1":  dayOne.Run,
		"2":  dayTwo.Run,
		"3":  dayThree.Run,
		"4":  dayFour.Run,
		"5":  dayFive.Run,
		"6":  daySix.Run,
		"7":  daySeven.Run,
		"8":  dayEight.Run,
		"9":  dayNine.Run,
		"10": dayTen.Run,
	}

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please specify the day to run after the path.")
	} else if len(args) > 2 {
		fmt.Println("Please only add one parameter, the day to run.")
	} else if _, exists := dayMap[args[1]]; len(args) == 2 && !exists {
		fmt.Println("Invalid parameter. Enter a valid number between 1 and 25.")
	} else {
		taskOne, taskTwo := dayMap[args[1]]()
		runTaskOne(taskOne)
		runTaskTwo(taskTwo)
		fmt.Println("--------------------------------------")
	}
}

func runTaskOne(task func() string) {
	defer timer()()
	fmt.Println("--------------------------------------")
	fmt.Printf("The result of task one is %v.\n", task())
}

func runTaskTwo(task func() string) {
	defer timer()()
	fmt.Println("--------------------------------------")
	fmt.Printf("The result of task two %v.\n", task())
}
