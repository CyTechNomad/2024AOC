package day3

import (
	"fmt"
	"strconv"
	"strings"
)

func Day3() {
	fmt.Println("Day 3:")
	fmt.Println("Part 1: ", part1())
	fmt.Println("Part 2: ", part2())
}

func getParams(i string) (int, int) {
	params := strings.Split(i, ",")
	if len(params) != 2 {
		return 0, 0
	}
	if len(params[0]) > 3 || len(params[1]) > 3 {
		return 0, 0
	}
	a, b := 0, 0
	a, err := strconv.Atoi(params[0])
	if err != nil {
		return 0, 0
	}
	b, err = strconv.Atoi(params[1])
	if err != nil {
		return 0, 0
	}

	return a, b
}
