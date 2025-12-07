package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func GetNextLoc(loc int, rotation int, direction string) (next_loc int, zero_counts int) {
	if direction == "L" {
		next_loc = loc - rotation
	} else {
		next_loc = loc + rotation
	}
	if next_loc < 0 {
		if loc == 0 {
			zero_counts -= 1
		}
		for {
			next_loc += 100
			zero_counts += 1
			if next_loc >= 0 {
				break
			}
		}
	} else if next_loc >= 100 {
		for {
			next_loc -= 100
			zero_counts += 1
			if next_loc == 0 {
				zero_counts -= 1
				break
			} else if next_loc < 100 {
				break
			}
		}
	}
	return next_loc, zero_counts
}

func main() {
	var next_loc int
	var zero_counts int
	loc := 50
	total_zero_counts := 0

	inputFile, err := os.Open("1/1.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		line := scanner.Text()
		direction := line[:1]
		rotation, _ := strconv.Atoi(line[1:])

		next_loc, zero_counts = GetNextLoc(loc, rotation, direction)
		total_zero_counts += zero_counts

		if next_loc == 0 {
			total_zero_counts += 1
		}

		loc = next_loc
	}
	fmt.Print(total_zero_counts, "\n")
}
