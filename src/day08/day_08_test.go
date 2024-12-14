package main

import "testing"

func TestGetAntinodes(t *testing.T) {
	tests := []struct {
		x1     int
		y1     int
		x2     int
		y2     int
		exp_x1 int
		exp_y1 int
		exp_x2 int
		exp_y2 int
	}{
		{
			x1:     1,
			y1:     1,
			x2:     2,
			y2:     2,
			exp_x1: 0,
			exp_y1: 0,
			exp_x2: 3,
			exp_y2: 3,
		},
		{
			x1:     2,
			y1:     1,
			x2:     1,
			y2:     2,
			exp_x1: 3,
			exp_y1: 0,
			exp_x2: 0,
			exp_y2: 3,
		},
	}

	for _, test := range tests {
		a1 := Antenna{'#', test.x1, test.y1}
		a2 := Antenna{'#', test.x2, test.y2}

		antinodes := GetAntinodes(a1, a2)

		if antinodes[0].x != test.exp_x1 || antinodes[0].y != test.exp_y1 {
			t.Fatalf("expected values (%d, %d), actual (%d, %d)", test.exp_x1, test.exp_y1, antinodes[0].x, antinodes[0].y)
		}
		if antinodes[1].x != test.exp_x2 || antinodes[1].y != test.exp_y2 {
			t.Fatalf("expected values (%d, %d), actual (%d, %d)", test.exp_x2, test.exp_y2, antinodes[1].x, antinodes[1].y)
		}
	}
}
