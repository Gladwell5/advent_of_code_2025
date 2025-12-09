package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func CalculateArea(point1 []int, point2 []int) (area int) {
	if point1[0] == point2[0] {
		area = 0
	} else if point1[1] == point2[1] {
		area = 0
	} else {
		x := int(math.Abs(float64(point1[0]-point2[0])) + 1)
		y := int(math.Abs(float64(point1[1]-point2[1])) + 1)
		area = x * y
	}
	return area
}

func MakeAreasList(red_tiles [][]int) (areas_list [][]int) {
	var this_area int
	var this_entry []int
	for idx1, tile1 := range red_tiles {
		for idx2, tile2 := range red_tiles {
			if idx1 >= idx2 {
				continue
			}
			this_area = CalculateArea(tile1, tile2)
			if this_area > 0 {
				this_entry = []int{this_area, idx1, idx2}
				areas_list = append(areas_list, this_entry)
			}
		}
	}
	sort.Slice(areas_list, func(i, j int) bool {
		return areas_list[i][0] > areas_list[j][0]
	})
	return areas_list
}

func MakeVerticalSeq(x int, y []int) (seq [][]int) {
	k := y[0]
	for k != y[1] {
		seq = append(seq, []int{x, k})
		if y[0] < y[1] {
			k++
		} else {
			k--
		}
	}
	return seq
}

func MakeHorizontalSeq(x []int, y int) (seq [][]int) {
	k := x[0]
	for k != x[1] {
		seq = append(seq, []int{k, y})
		if x[0] < x[1] {
			k++
		} else {
			k--
		}
	}
	return seq
}

func GetPerimeter(corners [][]int) (perimeter [][]int) {
	var next_loc []int
	corners_copy := append(corners, corners[0])
	loc := corners_copy[0]
	for idx, tile := range corners_copy {
		if tile[0] == loc[0] {
			next_loc = corners_copy[idx]
			perimeter = append(perimeter, MakeVerticalSeq(loc[0], []int{loc[1], next_loc[1]})...)
		} else if tile[1] == loc[1] {
			next_loc = corners_copy[idx]
			perimeter = append(perimeter, MakeHorizontalSeq([]int{loc[0], next_loc[0]}, loc[1])...)
		}
		loc = next_loc
	}
	return perimeter
}

func GetColBounds(red_tiles [][]int, idx1 int, idx2 int) (col_bounds []int) {
	if red_tiles[idx1][0] < red_tiles[idx2][0] {
		col_bounds = []int{red_tiles[idx1][0], red_tiles[idx2][0]}
	} else {
		col_bounds = []int{red_tiles[idx2][0], red_tiles[idx1][0]}
	}
	return col_bounds
}

func GetRowBounds(red_tiles [][]int, idx1 int, idx2 int) (row_bounds []int) {
	if red_tiles[idx1][1] < red_tiles[idx2][1] {
		row_bounds = []int{red_tiles[idx1][1], red_tiles[idx2][1]}
	} else {
		row_bounds = []int{red_tiles[idx2][1], red_tiles[idx1][1]}
	}
	return row_bounds
}

func PerimeterWithinBounds(col_bounds []int, row_bounds []int, perimeter [][]int) (contains_perimeter bool) {
	contains_perimeter = false
	for _, per_point := range perimeter {
		col_within := col_bounds[0] < per_point[0] && per_point[0] < col_bounds[1]
		row_within := row_bounds[0] < per_point[1] && per_point[1] < row_bounds[1]
		if col_within && row_within {
			contains_perimeter = true
			break
		}
	}
	return contains_perimeter
}

func main() {
	var red_tiles [][]int

	filename := "9/9.txt"

	inputFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		line := scanner.Text()
		this_tile := []int{}
		for num := range strings.SplitSeq(line, ",") {
			num_int, _ := strconv.Atoi(num)
			this_tile = append(this_tile, num_int)
		}
		red_tiles = append(red_tiles, this_tile)
	}
	areas_list := MakeAreasList(red_tiles)

	//part 1 - largest square with corners at 2 red tiles
	fmt.Print(areas_list[0][0], "\n")

	// part 2 - largest square not including a point on the perimeter
	perimeter := GetPerimeter(red_tiles)
	area := 0
	var col_bounds []int
	var row_bounds []int
	for _, area_idx1_idx2 := range areas_list {
		area = area_idx1_idx2[0]
		col_bounds = GetColBounds(red_tiles, area_idx1_idx2[1], area_idx1_idx2[2])
		row_bounds = GetRowBounds(red_tiles, area_idx1_idx2[1], area_idx1_idx2[2])
		if !PerimeterWithinBounds(col_bounds, row_bounds, perimeter) {
			break
		}
	}

	fmt.Print(area, "\n")
}
