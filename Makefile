output_path = $(PWD)/bin
all: compile

ifeq ($(OS),Windows_NT)
#Windows
compile:
		CGO_ENABLED=1 CGO_LDFLAGS="-lmingw32 -lSDL2" CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build -tags static -ldflags "-s -w" -o $(output_path)/chip8-windows-64.exe .
else
#Linux
compile:
	CGO_ENABLED=1 CGO_LDFLAGS="-lSDL2" GOOS=linux GOARCH=amd64 CC=gcc go build -tags static -ldflags -o $(output_path)/chip8-linux-amd64 .

endif