package utils

import "slices"

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

func Consolidate(circuits [][]string) (merged_circuits [][]string) {
	merged_circuits = [][]string{}
	for idx1 := 1; idx1 < len(circuits); idx1++ {
		for idx2 := idx1 - 1; idx2 >= 0; idx2-- {
			if len(Intersect(circuits[idx1], circuits[idx2])) > 0 {
				circuits[idx1] = Union(circuits[idx1], circuits[idx2])
				circuits[idx2] = []string{}
			}
		}
	}
	for _, circuit := range circuits {
		if len(circuit) > 0 {
			merged_circuits = append(merged_circuits, circuit)
		}
	}
	return merged_circuits
}
