package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputFile, err := os.Open("1/1.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	loc := 50
	var next_loc int
	zero_counts := 0

	for scanner.Scan() {
		line := scanner.Text()
		direction := line[:1]
		rotation, err := strconv.Atoi(line[1:])

		if err != nil {
			fmt.Println("Error converting string to integer:", err)
			return // Exit if there's an error
		}

		if direction == "L" {
			next_loc = loc - rotation
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
			}
		} else {
			next_loc = loc + rotation
			if next_loc >= 100 {
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
		}

		if next_loc == 0 {
			zero_counts += 1
		}

		loc = next_loc

		// fmt.Print(zero_counts, "\n")
	}
	fmt.Print(zero_counts, "\n")
}
