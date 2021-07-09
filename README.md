# jsdec-tiny
a tiny version for https://github.com/liulihaocai/JSDec/ | jsdec 缩小版
## 如何使用
### 在golang项目中使用
在你的项目下运行命令
```bash
go get -v -u -t github.com/sb-child/jsdec-tiny-mod@latest
```
然后在go代码中导入
```go
import jt "github.com/sb-child/jsdec-tiny-mod"
```
即可使用
```go
jsdec := jt.Jsdec{}
err := jsdec.ModInit()
# ...
```

### 编译成命令行工具
<需要事先安装`golang`>  
+ windows用户:
  ```bash
    # 运行 build.cmd
  ```
+ linux用户:
  ```bash
    ./build.sh
  ```
#### 运行
+ windows用户:
  ```cmd
    cd build
    main.exe
    (输入要解密的js的base64或使用管道, 以换行符作为结束符)
  ```
+ linux用户:
  ```bash
    cd build
    ./main
    (输入要解密的js的base64或使用管道, 以换行符作为结束符)
  ```
