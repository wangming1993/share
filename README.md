# 基于 go + grpc + consul 的微服务系统

[![Build Status](https://travis-ci.org/wangming1993/share.svg?branch=master)](https://travis-ci.org/wangming1993/share)

## 启动 consul

```shell
docker pull consul

docker run -d -p 8500:8500 consul

# 访问 http://127.0.0.1:8500/ui/#/dc1/services 
```
 
## 生成桩文件

```shell
cd $GOPATH/src/github.com/wangming1993/share/grpc/proto

docker pull registry.cn-hangzhou.aliyuncs.com/wangming/protoc:1.0

./gen-stub
```

## 启动 grpc 服务器

```shell
cd $GOPATH/src/github.com/wangming1993/share/grpc/server

# 启动hello service
go run hello/main.go

# 启动 member service
go run member/main.go

# 启动多个grpc service 以测试负载均衡
go run hello/main.go --port=1801
go run member/main.go --port=1802

```


## 客户端调用

```shell
cd $GOPATH/src/github.com/wangming1993/share/grpc/client

go run main.go
```

