package main

import (
	"math"
	// "rogue-go/internal/domain/geom"
	// "rogue-go/internal/domain/level"
)

func ComputeFOV(field *map_t, origin position_t, radius int) map[position_t]bool {
	visible := make(map[position_t]bool, (2*radius+1)*(2*radius+1))
	visible[origin] = true
	for y := origin.y - radius; y <= origin.y+radius; y++ {
		for x := origin.x - radius; x <= origin.x+radius; x++ {
			// p := geom.Point{X: x, Y: y}
			var pos position_t
			pos.y, pos.x = y, x
			if !InBounds(x, y) {
				continue
			}
			if distance(origin, pos) > float64(radius) {
				continue
			}
			if LineOfSight(field, origin, pos) {
				visible[pos] = true
			}
		}
	}
	return visible
}

func distance(a, b position_t) float64 {
	// dx, dy := float64(a.x-b.x), float64(a.y-b.y)
	distY := math.Abs(float64(a.y - b.y))
	distX := math.Abs(float64(a.x - b.x))
	// return math.Hypot(dx, dy)
	return math.Hypot(distY, distX)
}

func LineOfSight(field *map_t, a, b position_t) bool {
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
		if !InBounds(x0, y0) {
			return false
		}
		t := field.playground[y0][x0]
		// if t == 'x' {
		// 	return false
		// }
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

func InBounds(x, y int) bool {
	if x >= 0 && x < MAP_WIDTH && y >= 0 && y < MAP_HEIGHT {
		return true
	}
	return false
}
