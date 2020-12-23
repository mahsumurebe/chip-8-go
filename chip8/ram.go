package chip8

type RAM struct {
	ram [0x1000]uint8
}

func NewRAM() *RAM {
	return &RAM{
		ram: [0x1000]uint8{},
	}
}

func (ram *RAM) Write(address uint16, data uint8) {
	ram.ram[address] = data
}

func (ram *RAM) Read(address uint16) uint8 {
	return ram.ram[address]
}

func (ram *RAM) Clear() {
	ram.ram = [0x1000]uint8{}
}
