package day4

import (
	"aoc/utils"
	"strings"
)

const (
	searchTerm          = "XMAS"
	backwardsSearchTerm = "SAMX"
)

func part1() int {
	rows := utils.ReadInputLines("Day4")
	rowTotal := 0
	for _, row := range rows {
		if strings.Contains(row, searchTerm) || strings.Contains(row, backwardsSearchTerm) {
			// fmt.Println(row)
		}
		rowTotal += strings.Count(row, searchTerm)
		rowTotal += strings.Count(row, backwardsSearchTerm)
	}
	// fmt.Println("Found: ", rowTotal)

	// fmt.Println()

	columns := make([]string, len(rows[0]))
	columnTotal := 0
	for i := 0; i < len(rows[0]); i++ {
		column := ""
		for j := 0; j < len(rows); j++ {
			column += string(rows[j][i])
		}
		columns[i] = column
		if strings.Contains(column, searchTerm) || strings.Contains(column, backwardsSearchTerm) {
			// fmt.Println(column)
		}
		columnTotal += strings.Count(column, searchTerm)
		columnTotal += strings.Count(column, backwardsSearchTerm)
	}
	// fmt.Println("Found: ", columnTotal)

	fDiagnals := make([]string, len(rows)+len(columns)-1)
	// forwards diagnals
	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(columns); j++ {
			fDiagnals[i+j] += string(rows[i][j])
		}
	}

	// backwards diagnals
	bDiagnals := make([]string, len(rows)+len(columns)-1)
	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(columns); j++ {
			bDiagnals[i+j] += string(rows[i][len(columns)-j-1])
		}
	}

	diagnals := append(fDiagnals, bDiagnals...)

	// fmt.Println(diagnals)
	diagnalTotal := 0
	for _, diagnal := range diagnals {
		// if strings.Contains(diagnal, searchTerm) || strings.Contains(diagnal, backwardsSearchTerm) {
		// 	fmt.Println(diagnal)
		// }

		diagnalTotal += strings.Count(diagnal, searchTerm)
		diagnalTotal += strings.Count(diagnal, backwardsSearchTerm)
	}
	// fmt.Println("Found: ", diagnalTotal)

	return rowTotal + columnTotal + diagnalTotal
}
