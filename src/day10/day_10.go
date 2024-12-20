package main

import (
	"adventcode/reader"
	"fmt"
)

type Cell struct {
	height  int
	visited bool
}

type Coordinates struct {
	x int
	y int
}

func getTrailScore(matrix [][]int, head Coordinates) int {
	var score int

	matrix_cell := make([][]Cell, len(matrix))
	for i, row := range matrix {
		matrix_cell[i] = make([]Cell, len(row))
		for j, cell := range row {
			matrix_cell[i][j] = Cell{
				height: cell,
			}
		}
	}

	stack := make([]Coordinates, 1)
	stack[0] = head
	for len(stack) > 0 {
		// pop
		current := stack[0]
		stack = stack[1:]

		if matrix_cell[current.x][current.y].height == 9 {
			if !matrix_cell[current.x][current.y].visited {
				score++
			}
		} else {
			// find neighbours
			directions := []Coordinates{
				{-1, 0},
				{1, 0},
				{0, -1},
				{0, 1},
			}
			for _, dir := range directions {
				next_x := current.x + dir.x
				next_y := current.y + dir.y
				if next_x >= 0 && next_x < len(matrix) &&
					next_y >= 0 && next_y < len(matrix[0]) &&
					matrix_cell[next_x][next_y].height == matrix_cell[current.x][current.y].height+1 &&
					!matrix_cell[next_x][next_y].visited {
					// push
					stack = append(stack, Coordinates{next_x, next_y})
				}
			}
		}

		// part 1. uncomment this line
		// part 2. comment this line
		matrix_cell[current.x][current.y].visited = true
	}

	return score
}

func main() {
	var total_score int
	lines := reader.ReadLines("./day10/data/input1_2.txt")
	matrix := make([][]int, len(lines))
	for i, line := range lines {
		matrix[i] = make([]int, len(line))
		for j, cell := range line {
			matrix[i][j] = (int(cell) - '0')
		}
	}

	for i, row := range matrix {
		for j, val := range row {
			if val == 0 {
				head := Coordinates{i, j}
				total_score += getTrailScore(matrix, head)
			}
		}
	}

	fmt.Println("Total score is", total_score)
}
