package main

import (
	"adventcode/reader"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

func parseMultiplications(instructions string) int {
	var total int
	reg := regexp.MustCompile(`mul\((\d*),(\d*)\)`)
	matches_all := reg.FindAllSubmatch([]byte(instructions), -1)

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

func parseOperations(instructions string) int {
	var total int
	reg_mul := regexp.MustCompile(`mul\((\d*),(\d*)\)`)
	reg_do := regexp.MustCompile(`(do\(\))`)
	reg_dont := regexp.MustCompile(`(don't\(\))`)
	matches_all_mul := reg_mul.FindAllStringSubmatchIndex(instructions, -1)
	matches_all_do := reg_do.FindAllStringSubmatchIndex(instructions, -1)
	matches_all_dont := reg_dont.FindAllStringSubmatchIndex(instructions, -1)

	fmt.Printf("%v\n", matches_all_do)
	fmt.Printf("%v\n", matches_all_dont)

	// var prev_val int
	var idx_mul, idx_do, idx_dont int
	var operation int // 1 = mul, 2 = do, 3 = dont
	skip_mul := false
	for idx_mul < len(matches_all_mul) || idx_do < len(matches_all_do) || idx_dont < len(matches_all_dont) {

		// choose next operation
		operation = -1
		minInstrIdx := math.MaxInt
		if idx_mul < len(matches_all_mul) && minInstrIdx > matches_all_mul[idx_mul][0] {
			minInstrIdx = matches_all_mul[idx_mul][0]
			operation = 1
		}
		if idx_do < len(matches_all_do) && minInstrIdx > matches_all_do[idx_do][0] {
			minInstrIdx = matches_all_do[idx_do][0]
			operation = 2
		}
		if idx_dont < len(matches_all_dont) && minInstrIdx > matches_all_dont[idx_dont][0] {
			minInstrIdx = matches_all_dont[idx_dont][0]
			operation = 3
		}

		// exec operation
		if operation == 2 {
			fmt.Println("do op", matches_all_do[idx_do][0])
			skip_mul = false
			idx_do++
		} else if operation == 3 {
			fmt.Println("dont op", matches_all_dont[idx_dont][0])
			skip_mul = true
			idx_dont++
		} else if operation == 1 {
			fmt.Println("mul op", matches_all_mul[idx_mul][0])
			if !skip_mul {
				matches := matches_all_mul[idx_mul]
				if len(matches) >= 6 && matches[2] < matches[3] && matches[4] < matches[5] {
					val1 := instructions[matches[2]:matches[3]]
					intval1, err := strconv.Atoi(string(val1))
					if err != nil {
						panic(err)
					}

					val2 := instructions[matches[4]:matches[5]]
					intval2, err := strconv.Atoi(string(val2))
					if err != nil {
						panic(err)
					}

					fmt.Println("mul:", intval1, intval2)

					total += intval1 * intval2
				}
			}
			idx_mul++
		} else {
			panic("op not recognized")
		}
	}

	return total
}

func main() {
	var total int
	lines := reader.ReadLines("./day03/data/input2_1.txt")
	for _, line := range lines {
		total += parseOperations(line)
	}

	fmt.Println("total is", total)
}
