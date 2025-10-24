package main

import (
	"math"
)

func ComputeFOV(field *map_t, origin position_t, radius int) map[position_t]bool {
	visible := make(map[position_t]bool, (2*radius+1)*(2*radius+1))
	visible[origin] = true
	for y := origin.y - radius; y <= origin.y+radius; y++ {
		for x := origin.x - radius; x <= origin.x+radius; x++ {
			var pos position_t
			pos.y, pos.x = y, x
			if !inBounds(x, y) {
				continue
			}
			if computeDistance(origin, pos) > float64(radius) {
				continue
			}
			if lineOfSight(field, origin, pos) {
				visible[pos] = true
			}
		}
	}
	return visible
}

func computeDistance(a, b position_t) float64 {
	distY := math.Abs(float64(a.y - b.y))
	distX := math.Abs(float64(a.x - b.x))
	return math.Hypot(distY, distX)
}

func lineOfSight(field *map_t, a, b position_t) bool {
	x0, y0 := a.x, a.y
	x1, y1 := b.x, b.y
	dx := int(math.Abs(float64(x1 - x0)))
	sx := -1
	if x0 < x1 {
		sx = 1
	}
	dy := -int(math.Abs(float64(y1 - y0)))
	sy := -1
	if y0 < y1 {
		sy = 1
	}
	err := dx + dy
	for {
		if x0 == x1 && y0 == y1 {
			return true
		}
		e2 := 2 * err
		if e2 >= dy {
			err += dy
			x0 += sx
		}
		if e2 <= dx {
			err += dx
			y0 += sy
		}
		if !inBounds(x0, y0) {
			return false
		}
		t := field.playground[y0][x0]
		if t == '#' {
			if field.playground[a.y][a.x] == '@' {
				var pos position_t
				pos.y, pos.x = y0, x0
				field.visited[pos] = true
			}
			return false
		}

	}
}

func inBounds(x, y int) bool {
	if x >= 0 && x < MAP_WIDTH && y >= 0 && y < MAP_HEIGHT {
		return true
	}
	return false
}
