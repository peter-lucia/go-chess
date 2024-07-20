#!/bin/bash

cd ../
echo "Setting up ./engine"
mkdir -p engine
cd engine/
go mod init github.com/peter-lucia/go-chess/engine
go mod edit -replace github.com/peter-lucia/go-chess/engine=../engine


cd ../
echo "Setting up ./ui"
mkdir -p ui
go mod init github.com/peter-lucia/go-chess/ui
go mod edit -replace github.com/peter-lucia/go-chess/ui=../ui

cd ../
echo "Setting up ./engine"
mkdir -p runner
cd runner/
go mod init github.com/peter-lucia/go-chess/runner
go get ../engine
go get ../ui
go mod tidy

cd ../

go work init ./engine ./ui ./runner

go run ./runner