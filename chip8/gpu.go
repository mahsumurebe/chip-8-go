package chip8

type GPU struct {
	w   uint8
	h   uint8
	ram *RAM
}

func NewGPU() *GPU {
	return &GPU{
		w:   64,
		h:   32,
		ram: NewRAM(),
	}
}

func (gpu *GPU) SetPixel(x uint8, y uint8) bool {
	if x > gpu.w {
		x = 0
	} else if x < 0 {
		x = gpu.w - 1
	}
	if y > gpu.h {
		y = 0
	} else if y < 0 {
		y = gpu.h - 1
	}

	offset := uint16(y)*uint16(gpu.w) + uint16(x)

	data := gpu.ram.Read(offset)
	value := data ^ 1
	gpu.ram.Write(offset, value)

	return value == 0x0
}

func (gpu *GPU) GetMemory() [0x1000]uint8 {
	return gpu.ram.ram
}

func (gpu *GPU) Clear() {
	gpu.ram.Clear()
}