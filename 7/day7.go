package main

import (
	utils "advent_of_code_2025/pkg/utils"
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func GetNextLocs(loc int, splitters []int) (next_beam_locs []int, is_split bool) {
	if slices.Contains(splitters, loc) {
		for _, adj := range []int{-1, 1} {
			if !slices.Contains(next_beam_locs, loc+adj) {
				next_beam_locs = append(next_beam_locs, loc+adj)
			}
		}
		is_split = true
	} else {
		if !slices.Contains(next_beam_locs, loc) {
			next_beam_locs = append(next_beam_locs, loc)
		}
		is_split = false
	}
	return next_beam_locs, is_split
}

func GetNextLevel(beam_locs []int, last_level map[int]int, splitters []int) (level map[int]int, split_count int) {
	level = make(map[int]int) // maps need to be initialised
	for _, beam_loc := range beam_locs {
		path_count := last_level[beam_loc]
		next_locs, is_split := GetNextLocs(beam_loc, splitters)
		if is_split {
			split_count += 1
		}
		for _, next_loc := range next_locs {
			// if next_loc already in level, add to the count
			if slices.Contains(utils.GetKeys(level), next_loc) {
				level[next_loc] += path_count
			} else {
				level[next_loc] = path_count
			}
		}
	}
	return level, split_count
}

func main() {
	var beam_locs []int
	var splitters []int
	var total_paths int
	var total_split_count int
	var paths []map[int]int
	filename := "7/7.txt"

	inputFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "S") {
			level := make(map[int]int)
			level[utils.FindAllIndexes(line, "S")[0]] = 1
			paths = append(paths, level)
		} else if strings.Contains(line, "^") {
			splitters = utils.FindAllIndexes(line, "^")
			beam_locs = utils.GetKeys(paths[len(paths)-1])
			level, split_count := GetNextLevel(beam_locs, paths[len(paths)-1], splitters)
			total_split_count += split_count
			paths = append(paths, level)
		}
	}

	for _, value := range paths[len(paths)-1] {
		total_paths += value
	}

	fmt.Print(total_split_count, "\n")
	fmt.Print(total_paths, "\n")
}
