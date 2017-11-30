title: 使用 go + grpc 打造微服务系统
speaker: Mike Wang
url: https://github.com/ksky521/nodeppt
transition: cards
files: /js/demo.js,/css/demo.css
theme: moon

[slide]

# 基于 go + grpc + consul 的微服务系统
## 演讲者：Mike Wang

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

import "fmt"

func main() {
    fmt.Println("hello world")
}
```

```go
$ go build hello-world.go
$ ls
hello-world hello-world.go

$ ./hello-world
hello world
```

[slide]

# go 的基本概念

- package
- interface
- 匿名函数，闭包
- 结构体，指针
- 管道
- 协程

[slide]

# 什么是微服务

微服务是一种架构风格，一个大型复杂软件应用由一个或多个微服务组成。系统中的各个微服务可被独立部署，各个微服务之间是松耦合的。每个微服务仅关注于完成一件任务并很好地完成该任务。在所有情况下，每个任务代表着一个小的业务能力

[slide]

## 微服务架构的优点：

- 每个服务都比较简单，只关注于一个业务功能。
- 微服务架构方式是松耦合的，可以提供更高的灵活性。
- 微服务可通过最佳及最合适的不同的编程语言与工具进行开发，能够做到有的放矢地解决针对性问题。
- 每个微服务可由不同团队独立开发，互不影响，加快推出市场的速度。
- 微服务架构是持续交付(CD)的巨大推动力，允许在频繁发布不同服务的同时保持系统其他部分的可用性和稳定性

[slide]
[magic data-transition="cover-circle"]
-----

## 单体架构图
<div>
    <img src="/img/monolithic.png" width="600px;">
</div>

====
## 微服务架构图
<div>
    <img src="/img/microservice.png" width="800px;">
</div>

[/magic]

[slide]

# WHY GRPC!!!

[slide]

 ## Core Features

    - Idiomatic client libraries in 10 languages
    - Highly efficient on wire and with a simple service definition framework
    - Bi-directional streaming with http/2 based transport
    - Pluggable auth, tracing, load balancing and health checking

[slide]

## 使用 consul 实现服务注册与服务发现

[slide]

## middleware

- auth
- tracing (opentracing)
- load balance
- 熔断

[slide]

# THANKS !!!

## Mike Wang
