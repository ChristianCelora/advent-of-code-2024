package main

import (
	"adventcode/reader"
	"fmt"
	"strconv"
	"strings"
)

type Instruction struct {
	prev int
	next int
}

func isUpdateCorrect(updates []int, instructions []Instruction) bool {
	var update_idx map[int]int
	update_idx = make(map[int]int, len(updates))
	for i, val := range updates {
		update_idx[val] = i
	}

	ok, _, _ := checkInstructions(instructions, update_idx)
	return ok
}

func checkInstructions(instructions []Instruction, update_idx map[int]int) (bool, int, int) {
	for _, instr := range instructions {
		val_prev, ok1 := update_idx[instr.prev]
		val_next, ok2 := update_idx[instr.next]

		if ok1 && ok2 {
			if val_prev >= val_next {
				return false, instr.prev, instr.next
			}
		}
	}
	return true, -1, -1
}

// The idea is to apply some sort of bubble sort until is corrected
func correctUpdate(updates []int, instructions []Instruction) []int {
	var correct_update []int
	var update_idx map[int]int
	update_idx = make(map[int]int, len(updates))
	for i, val := range updates {
		update_idx[val] = i
	}

	var is_ok bool
	var idx1, idx2 int
	for !is_ok {
		is_ok, idx1, idx2 = checkInstructions(instructions, update_idx)
		if !is_ok {
			aux := update_idx[idx1]
			update_idx[idx1] = update_idx[idx2]
			update_idx[idx2] = aux
		}
	}

	correct_update = make([]int, len(updates))
	for val, idx := range update_idx {
		correct_update[idx] = val
	}

	return correct_update
}

func main() {
	lines := reader.ReadLines("./day05/data/input1_2.txt")

	var i int
	var instruction []Instruction

	// instructions
	for lines[i] != "" {
		instr := strings.Split(lines[i], "|")
		if len(instr) != 2 {
			panic("error reading " + lines[i])
		}

		prev_inst, _ := strconv.Atoi(instr[0])
		next_inst, _ := strconv.Atoi(instr[1])
		instruction = append(instruction, Instruction{prev_inst, next_inst})

		i++
	}

	var updates []int
	var correct_updates_counter, fixed_updates_counter int
	for _, line := range lines {
		line_vals := strings.Split(line, ",")
		updates = make([]int, len(line_vals))
		for i, update := range line_vals {
			updates[i], _ = strconv.Atoi(update)
		}

		if isUpdateCorrect(updates, instruction) {
			correct_updates_counter += updates[len(updates)/2]
		} else {
			// fmt.Println("incorrect update", updates)
			correct_update := correctUpdate(updates, instruction)
			// fmt.Println("fixed update", correct_update)
			fixed_updates_counter += correct_update[len(updates)/2]
		}
	}

	fmt.Println("correct update counter", correct_updates_counter)
	fmt.Println("fixed update counter", fixed_updates_counter)
}
