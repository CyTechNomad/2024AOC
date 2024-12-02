package day2

import (
	"fmt"
	"math"
)

func Day2() {
	fmt.Println("Day 2:")
	fmt.Println("Part 1: ", part1())
	fmt.Println("Part 2: ", part2())
}

func isAccending(report []int) bool {
	aPoints := 0
	dPoints := 0
	for i := 1; i < len(report); i++ {
		if report[i] > report[i-1] {
			aPoints++
		} else if report[i] < report[i-1] {
			dPoints++
		}
	}
	return aPoints > dPoints
}

func atLeastOneStepAppart(report []int) bool {
	for i := 1; i < len(report); i++ {
		if report[i] == report[i-1] {
			return false
		}
	}
	return true
}

func isConsistent(report []int) bool {
	Accending := isAccending(report)
	for i := 1; i < len(report); i++ {
		if Accending != (report[i] > report[i-1]) {
			return false
		}
	}

	return true
}

func isWithinRange(report []int, r int) bool {
	for i := 1; i < len(report); i++ {
		if int(math.Abs(float64(report[i]-report[i-1]))) > r {
			return false
		}
	}
	return true
}

func isSafe(report []int) bool {
	return atLeastOneStepAppart(report) && isConsistent(report) && isWithinRange(report, 3)
}
