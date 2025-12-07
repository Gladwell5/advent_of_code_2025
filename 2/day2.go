package main

import (
	"fmt"
	"log"
	"math"
	"os"
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
		num_set := map[int]struct{}{}

		m := strings.Split(line, "-")

		upper_bound, _ := strconv.Atoi(m[1])
		lower_bound, _ := strconv.Atoi(m[0])

		// start_stem := m[0][:int(math.Floor(float64(len(m[0]))/2))]
		// start_int, _ := strconv.Atoi(start_stem)
		start_int := 1

		end_stem := m[1][:int(math.Ceil(float64(len(m[1]))/2))]
		end_int, _ := strconv.Atoi(end_stem)

		// fmt.Print(m, "\n")

		for i := start_int; i <= end_int; i++ {
			i_str := strconv.Itoa(i)
			for n := 2; n <= len(m[1]); n++ {
				i_rep, _ := strconv.Atoi(strings.Repeat(i_str, n))
				if lower_bound <= i_rep {
					if i_rep <= upper_bound {
						_, found := num_set[i_rep]
						if !found {
							// fmt.Print(i_rep, "\n")
							num_set[i_rep] = struct{}{}
							invalid_sum += i_rep
						}
					}
				}
			}
		}
	}
	fmt.Print(invalid_sum, "\n")
}
