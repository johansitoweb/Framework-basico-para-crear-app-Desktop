package main

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

type Style struct {
	TextColor sdl.Color
	BackColor sdl.Color
}

type Window struct {
	sdlWindow   *sdl.Window
	sdlRenderer *sdl.Renderer
	style       Style
}

type Button struct {
	rect    sdl.Rect
	text    string
	style   Style
	onClick func()
}

type Label struct {
	rect  sdl.Rect
	text  string
	style Style
}

type VBox struct {
	children []interface{}
	rect     sdl.Rect
}

func NewWindow(title string, width, height int, style Style) (*Window, error) {
	window, err := sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(width), int32(height), sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, err
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return nil, err
	}

	return &Window{sdlWindow: window, sdlRenderer: renderer, style: style}, nil
}

func (w *Window) DrawButton(button *Button) error {
	w.sdlRenderer.SetDrawColor(button.style.BackColor.R, button.style.BackColor.G, button.style.BackColor.B, button.style.BackColor.A)
	w.sdlRenderer.FillRect(&button.rect)
	return nil
}

func (w *Window) DrawLabel(label *Label) error {
	return nil
}

func (w *Window) DrawVBox(vbox *VBox) error {
	y := vbox.rect.Y
	for _, child := range vbox.children {
		switch c := child.(type) {
		case *Button:
			c.rect.Y = y
			w.DrawButton(c)
			y += c.rect.H
		case *Label:
			c.rect.Y = y
			w.DrawLabel(c)
			y += c.rect.H
		}
	}
	return nil
}

func main() {
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		fmt.Fprintf(os.Stderr, "Error al inicializar SDL: %s\n", err)
		return
	}
	defer sdl.Quit()

	style := Style{
		TextColor: sdl.Color{R: 255, G: 255, B: 255, A: 255},
		BackColor: sdl.Color{R: 0, G: 0, B: 150, A: 255},
	}

	window, err := NewWindow("Mi Framework", 640, 480, style)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error al crear la ventana: %s\n", err)
		return
	}
	defer window.sdlWindow.Destroy()
	defer window.sdlRenderer.Destroy()

	button := &Button{
		rect: sdl.Rect{X: 100, Y: 100, W: 100, H: 50},
		text: "Clic",
		style: Style{
			TextColor: sdl.Color{R: 255, G: 255, B: 255, A: 255},
			BackColor: sdl.Color{R: 0, G: 200, B: 0, A: 255},
		},
		onClick: func() {
			fmt.Println("Botón clicado")
		},
	}

	label := &Label{
		rect:  sdl.Rect{X: 200, Y: 200, W: 100, H: 20},
		text:  "Texto",
		style: style,
	}

	vbox := &VBox{
		children: []interface{}{label, button},
		rect:     sdl.Rect{X: 0, Y: 0, W: 640, H: 480},
	}

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			}
		}

		window.sdlRenderer.SetDrawColor(window.style.BackColor.R, window.style.BackColor.G, window.style.BackColor.B, window.style.BackColor.A)
		window.sdlRenderer.Clear()

		window.DrawVBox(vbox)

		window.sdlRenderer.Present()
		sdl.Delay(16)
	}
}
