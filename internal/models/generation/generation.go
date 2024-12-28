package generation

import "math/rand/v2"

type Generation struct {
	Cells map[[2]int]struct{}
}

func (generation *Generation) IsAlive(x, y int) bool {
	_, ok := generation.Cells[[2]int{x, y}]
	return ok
}

func (generation *Generation) SetAlive(x, y int) {
	generation.Cells[[2]int{x, y}] = struct{}{}
}

func (generation *Generation) CountNeighbors(x, y int) int {
	count := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			if generation.IsAlive(x+dx, y+dy) {
				count++
			}
		}
	}
	return count
}

func (generation *Generation) ShouldLive(x, y int) bool {
	neighbors := generation.CountNeighbors(x, y)
	return neighbors == 3 || (generation.IsAlive(x, y) && neighbors == 2)
}

func Seed(width, height int) Generation {
	generation := Generation{Cells: make(map[[2]int]struct{})}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if randFloat := rand.Float64(); randFloat < 0.5 {
				generation.SetAlive(x, y)
			}
		}
	}
	return generation
}
