package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

type Task struct {
	Numbers []int
	Op      string
}

func (t *Task) Apply() int {
	result := 0

	switch t.Op {
	case "+":
		for _, num := range t.Numbers {
			result += num
		}
	case "-":
		for _, num := range t.Numbers {
			result -= num
		}
	case "*":
		result = t.Numbers[0]
		for _, num := range t.Numbers[1:] {
			result *= num
		}
	case "/":
		result = t.Numbers[0]
		for _, num := range t.Numbers[1:] {
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

func DaySix() {
	rawData := getData("./challenge-input/day-6.txt", "\n")

	blocks := splitVerticalBlocks(rawData)
	part1Tasks := part1TaskParser(blocks)
	part2Tasks := part2TaskParser(blocks)

	part1 := 0
	for _, task := range part1Tasks {
		part1 += task.Apply()
	}

	part2 := 0
	for _, task := range part2Tasks {
		part2 += task.Apply()
	}

	fmt.Println("Final Result:", part1)
	fmt.Println("Part 2 Result:", part2)
}

func part1TaskParser(blocks [][]string) []Task {
	tasks := []Task{}
	for _, block := range blocks {
		nums, op := parseSequentially(block)
		tasks = append(tasks, Task{nums, op})
	}
	return tasks
}

func part2TaskParser(blocks [][]string) []Task {
	nums := [][]int{}
	ops := []string{}

	for _, block := range blocks {
		op := block[len(block)-1]
		op = strings.TrimSpace(op)
		ops = append(ops, op)
		numbers := extractNumbersFromBlock(block)
		nums = append(nums, numbers)
	}

	return createTasks(nums, ops)
}

func splitVerticalBlocks(lines []string) [][]string {
	if len(lines) == 0 {
		return nil
	}

	var grid [][]rune
	maxWidth := 0
	for _, line := range lines {
		r := []rune(line)
		if len(r) > maxWidth {
			maxWidth = len(r)
		}
		grid = append(grid, r)
	}

	for i := range grid {
		for len(grid[i]) < maxWidth {
			grid[i] = append(grid[i], ' ')
		}
	}

	var blocks [][]string
	startCol := -1

	for col := 0; col < maxWidth; col++ {
		isGap := true
		for row := 0; row < len(grid); row++ {
			if grid[row][col] != ' ' {
				isGap = false
				break
			}
		}

		if !isGap {
			if startCol == -1 {
				startCol = col
			}
		} else {
			if startCol != -1 {
				blocks = append(blocks, extractSlice(grid, startCol, col))
				startCol = -1
			}
		}
	}

	if startCol != -1 {
		blocks = append(blocks, extractSlice(grid, startCol, maxWidth))
	}

	return blocks
}

func parseSequentially(block []string) ([]int, string) {
	nums := []int{}
	op := strings.TrimSpace(block[len(block)-1])
	for _, line := range block[:len(block)-1] {
		trimmed := strings.TrimSpace(line)
		converted, err := strconv.Atoi(trimmed)
		if err == nil {
			nums = append(nums, converted)
		}
	}
	return nums, op
}

func extractNumbersFromBlock(block []string) []int {
	if len(block) == 0 {
		return nil
	}

	maxWidth := 0
	for _, line := range block {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}

	var numbers []int

	for col := 0; col < maxWidth; col++ {
		var digitBuffer strings.Builder

		for _, line := range block {
			if col < len(line) {
				char := rune(line[col])
				if unicode.IsDigit(char) {
					digitBuffer.WriteRune(char)
				}
			}
		}

		if digitBuffer.Len() > 0 {
			num, err := strconv.Atoi(digitBuffer.String())
			if err == nil {
				numbers = append(numbers, num)
			}
		}
	}

	slices.Reverse(numbers)
	return numbers
}

func createTasks(nums [][]int, ops []string) []Task {
	tasks := []Task{}

	for i := range nums {
		task := Task{
			Numbers: nums[i],
			Op:      ops[i],
		}
		tasks = append(tasks, task)
	}
	return tasks
}

func extractSlice(grid [][]rune, start, end int) []string {
	var block []string
	for _, row := range grid {
		block = append(block, string(row[start:end]))
	}
	return block
}
