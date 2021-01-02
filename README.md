# CHIP-8 Interpreter/Emulator
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/mahsumurebe/chip-8-go?style=for-the-badge)
![GitHub all releases](https://img.shields.io/github/downloads/mahsumurebe/chip-8-go/total?style=for-the-badge)

![issues-open](https://img.shields.io/github/issues/mahsumurebe/chip-8-go?style=for-the-badge)
![issues-closed](https://img.shields.io/github/issues-closed/mahsumurebe/chip-8-go?style=for-the-badge)
![license](https://img.shields.io/github/license/mahsumurebe/chip-8-go?style=for-the-badge)

> CHIP-8 is an interpreted programming language, developed by Joseph Weisbecker. It was initially used on the COSMAC VIP and Telmac 1800 8-bit microcomputers in the mid-1970s. CHIP-8 programs are run on a CHIP-8 virtual machine. It was made to allow video games to be more easily programmed for these computers. [Wiki](https://en.wikipedia.org/wiki/CHIP-8)

## Building

### Requiretments

In order for you to compile, the SDL library must be defined on your computer. You must use `mingw` compiler to building
for windows.

[Install SDL](https://github.com/veandco/go-sdl2#requirements)

### Windows

```shell script
CGO_ENABLED=1 CGO_LDFLAGS="-lmingw32 -lSDL2" CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build -tags static -ldflags "-s -w" -o ./bin/chip8.exe .
```

### Linux

```shell script
CGO_ENABLED=1 CGO_LDFLAGS="-lSDL2" GOOS=linux GOARCH=amd64 CC=gcc go build -tags static -ldflags -o ./bin/chip8 .
```

## Games

You can download [here](https://www.zophar.net/pdroms/chip8/chip-8-games-pack.html)


## How To Run ?
To start a game, define the rom path to the first argument of the compiled binary.

***Example:***
```shell script
./bin/chip8.exe games/PONG
```

***PS:*** ROM files should be located under the directory where you run the binary file. Below is an example directory tree for the example above.

```text
.
|-bin
|---chip8.exe <-- Binary file
|-chip8
|-games
|---PONG <-- ROM
```

## Thanks

Thank you very much to [Mehmet Okan Taştan](https://github.com/motastan95), who was with me during the development process.