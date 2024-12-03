package day3

import (
	"aoc/utils"
	"strings"
)

func part1() int {
	lines := utils.ReadInputLines("Day3")
	sum := 0
	for _, line := range lines {
		instructions := strings.Split(line, "mul(")
		for _, instruction := range instructions {
			if strings.Contains(instruction, ")") {
				params := strings.Split(instruction, ")")[0]
				i, j := getParams(params)
				sum += i * j
			}
		}
	}

	return sum
}
