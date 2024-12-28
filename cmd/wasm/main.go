package main

import (
	"syscall/js"
	"time"

	"github.com/jskrd/go-life/internal/models/generation"
	"github.com/jskrd/go-life/internal/models/universe"
)

func main() {
	document := js.Global().Get("document")
	canvas := document.Call("getElementById", "app")
	if !canvas.InstanceOf(js.Global().Get("HTMLCanvasElement")) {
		panic("Canvas element not found")
	}

	const scale = 4

	resizeCanvasToViewport(canvas, scale)
	js.Global().Call(
		"addEventListener",
		"resize",
		js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			resizeCanvasToViewport(canvas, scale)
			return nil
		}),
	)

	generation := generation.Seed(
		canvas.Get("width").Int()/scale,
		canvas.Get("height").Int()/scale,
	)
	universe := universe.Universe{Generation: generation}

	context := canvas.Call("getContext", "2d")
	if !context.InstanceOf(js.Global().Get("CanvasRenderingContext2D")) {
		panic("2D context not found")
	}

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		js.Global().Call(
			"requestAnimationFrame",
			js.FuncOf(func(this js.Value, args []js.Value) interface{} {
				context.Call(
					"clearRect",
					0,
					0,
					canvas.Get("width").Int(),
					canvas.Get("height").Int(),
				)
				context.Call(
					"drawImage",
					universe.ToCanvas(
						canvas.Get("width").Int(),
						canvas.Get("height").Int(),
					),
					0,
					0,
				)

				universe.Tick()
				return nil
			}),
		)
	}

	// Prevent program from exiting
	c := make(chan struct{})
	<-c
}

func resizeCanvasToViewport(canvas js.Value, scale int) {
	canvas.Set("width", js.Global().Get("innerWidth").Int())
	canvas.Set("height", js.Global().Get("innerHeight").Int())

	context := canvas.Call("getContext", "2d")
	if !context.InstanceOf(js.Global().Get("CanvasRenderingContext2D")) {
		panic("2D context not found")
	}

	context.Set("imageSmoothingEnabled", false)
	context.Call("scale", scale, scale)
}
