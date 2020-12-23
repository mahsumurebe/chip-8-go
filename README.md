# CHIP-8 Interpreter/Emulator

> CHIP-8 is an interpreted programming language, developed by Joseph Weisbecker. It was initially used on the COSMAC VIP and Telmac 1800 8-bit microcomputers in the mid-1970s. CHIP-8 programs are run on a CHIP-8 virtual machine. It was made to allow video games to be more easily programmed for these computers. [Wiki](https://en.wikipedia.org/wiki/CHIP-8)

## Building

## Requiretments

In order for you to compile, the SDL library must be defined on your computer. You must use `mingw` compiler to building
for windows

[Install SDL](https://github.com/veandco/go-sdl2#requirements)

### Windows

```shell script
CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build -tags static -ldflags "-s -w" -i -o ./bin/chip8.exe .
```

### Linux

```shell script
CGO_ENABLED=1 CC=gcc GOOS=linux GOARCH=amd64 go build -tags static -ldflags "-s -w" -i -o ./bin/chip8 .
```

## Games

You can download [here](https://www.zophar.net/pdroms/chip8/chip-8-games-pack.html)


## Thanks

Thank you very much to [Mehmet Okan Taştan](https://github.com/motastan), who was with me during the development process.