package main

import (
	"adventcode/reader"
	"fmt"
	"strconv"
)

type DiskCell struct {
	file_id int // -1 = free space
}

// part 2
type DiskBlock struct {
	file_id int // -1 = free space
	start   int
	end     int
}

func (d *DiskBlock) getBlockSize() int {
	return d.end - d.start + 1
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

func calcChecksumBlock(compressed_disk_map []DiskBlock) int {
	var checksum, idx int
	for _, block := range compressed_disk_map {
		for i := 0; i < block.getBlockSize(); i++ {
			if block.file_id != -1 {
				checksum += idx * block.file_id
			}
			idx++
		}
	}

	return checksum
}

func main() {
	lines := reader.ReadLines("./day09/data/input1_2.txt")
	disk_map := lines[0]

	var file_id, disk_size int
	var expanded_disk_map []DiskCell
	var disk_block_map []DiskBlock
	for _, cell := range disk_map {
		val, _ := strconv.Atoi(string(cell))
		disk_size += val
	}

	expanded_disk_map = make([]DiskCell, disk_size)
	disk_block_map = make([]DiskBlock, 0)

	var idx, idx2 int
	for i, cell := range disk_map {
		val, _ := strconv.Atoi(string(cell))
		if i%2 == 0 {
			// file
			for j := 0; j < val; j++ {
				expanded_disk_map[idx].file_id = file_id
				idx++
			}

			// part 2. file
			disk_block := DiskBlock{
				file_id: file_id,
				start:   idx2,
				end:     idx2 + val - 1,
			}
			disk_block_map = append(disk_block_map, disk_block)
			idx2 += val

			file_id++
		} else {
			// free space
			for j := 0; j < val; j++ {
				expanded_disk_map[idx].file_id = -1
				idx++
			}

			// part 2. free space
			disk_block := DiskBlock{
				file_id: -1,
				start:   idx2,
				end:     idx2 + val - 1,
			}
			disk_block_map = append(disk_block_map, disk_block)
			idx2 += val
		}
	}

	compressed_disk_map := compressDiskMapByCell(expanded_disk_map)
	// printDiskMap(compressed_disk_map)
	fmt.Println("checksum is", calcChecksum(compressed_disk_map))

	// printDiskMapByBlocks(disk_block_map)
	compressed_disk_block_map := compressDiskMapByBlock(disk_block_map)
	// printDiskMapByBlocks(compressed_disk_block_map)
	fmt.Println("checksum for part 2 is", calcChecksumBlock(compressed_disk_block_map))

}

func printDiskMap(compressed_disk_map []DiskCell) {
	for _, cell := range compressed_disk_map {
		if cell.file_id == -1 {
			fmt.Print(".")
		} else {
			fmt.Print(cell.file_id)
		}
	}
	fmt.Println()
}

func printDiskMapByBlocks(compressed_disk_map []DiskBlock) {
	for _, block := range compressed_disk_map {
		for i := 0; i < block.getBlockSize(); i++ {
			if block.file_id == -1 {
				fmt.Print(".")
			} else {
				fmt.Print(block.file_id)
			}
		}
	}
	fmt.Println()
}

func compressDiskMapByCell(expanded_disk_map []DiskCell) []DiskCell {
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
	return compressed_disk_map
}

func compressDiskMapByBlock(expanded_disk_block_map []DiskBlock) []DiskBlock {
	var i, j int

	for j = len(expanded_disk_block_map) - 1; j >= 0; j-- {
		if expanded_disk_block_map[j].file_id != -1 {
			// fmt.Println("file_id", expanded_disk_block_map[j].file_id)
			for i = 0; i <= j; i++ {
				// find first free space block
				if expanded_disk_block_map[i].file_id == -1 {
					if expanded_disk_block_map[i].getBlockSize() == expanded_disk_block_map[j].getBlockSize() {
						// free space is perfect size. change file_id
						expanded_disk_block_map[i].file_id = expanded_disk_block_map[j].file_id
						// free space where file was
						expanded_disk_block_map[j].file_id = -1
						break
					} else if expanded_disk_block_map[i].getBlockSize() > expanded_disk_block_map[j].getBlockSize() {
						// free space is greater. split in 2 blocks one with free space, one with file
						new_file_block := DiskBlock{
							file_id: expanded_disk_block_map[j].file_id,
							start:   expanded_disk_block_map[i].start,
							end:     expanded_disk_block_map[i].start + expanded_disk_block_map[j].getBlockSize() - 1,
						}
						new_free_space := DiskBlock{
							file_id: -1,
							start:   expanded_disk_block_map[i].start + expanded_disk_block_map[j].getBlockSize(),
							end:     expanded_disk_block_map[i].start + expanded_disk_block_map[j].getBlockSize() + (expanded_disk_block_map[i].getBlockSize() - expanded_disk_block_map[j].getBlockSize()) - 1,
						}

						// free space where file was
						expanded_disk_block_map[j].file_id = -1

						// is there a better way?
						new_expanded_disk_block_map := make([]DiskBlock, len(expanded_disk_block_map)+1)
						copy(new_expanded_disk_block_map, expanded_disk_block_map)
						new_expanded_disk_block_map = append(new_expanded_disk_block_map[:i], new_file_block, new_free_space)
						if i < len(expanded_disk_block_map)-1 {
							new_expanded_disk_block_map = append(new_expanded_disk_block_map, expanded_disk_block_map[i+1:]...)
						}
						copy(expanded_disk_block_map, new_expanded_disk_block_map)
						break
					}
				}
			}
			// printDiskMapByBlocks(expanded_disk_block_map)
		}
	}

	return expanded_disk_block_map
}
