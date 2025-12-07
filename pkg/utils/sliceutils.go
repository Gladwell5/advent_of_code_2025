package utils

import (
	"strconv"
)

func Overlap(range1 []int, range2 []int) (overlapping bool, combined_range []int) {
	overlapping = true
	if range1[0] > range2[1] || range1[1] < range2[0] {
		overlapping = false
	}
	if overlapping {
		lower_bound := min(range1[0], range2[0])
		upper_bound := max(range1[1], range2[1])
		combined_range = append(combined_range, lower_bound, upper_bound)
	}
	return overlapping, combined_range
}

func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func DropIndexes(in_list [][]int, indexes []int) (out_list [][]int) {
	for i, x := range in_list {
		if Contains(indexes, i) {
			continue
		} else {
			out_list = append(out_list, x)
		}
	}
	return out_list
}

func GetMaxIndex(x string) (max_x_i int) {
	max_x := 0
	max_x_i = -1
	for i := 0; i < len(x); i++ {
		x_int, _ := strconv.Atoi(string(x[i]))
		if x_int > max_x {
			max_x = x_int
			max_x_i = i
		}
	}
	return max_x_i
}

func Product(num_list []int) int {
	prod_num := 1
	for _, num := range num_list {
		prod_num *= num
	}
	return prod_num
}

func Sum(num_list []int) int {
	sum_num := 0
	for _, num := range num_list {
		sum_num += num
	}
	return sum_num
}
