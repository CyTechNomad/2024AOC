package day4

import (
	"aoc/utils"
	"fmt"
)

func part2() int {
	rows := utils.ReadInputLines("Day4")

	columns := make([]string, len(rows[0]))
	for i := 0; i < len(rows[0]); i++ {
		column := ""
		for j := 0; j < len(rows); j++ {
			column += string(rows[j][i])
		}
		columns[i] = column
	}

	fDiagnals := make([]string, len(rows)+len(columns)-1)
	bDiagnals := make([]string, len(rows)+len(columns)-1)
	// diagnals
	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(columns); j++ {
			fDiagnals[i+j] += string(rows[i][j])
			bDiagnals[i+j] += string(rows[i][len(columns)-j-1])
		}
		fmt.Println(fDiagnals[i], bDiagnals[i])
	}

	return 0
}
