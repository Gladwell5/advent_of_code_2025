package utils

import (
	"math"
	"strconv"
)

func CalculateDistance(box1 []int, box2 []int) (distance float64) {
	for idx := range 3 {
		distance += math.Pow(float64(box1[idx]-box2[idx]), 2)
	}
	distance = math.Sqrt(distance)
	return distance
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
	// gets AND REMOVES key-value for min value
	delete(distance_map, key)
	return key, value
}
