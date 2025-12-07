package main

import (
	utils "advent_of_code_2025/pkg/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RunOp(values []int, op string) (result int) {
	if strings.TrimSpace(op) == "+" {
		result = utils.Sum(values)
	} else if strings.TrimSpace(op) == "*" {
		result = utils.Product(values)
	}
	return result
}

func GetVerticalNumber(verticals []string, idx int) (int, error) {
	vert_num := ""
	for _, vertical := range verticals {
		vert_num += string(vertical[idx])
	}
	return strconv.Atoi(strings.TrimSpace(vert_num))
}

func main() {
	filename := "6/6.txt"
	var operators string
	var verticals []string

	inputFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	max_len := 0
	for scanner.Scan() {
		line := scanner.Text()
		max_len = max(len(line), max_len)
		if strings.Contains(line, "+") || strings.Contains(line, "*") {
			operators = line
		} else {
			verticals = append(verticals, line)
		}
	}
	for len(operators) < max_len {
		operators += " "
	}

	running_total := 0
	var values []int
	op := ""
	for idx := max_len - 1; idx >= 0; idx-- {
		op += string(operators[idx])
		vertical_int, err := GetVerticalNumber(verticals, idx)
		// err in converting to int happens when entire vertical is space
		// this is the signal to run the current operation and reset op/values
		if err != nil {
			running_total += RunOp(values, op)
			values = []int{}
			op = ""
		} else {
			values = append(values, vertical_int)
		}
	}
	running_total += RunOp(values, op)

	fmt.Print(running_total, "\n")
}
