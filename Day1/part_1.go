package day1

import (
	"aoc/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func part1() int {
	fmt.Println(utils.ReadInputLines("Day1"))
	listOne := make([]int, 0)
	listTwo := make([]int, 0)
	for _, line := range utils.ReadInputLines("Day1") {

		pair := strings.Split(line, "   ")
		i, err := strconv.Atoi(pair[0])
		if err != nil {
			panic(err)
		}
		listOne = append(listOne, i)
		i, err = strconv.Atoi(pair[1])
		if err != nil {
			panic(err)
		}
		listTwo = append(listTwo, i)
	}

	fmt.Println(listOne)
	fmt.Println(listTwo)

	sort.Slice(listOne, func(i, j int) bool {
		return listOne[i] < listOne[j]
	})

	sort.Slice(listTwo, func(i, j int) bool {
		return listTwo[i] < listTwo[j]
	})

	total := 0
	for i := 0; i < len(listOne); i++ {
		fmt.Println(listOne[i], " ", listTwo[i])
		total += int(math.Abs(float64(listOne[i] - listTwo[i])))
	}

	return total
}
