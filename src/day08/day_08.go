package main

import (
	"adventcode/reader"
	"fmt"
	"math"
)

type Antenna struct {
	signal rune
	x      int
	y      int
}

type Antinode struct {
	signal rune
	x      int
	y      int
}

func GetAntinodes(a1 Antenna, a2 Antenna) []Antinode {
	antinodes := make([]Antinode, 2)

	delta_x := int(math.Abs(float64(a1.x - a2.x)))
	delta_y := int(math.Abs(float64(a1.y - a2.y)))

	// antinodes[0] = Antinode{
	// 	signal: '#',
	// 	x:      min(a1.x, a2.x) - delta_x,
	// 	y:      min(a1.y, a2.y) - delta_y,
	// }''
	// var new_x, new_y int
	// if a1.x < a2.x {
	// 	if a1.y < a2.y {
	// 		new_x = a1.x - delta_x
	// 		new_y = a1.y - delta_x
	// 	} else {
	// 		new_x = a1.x - delta_x
	// 		new_y = a1.y + delta_x
	// 	}
	// } else {
	// 	if a1.y < a2.y {
	// 		new_x = a1.x - delta_x
	// 		new_y = a1.y - delta_x
	// 	} else {
	// 		new_x = a1.x - delta_x
	// 		new_y = a1.y + delta_x
	// 	}
	// }

	var new_y, new_x int
	if a1.y < a2.y {
		new_y = a1.y - delta_y
	} else {
		new_y = a1.y + delta_y
	}
	if a1.x < a2.x {
		new_x = a1.x - delta_x
	} else {
		new_x = a1.x + delta_x
	}
	antinodes[0] = Antinode{
		signal: '#',
		x:      new_x,
		y:      new_y,
	}

	// antinodes[1] = Antinode{
	// 	signal: '#',
	// 	x:      max(a1.x, a2.x) + delta_x,
	// 	y:      max(a1.y, a2.y) + delta_y,
	// }

	if a1.y < a2.y {
		new_y = a2.y + delta_y
	} else {
		new_y = a2.y - delta_y
	}
	if a1.x < a2.x {
		new_x = a2.x + delta_x
	} else {
		new_x = a2.x - delta_x
	}
	antinodes[1] = Antinode{
		signal: '#',
		x:      new_x,
		y:      new_y,
	}

	return antinodes
}

func main() {
	lines := reader.ReadLines("./day08/data/input1_2.txt")

	antennas := make(map[rune][]Antenna)
	matrix := make([][]Antenna, len(lines))
	matrix_antinodes := make([][]Antinode, len(lines))
	for i, line := range lines {
		matrix[i] = make([]Antenna, len(line))
		matrix_antinodes[i] = make([]Antinode, len(line))
	}

	for i, line := range lines {
		for j, cell := range line {
			if cell != '.' {
				matrix[i][j] = Antenna{
					signal: cell,
					x:      i,
					y:      j,
				}

				_, ok := antennas[cell]
				if !ok {
					antennas[cell] = make([]Antenna, 0)
				}

				if len(antennas[cell]) > 0 {
					for _, antenna := range antennas[cell] {
						antinodes := GetAntinodes(antenna, matrix[i][j])
						for _, antinode := range antinodes {
							if antinode.x >= 0 && antinode.x < len(matrix_antinodes) &&
								antinode.y >= 0 && antinode.y < len(matrix_antinodes[0]) {
								matrix_antinodes[antinode.x][antinode.y] = antinode

								// part 2. we need to recalculate resonant harmonics here
							}
						}
					}
				}

				antennas[cell] = append(antennas[cell], matrix[i][j])
			}
		}
	}

	var count_antinodes int
	for _, row := range matrix_antinodes {
		for _, cell := range row {
			if cell.signal != 0 {
				count_antinodes++
			}

			// if cell.signal == 0 {
			// 	fmt.Printf("%c", '.')
			// } else {
			// 	fmt.Printf("%c", cell.signal)
			// }
		}
		// fmt.Println("")
	}

	fmt.Println("There are", count_antinodes, "unique antinodes")
}
