package main

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"strings"

	"github.com/robertkrimen/otto"
)

func main() {
	if len(os.Args) == 1 {
		return
	}
	inp := os.Args[1]
	inp = strings.ReplaceAll(inp, "\n", "")
	out, err := base64.StdEncoding.DecodeString(inp)
	outStr := string(out)
	if err != nil {
		panic(err)
	}
	file, err := os.Open("dec.js")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileByte, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fileString := string(fileByte)
	vm := otto.New()
	_, err = vm.Run(fileString)
	if err != nil {
		panic(err)
	}
	ret, err := vm.Call("autoscan", nil, outStr)
	if err != nil {
		panic(err)
	}
	print(ret.String())
}
