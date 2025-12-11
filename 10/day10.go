package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseLine(line string) (lights []int, buttons [][]int, joltages []int) {
	for _, element := range strings.Split(line, " ") {
		if string(element[0]) == "[" {
			for _, item := range element[1 : len(element)-1] {
				if string(item) == "#" {
					lights = append(lights, 1)
				} else {
					lights = append(lights, 0)
				}
			}
		} else if string(element[0]) == "(" {
			button := []int{}
			for _, item := range strings.Split(element[1:len(element)-1], ",") {
				item_int, _ := strconv.Atoi(item)
				button = append(button, item_int)
			}
			buttons = append(buttons, button)
		} else if string(element[0]) == "{" {
			for _, item := range strings.Split(element[1:len(element)-1], ",") {
				item_int, _ := strconv.Atoi(item)
				joltages = append(joltages, item_int)
			}
		}
	}
	return lights, buttons, joltages
}

func StateAsKey(state []int) (key string) {
	str_items := []string{}
	for _, item := range state {
		str_items = append(str_items, strconv.Itoa(item))
	}
	key = strings.Join(str_items, ";")
	return key
}

func KeyAsState(key string) (state []int) {
	for item := range strings.SplitSeq(key, ";") {
		item_int, _ := strconv.Atoi(string(item))
		state = append(state, item_int)
	}
	return state
}

func ToggleLights(lights []int, button []int) []int {
	for _, index := range button {
		lights[index] = 1 - lights[index]
	}
	return lights
}

func main() {
	filename := "10/10.txt"
	total_presses := 0

	inputFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	var state []int
	var target []int
	var new_state []int
	var buttons [][]int

	line_num := 0
	for scanner.Scan() {
		line := scanner.Text()
		target, buttons, _ = ParseLine(line)
		states := make(map[string]int)
		start := make([]int, len(target))
		states[StateAsKey((start))] = 0
		for {
			updates := make(map[string]int)
			for state_key, presses := range states {
				for _, button := range buttons {
					state = KeyAsState(state_key)
					new_state = ToggleLights(state, button)
					overshot := false
					for idx, item := range new_state {
						if item > target[idx] {
							overshot = true
							break
						}
					}
					if !overshot {
						updates[StateAsKey(new_state)] = presses + 1
					}
				}
			}
			for state_key, presses := range updates {
				if states[state_key] == 0 {
					states[state_key] = presses
				}
			}
			if states[StateAsKey(target)] > 0 {
				break
			}
		}
		total_presses += states[StateAsKey(target)]
		line_num += 1
		fmt.Print("\r", line_num)
	}
	fmt.Print("\n", total_presses, "\n")
}
