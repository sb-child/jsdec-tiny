package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/robertkrimen/otto"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("error: 需要参数")
		panic("need more args")
	}
	inp := os.Args[1]
	inp = strings.ReplaceAll(inp, "\n", "")
	out, err := base64.StdEncoding.DecodeString(inp)
	outStr := string(out)
	if err != nil {
		fmt.Println("error: base64解码失败")
		panic(err)
	}
	file, err := os.Open("dec.js")
	if err != nil {
		fmt.Println("error: 无法打开 dec.js")
		panic(err)
	}
	defer file.Close()
	fileByte, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("error: 无法读取 dec.js")
		panic(err)
	}
	fileString := string(fileByte)
	vm := otto.New()
	_, err = vm.Run(fileString)
	if err != nil {
		fmt.Printf("error: 加载 dec.js 时出错: %s", err.Error())
		panic(err)
	}
	ret, err := vm.Call("autoscan", nil, outStr)
	if err != nil {
		fmt.Printf("error: 解密时出错: %s", err.Error())
		panic(err)
	}
	print(ret.String())
}
