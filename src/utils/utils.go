package utils

import "fmt"

func PrintMatrix[T any](matrix [][]T) {
	for _, rows := range matrix {
		fmt.Printf("%q\n", rows)
	}
}
