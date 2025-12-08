package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	daySolutions := map[int]func(){
		1: DayOne,
		2: DayTwo,
		3: DayThree,
		4: DayFour,
		5: DayFive,
		6: DaySix,
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <day_number>")
		fmt.Println("Example: go run . 1")
		os.Exit(1)
	}

	dayNum, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Error: '%s' is not a valid day number\n", os.Args[1])
		os.Exit(1)
	}

	if solution, exists := daySolutions[dayNum]; exists {
		fmt.Printf("Running solution for Day %d...\n\n", dayNum)
		solution()
	} else {
		fmt.Println("Oops! Seems like that problem doesn't exist or I didn't solve it!")
		os.Exit(1)
	}
}
