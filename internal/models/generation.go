package models

import "math/rand/v2"

type Generation struct {
	Cells map[[2]int]struct{}
}

func (g *Generation) CountNeighbors(x, y int) uint {
	count := uint(0)
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			if g.IsAlive(x+dx, y+dy) {
				count++
			}
		}
	}
	return count
}

func (g *Generation) IsAlive(x, y int) bool {
	_, ok := g.Cells[[2]int{x, y}]
	return ok
}

func (g *Generation) Next() Generation {
	next := Generation{Cells: make(map[[2]int]struct{})}
	for coord := range g.Cells {
		x, y := coord[0], coord[1]
		if g.ShouldLive(x, y) {
			next.SetAlive(x, y)
		}
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				if dx == 0 && dy == 0 {
					continue
				}
				if g.ShouldLive(x+dx, y+dy) {
					next.SetAlive(x+dx, y+dy)
				}
			}
		}
	}
	return next
}

func NewGeneration(width, height uint) Generation {
	g := Generation{Cells: make(map[[2]int]struct{})}
	for x := 0; x < int(width); x++ {
		for y := 0; y < int(height); y++ {
			if randFloat := rand.Float64(); randFloat < 0.5 {
				g.SetAlive(x, y)
			}
		}
	}
	return g
}

func (g *Generation) SetAlive(x, y int) {
	g.Cells[[2]int{x, y}] = struct{}{}
}

func (g *Generation) ShouldLive(x, y int) bool {
	neighbors := g.CountNeighbors(x, y)
	return neighbors == 3 || (g.IsAlive(x, y) && neighbors == 2)
}
