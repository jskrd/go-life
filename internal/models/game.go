package models

import (
	"strconv"
	"syscall/js"
)

type Game struct {
	Canvas     Canvas
	frames     Frames
	Generation Generation
	Height     uint
	Width      uint
}

func (g *Game) clearCanvas() {
	ctx := g.Canvas.Context
	ctx.Call("clearRect", 0, 0, g.Width, g.Height)
}

func (g *Game) drawFramesPerSecond() {
	ctx := g.Canvas.Context
	ctx.Set("fillStyle", "magenta")
	ctx.Set("font", "8px Arial")
	ctx.Call(
		"fillText",
		"FPS: "+strconv.Itoa(int(g.frames.perSecond)),
		8,
		8,
	)
}

func (g *Game) drawGeneration() {
	ctx := g.Canvas.Context
	ctx.Set("fillStyle", "white")
	for coord := range g.Generation.Cells {
		x, y := coord[0], coord[1]
		ctx.Call("fillRect", x, y, 1, 1)
	}
}

func (g *Game) loop() {
	g.frames.calculate()

	g.clearCanvas()
	g.drawGeneration()
	g.drawFramesPerSecond()

	g.Generation = g.Generation.Next()

	js.Global().Call(
		"requestAnimationFrame",
		js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			g.loop()
			return nil
		}),
	)
}

func NewGame() Game {
	document := js.Global().Get("document")

	canvas := document.Call("querySelector", "canvas")
	if !canvas.InstanceOf(js.Global().Get("HTMLCanvasElement")) {
		panic("Canvas element not found")
	}

	context := canvas.Call("getContext", "2d")
	if !context.InstanceOf(js.Global().Get("CanvasRenderingContext2D")) {
		panic("2D context not found")
	}

	return Game{
		Canvas: Canvas{Context: context, Element: canvas},
	}
}

func (g *Game) Resize() {
	const scale float64 = 4

	rect := g.Canvas.Element.Call("getBoundingClientRect")
	width := rect.Get("width").Float()
	height := rect.Get("height").Float()
	ratio := js.Global().Get("devicePixelRatio").Float()

	g.Width = uint(width / scale)
	g.Height = uint(height / scale)

	g.Canvas.Element.Set("width", width*ratio)
	g.Canvas.Element.Set("height", height*ratio)

	g.Canvas.Context.Call("scale", ratio*scale, ratio*scale)
}

func (g *Game) Start() {
	g.Resize()
	js.Global().Get("window").Call(
		"addEventListener",
		"resize",
		js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			g.Resize()
			return nil
		}),
	)

	g.Generation = NewGeneration(g.Width, g.Height)

	js.Global().Call(
		"requestAnimationFrame",
		js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			g.loop()
			return nil
		}),
	)
}
