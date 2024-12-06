package main

import (
	"adventcode/reader"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func isLevelSafeWithTollerace(levels []int, tollerance int) bool {
	var prev_level int
	var is_increasing, is_safe bool
	fmt.Println(levels, tollerance)

	for i, level := range levels {
		if i != 0 && level != levels[0] {
			is_increasing = levels[0] < level
			break
		}
	}

	for i, level := range levels {
		is_safe = true
		if i != 0 {
			if is_increasing && prev_level > level {
				is_safe = false
			} else if !is_increasing && prev_level < level {
				is_safe = false
			}

			delta := int(math.Abs(float64(prev_level) - float64(level)))
			if delta < 1 || delta > 3 {
				is_safe = false
			}
		}

		if !is_safe {
			if tollerance > 0 {
				fmt.Println(prev_level, level, i)
				// exclude current level from levels
				var is_safe_current bool
				new_levels_1 := make([]int, len(levels))
				copy(new_levels_1, levels)
				if i == len(new_levels_1)-1 {
					is_safe_current = isLevelSafeWithTollerace(new_levels_1[0:i], tollerance-1)
				} else {
					is_safe_current = isLevelSafeWithTollerace(append(new_levels_1[:i], new_levels_1[i+1:]...), tollerance-1)
				}

				// exclude previous level (i - 1) from levels
				var is_safe_prev bool
				new_levels_2 := make([]int, len(levels))
				copy(new_levels_2, levels)
				if i-1 == 0 {
					is_safe_prev = isLevelSafeWithTollerace(new_levels_2[i:], tollerance-1)
				} else {
					is_safe_prev = isLevelSafeWithTollerace(append(new_levels_2[:i-1], new_levels_2[i:]...), tollerance-1)
				}

				return is_safe_prev || is_safe_current
			} else {
				return false
			}
		} else {
			prev_level = level
		}
	}

	return true
}

func isLevelSafe(levels []int) bool {
	var prev_level int
	var is_increasing, is_safe bool
	// fmt.Println(levels)

	for i, level := range levels {
		if i != 0 && level != levels[0] {
			is_increasing = levels[0] < level
			break
		}
	}

	for i, level := range levels {
		is_safe = true
		if i != 0 {
			if is_increasing && prev_level > level {
				is_safe = false
			} else if !is_increasing && prev_level < level {
				is_safe = false
			}

			delta := int(math.Abs(float64(prev_level) - float64(level)))
			if delta < 1 || delta > 3 {
				is_safe = false
			}
		}

		if !is_safe {
			return false
		} else {
			prev_level = level
		}
	}

	return true
}

func main() {
	var safe_levels, safe_levels_tollerance int

	lines := reader.ReadLines("./day02/data/input_test.txt")
	for _, line := range lines {
		levels := strings.Split(line, " ")
		levels_int := make([]int, len(levels))
		for i, level := range levels {
			lev, err := strconv.Atoi(level)
			if err != nil {
				panic(err)
			}
			levels_int[i] = lev
		}

		if isLevelSafe(levels_int) {
			safe_levels++
		} else {
			var test1, test2 bool
			test1 = false
			test2 = false
			// brute force...
			for i := 0; i < len(levels_int); i++ {
				tmp := make([]int, len(levels_int))
				copy(tmp, levels_int)
				if i == 0 {
					tmp = tmp[1:]
				} else if i == len(levels_int)-1 {
					tmp = tmp[:i]
				} else {
					tmp = append(tmp[0:i], tmp[i+1:]...)
				}

				if isLevelSafe(tmp) {
					test1 = true
					safe_levels_tollerance++
					break
				}
			}

			if isLevelSafeWithTollerace(levels_int, 1) {
				test2 = true
				safe_levels_tollerance++
			}

			if test1 != test2 {
				fmt.Println(levels_int)
			}
		}
	}

	fmt.Printf("safe levels: %d\n", safe_levels)
	fmt.Printf("safe levels with tollerance 1: %d\n", safe_levels+safe_levels_tollerance)
}
