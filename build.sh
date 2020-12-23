#!/usr/bin/env bash

if [[ "$(uname)" == "Darwin" ]]; then
  echo "Cannot be built in Darwin."
  exit 1
elif [[ "$(expr substr $(uname -s) 1 5)" == "Linux" ]]; then
  echo "Building for Linux"
  CGO_ENABLED=1 CC=gcc GOOS=linux GOARCH=amd64 go build -tags static -ldflags "-s -w" -i -o ./bin/chip8 .
elif [[ "$(expr substr $(uname -s) 1 10)" == "MINGW32_NT" ]]; then
  echo "Building for Win32"
  CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=386 go build -tags static -ldflags "-s -w" -i -o ./bin/chip8.exe .
elif [[ "$(expr substr $(uname -s) 1 10)" == "MINGW64_NT" ]] || [[ "$(expr substr $(uname -s) 1 7)" == "MSYS_NT" ]]; then
  echo "Building for Win64"
  CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build -tags static -ldflags "-s -w" -i -o ./bin/chip8.exe .
fi
echo "Build end"
