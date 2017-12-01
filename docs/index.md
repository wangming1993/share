title: 基于 grpc + consul 的微服务系统
speaker: Mike Wang
url: https://github.com/ksky521/nodeppt
transition: cards
files: /js/demo.js,/css/demo.css
theme: moon

[slide]

# 基于 grpc + consul 的微服务系统
## 演讲者：Mike Wang


[slide]

# 什么是微服务

系统由N多可独立部署的服务组成，各个微服务之间是松耦合的。每个微服务仅关注于完成一件任务。在所有情况下，每个任务代表着一个小的业务能力

[slide]

## 微服务架构的优点：

- 每个服务都比较简单，只关注于一个业务功能。
- 微服务架构方式是松耦合的，可以提供更高的灵活性。
- 微服务在语言选择上更灵活，不同服务的团队可以选择自己熟练的语言。
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

## https://grpc.io/

[slide]

![](/img/grpc-features.png)

[slide]

# 服务定义简单, 基于 [Protocol Buffer](https://developers.google.com/protocol-buffers/docs/overview)

- protocol buffers 是一种与语言，平台无关的数据序列化协议，由 google 开源
- 它灵活， 高效， 自动化， 类似于 XML, JSON, 但是更小，更快，更简单

[slide]

```proto
syntax = "proto3";

package proto;

// protoc -I=. *.proto --go_out=plugins=grpc:.

message HelloRequest {
    string greeting = 1;
}

message HelloResponse {
    string reply = 1;
}

service HelloService {
    rpc SayHello(HelloRequest) returns (HelloResponse);
}

```

[slide]

## 多语言支持

![](/img/grpc-languages.png)

[slide]

# 基于 http2 的数据传输

- 多路复用
- 二进制分帧
- 首部压缩
- 服务端推送

[slide]

![](/img/http2.png)

[slide]

## 使用 consul 实现服务注册与服务发现

- Consul 是一个支持多数据中心分布式高可用的服务发现和配置共享的服务软件,用 Go 语言开发
- 一致性协议采用 Raft 算法,用来保证服务的高可用

[slide]

基于[consul-template](https://github.com/hashicorp/consul-template)的配置管理

[slide]

# 与 docker 的完美配合

- [containerpilot](https://github.com/joyent/containerpilot/blob/master/docs/30-configuration/32-configuration-file.md#consul)

[slide]

# 日志服务

## `sls` + `TraceId`

- 每一个 request 会携带一个 `X-Req-ID` 的 `Header`, grpc 调用时会携带在 `context` metadata 中，记录整个链路调用过程 

[slide]

## 监控服务

[kibana](https://www.elastic.co/products/kibana) + [influxdb](https://www.influxdata.com/)

[slide]

## middleware

- go-grpc-middleware 
    - https://github.com/grpc-ecosystem/go-grpc-middleware
- tracing  
    - https://github.com/opentracing/opentracing-go
- 熔断和服务降级 
    - https://github.com/Netflix/Hystrix
- grpc-gateway 
    - https://github.com/grpc-ecosystem/grpc-gateway

[slide]

# THANKS !!!

## Mike Wang
