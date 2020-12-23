package chip8

type Keyboard struct {
	KeyCode     map[uint8]uint8
	PressedKeys map[uint8]bool
}

func NewKeyboard() *Keyboard {
	keyboard := &Keyboard{
		KeyCode: map[uint8]uint8{
			0x1: 49,  // 1
			0x2: 50,  // 2
			0x3: 51,  // 3
			0xc: 52,  // 4
			0x4: 113, // Q
			0x5: 119, // W
			0x6: 101, // E
			0xD: 114, // R
			0x7: 97,  // A
			0x8: 115, // S
			0x9: 100, // D
			0xE: 102, // F
			0xA: 122, // Z
			0x0: 120, // X
			0xB: 99,  // C
			0xF: 118, // V
		},
		PressedKeys: map[uint8]bool{
			0x1: false,
			0x2: false,
			0x3: false,
			0xc: false,
			0x4: false,
			0x5: false,
			0x6: false,
			0xD: false,
			0x7: false,
			0x8: false,
			0x9: false,
			0xE: false,
			0xA: false,
			0x0: false,
			0xB: false,
			0xF: false,
		},
	}

	return keyboard
}

func (keyboard Keyboard) IsPressed(key uint8) bool {
	return keyboard.PressedKeys[key]
}
