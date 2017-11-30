FROM ubuntu:14.04

WORKDIR /mike

COPY bin/protoc /usr/bin/protoc
COPY bin/protoc-gen-go /usr/bin/protoc-gen-go

RUN chmod +x /usr/bin/protoc
RUN chmod +x /usr/bin/protoc-gen-go

