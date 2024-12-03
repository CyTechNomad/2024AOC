package day3

import (
	"aoc/utils"
	"fmt"
	"strings"
)

func part2() int {
	lines := utils.ReadInputLines("Day3")
	sum := 0
	do := true
	for _, line := range lines {
		instructions := strings.Split(line, "mul(")
		for _, instruction := range instructions {
			commands := strings.Split(instruction, ")")
			for _, command := range commands {
				fmt.Println(command)
				if strings.Contains(command, "do(") || strings.Contains(command, "don't(") {
					i := strings.LastIndex(command, "do(")
					j := strings.LastIndex(command, "don't(")
					if i > -1 || j > -1 {
						command = command[max(i, j):]
						do = i > j
						// fmt.Println(do)
					}
				}
				if do {
					fmt.Println(command)
					params := strings.Split(command, ")")[0]
					i, j := getParams(params)
					// fmt.Println(i, j, sum)
					sum += i * j
				}
			}
		}
	}

	return sum
}
