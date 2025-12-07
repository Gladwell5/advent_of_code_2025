package main

import (
	utils "advent_of_code_2025/pkg/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	filename := "3/3.txt"

	inputFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	total_joltage := 0
	n_batteries := 12

	for scanner.Scan() {
		line := scanner.Text()
		joltage_str := string("")
		joltage := 0
		from_index := 0
		to_index := 0
		last_index := -1
		battery_index := 0
		for i := 1; i <= n_batteries; i++ {
			from_index = (last_index + 1)
			to_index = len(line) - (n_batteries - i)
			battery_index = utils.GetMaxIndex(string(line[from_index:to_index])) + from_index
			joltage_str += string(line[battery_index])
			last_index = battery_index
		}
		// fmt.Print(joltage_str, "\n")
		joltage, _ = strconv.Atoi(joltage_str)
		total_joltage += joltage
	}
	fmt.Print(total_joltage, "\n")
}
