package main

import (
	"adventcode/reader"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Operation struct {
	op     int // 0 = AND, 1 = OR, 2 = XOR
	left   string
	right  string
	result string
}

func executeOperations(operations_queue []Operation, values map[string]int) {
	var val_result int

	for len(operations_queue) > 0 {
		// dequeue
		operation := operations_queue[0]
		operations_queue = operations_queue[1:]

		val_left, ok_left := values[operation.left]
		val_right, ok_right := values[operation.right]
		// fmt.Println(operation)
		// fmt.Println("left", val_left, ok_left, "right", val_right, ok_right)

		if ok_left && ok_right {
			if operation.op == 0 {
				// AND
				val_result = val_left & val_right
			} else if operation.op == 1 {
				// OR
				val_result = val_left | val_right
			} else {
				// XOR
				if val_left != val_right {
					val_result = 1
				} else {
					val_result = 0
				}
			}
			values[operation.result] = val_result
		} else {
			// we don't have the values yet
			// enqueue
			operations_queue = append(operations_queue, operation)
		}
	}
}

func getVariableBytes(values_map map[string]int, variable byte) string {
	values_map_keys := make([]string, 0)
	for k, _ := range values_map {
		if k[0] == variable {
			values_map_keys = append(values_map_keys, k)
		}
	}

	sort.Strings(values_map_keys)
	var final_result string
	for _, key := range values_map_keys {
		final_result = strconv.Itoa(values_map[key]) + final_result
	}
	return final_result
}

func main() {
	lines := reader.ReadLines("./day24/data/input1_1.txt")

	values_map := make(map[string]int)
	operations := make([]Operation, 0)
	var read_operations bool
	var operation int
	for _, line := range lines {
		if line == "" {
			read_operations = true
			continue
		}

		if read_operations {
			operation_line := strings.Split(line, "->")
			operation_line_left := strings.Split(operation_line[0], " ")
			switch operation_line_left[1] {
			case "AND":
				operation = 0
			case "OR":
				operation = 1
			case "XOR":
				operation = 2

			}
			operations = append(operations, Operation{
				op:     operation,
				left:   operation_line_left[0],
				right:  strings.Trim(operation_line_left[2], " "),
				result: strings.Trim(operation_line[1], " "),
			})
		} else {
			// init values map
			val := strings.Split(line, ":")
			values_map[val[0]], _ = strconv.Atoi(strings.Trim(val[1], " "))
		}
	}

	executeOperations(operations, values_map)

	final_result := getVariableBytes(values_map, 'z')

	final_res, _ := strconv.ParseInt(final_result, 2, 64)
	fmt.Println("final result is", final_res)

	// part 2
	x, _ := strconv.ParseInt(getVariableBytes(values_map, 'x'), 2, 64)
	y, _ := strconv.ParseInt(getVariableBytes(values_map, 'y'), 2, 64)
	z := x + y
	expected_res := fmt.Sprintf("%b", z)

	fmt.Println(x, y, z, expected_res, final_result)

	// mutations: 4 swaps not repeated

	// stupid solution: O(nˆ4)
	for i := 0; i < len(operations); i++ {
		for j := 0; j < len(operations); j++ {
			if i != j {
				for k := 0; k < len(operations); k++ {
					if k != j && k != i {
						for m := 0; m < len(operations); m++ {
							if m != j && m != i && m != k {

							}
						}
					}
				}
			}
		}
	}

	// Analize output. which bits are wrong? Start swapping from those
	// Create a subset of operations
	// The subset should be smaller and we can brute force from there all combinations

}
