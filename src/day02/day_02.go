package main

import (
	"adventcode/reader"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func isLevelSafeWithTollerace(levels []int, tollerance int) bool {
	var prev_level, idx int
	var is_safe bool

	diffs := make([]int, len(levels)-1)

	for i, level := range levels {
		if i != 0 {
			diffs[i-1] = prev_level - level
		}

		prev_level = level
	}

	is_safe, idx = isDeltaSafe(diffs, 1, 3)
	if !is_safe && tollerance > 0 {
		is_safe_current := isLevelSafeWithTollerace(copyAndRemoveElement(levels, idx+1), tollerance-1)
		is_safe_prev := isLevelSafeWithTollerace(copyAndRemoveElement(levels, idx), tollerance-1)
		is_safe = is_safe_prev || is_safe_current
	}
	if is_safe {
		return true
	}

	is_safe, idx = isDeltaSafe(diffs, -3, -1)
	if !is_safe && tollerance > 0 {
		is_safe_current := isLevelSafeWithTollerace(copyAndRemoveElement(levels, idx+1), tollerance-1)
		is_safe_prev := isLevelSafeWithTollerace(copyAndRemoveElement(levels, idx), tollerance-1)
		is_safe = is_safe_prev || is_safe_current
	}

	return is_safe
}

func isDeltaSafe(diffs []int, bottom int, ceil int) (bool, int) {
	for i, d := range diffs {
		if d < bottom || d > ceil {
			return false, i
		}
	}

	return true, -1
}

// should have used generics
func copyAndRemoveElement[T any](elements []T, index int) []T {
	var elements_copy []T
	elements_copy = make([]T, len(elements))
	copy(elements_copy, elements)

	if index == 0 {
		return elements_copy[1:]
	} else if index == len(elements)-1 {
		return elements_copy[:index]
	} else {
		return append(elements_copy[:index], elements_copy[index+1:]...)
	}
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

	lines := reader.ReadLines("./day02/data/input1_2.txt")
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

			// brute force...
			// for i := 0; i < len(levels_int); i++ {
			// 	tmp := make([]int, len(levels_int))
			// 	copy(tmp, levels_int)
			// 	if i == 0 {
			// 		tmp = tmp[1:]
			// 	} else if i == len(levels_int)-1 {
			// 		tmp = tmp[:i]
			// 	} else {
			// 		tmp = append(tmp[0:i], tmp[i+1:]...)
			// 	}

			// 	if isLevelSafe(tmp) {
			// 		safe_levels_tollerance++
			// 		break
			// 	}
			// }

			if isLevelSafeWithTollerace(levels_int, 1) {
				safe_levels_tollerance++
			}
		}
	}

	fmt.Printf("safe levels: %d\n", safe_levels)
	fmt.Printf("safe levels with tollerance 1: %d\n", safe_levels+safe_levels_tollerance)
}
