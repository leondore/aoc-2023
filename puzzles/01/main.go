package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/leondore/aoc-2023/utils"
)

func main() {
	inputData := utils.ReadFile("puzzles/01/input.txt")
	lines := strings.Split(inputData, "\n")

	re := regexp.MustCompile(`[^0-9]`)

	var result int64

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		digits := re.ReplaceAllString(line, "")

		calibrationValueS := string(digits[0]) + string(digits[len(digits)-1])
		calibrationValue, err := strconv.ParseInt(calibrationValueS, 10, 64)
		utils.CheckError(err)

		result += calibrationValue
	}

	fmt.Println(result)
}
