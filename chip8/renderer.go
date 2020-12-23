package chip8

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

type Renderer struct {
	chip *CHIP8
	W    int32
	H    int32
}

func NewRenderer() *Renderer {
	return &Renderer{
		chip: nil,
		W:    int32(800),
		H:    int32(600),
	}
}

var (
	window *sdl.Window
	r      *sdl.Renderer
)

func (renderer Renderer) Init() {
	// SDL Init

	//if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
	//	panic(err)
	//}

	_window, err := sdl.CreateWindow("CHIP-8", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		renderer.W, renderer.H, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatalf("Failed to create window: %s\n", err)
	}
	window = _window

	_r, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		log.Fatalf("Failed to create renderer: %s\n", err)
	}
	r = _r
}

func (renderer Renderer) CheckEvents() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.KeyboardEvent:
			for i := 0x0; i < len(renderer.chip.KB.KeyCode); i++ {
				if t.State != sdl.PRESSED {
					if uint8(t.Keysym.Sym) == renderer.chip.KB.KeyCode[uint8(i)] {
						renderer.chip.KB.PressedKeys[uint8(i)] = false
						break
					}
				} else {
					if uint8(t.Keysym.Sym) == renderer.chip.KB.KeyCode[uint8(i)] {
						renderer.chip.KB.PressedKeys[uint8(i)] = true
						break
					}
				}
			}
		case *sdl.QuitEvent:
			renderer.chip.CPU.stop = true
			break
		}
	}
}

func (renderer Renderer) Render() {
	_ = r.SetDrawColor(0, 0, 0, 255)
	_ = r.Clear()
	mem := renderer.chip.GPU.GetMemory()
	gpuW := uint16(renderer.chip.GPU.w)
	gpuH := uint16(renderer.chip.GPU.h)
	scaleW := renderer.W / int32(gpuW)
	scaleH := renderer.H / int32(gpuH)
	for x := uint16(0); x < gpuW; x++ {
		for y := uint16(0); y < gpuH; y++ {
			offset := (y * gpuW) + x

			if mem[offset] == 0x1 {
				rect := sdl.Rect{
					X: scaleW * int32(x),
					Y: scaleH * int32(y),
					W: scaleW,
					H: scaleH,
				}
				_ = r.SetDrawColor(255, 255, 255, 255)
				_ = r.FillRect(&rect)

			}
		}
	}
	r.Present()
	sdl.PollEvent()
}

func (renderer Renderer) Dispose() {
	defer sdl.Quit()
	defer window.Destroy()
	defer r.Destroy()
}
