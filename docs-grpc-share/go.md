[slide data-transition="vertical3d"]
# WHY GO !!!

[slide]

## go 的优势 

- 开源的，大公司做背书 
- 编译快，静态编译，极易部署
- 跨平台
- 语法简单
- 天然的并发

[slide]

# Go 的基本使用

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    go fmt.Println("go")
    time.sleep(time.Second)

    fmt.Println("hello world")
}
```

```go
$ go build hello-world.go
$ ls
hello-world hello-world.go

$ ./hello-world
go
hello world
```

[slide]
[magic data-transition="cover-circle"]
-----

# go 的基本概念

- package
- interface
- 匿名函数，闭包
- 结构体，指针
- 管道
- 协程

====

<div>
    <img src="/img/go-import-order.jpeg" width="1000px;">
</div>

[/magic]
