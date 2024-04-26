#!/bin/bash

mkdir -p bin

echo Building...
env GOOS="linux" GOARCH="amd64" go build -trimpath -ldflags "-s -w" -a -o bin/nlrme .
env GOOS="windows" GOARCH="amd64" go build -trimpath -ldflags "-s -w" -a -o bin/nlrme.exe .

# echo Packing...
# rm -f bin/nlrme-small
# upx --best --lzma -o bin/nlrme-small bin/nlrme

echo Done.

echo
ls -hl bin/nlrme
ls -hl bin/nlrme.exe
# ls -hl bin/nlrme-small

echo
file bin/nlrme
file bin/nlrme.exe
# file bin/nlrme-small
