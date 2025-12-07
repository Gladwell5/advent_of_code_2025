package main

import (
	utils "advent_of_code_2025/pkg/utils"
	"bufio"
	"fmt"
	"os"
)

func main() {
	rolls := []utils.Loc{}
	line_idx := 0
	max_line_length := 0
	n_accessible := 0
	new_accessible := true

	filename := "4/4.txt"

	inputFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		line := scanner.Text()
		max_line_length = max(len(line), max_line_length)
		at_indexes := utils.FindAllIndexes(line, "@")
		for _, at_idx := range at_indexes {
			loc := utils.Loc{Row: line_idx, Col: at_idx}
			rolls = append(rolls, loc)
		}
		line_idx++
	}
	grid_shape := []int{line_idx, max_line_length}

	accessible := utils.GetAccessible(rolls, grid_shape)
	for {
		new_accessible = false
		for _, loc := range accessible {
			rolls = utils.RemoveStruct(rolls, loc)
			n_accessible += 1
			new_accessible = true
		}
		if !new_accessible {
			break
		} else {
			accessible = utils.GetAccessible(rolls, grid_shape)
		}
	}
	fmt.Print(n_accessible, "\n")
}
