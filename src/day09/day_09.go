package main

import (
	"adventcode/reader"
	"fmt"
	"strconv"
)

type DiskCell struct {
	file_id int
}

func calcChecksum(compressed_disk_map []DiskCell) int {
	var checksum int

	for i, cell := range compressed_disk_map {
		if cell.file_id == -1 {
			break
		}

		checksum += i * cell.file_id
	}

	return checksum
}

func main() {
	lines := reader.ReadLines("./day09/data/input1_2.txt")
	disk_map := lines[0]

	var file_id, disk_size int
	var expanded_disk_map []DiskCell
	for _, cell := range disk_map {
		val, _ := strconv.Atoi(string(cell))
		disk_size += val
	}

	expanded_disk_map = make([]DiskCell, disk_size)

	var idx int
	for i, cell := range disk_map {
		val, _ := strconv.Atoi(string(cell))
		if i%2 == 0 {
			// file
			for j := 0; j < val; j++ {
				expanded_disk_map[idx].file_id = file_id
				idx++
			}
			file_id++
		} else {
			// free space
			for j := 0; j < val; j++ {
				expanded_disk_map[idx].file_id = -1
				idx++
			}
		}
	}

	var i, j int
	compressed_disk_map := make([]DiskCell, len(expanded_disk_map))
	j = len(expanded_disk_map) - 1

	for i <= j {
		if expanded_disk_map[i].file_id != -1 {
			compressed_disk_map[i].file_id = expanded_disk_map[i].file_id
			i++
		} else {
			if expanded_disk_map[j].file_id == -1 {
				compressed_disk_map[j].file_id = expanded_disk_map[j].file_id
				j--
			} else {
				// i points to a free space
				// j points to a occupied space
				// swap
				compressed_disk_map[i].file_id = expanded_disk_map[j].file_id
				compressed_disk_map[j].file_id = expanded_disk_map[i].file_id
				i++
				j--
			}
		}
	}

	// fmt.Println("compressed map", string(compressed_disk_map))
	// for _, cell := range compressed_disk_map {
	// 	if cell.file_id == -1 {
	// 		fmt.Print('.')
	// 	} else {
	// 		fmt.Print(cell.file_id)
	// 	}
	// }
	// fmt.Println()
	fmt.Println("checksum is", calcChecksum(compressed_disk_map))
}
