package day2

import (
	"aoc/utils"
	"strconv"
	"strings"
)

func part2() int {
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

		if isSafe(report) {
			safe++
			continue
		}

		for i := 0; i < len(report); i++ {
			modifiedReport := append([]int{}, report[:i]...)
			modifiedReport = append(modifiedReport, report[i+1:]...)

			if isSafe(modifiedReport) {
				safe++
				break
			}
		}
	}

	return safe
}
