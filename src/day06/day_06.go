package main

import (
	"adventcode/reader"
	"fmt"
)

type Guard struct {
	x         int
	y         int
	direction int // 0 = up, 1 = right, 2 = down, 3 = left
}


func moveGuard(matrix *[][]rune, x int, y int) {
	visited_matrix := make([][])

	guard := Guard{x, y, 0}
	i := guard.x
	j := guard.y
	for i >= 0 && i < len(*matrix) && j >= 0 && j < len((*matrix)[0]) {
		// fmt.Println("-------------")
		// utils.PrintMatrix[rune](*matrix)
		(*matrix)[i][j] = 'X'

		var move_x, move_y int
		switch guard.direction {
		case 0:
			move_x = -1
			move_y = 0
		case 1:
			move_x = 0
			move_y = 1
		case 2:
			move_x = 1
			move_y = 0
		case 3:
			move_x = 0
			move_y = -1
		}

		if i+move_x >= 0 && i+move_x < len(*matrix) &&
			j+move_y >= 0 && j+move_y < len((*matrix)[0]) &&
			(*matrix)[i+move_x][j+move_y] == '#' {
			guard.direction = (guard.direction + 1) % 4
		} else {
			i += move_x
			j += move_y
		}
	}
}

func main() {
	lines := reader.ReadLines("./day06/data/input1_2.txt")
	matrix := make([][]rune, len(lines))

	var start_x, start_y int
	for i, line := range lines {
		matrix[i] = make([]rune, len(line))
		for j, l := range line {
			matrix[i][j] = l
			if l == '^' {
				start_x = i
				start_y = j
			}
		}
	}

	moveGuard(&matrix, start_x, start_y)

	var counter_visited int
	for _, row := range matrix {
		for _, cell := range row {
			if cell == 'X' {
				counter_visited++
			}
		}
	}

	fmt.Println("guard has visited", counter_visited, "lines")
}
