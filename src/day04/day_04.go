package main

import (
	"adventcode/reader"
	"fmt"
)

func countAllXmasOccurences(matrix [][]byte) (int, int) {
	var counterWords, counterXshaped int

	for i, row := range matrix {
		for j, cell := range row {
			if cell == 'X' {
				counterWords += countXmasOccurencesFromPoint(matrix, i, j)
			} else if cell == 'A' {
				counterXshaped += countXShapedOccurencesFromPoint(matrix, i, j)

			}
		}
	}

	return counterWords, counterXshaped
}

func countXmasOccurencesFromPoint(matrix [][]byte, i int, j int) int {
	var counter int
	var xmas_length int
	xmas_length = 4

	// down
	if i+3 <= len(matrix)-1 {
		res := ""
		for k := 0; k < xmas_length; k++ {
			key := i + k
			res += string(matrix[key][j])
		}
		if res == "XMAS" {
			// fmt.Println("found down", i, j)
			counter++
		}
	}

	// up
	if i-3 >= 0 {
		res := ""
		for k := 0; k < xmas_length; k++ {
			key := i - k
			res += string(matrix[key][j])
		}
		if res == "XMAS" {
			// fmt.Println("found up", i, j)
			counter++
		}
	}

	// backwards
	if j-3 >= 0 {
		res := ""
		for k := 0; k < xmas_length; k++ {
			key := j - k
			res += string(matrix[i][key])
		}
		if res == "XMAS" {
			// fmt.Println("found backwards", i, j)
			counter++
		}
	}

	// forward
	if j+3 <= len(matrix[0])-1 {
		res := ""
		for k := 0; k < xmas_length; k++ {
			key := j + k
			res += string(matrix[i][key])
		}
		if res == "XMAS" {
			// fmt.Println("found forward", i, j)
			counter++
		}
	}

	// diagonal 1
	if i+3 <= len(matrix)-1 && j+3 <= len(matrix[0])-1 {
		res := ""
		for k := 0; k < xmas_length; k++ {
			key1 := i + k
			key2 := j + k
			res += string(matrix[key1][key2])
		}
		if res == "XMAS" {
			// fmt.Println("found diagonal 1", i, j)
			counter++
		}
	}

	// diagonal 2
	if i-3 >= 0 && j+3 <= len(matrix[0])-1 {
		res := ""
		for k := 0; k < xmas_length; k++ {
			key1 := i - k
			key2 := j + k
			res += string(matrix[key1][key2])
		}
		if res == "XMAS" {
			// fmt.Println("found diagonal 2", i, j)
			counter++
		}
	}

	// diagonal 3
	if i+3 <= len(matrix)-1 && j-3 >= 0 {
		res := ""
		for k := 0; k < xmas_length; k++ {
			key1 := i + k
			key2 := j - k
			res += string(matrix[key1][key2])
		}
		if res == "XMAS" {
			// fmt.Println("found diagonal 3", i, j)
			counter++
		}
	}

	// diagonal 4
	if j-3 >= 0 && i-3 >= 0 {
		res := ""
		for k := 0; k < xmas_length; k++ {
			key1 := i - k
			key2 := j - k
			res += string(matrix[key1][key2])
		}
		if res == "XMAS" {
			// fmt.Println("found diagonal 4", i, j)
			counter++
		}
	}

	return counter
}

func countXShapedOccurencesFromPoint(matrix [][]byte, i int, j int) int {
	var counter int
	var diag1_ok, diag2_ok bool

	// diagonal 1
	if j-1 >= 0 && i-1 >= 0 && i+1 <= len(matrix)-1 && j+1 <= len(matrix[0])-1 {
		if matrix[i-1][j-1] == 'M' && matrix[i+1][j+1] == 'S' {
			diag1_ok = true
		}

		if matrix[i-1][j-1] == 'S' && matrix[i+1][j+1] == 'M' {
			diag1_ok = true
		}
	}

	// diagonal 2
	if j-1 >= 0 && i-1 >= 0 && i+1 <= len(matrix)-1 && j+1 <= len(matrix[0])-1 {
		if matrix[i-1][j+1] == 'M' && matrix[i+1][j-1] == 'S' {
			diag2_ok = true

		}

		if matrix[i-1][j+1] == 'S' && matrix[i+1][j-1] == 'M' {
			diag2_ok = true

		}
	}

	if diag1_ok && diag2_ok {
		counter++
	}

	return counter
}

func main() {
	var matrix [][]byte
	lines := reader.ReadLines("./day04/data/input1_2.txt")

	matrix = make([][]byte, len(lines))
	for i, line := range lines {
		matrix[i] = make([]byte, len(line))
		for j, c := range line {
			matrix[i][j] = byte(c)
		}
	}

	counterXmas, counterXshaped := countAllXmasOccurences(matrix)
	fmt.Println("XMAS count:", counterXmas)
	fmt.Println("X-shaped count:", counterXshaped)

}
