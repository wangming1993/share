FROM ubuntu:14.04

WORKDIR /mike

COPY bin/protoc /usr/bin/protoc
COPY bin/protoc-gen-go /usr/bin/protoc-gen-go
COPY bin/protoc-gen-grpc-gateway /usr/bin/protoc-gen-grpc-gateway
COPY bin/protoc-gen-swagger /usr/bin/protoc-gen-swagger

RUN chmod +x /usr/bin/protoc
RUN chmod +x /usr/bin/protoc-gen-go
RUN chmod +x /usr/bin/protoc-gen-grpc-gateway
RUN chmod +x /usr/bin/protoc-gen-swagger

