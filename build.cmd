go get -d -v ./...
go build main.go
mkdir build
move main.exe build\
copy pkg\jsdec_tiny\dec.js build\
