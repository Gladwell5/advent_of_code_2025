package main

import (
	utils "advent_of_code_2025/pkg/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := "5/5.txt"
	var fresh [][]int
	var ingredients []int

	inputFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "-") {
			from_to := strings.Split(line, "-")
			from, _ := strconv.Atoi(from_to[0])
			to, _ := strconv.Atoi(from_to[1])
			this_range := []int{from, to}
			fresh = append(fresh, this_range)
		} else if line != "" {
			ingredient, _ := strconv.Atoi(line)
			ingredients = append(ingredients, ingredient)
		}
	}

	for {
		last_len := len(fresh)
		for idx1 := range last_len {
			// since fresh can reduce in size
			// skip a loop if indexes are no longer valid
			if idx1 >= len(fresh) {
				continue
			}
			for idx2 := range last_len {
				// skip also when indexes are the same
				if idx2 >= len(fresh) || idx1 == idx2 {
					continue
				}
				// if ranges overlap, drop them and add a new range
				// with bounds that cover the full combined range
				is_overlapping, combined_range := utils.Overlap(fresh[idx1], fresh[idx2])
				if is_overlapping {
					fresh = utils.DropIndexes(fresh, []int{idx1, idx2})
					fresh = append(fresh, combined_range)
				}
			}
		}
		// no change in number of ranges indicates no further
		// range merges are possible
		if len(fresh) == last_len {
			break
		}
	}

	fresh_ids := 0
	for _, rng := range fresh {
		fresh_ids += rng[1] - rng[0] + 1
	}

	n_fresh := 0
	for _, ingredient := range ingredients {
		for _, rng := range fresh {
			if rng[0] <= ingredient && ingredient <= rng[1] {
				n_fresh += 1
				break
			}
		}
	}

	fmt.Print(fresh_ids, "\n")
	fmt.Print(n_fresh, "\n")
}
