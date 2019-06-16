package ui

import (
	"image"
	"log"
	_ "time"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	Up = iota
	Down
	Left
	Right
)

type Renderer interface {
	RenderTo(image *image.RGBA)
	Input(keys [4]bool)
}

type UI struct {
	window   *glfw.Window
	renderer Renderer
	image    *image.RGBA
}

const (
	width  = 240
	height = 150
)

func StartUI(renderer Renderer) {
	ui := &UI{
		renderer: renderer,
		image:    image.NewRGBA(image.Rect(0, 0, width, height)),
	}

	if err := glfw.Init(); err != nil {
		log.Fatalln(err)
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)

	window, err := glfw.CreateWindow(640, 400, "Ray Tracer", nil, nil)

	if err != nil {
		log.Fatalln(err)
	}

	window.MakeContextCurrent()

	ui.window = window

	if err := gl.Init(); err != nil {
		log.Fatalln(err)
	}

	gl.Enable(gl.TEXTURE_2D)
	gl.ClearColor(1, 0, 1, 1)
	gl.BindTexture(gl.TEXTURE_2D, createTexture())

	ui.StartRunLoop()
}

func (ui *UI) Shutdown() {
	glfw.Terminate()
}

func (ui *UI) StartRunLoop() {
	for !ui.window.ShouldClose() {
		ui.renderer.RenderTo(ui.image)
		ui.renderer.Input(readKeys(ui.window))

		ui.setTexture(ui.image)
		drawFrame()
		ui.window.SwapBuffers()
		glfw.PollEvents()
	}
}

func (ui *UI) setTexture(im *image.RGBA) {
	size := im.Rect.Size()

	gl.TexImage2D(
		gl.TEXTURE_2D, 0, gl.RGBA, int32(size.X), int32(size.Y),
		0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(ui.image.Pix))
}

func createTexture() uint32 {
	var texture uint32
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.BindTexture(gl.TEXTURE_2D, 0)
	return texture
}

func drawFrame() {
	gl.Clear(gl.COLOR_BUFFER_BIT)

	x := float32(1)
	y := float32(1)

	gl.Begin(gl.QUADS)
	gl.TexCoord2f(0, 1)
	gl.Vertex2f(-x, -y)
	gl.TexCoord2f(1, 1)
	gl.Vertex2f(x, -y)
	gl.TexCoord2f(1, 0)
	gl.Vertex2f(x, y)
	gl.TexCoord2f(0, 0)
	gl.Vertex2f(-x, y)
	gl.End()
}

func readKeys(window *glfw.Window) [4]bool {
	var keys [4]bool

	keys[Up] = window.GetKey(glfw.KeyUp) == glfw.Press
	keys[Down] = window.GetKey(glfw.KeyDown) == glfw.Press
	keys[Left] = window.GetKey(glfw.KeyLeft) == glfw.Press
	keys[Right] = window.GetKey(glfw.KeyRight) == glfw.Press

	return keys
}
