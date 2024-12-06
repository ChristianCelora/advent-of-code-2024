package main

import (
	"adventcode/reader"
	"fmt"
	"regexp"
	"strconv"
)

func parseMultiplications(instructions string) int {
	var total int
	reg := regexp.MustCompile(`mul\((\d*),(\d*)\)`)
	matches_all := reg.FindAllSubmatch([]byte(instructions), -1)

	// fmt.Printf("%q\n", matches_all)

	var prev_val int
	for _, matches := range matches_all {
		for i, val := range matches[1:] {
			intval, err := strconv.Atoi(string(val))
			if err != nil {
				panic(err)
			}

			if i%2 == 0 {
				prev_val = intval
			} else {
				total += prev_val * intval
			}
		}
	}

	return total
}

func main() {
	var total int
	lines := reader.ReadLines("./day03/data/input1_2.txt")
	for _, line := range lines {
		total += parseMultiplications(line)
	}

	fmt.Println("total is", total)
}
