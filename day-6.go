package main

import (
	"fmt"
	"strings"
	"strconv"
)

type Task struct {
	Numbers []int
	Op      string
}

func applyOp(task Task) int {
	result := 0

	switch task.Op {
	case "+":
		for _, num := range task.Numbers {
			result += num
		}
	case "-":
		for _, num := range task.Numbers {
			result -= num
		}
	case "*":
		result = task.Numbers[0]
		for _, num := range task.Numbers[1:] {
			result *= num
		}
	case "/":
		result = task.Numbers[0]
		for _, num := range task.Numbers[1:] {
			if num != 0 {
				result /= num
			} else {
				result = 0 // handle division by zero
			}
		}
	default:
		// handle unknown operation
		result = 0
	}
	return result
}

func columnsToRows(data [][]string) [][]string {
	if len(data) == 0 {
		return [][]string{}
	}
	numCols := len(data[0])
	numRows := len(data)
	transposed := make([][]string, numCols)
	for i := range transposed {
		transposed[i] = make([]string, numRows)
	}
	for i := range numRows {
		for j := range numCols {
			transposed[j][i] = data[i][j]
		}
	}
	return transposed
}

func parseData(data []string) []Task {
	tasks := []Task{}
	splitData := [][]string{}
	for _, line := range data {
		splitLine := strings.Fields(line)
		splitData = append(splitData, splitLine)
	}

  columnData := columnsToRows(splitData)

	// Col is: [numbers..., op]
	for _, col := range columnData {
		numbers := []int{}
		for _, strNum := range col[:len(col)-1] {
			num, err := strconv.Atoi(strNum)
			if err == nil {
				numbers = append(numbers, num)
			}
		}
		op := col[len(col)-1]
		task := Task{
			Numbers: numbers,
			Op:      op,
		}
		tasks = append(tasks, task)
	}

	return tasks
}

func DaySix() {
	rawData := getData("./challenge-input/day-6.txt", "\n")
	tasks := parseData(rawData)

	result := 0
	for _, task := range tasks {
		result += applyOp(task)
	}
	fmt.Println("Final Result:", result)
}
