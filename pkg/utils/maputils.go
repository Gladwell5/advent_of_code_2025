package utils

func GetKeys(myMap map[int]int) []int {
	keys := make([]int, 0, len(myMap))
	for key := range myMap {
		keys = append(keys, key)
	}
	return keys
}
