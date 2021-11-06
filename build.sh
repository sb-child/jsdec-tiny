#! /bin/bash
go get -d -v ./...
go build main.go
mkdir build
mv main build/
# cp dec.js build/
