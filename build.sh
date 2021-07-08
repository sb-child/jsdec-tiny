#! /bin/bash
go get -d -v ./...
go build main.go
mkdir build
mv main build/
cp pkg/jsdec_tiny/dec.js build/
