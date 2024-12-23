package day1

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func Day1Part1(input string) int {
	lines := strings.Split(input, "\n")
	lefts := []int{}
	rights := []int{}
	for _, line := range lines {
		numbers := strings.Fields(line)
		if len(numbers) == 0 {
			break
		}
		left, _ := strconv.Atoi(numbers[0])
		lefts = append(lefts, left)
		right, _ := strconv.Atoi(numbers[1])
		rights = append(rights, right)
	}
	sort.Ints(lefts)
	sort.Ints(rights)
	distances := []int{}
	for i, left := range lefts {
		right := rights[i]
		distances = append(distances, int(math.Abs(float64(left-right))))
	}
	total := 0
	for _, distance := range distances {
		total += distance
	}
	return total
}

func Day1Part2(input string) int {
	lines := strings.Split(input, "\n")
	lefts := make([]int, 0)
	rights := make([]int, 0)
	for _, line := range lines {
		numbers := strings.Fields(line)
		if len(numbers) == 0 {
			break
		}
		left, _ := strconv.Atoi(numbers[0])
		lefts = append(lefts, left)
		right, _ := strconv.Atoi(numbers[1])
		rights = append(rights, right)
	}
	sort.Ints(lefts)
	sort.Ints(rights)
	total := 0
	for _, left := range lefts {
		rightCount := 0
		for _, right := range rights {
			if right == left {
				rightCount++
			}
		}
		total += left * rightCount
	}
	return total
}
