package utils

import (
	"strings"
)

type Loc struct {
	Row int
	Col int
}

func GetNeighbourLocs(loc Loc, size []int) (neighbour_locs []Loc) {
	for x_adj := -1; x_adj < 2; x_adj++ {
		adj_x := loc.Row + x_adj
		x_valid := 0 <= adj_x && adj_x < size[0]
		if !x_valid {
			continue
		}
		for y_adj := -1; y_adj < 2; y_adj++ {
			if x_adj == 0 && y_adj == 0 {
				continue
			}
			adj_y := loc.Col + y_adj
			y_valid := 0 <= adj_y && adj_y < size[1]
			if !y_valid {
				continue
			}
			neighbour_locs = append(neighbour_locs, Loc{adj_x, adj_y})
		}
	}
	return neighbour_locs
}

func FindAllIndexes(str string, char string) (indexes []int) {
	var idx int
	if strings.Contains(str, char) {
		idx = strings.Index(str, char)
		indexes = append(indexes, idx)
		for {
			tmp_idx := strings.Index(str[idx+1:], char)
			if tmp_idx == -1 {
				break
			}
			idx += tmp_idx + 1
			indexes = append(indexes, idx)
		}
	}
	return indexes
}

func ContainsStruct(s []Loc, e Loc) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func RemoveStruct(list []Loc, loc Loc) []Loc {
	idx := 0
	for idx < len(list) {
		if list[idx] == loc {
			break
		}
		idx++
	}
	if idx == len(list) {
		return list
	} else {
		return append(list[:idx], list[idx+1:]...)
	}
}

func GetAccessible(rolls []Loc, grid_shape []int) (accessible []Loc) {
	neighbour_count := 0
	for _, loc := range rolls {
		neighbour_count = 0
		for _, neigh_loc := range GetNeighbourLocs(loc, grid_shape) {
			if ContainsStruct(rolls, neigh_loc) {
				neighbour_count += 1
			}
		}
		if neighbour_count < 4 {
			accessible = append(accessible, loc)
		}
	}
	return accessible
}
