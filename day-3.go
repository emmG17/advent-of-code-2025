package main

import (
	"fmt"
	"strconv"
)

func DayThree() {
	banks := getData("./challenge-input/day-3.txt", "\n")

	sum := 0
	for j, bank := range banks {
		maxB := 0
		maxA := 0
		for i := 0; i < len(bank) - 1; i++ {
			numA, err := strconv.Atoi(string(bank[i]))
			check(err)
			numB, err := strconv.Atoi(string(bank[i + 1]))
			check(err)

			if maxA < numA {
				maxA = numA
				maxB = 0
			}

			if maxB < numB {
				maxB = numB
			}
		}

		sum += maxA * 10 + maxB
		fmt.Printf("%d: %d\n", j, maxA * 10 + maxB)
	}

	fmt.Printf("Day 3 results, Part A: %d", sum)
}
