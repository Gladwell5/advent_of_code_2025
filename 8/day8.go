package main

import (
	"advent_of_code_2025/pkg/utils"
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	var boxes [][]int
	var circuits [][]string
	var key string
	var distance_map map[string]float64

	n_connections := 1000
	filename := "8/8.txt"

	inputFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	line_idx := 0
	for scanner.Scan() {
		line := scanner.Text()
		boxes = append(boxes, utils.AsPoint(line))
		circuits = append(circuits, []string{strconv.Itoa(line_idx)})
		line_idx += 1
	}
	distance_map = utils.MakeDistMatrix(boxes)

	// part 1
	// do first n_connections connections
	for range n_connections {
		key, _ = utils.GetMinDistance(distance_map)
		circuits = append(circuits, strings.Split(key, ";"))
	}
	circuits = utils.Consolidate(circuits)

	//product of top 3 lengths of the consolidated circuits
	top3_len := []int{0, 0, 0}
	min_value := 0
	for _, circuit := range circuits {
		circuit_len := len(circuit)
		if circuit_len > min_value {
			top3_len[slices.Index(top3_len, min_value)] = circuit_len
			min_value = slices.Min(top3_len)
		}
	}
	fmt.Print(utils.Product(top3_len), "\n")

	// part 2
	last_x_product := 1
	for {
		key, _ = utils.GetMinDistance(distance_map)
		circuits = utils.Consolidate(append(circuits, strings.Split(key, ";")))
		if len(circuits) == 1 {
			for key := range strings.SplitSeq(key, ";") {
				key_int, _ := strconv.Atoi(key)
				last_x_product *= boxes[key_int][0]
			}
			break
		}
	}
	fmt.Print(last_x_product, "\n")
}
