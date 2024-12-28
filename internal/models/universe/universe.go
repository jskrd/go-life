package universe

import (
	"syscall/js"

	"github.com/jskrd/go-life/internal/models/generation"
)

type Universe struct {
	Generation generation.Generation
}

func (universe *Universe) Tick() {
	newGeneration := generation.Generation{Cells: make(map[[2]int]struct{})}
	for coord := range universe.Generation.Cells {
		x, y := coord[0], coord[1]
		if universe.Generation.ShouldLive(x, y) {
			newGeneration.SetAlive(x, y)
		}
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				if dx == 0 && dy == 0 {
					continue
				}
				if universe.Generation.ShouldLive(x+dx, y+dy) {
					newGeneration.SetAlive(x+dx, y+dy)
				}
			}
		}
	}
	universe.Generation = newGeneration
}

func (universe *Universe) ToCanvas(width, height int) js.Value {
	canvas := js.Global().Get("document").Call("createElement", "canvas")
	canvas.Set("width", width)
	canvas.Set("height", height)

	context := canvas.Call("getContext", "2d")
	if !context.InstanceOf(js.Global().Get("CanvasRenderingContext2D")) {
		panic("2D context not found")
	}

	for coord := range universe.Generation.Cells {
		x, y := coord[0], coord[1]

		if x < 0 || x >= width || y < 0 || y >= height {
			continue
		}

		context.Set("fillStyle", "white")
		context.Call("fillRect", x, y, 1, 1)
	}

	return canvas
}
