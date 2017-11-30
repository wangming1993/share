# 基于 go + grpc + consul 的微服务系统

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

docker run -i --rm \
    -v $GOPATH/src/github.com/wangming1993/share/grpc/proto:/srv/proto \
    -w /srv/proto \
    registry.cn-hangzhou.aliyuncs.com/wangming/protoc:1.0 \
    bash -c "protoc -I=. *.proto --go_out=plugins=grpc:."
```

## 启动 grpc 服务器

```shell
cd $GOPATH/src/github.com/wangming1993/share/grpc/server

go run main.go --port=1701

# 启动多个grpc service 以测试负载均衡
go run main.go --port=1702

```


## 客户端调用

```shell
cd $GOPATH/src/github.com/wangming1993/share/grpc/client

go run main.go
```

