package chip8

import (
	"fmt"
	"log"
	"math/rand"
)

type CPU struct {
	pause bool
	stop  bool
	chip  *CHIP8
	// Program counter
	PC uint16
	// Instruction register
	I uint16
	// V
	V [0xF + 0x1]byte
	// Stack
	Stack [0xF + 0x1]uint16
	// Stack pointer
	SP int8
}

func NewCPU() *CPU {
	return &CPU{
		pause: false,
		stop:  false,
		PC:    0x200,
		I:     0x0,
		SP:    -1,
		Stack: [0xF + 0x1]uint16{},
		V:     [0xF + 0x1]byte{},
	}
}

const CarryFlag = 0xF

func (cpu *CPU) fetch() uint16 {
	opcode := (uint16(cpu.chip.RAM.Read(cpu.PC)) << 8) | uint16(cpu.chip.RAM.Read(cpu.PC+1))

	//fmt.Printf("Fetchin from %02x to %02x: %04x\n", cpu.PC, cpu.PC+1, opcode)
	return opcode
}

func unknownOpcode(opcode *uint16) {
	log.Fatal("Unknown opcode: ", fmt.Sprintf("%04x", *opcode))
}

func (cpu *CPU) execute(opcode *uint16) {
	cpu.PC += 2
	//fmt.Printf("Execute %004x opcode.\n", opcode)

	X := byte(*opcode & 0x0F00 >> 8)
	Y := byte((*opcode & 0x00F0) >> 4)
	NNN := *opcode & 0x0FFF
	NN := *opcode & 0x00FF
	N := *opcode & 0x000F

	A := *opcode & 0xF000

	switch A {
	case 0x0000:
		switch *opcode {
		case 0x00E0:
			//00E0 - CLS
			cpu.chip.GPU.ram.Clear()
		case 0x00EE:
			//00EE - RET
			cpu.PC = cpu.Stack[cpu.SP]

			cpu.Stack[cpu.SP] = 0
			cpu.SP -= 1
		default:
			unknownOpcode(opcode)
		}
	case 0x1000:
		cpu.PC = NNN
	case 0x2000:
		cpu.SP += 1
		cpu.Stack[cpu.SP] = cpu.PC
		cpu.PC = NNN
	case 0x3000:
		if cpu.V[X] == uint8(NN) {
			cpu.PC += 2
		}
	case 0x4000:
		if cpu.V[X] != uint8(NN) {
			cpu.PC += 2
		}
	case 0x5000:
		if cpu.V[X] == cpu.V[Y] {
			cpu.PC += 2
		}
	case 0x6000:
		cpu.V[X] = uint8(NN)
	case 0x7000:
		cpu.V[X] += uint8(NN)
	case 0x8000:
		switch N {
		case 0x0000:
			cpu.V[X] = cpu.V[Y]
		case 0x0001:
			cpu.V[X] = cpu.V[X] | cpu.V[Y]
		case 0x0002:
			cpu.V[X] = cpu.V[X] & cpu.V[Y]
		case 0x0003:
			cpu.V[X] = cpu.V[X] ^ cpu.V[Y]
		case 0x0004:
			sum := uint16(cpu.V[X]) + uint16(cpu.V[Y])

			if sum > 0xFF {
				cpu.V[CarryFlag] = 0x1
			} else {
				cpu.V[CarryFlag] = 0x0
			}

			cpu.V[X] = uint8(sum)

		case 0x5:
			if cpu.V[X] > cpu.V[Y] {
				cpu.V[CarryFlag] = 0x1
			} else {
				cpu.V[CarryFlag] = 0x0
			}

			cpu.V[X] -= cpu.V[Y]
		case 0x6:
			cpu.V[CarryFlag] = cpu.V[X] & 0x01
			cpu.V[X] >>= 1
		case 0x7:
			if cpu.V[X] > cpu.V[Y] {
				cpu.V[CarryFlag] = 0x0
			} else {
				cpu.V[CarryFlag] = 0x1
			}

			cpu.V[X] = cpu.V[Y] - cpu.V[X]
		case 0xE:
			cpu.V[CarryFlag] = cpu.V[X] & 0x80
			cpu.V[X] <<= 1
		default:
			unknownOpcode(opcode)
		}

	case 0x9000:
		if cpu.V[X] != cpu.V[Y] {
			cpu.PC += 2
		}
	case 0xA000:
		cpu.I = NNN
	case 0xB000:
		cpu.PC = uint16(cpu.V[0x0]) + NNN
	case 0xC000:
		cpu.V[X] = uint8(rand.Int()%(0xFF)) & uint8(NN)
	case 0xD000:
		width := byte(8)
		height := N

		cpu.V[CarryFlag] = 0
		for row := uint16(0); row < height; row++ {
			spriteRow := cpu.chip.RAM.Read(cpu.I + row)
			for col := byte(0); col < width; col++ {
				spriteCol := spriteRow & 0x80 // 0X80 = 1000 0000

				if spriteCol != 0x0 {
					if cpu.chip.GPU.SetPixel(cpu.V[X]+col, cpu.V[Y]+byte(row)) == true {
						cpu.V[CarryFlag] = 1
					}
				}

				spriteRow = spriteRow << 1
			}
		}

	case 0xE000:
		switch NN {
		case 0x9E:
			if cpu.chip.KB.IsPressed(cpu.V[X]) {
				cpu.PC += 2
			}
		case 0xA1:
			if !cpu.chip.KB.IsPressed(cpu.V[X]) {
				cpu.PC += 2
			}
		}
	case 0xF000:
		switch NN {
		case 0x07:
			cpu.V[X] = cpu.chip.delayTimer
		case 0x0A:
			//cpu.pause = true
			// until catching pause catch key
		case 0x15:
			cpu.chip.delayTimer = cpu.V[X]
		case 0x18:
			cpu.chip.soundTimer = cpu.V[X]
		case 0x1E:
			cpu.I += uint16(cpu.V[X])
		case 0x29:
			cpu.I = uint16(cpu.V[X]) * 0x5
		case 0x33:
			number := &cpu.V[X]
			hundredsDigit := *number / 100
			tensDigit := (*number % 100) / 10
			onesDigit := *number % 10

			cpu.chip.RAM.Write(cpu.I, hundredsDigit)
			cpu.chip.RAM.Write(cpu.I+1, tensDigit)
			cpu.chip.RAM.Write(cpu.I+2, onesDigit)
		case 0x55:
			for i := byte(0); i <= X; i++ {
				cpu.chip.RAM.Write(cpu.I+uint16(i), cpu.V[i])
			}
		case 0x65:
			for i := byte(0); i <= X; i++ {
				cpu.V[i] = cpu.chip.RAM.Read(cpu.I + uint16(i))
			}
		default:
			unknownOpcode(opcode)
		}
	default:
		unknownOpcode(opcode)
	}
}

func (cpu *CPU) beep() {

}

func (cpu *CPU) updateTimers() {
	if cpu.chip.delayTimer > 0 {
		cpu.chip.delayTimer -= 1
	}
	if cpu.chip.soundTimer > 0 {
		cpu.chip.soundTimer -= 1

		go cpu.beep()
	}
}
func (cpu *CPU) cycle() {
	opcode := cpu.fetch()
	cpu.execute(&opcode)
}
