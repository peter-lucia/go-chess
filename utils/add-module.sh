	#!/bin/bash

mkdir -p $1 && cd $1
go mod init github.com/peter-lucia/go-chess/$1
touch $1.go
cd ../
go work use $1
cd ./runner
go get github.com/peter-lucia/go-chess/$1
go mod edit -replace github.com/peter-lucia/go-chess/$1=../$1
go mod tidy  # add missing and remove unused modules
cd ../
go work sync
echo "Finished creating new module $1 and added it to runner"
