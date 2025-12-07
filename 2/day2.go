package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	filename := "2/2.txt"

	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), ",")
	invalid_sum := 0
	for _, line := range lines {
		var num_set []int

		rng := strings.Split(line, "-")
		upper_bound, _ := strconv.Atoi(rng[1])
		lower_bound, _ := strconv.Atoi(rng[0])

		// max stem is first half of upper range value
		// max value is max stam repeated twice
		end_stem := rng[1][:int(math.Ceil(float64(len(rng[1]))/2))]
		end_int, _ := strconv.Atoi(end_stem)

		for i := 1; i <= end_int; i++ {
			for n := 2; n <= len(rng[1]); n++ {
				i_rep, _ := strconv.Atoi(strings.Repeat(strconv.Itoa(i), n))
				if lower_bound <= i_rep && i_rep <= upper_bound {
					if !slices.Contains(num_set, i_rep) {
						num_set = append(num_set, i_rep)
						invalid_sum += i_rep
					}
				}
			}
		}
	}
	fmt.Print(invalid_sum, "\n")
}
