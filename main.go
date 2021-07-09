package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	jt "github.com/sb-child/jsdec-tiny-mod"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 40960)
	inp, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error: 无法从stdin读取: ", err)
	}
	inp = strings.ReplaceAll(inp, "\n", "")
	out, err := base64.StdEncoding.DecodeString(inp)
	outStr := string(out)
	if err != nil {
		fmt.Println("error: base64解码失败")
		panic(err)
	}
	jsdec := jt.Jsdec{}
	err = jsdec.ModInit()
	if err != nil {
		fmt.Println("error: 无法打开 dec.js")
		panic(err)
	}
	err = jsdec.LoadJS()
	if err != nil {
		fmt.Printf("error: 加载 dec.js 时出错: %s", err.Error())
		panic(err)
	}
	ret, err := jsdec.Decrypt(outStr)
	if err != nil {
		fmt.Printf("error: 解密时出错: %s", err.Error())
		panic(err)
	}
	print(ret)
}
