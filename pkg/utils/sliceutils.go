package utils

import (
	"slices"
	"strconv"
	"strings"
)

func Overlap(range1 []int, range2 []int) (overlapping bool, combined_range []int) {
	overlapping = true
	if range1[0] > range2[1] || range1[1] < range2[0] {
		overlapping = false
	}
	if overlapping {
		lower_bound := min(range1[0], range2[0])
		upper_bound := max(range1[1], range2[1])
		combined_range = []int{lower_bound, upper_bound}
	}
	return overlapping, combined_range
}

func DropIndexes(in_list [][]int, indexes []int) (out_list [][]int) {
	for i, x := range in_list {
		if slices.Contains(indexes, i) {
			continue
		} else {
			out_list = append(out_list, x)
		}
	}
	return out_list
}

func ListAsKey(list []int) (key string) {
	str_items := []string{}
	for _, item := range list {
		str_items = append(str_items, strconv.Itoa(item))
	}
	key = strings.Join(str_items, ";")
	return key
}

func Deduplicate(in_list [][]int) (out_list [][]int) {
	added := []string{}
	for _, x := range in_list {
		x_key := ListAsKey(x)
		if slices.Contains(added, x_key) {
			continue
		} else {
			out_list = append(out_list, x)
			added = append(added, x_key)
		}
	}
	return out_list
}

func ContainsElement(list [][]int, element []int) bool {
	x_keys := []string{}
	for _, x := range list {
		x_keys = append(x_keys, ListAsKey(x))
	}
	return slices.Contains(x_keys, ListAsKey(element))
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

func Product(num_list []int) (prod_num int) {
	prod_num = 1
	for _, num := range num_list {
		prod_num *= num
	}
	return prod_num
}

func Sum(num_list []int) (sum_num int) {
	for _, num := range num_list {
		sum_num += num
	}
	return sum_num
}

func AsPoint(line string) (point []int) {
	var num_int int
	for num := range strings.SplitSeq(line, ",") {
		num_int, _ = strconv.Atoi(num)
		point = append(point, num_int)
	}
	return point
}
