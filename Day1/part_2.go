package day1

import (
	"aoc/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func part2() int {
	fmt.Println(utils.ReadInputLines("Day1"))
	listOne := make([]int, 0)
	listTwo := make([]string, 0)
	for _, line := range utils.ReadInputLines("Day1") {

		pair := strings.Split(line, "   ")
		i, err := strconv.Atoi(pair[0])
		if err != nil {
			panic(err)
		}
		listOne = append(listOne, i)
		listTwo = append(listTwo, pair[1])
	}

	fmt.Println(listOne)
	fmt.Println(listTwo)

	sort.Slice(listOne, func(i, j int) bool {
		return listOne[i] < listOne[j]
	})

	total := 0

	for i := 0; i < len(listOne); i++ {
		val := strconv.Itoa(listOne[i])
		list := strings.Join(listTwo, ",")
		c := strings.Count(list, val)
		total += (c * listOne[i])
		fmt.Println(listOne[i], " ", listTwo[i], " ", c)
	}

	return total
}
