package main

import (
	"adventcode/reader"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func canBeCombined(expected int, values []int) bool {
	operations := [2]string{"add", "mul"}
	n := int(math.Pow(float64(len(operations)), float64(len(values)-1)))

	// make combinations
	combinations := make([]string, n)
	pad_format := fmt.Sprintf("%%0%db", len(values)-1) // convert in bit representation and left pad zeroes
	for i := 0; i < n; i++ {
		combinations[i] = fmt.Sprintf(pad_format, i)
	}

	// fmt.Println(values)
	// fmt.Println("------")
	for _, comb := range combinations {
		var res, i int
		res = values[i]
		// fmt.Printf("%s (%d)\n", comb, len(comb))
		for _, op := range comb {
			i++
			// 0 = add,  1 = mul
			if op == '0' {
				res += values[i]
			} else {
				res *= values[i]
			}
		}

		if res == expected {
			return true
		}
	}
	return false
}

func main() {
	var total_calibration int
	lines := reader.ReadLines("./day07/data/input1_2.txt")

	for _, line := range lines {
		line_slice := strings.Split(line, ":")

		test_val, _ := strconv.Atoi(line_slice[0])
		values_str := strings.Split(strings.Trim(line_slice[1], " "), " ")
		values := make([]int, len(values_str))
		for i, val := range values_str {
			v, _ := strconv.Atoi(val)
			values[i] = v
		}

		if canBeCombined(test_val, values) {
			total_calibration += test_val
		}

	}

	fmt.Println("total calibration value is", total_calibration)
}
