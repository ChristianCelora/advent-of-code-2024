package main

import (
	"adventcode/reader"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func CalcDistance(list_sx []int, list_dx []int) int {
	var distance int

	sort.Sort(sort.IntSlice(list_sx))
	sort.Sort(sort.IntSlice(list_dx))

	for idx := range list_sx {
		distance += int(math.Abs(float64(list_dx[idx] - list_sx[idx])))
	}

	return distance
}

func CalcSimilarity(list_sx []int, list_dx []int) int {
	var similarity int
	occurrences := make(map[int]int)

	for _, dx := range list_dx {
		if _, ok := occurrences[dx]; ok {
			occurrences[dx] += 1
		} else {
			occurrences[dx] = 1
		}
	}

	for _, sx := range list_sx {
		if _, ok := occurrences[sx]; ok {
			similarity += sx * occurrences[sx]
		}
	}

	return similarity
}

func main() {
	var list_sx, list_dx []int
	var distance, similarity int
	lines := reader.ReadLines("./day01/data/input1_2.txt")
	for _, line := range lines {
		values := strings.Split(line, "   ")

		sx, err := strconv.Atoi(values[0])
		if err != nil {
			panic(err)
		}
		dx, err := strconv.Atoi(values[1])
		if err != nil {
			panic(err)
		}

		list_sx = append(list_sx, sx)
		list_dx = append(list_dx, dx)
	}

	distance = CalcDistance(list_sx, list_dx)
	fmt.Printf("Distance is %d\n", distance)

	similarity = CalcSimilarity(list_sx, list_dx)
	fmt.Printf("similarity is %d\n", similarity)
}
