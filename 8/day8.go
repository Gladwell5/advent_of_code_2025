package main

import (
	"advent_of_code_2025/pkg/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func CalculateDistance(box1 []int, box2 []int) (distance float64) {
	for idx := range 3 {
		distance += math.Pow(float64(box1[idx]-box2[idx]), 2)
	}
	distance = math.Sqrt(distance)
	return distance
}

func MakeBox(line string) (box []int) {
	var num_int int
	for _, num := range strings.Split(line, ",") {
		num_int, _ = strconv.Atoi(num)
		box = append(box, num_int)
	}
	return box
}

func MakeDistMatrix(boxes [][]int) (distance_map map[string]float64) {
	distance_map = make(map[string]float64)
	var distance float64
	var key string
	for idx1 := range len(boxes) {
		for idx2 := range len(boxes) {
			if idx1 >= idx2 {
				continue
			}
			distance = CalculateDistance(boxes[idx1], boxes[idx2])
			key = strconv.Itoa(idx1) + ";" + strconv.Itoa(idx2)
			distance_map[key] = distance
		}
	}
	return distance_map
}

func GetMinDistance(distance_map map[string]float64) (key string, value float64) {
	value = math.Inf(1)
	for k, v := range distance_map {
		if v < value {
			key = k
			value = v
		}
	}
	delete(distance_map, key)
	return key, value
}

func Intersect(set1 []string, set2 []string) (intersection []string) {
	for _, element1 := range set1 {
		for _, element2 := range set2 {
			if element1 == element2 {
				intersection = append(intersection, element1)
			}
		}
	}
	return intersection
}

func Union(set1 []string, set2 []string) (union []string) {
	for _, element1 := range set1 {
		if !slices.Contains(union, element1) {
			union = append(union, element1)
		}
	}
	for _, element2 := range set2 {
		if !slices.Contains(union, element2) {
			union = append(union, element2)
		}
	}
	return union
}

func MergeCircuits(circuits [][]string) (merged_circuits [][]string) {
	idx1 := 1
	merged_circuits = [][]string{}
	for idx1 < len(circuits) {
		for idx2 := idx1 - 1; idx2 >= 0; idx2-- {
			intersection := Intersect(circuits[idx1], circuits[idx2])
			if len(intersection) > 0 {
				circuits[idx1] = Union(circuits[idx1], circuits[idx2])
				circuits[idx2] = []string{}
			}
		}
		idx1 += 1
	}
	for _, circuit := range circuits {
		if len(circuit) > 0 {
			merged_circuits = append(merged_circuits, circuit)
		}
	}
	return merged_circuits
}

func main() {
	var boxes [][]int
	var circuits [][]string
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
		boxes = append(boxes, MakeBox(line))
		circuits = append(circuits, []string{strconv.Itoa(line_idx)})
		line_idx += 1
	}
	// n_lines := line_idx
	// fmt.Print(n_lines, "\n")

	distance_map := MakeDistMatrix(boxes)
	for range n_connections {
		min_key, _ := GetMinDistance(distance_map)
		circuits = append(circuits, strings.Split(min_key, ";"))
	}
	merged_circuits := MergeCircuits(circuits)
	circuit_lengths := []int{}
	for _, circuit := range merged_circuits {
		circuit_lengths = append(circuit_lengths, len(circuit))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(circuit_lengths)))
	fmt.Print(utils.Product(circuit_lengths[:3]), "\n")

	iter_count := 1
	copy(circuits, merged_circuits)
	// circuits = [][]string{}
	// for idx := range n_lines {
	// 	circuits = append(circuits, []string{strconv.Itoa(idx)})
	// }
	// print(len(circuits), "\n")
	last_x_product := 1
	for {
		min_key, _ := GetMinDistance(distance_map)
		circuits = append(circuits, strings.Split(min_key, ";"))
		merged_circuits = MergeCircuits(circuits)
		fmt.Print("\r", iter_count, " ", len(merged_circuits), "    ")
		if len(merged_circuits) > 1 {
			copy(circuits, merged_circuits)
		} else {
			// fmt.Print("\n", min_key, "\n")
			for _, key := range strings.Split(min_key, ";") {
				key_int, _ := strconv.Atoi(key)
				last_x_product *= boxes[key_int][0]
			}
			break
		}
		iter_count += 1
	}
	fmt.Print("x product of last two boxes:", last_x_product, "\n")
}
