package chip8

import (
	"github.com/veandco/go-sdl2/sdl"
)

type CHIP8 struct {
	CPU      *CPU
	GPU      *GPU
	Renderer *Renderer
	RAM      *RAM
	KB       *Keyboard
	Buzzer   *Buzzer

	delayTimer uint8
	soundTimer uint8
}

func NewCHIP8() *CHIP8 {
	chip := &CHIP8{
		CPU:        NewCPU(),
		GPU:        NewGPU(),
		Renderer:   NewRenderer(),
		RAM:        NewRAM(),
		KB:         NewKeyboard(),
		Buzzer:     NewBuzzer(),
		delayTimer: 0x0,
		soundTimer: 0x0,
	}
	chip.CPU.chip = chip
	chip.Renderer.chip = chip
	chip.Buzzer.chip = chip
	return chip
}

func (chip *CHIP8) loadFonts() {
	fonts := [16 * 5]byte{
		0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
		0x20, 0x60, 0x20, 0x20, 0x70, // 1
		0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
		0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
		0x90, 0x90, 0xF0, 0x10, 0x10, // 4
		0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
		0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
		0xF0, 0x10, 0x20, 0x40, 0x40, // 7
		0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
		0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
		0xF0, 0x90, 0xF0, 0x90, 0x90, // A
		0xE0, 0x90, 0xE0, 0x90, 0xE0, // B
		0xF0, 0x80, 0x80, 0x80, 0xF0, // C
		0xE0, 0x90, 0x90, 0x90, 0xE0, // D
		0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
		0xF0, 0x80, 0xF0, 0x80, 0x80, // F
	}

	for i := byte(0); i < byte(len(fonts)); i++ {
		chip.RAM.Write(uint16(i), fonts[i])
	}
}

func (chip *CHIP8) LoadRom(data *[]uint8) {
	chip.RAM.Clear()
	chip.loadFonts()

	for i := 0; i < len(*data); i++ {
		chip.RAM.Write(uint16(i)+0x200, (*data)[i])
	}
}

func (chip *CHIP8) Run() {
	chip.Renderer.Init()
	go func() {
		for !chip.CPU.stop {
			chip.Renderer.CheckEvents()
		}
	}()
	for !chip.CPU.stop {
		for i := 0; i < 10; i++ {
			chip.CPU.cycle()
		}
		chip.CPU.updateTimers()
		chip.Renderer.Render()
		sdl.Delay(1000 / 60)
	}
}
