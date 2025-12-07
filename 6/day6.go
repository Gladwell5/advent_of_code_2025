package main

import (
	utils "advent_of_code_2025/pkg/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RunOperation(values []int, operator string) (result int) {
	if strings.TrimSpace(operator) == "+" {
		result = utils.Sum(values)
	} else if strings.TrimSpace(operator) == "*" {
		result = utils.Product(values)
	}
	return result
}

func GetVerticalNumber(verticals []string, idx int) (int, error) {
	vertical_num := ""
	for _, vertical := range verticals {
		vertical_num += string(vertical[idx])
	}
	return strconv.Atoi(strings.TrimSpace(vertical_num))
}

func main() {
	var operators string
	var verticals []string
	var values []int

	filename := "6/6.txt"
	max_line_len := 0
	operator := ""
	running_total := 0

	inputFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		line := scanner.Text()
		max_line_len = max(len(line), max_line_len)
		if strings.Contains(line, "+") || strings.Contains(line, "*") {
			operators = line
		} else {
			verticals = append(verticals, line)
		}
	}
	// pad operations if shorter than verticals
	for len(operators) < max_line_len {
		operators += " "
	}

	// work from right to left with idx and then down
	// the rows with the GetVerticalNumber function
	for idx := max_line_len - 1; idx >= 0; idx-- {
		vertical_int, err := GetVerticalNumber(verticals, idx)
		// err in converting to int happens when entire vertical is space
		// this triggers the run of the current operation and a reset of operator/values
		if err != nil {
			running_total += RunOperation(values, operator)
			values = []int{}
			operator = ""
		} else {
			operator += string(operators[idx])
			values = append(values, vertical_int)
		}
	}
	running_total += RunOperation(values, operator)

	fmt.Print(running_total, "\n")
}
