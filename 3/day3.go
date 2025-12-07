package main

import (
	utils "advent_of_code_2025/pkg/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func GetMaxJoltage(joltages string, n_batteries int) (joltage int) {
	joltage_str := ""
	from_idx := 0
	to_idx := 0
	last_idx := -1
	battery_idx := 0
	for i := 1; i <= n_batteries; i++ {
		from_idx = (last_idx + 1)
		to_idx = len(joltages) - (n_batteries - i)
		battery_idx = utils.GetMaxIndex(string(joltages[from_idx:to_idx])) + from_idx
		joltage_str += string(joltages[battery_idx])
		last_idx = battery_idx
	}
	joltage, _ = strconv.Atoi(joltage_str)
	return joltage
}

func main() {
	var joltage int
	total_joltage := 0
	n_batteries := 12

	filename := "3/3.txt"

	inputFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		line := scanner.Text()
		joltage = GetMaxJoltage(line, n_batteries)
		total_joltage += joltage
	}
	fmt.Print(total_joltage, "\n")
}
