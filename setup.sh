#!/bin/bash
echo "go build main.go"
go build main.go

# setup a chess library
mkdir -p engine
cd engine/
go mod init github.com/peter-lucia/go-chess/engine
go mod edit -replace github.com/peter-lucia/go-chess/engine=../engine

cd ../

mkdir -p runner
cd runner/
go mod init github.com/peter-lucia/go-chess/runner
go get ../engine
go mod tidy

cd ../