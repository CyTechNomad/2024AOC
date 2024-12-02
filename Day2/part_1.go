package day2

import (
	"aoc/utils"
	"strconv"
	"strings"
)

func part1() int {
	safe := 0
	for _, line := range utils.ReadInputLines("Day2") {
		reportS := strings.Split(line, " ")

		if len(reportS) < 2 {
			safe++
			continue
		}

		report := make([]int, 0)
		for _, r := range reportS {
			if i, err := strconv.Atoi(r); err == nil {
				report = append(report, i)
			}
		}

		if !isSafe(report) {
			continue
		}

		safe++
	}

	return safe
}
