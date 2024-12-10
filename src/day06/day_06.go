package main

import (
	"adventcode/reader"
	"adventcode/utils"
	"fmt"
)

type Guard struct {
	x         int
	y         int
	direction int // 0 = up, 1 = right, 2 = down, 3 = left
}

type VisitedDirections struct {
	dir [4]bool
}

func moveGuard(matrix *[][]rune, x int, y int) bool {
	visited_matrix := make([][]VisitedDirections, len((*matrix)))
	for i, row := range *matrix {
		visited_matrix[i] = make([]VisitedDirections, len(row))
	}

	guard := Guard{x, y, 0}
	i := guard.x
	j := guard.y
	for i >= 0 && i < len(*matrix) && j >= 0 && j < len((*matrix)[0]) {
		(*matrix)[i][j] = 'X'
		visited_matrix[i][j].dir[guard.direction] = true

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

		if i+move_x >= 0 && i+move_x < len(*matrix) && j+move_y >= 0 && j+move_y < len((*matrix)[0]) {
			// peek next position. if it is obstacle rotate direction 90 deg clockwise
			if isObstacle(matrix, i+move_x, j+move_y) {
				guard.direction = (guard.direction + 1) % 4
			} else {
				i += move_x
				j += move_y
			}

			// reached loop
			if visited_matrix[i][j].dir[guard.direction] {
				return true
			}
		} else {
			i += move_x
			j += move_y
		}
	}

	return false
}

func isObstacle(matrix *[][]rune, i int, j int) bool {
	return (*matrix)[i][j] == '#' || (*matrix)[i][j] == 'O'
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
	var counter_stuck_in_loop int
	for i, row := range matrix {
		for j, cell := range row {
			if cell == 'X' {
				counter_visited++

				if i != start_x || j != start_y {
					// mutate matrix based on path (we add only one obstacle. Not visited cells can be ignored)
					new_matrix := utils.CopyMatrix(matrix)
					new_matrix[i][j] = 'O'
					is_stuck_in_loop := moveGuard(&new_matrix, start_x, start_y)
					if is_stuck_in_loop {
						// fmt.Println("-------------")
						// utils.PrintMatrix[rune](new_matrix)
						counter_stuck_in_loop++
					}
				}
			}
		}
	}

	fmt.Println("guard has visited", counter_visited, "lines")
	fmt.Println("guard is stuck in loop in", counter_stuck_in_loop, "mutations")
}
