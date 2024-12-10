package utils

import "fmt"

func PrintMatrix[T any](matrix [][]T) {
	for _, rows := range matrix {
		fmt.Printf("%q\n", rows)
	}
}

func CopyMatrix[T any](matrix [][]T) [][]T {
	duplicate := make([][]T, len(matrix))
	for i := range matrix {
		duplicate[i] = make([]T, len(matrix[i]))
		copy(duplicate[i], matrix[i])
	}

	return duplicate
}
