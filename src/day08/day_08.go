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
	origin string
	signal rune
	x      int
	y      int
}

type AntennaPair struct {
	origin string
	a1     Antenna
	a2     Antenna
}

func GetAntinodes(a1 Antenna, a2 Antenna, origin string) []Antinode {
	antinodes := make([]Antinode, 2)

	delta_x := int(math.Abs(float64(a1.x - a2.x)))
	delta_y := int(math.Abs(float64(a1.y - a2.y)))

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
		origin: origin,
		signal: a1.signal,
		x:      new_x,
		y:      new_y,
	}

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
		origin: origin,
		signal: a2.signal,
		x:      new_x,
		y:      new_y,
	}

	return antinodes
}

func getOriginUuid(a1 Antenna, a2 Antenna) string {
	return string(a1.signal) + string(a1.x) + string(a1.y) + string(a2.signal) + string(a2.x) + string(a2.y)
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
						// part 2. antennas are antinodes.
						// Should not be necessary
						// antinode_antenna := Antinode(antenna)
						// matrix_antinodes[antenna.x][antenna.y] = antinode_antenna

						antinode_stack := make([]AntennaPair, 0)
						starting_pair := AntennaPair{
							a1:     antenna,
							a2:     matrix[i][j],
							origin: getOriginUuid(antenna, matrix[i][j]),
						}
						antinode_stack = append(antinode_stack, starting_pair)

						for len(antinode_stack) > 0 {
							// pop first
							antennaPair := antinode_stack[0]
							antinode_stack = antinode_stack[1:]

							antinodes := GetAntinodes(antennaPair.a1, antennaPair.a2, antennaPair.origin)
							pairs := []AntennaPair{
								{
									origin: antennaPair.origin,
									a1:     antennaPair.a1,
									a2: Antenna{
										signal: antinodes[0].signal,
										x:      antinodes[0].x,
										y:      antinodes[0].y,
									},
								},
								{
									origin: antennaPair.origin,
									a1:     antennaPair.a2,
									a2: Antenna{
										signal: antinodes[1].signal,
										x:      antinodes[1].x,
										y:      antinodes[1].y,
									},
								},
							}
							for _, pair := range pairs {
								antinode := pair.a2
								if antinode.x >= 0 && antinode.x < len(matrix_antinodes) &&
									antinode.y >= 0 && antinode.y < len(matrix_antinodes[0]) {
									// part 2. we need to recalculate resonant harmonics here
									// if matrix_antinodes[antinode.x][antinode.y].signal != pair.a1.signal { // eval only new signals
									if matrix_antinodes[antinode.x][antinode.y].origin != pair.origin { // eval only new signals
										antinode_stack = append(antinode_stack, pair)
									}

									matrix_antinodes[antinode.x][antinode.y] = Antinode{
										origin: pair.origin,
										signal: antinode.signal,
										x:      antinode.x,
										y:      antinode.y,
									}
								}
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

			if cell.signal == 0 {
				fmt.Printf("%c", '.')
			} else {
				fmt.Printf("%c", '#')
			}
		}
		fmt.Println("")
	}

	fmt.Println("There are", count_antinodes, "unique antinodes")
}
