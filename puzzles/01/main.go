package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/leondore/aoc-2023/utils"
)

func main() {
	inputData := utils.ReadFile("puzzles/01/input.txt")
	lines := strings.Split(inputData, "\n")

	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	var result int64

	type numidx struct {
		Number string
		Index  int
	}

	for _, line := range lines {
		var numsFound []numidx
		for _, number := range numbers {
			idxf := strings.Index(line, number)
			if idxf != -1 {
				numsFound = append(numsFound, numidx{Number: number, Index: idxf})
			}
			idxl := strings.LastIndex(line, number)
			if idxl != -1 {
				numsFound = append(numsFound, numidx{Number: number, Index: idxl})
			}
		}

		if len(numsFound) == 0 {
			continue
		}

		sort.Slice(numsFound, func(i, j int) bool {
			return numsFound[i].Index < numsFound[j].Index
		})

		calibrationValueS := convertToDigit(numsFound[0].Number) + convertToDigit(numsFound[len(numsFound)-1].Number)
		calibrationValue, err := strconv.ParseInt(calibrationValueS, 10, 64)
		utils.CheckError(err)

		result += calibrationValue
	}

	fmt.Println(result)
}

func convertToDigit(number string) string {
	numbersMap := make(map[string]string)
	numbersMap["one"] = "1"
	numbersMap["two"] = "2"
	numbersMap["three"] = "3"
	numbersMap["four"] = "4"
	numbersMap["five"] = "5"
	numbersMap["six"] = "6"
	numbersMap["seven"] = "7"
	numbersMap["eight"] = "8"
	numbersMap["nine"] = "9"

	if _, ok := numbersMap[number]; !ok {
		return number
	}

	return numbersMap[number]
}
