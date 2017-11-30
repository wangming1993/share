## docker image

### protoc

1. 从 https://github.com/google/protobuf/releases 下载可执行文件 **protoc**

2. 通过命令：
```shell
$ go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
```

生成可执行文件 **protoc-gen-go**

在上面得到的可执行文件拷贝到 `protoc/bin` 目录， 然后执行脚本：

```shell
$ bash buildimage.sh
```
