package main

import (
	"fmt"
	"strings"

	"github.com/leondore/aoc-2023/utils"
)

func main() {
	inputData := utils.ReadFile("puzzles/01/test.txt")
	lines := strings.Split(inputData, "\n")

	for _, line := range lines {
		fmt.Println(line)
	}
}
