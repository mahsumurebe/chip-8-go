package main

import (
	"chip-8-go/chip8"
	"io/ioutil"
	"log"
	"os"
	p "path/filepath"
	"runtime"
)

var chip *chip8.CHIP8

func readProgram() *[]uint8 {
	pwd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}
	if len(os.Args) < 2 {
		panic("Please specify the ROM path.")
	}

	path := p.Join(pwd, os.Args[1])

	if data, err := ioutil.ReadFile(path); err == nil {
		return &data
	} else {
		panic(err)
	}

}

func dispose() {
	// Clear SDL
	chip.Renderer.Dispose()
}

func main() {
	runtime.LockOSThread()
	chip = chip8.NewCHIP8()
	defer dispose()
	chip.LoadRom(readProgram())
	chip.Run()
}
